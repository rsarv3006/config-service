package handler

import (
	"config-service/model"
	"context"
	"errors"
	"log"
	"strconv"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/apialerts/apialerts-go"
	"github.com/go-fuego/fuego"
	"github.com/google/uuid"
	"github.com/patrickmn/go-cache"
	"google.golang.org/api/iterator"
)

const appConfigsCollection = "ConfigService-app_configs"

type ConfigResources struct {
	Client          *firestore.Client
	ApiAlertsClient apialerts.Client
	Cache           *cache.Cache
}

type ConfigToCreate struct{}

type CreateConfigReturn struct{}

type ConfigReturn struct {
	Message string           `json:"message"`
	Data    *model.AppConfig `json:"data"`
}

type GetConfigAppsReturn struct {
	AppNames []string `json:"appNames"`
}

type CreateConfigDto struct {
	Config map[string]any `json:"config"`
}

type GetAllConfigsForAppReturn struct {
	Configs []*model.AppConfig `json:"configs"`
}

func (h *ConfigResources) ActivateConfig(c fuego.ContextNoBody) (*ConfigReturn, error) {
	currentUser, ok := c.Context().Value("currentUser").(*model.User)
	if !ok {
		return nil, errors.New("current user not found in context")
	}

	if currentUser.Role != "admin" {
		return nil, fuego.ForbiddenError{}
	}

	appName := c.PathParam("AppName")
	if appName == "" {
		return nil, fuego.BadRequestError{
			Detail: "AppName parameter is required",
		}
	}

	version := c.PathParam("Version")
	if version == "" {
		return nil, fuego.BadRequestError{
			Detail: "Version parameter is required",
		}
	}

	versionInt, err := strconv.Atoi(version)
	if err != nil {
		return nil, fuego.BadRequestError{
			Detail: "Invalid version number submitted",
			Err:    err,
		}
	}

	ctx := context.Background()

	// Find config to activate
	iter := h.Client.Collection(appConfigsCollection).
		Where("app_name", "==", appName).
		Where("version", "==", versionInt).
		Limit(1).
		Documents(ctx)
	defer iter.Stop()

	doc, err := iter.Next()
	if err == iterator.Done {
		return nil, fuego.NotFoundError{}
	}
	if err != nil {
		return nil, fuego.InternalServerError{
			Detail: "Failed to fetch config to activate",
			Err:    err,
		}
	}

	var config model.AppConfig
	if err := doc.DataTo(&config); err != nil {
		return nil, fuego.InternalServerError{
			Detail: "Failed to parse config",
			Err:    err,
		}
	}

	if config.Status != "inactive" {
		return nil, fuego.BadRequestError{
			Detail: "Config is already active",
		}
	}

	// Deactivate old config
	oldConfig, err := h.getCurrentConfigForApp(appName)
	if err != nil {
		return nil, fuego.InternalServerError{
			Detail: "Failed to fetch current config",
			Err:    err,
		}
	}

	if oldConfig != nil {
		_, err = h.Client.Collection(appConfigsCollection).Doc(oldConfig.ID).
			Update(ctx, []firestore.Update{
				{Path: "status", Value: "inactive"},
			})
		if err != nil {
			return nil, fuego.InternalServerError{
				Detail: "Error deactivating old config",
				Err:    err,
			}
		}
	}

	// Activate new config
	_, err = h.Client.Collection(appConfigsCollection).Doc(config.ID).
		Update(ctx, []firestore.Update{
			{Path: "status", Value: "active"},
		})
	if err != nil {
		return nil, fuego.InternalServerError{
			Detail: "Error activating config",
			Err:    err,
		}
	}

	config.Status = "active"

	go h.ApiAlertsClient.Send("Config activated for app: "+appName,
		[]string{appName, version},
		"v1/config/"+appName)

	return &ConfigReturn{
		Message: "Success",
		Data:    &config,
	}, nil
}

func (h *ConfigResources) GetAllConfigsForApp(c fuego.ContextNoBody) (*GetAllConfigsForAppReturn, error) {
	currentUser, ok := c.Context().Value("currentUser").(*model.User)
	if !ok {
		return nil, errors.New("current user not found in context")
	}

	if currentUser.Role != "admin" {
		return nil, fuego.ForbiddenError{}
	}

	appName := c.PathParam("AppName")
	if appName == "" {
		return nil, fuego.BadRequestError{
			Detail: "AppName parameter is required",
		}
	}

	ctx := context.Background()
	iter := h.Client.Collection(appConfigsCollection).
		Where("app_name", "==", appName).
		Documents(ctx)
	defer iter.Stop()

	var configs []*model.AppConfig
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fuego.InternalServerError{
				Detail: "Failed to fetch all configs",
				Err:    err,
			}
		}

		var config model.AppConfig
		if err := doc.DataTo(&config); err != nil {
			return nil, fuego.InternalServerError{
				Detail: "Failed to parse config",
				Err:    err,
			}
		}
		configs = append(configs, &config)
	}

	return &GetAllConfigsForAppReturn{
		Configs: configs,
	}, nil
}

func (h *ConfigResources) CreateConfig(c fuego.ContextWithBody[CreateConfigDto]) (*ConfigReturn, error) {
	currentUser, ok := c.Context().Value("currentUser").(*model.User)
	if !ok {
		return nil, errors.New("current user not found in context")
	}

	if currentUser.Role != "admin" {
		return nil, fuego.ForbiddenError{}
	}

	appName := c.PathParam("AppName")
	if appName == "" {
		return nil, fuego.BadRequestError{
			Detail: "AppName parameter is required",
		}
	}

	body, err := c.Body()
	if err != nil {
		return nil, err
	}

	configVersion := 1
	status := "active"

	previousConfig, err := h.getCurrentConfigForApp(appName)
	if err != nil {
		log.Println(err)
		return nil, fuego.BadRequestError{
			Detail: "Error fetching previous config",
		}
	}

	if previousConfig != nil {
		configVersion = previousConfig.Version + 1
		status = "inactive"
	}

	config := model.AppConfig{
		ID:        uuid.New().String(),
		CreatedAt: time.Now(),
		AppName:   appName,
		Version:   configVersion,
		Status:    status,
		Config:    body.Config,
	}

	ctx := context.Background()
	_, err = h.Client.Collection(appConfigsCollection).Doc(config.ID).Set(ctx, config)
	if err != nil {
		return nil, fuego.BadRequestError{
			Detail: "Error Creating the App Config",
			Err:    err,
		}
	}

	versionString := strconv.Itoa(configVersion)
	go h.ApiAlertsClient.Send("Config created for app: "+appName,
		[]string{appName, versionString},
		"v1/config/"+appName)

	return &ConfigReturn{
		Message: "Success",
		Data:    &config,
	}, nil
}

func (h *ConfigResources) GetConfigApps(c fuego.ContextNoBody) (*GetConfigAppsReturn, error) {
	currentUser, ok := c.Context().Value("currentUser").(*model.User)
	if !ok {
		return nil, errors.New("current user not found in context")
	}

	if currentUser.Role != "admin" {
		return nil, fuego.ForbiddenError{}
	}

	ctx := context.Background()
	iter := h.Client.Collection(appConfigsCollection).Documents(ctx)
	defer iter.Stop()

	appNamesMap := make(map[string]bool)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Println(err)
			return nil, fuego.InternalServerError{
				Err:    err,
				Detail: "Failed to fetch configs",
			}
		}

		var config model.AppConfig
		if err := doc.DataTo(&config); err != nil {
			continue
		}
		appNamesMap[config.AppName] = true
	}

	appNames := make([]string, 0, len(appNamesMap))
	for name := range appNamesMap {
		appNames = append(appNames, name)
	}

	return &GetConfigAppsReturn{
		AppNames: appNames,
	}, nil
}

func (h *ConfigResources) GetConfig(c fuego.ContextNoBody) (*ConfigReturn, error) {
	currentUser, ok := c.Context().Value("currentUser").(*model.User)
	if !ok {
		return nil, errors.New("current user not found in context")
	}

	appName := c.PathParam("AppName")
	if appName == "" {
		return nil, fuego.BadRequestError{
			Detail: "AppName parameter is required",
		}
	}

	if currentUser.Role != "admin" && currentUser.AppName != appName {
		return nil, fuego.ForbiddenError{}
	}

	if cached, found := h.Cache.Get(appName); found {
		return &ConfigReturn{
			Message: "Success",
			Data:    cached.(*model.AppConfig),
		}, nil
	}

	config, err := h.getCurrentConfigForApp(appName)
	if err != nil {
		log.Println(err)
		return nil, fuego.InternalServerError{
			Err:    err,
			Detail: "Failed to fetch config",
		}
	}

	if config == nil {
		return nil, fuego.NotFoundError{}
	}

	h.Cache.Set(appName, config, cache.DefaultExpiration)

	return &ConfigReturn{
		Message: "Success",
		Data:    config,
	}, nil
}

func (h *ConfigResources) getCurrentConfigForApp(appName string) (*model.AppConfig, error) {
	ctx := context.Background()
	iter := h.Client.Collection(appConfigsCollection).
		Where("app_name", "==", appName).
		Where("status", "==", "active").
		OrderBy("version", firestore.Asc).
		Documents(ctx)
	defer iter.Stop()

	var configs []*model.AppConfig
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		var config model.AppConfig
		if err := doc.DataTo(&config); err != nil {
			return nil, err
		}
		configs = append(configs, &config)
	}

	if len(configs) == 0 {
		return nil, nil
	}

	return configs[len(configs)-1], nil
}
