package handler

import (
	"config-service/ent"
	"config-service/ent/appconfig"
	"context"
	"errors"
	"log"
	"strconv"

	"entgo.io/ent/dialect/sql"
	"github.com/apialerts/apialerts-go"
	"github.com/go-fuego/fuego"
	"github.com/patrickmn/go-cache"
)

type ConfigResources struct {
	DbClient        *ent.Client
	ApiAlertsClient apialerts.Client
	Cache           *cache.Cache
}

type ConfigToCreate struct{}

type CreateConfigReturn struct{}

type ConfigReturn struct {
	Message string         `json:"message"`
	Data    *ent.AppConfig `json:"data"`
}

type GetConfigAppsReturn struct {
	AppNames []string `json:"appNames"`
}

type CreateConfigDto struct {
	Config map[string]any `json:"config"`
}

type GetAllConfigsForAppReturn struct {
	Configs []*ent.AppConfig `json:"configs"`
}

// fuego.Post(config, "/{AppName}/activate/{Version}")

func (h *ConfigResources) ActivateConfig(c fuego.ContextNoBody) (*ConfigReturn, error) {
	currentUser, ok := c.Context().Value("currentUser").(*ent.User)
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
	if appName == "" {
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

	config, err := h.DbClient.
		AppConfig.
		Query().
		Where(
			appconfig.AppName(appName),
			appconfig.Version(versionInt),
		).First(context.Background())

	if err != nil {
		return nil, fuego.InternalServerError{
			Detail: "Failed to fetch config to activate",
			Err:    err,
		}
	}

	if config.Status != "inactive" {
		return nil, fuego.BadRequestError{
			Detail: "Config is already active",
		}
	}

	oldConfig, err := h.getCurrentConfigForApp(appName)

	if err != nil {
		return nil, fuego.InternalServerError{
			Detail: "Failed to fetch config to activate",
			Err:    err,
		}
	}

	if oldConfig != nil {
		oldConfig.Status = "inactive"
		if err := h.DbClient.AppConfig.UpdateOne(oldConfig).SetStatus("inactive").Exec(context.Background()); err != nil {
			return nil, fuego.InternalServerError{
				Detail: "Error deactivating old config",
				Err:    err,
			}
		}
	}

	config.Status = "active"

	if err := h.DbClient.AppConfig.UpdateOne(config).SetStatus("active").Exec(context.Background()); err != nil {
		return nil, fuego.InternalServerError{
			Detail: "Error activating config",
			Err:    err,
		}
	}

	go h.ApiAlertsClient.Send("Config activated for app: "+appName,
		[]string{appName, version},
		"v1/config/"+appName)

	return &ConfigReturn{
		Message: "Success",
		Data:    config,
	}, nil

}

func (h *ConfigResources) GetAllConfigsForApp(c fuego.ContextNoBody) (*GetAllConfigsForAppReturn, error) {
	currentUser, ok := c.Context().Value("currentUser").(*ent.User)
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

	configs, err := h.DbClient.AppConfig.
		Query().
		Where(appconfig.AppName(appName)).
		All(context.Background())

	if err != nil {
		return nil, fuego.InternalServerError{
			Detail: "Failed to fetch all configs",
			Err:    err,
		}
	}

	return &GetAllConfigsForAppReturn{
		Configs: configs,
	}, nil
}

func (h *ConfigResources) CreateConfig(c fuego.ContextWithBody[CreateConfigDto]) (*ConfigReturn, error) {
	currentUser, ok := c.Context().Value("currentUser").(*ent.User)
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

	previousConfig, err := h.getCurrentConfigForApp(appName)

	if err != nil {
		log.Println(err)
		return nil, fuego.BadRequestError{
			Detail: "Error fetching previous config",
		}
	}

	status := "active"

	if previousConfig != nil {
		configVersion = previousConfig.Version + 1
		status = "inactive"
	}

	config, err := h.DbClient.AppConfig.
		Create().
		SetAppName(appName).
		SetVersion(configVersion).
		SetStatus(status).
		SetConfig(body.Config).
		Save(context.Background())

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
		Data:    config,
	}, nil
}

func (h *ConfigResources) GetConfigApps(c fuego.ContextNoBody) (*GetConfigAppsReturn, error) {
	currentUser, ok := c.Context().Value("currentUser").(*ent.User)
	if !ok {
		return nil, errors.New("current user not found in context")
	}

	if currentUser.Role != "admin" {
		return nil, fuego.ForbiddenError{}
	}
	appNames, err := h.DbClient.AppConfig.
		Query().
		Select(appconfig.FieldAppName).
		GroupBy(appconfig.FieldAppName).
		Strings(context.Background())

	if err != nil {
		log.Println(err)
		return nil, fuego.InternalServerError{
			Err:    err,
			Detail: "Failed to fetch config",
		}
	}

	return &GetConfigAppsReturn{
		AppNames: appNames,
	}, nil
}

func (h *ConfigResources) GetConfig(c fuego.ContextNoBody) (*ConfigReturn, error) {
	currentUser, ok := c.Context().Value("currentUser").(*ent.User)
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
			Data:    cached.(*ent.AppConfig),
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

func (h *ConfigResources) getCurrentConfigForApp(appName string) (*ent.AppConfig, error) {
	config, err := h.DbClient.AppConfig.
		Query().
		Where(
			appconfig.And(
				appconfig.AppName(appName),
				appconfig.Status("active"),
			),
		).
		Order(appconfig.ByVersion(sql.OrderAsc())).
		All(context.Background())

	if err != nil {
		return nil, err
	}

	if len(config) == 0 {
		return nil, nil
	}

	return config[len(config)-1], nil
}
