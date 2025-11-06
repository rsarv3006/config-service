package handler

import (
	"config-service/auth"
	"config-service/config"
	"config-service/model"
	"context"
	"errors"
	"log"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/apialerts/apialerts-go"
	"github.com/go-fuego/fuego"
	"github.com/google/uuid"
	"google.golang.org/api/iterator"
)

const usersCollection = "users"

type UserResources struct {
	Client          *firestore.Client
	ApiAlertsClient apialerts.Client
}

type UserToCreate struct{}

type CreateUserReturn struct {
	Message string      `json:"message"`
	Data    *model.User `json:"data"`
	Token   string      `json:"token"`
}

type GetAllUsersReturn struct {
	UserNames []string `json:"userNames"`
}

func (h *UserResources) GetAllUsers(c fuego.ContextNoBody) (*GetAllUsersReturn, error) {
	currentUser, ok := c.Context().Value("currentUser").(*model.User)
	if !ok {
		return nil, errors.New("current user not found in context")
	}

	if currentUser.Role != "admin" {
		return nil, fuego.ForbiddenError{
			Detail: "Forbidden",
		}
	}

	ctx := context.Background()
	iter := h.Client.Collection(usersCollection).Documents(ctx)
	defer iter.Stop()

	var userNames []string
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fuego.InternalServerError{
				Detail: "Failed to get app names",
				Err:    err,
			}
		}

		var user model.User
		if err := doc.DataTo(&user); err != nil {
			continue
		}
		if user.AppName != "" {
			userNames = append(userNames, user.AppName)
		}
	}

	return &GetAllUsersReturn{
		UserNames: userNames,
	}, nil
}

func (h *UserResources) CreateUser(c fuego.ContextWithBody[UserToCreate]) (*CreateUserReturn, error) {
	currentUser, ok := c.Context().Value("currentUser").(*model.User)
	if !ok {
		return nil, errors.New("current user not found in context")
	}

	if currentUser.Role != "admin" {
		return nil, fuego.ForbiddenError{
			Detail: "Forbidden",
		}
	}

	// Get AppName from URL params instead of body
	appName := c.PathParam("AppName")
	if appName == "" {
		return nil, fuego.BadRequestError{
			Detail: "AppName parameter is required",
		}
	}

	ctx := context.Background()

	// Check if user already exists
	iter := h.Client.Collection(usersCollection).
		Where("app_name", "==", appName).
		Limit(1).
		Documents(ctx)
	defer iter.Stop()

	_, err := iter.Next()
	if err != iterator.Done {
		// User exists or there was an error
		if err != nil {
			return nil, fuego.InternalServerError{
				Detail: "Error checking for existing user",
				Err:    err,
			}
		}
		return nil, fuego.ConflictError{
			Detail: "A user with the provided app name already exists.",
		}
	}

	// Create new user
	user := model.User{
		ID:        uuid.New().String(),
		CreatedAt: time.Now(),
		Role:      "user",
		AppName:   appName,
		Status:    "active",
	}

	_, err = h.Client.Collection(usersCollection).Doc(user.ID).Set(ctx, user)
	if err != nil {
		log.Println(err)
		return nil, fuego.InternalServerError{
			Detail: "Error Creating User",
			Err:    err,
		}
	}

	jwtSecret := config.Config("JWT_SECRET")
	log.Println("jwtSecret: " + jwtSecret)

	token, err := auth.GenerateJWTFromSecret(&user, jwtSecret)
	if err != nil {
		log.Println(err)
		return nil, fuego.InternalServerError{
			Detail: "Error Generating JWT for user",
		}
	}

	log.Println("token: " + token)
	log.Println("apiAlertsClient: ")

	go h.ApiAlertsClient.Send("User created for app: "+appName,
		[]string{appName},
		"v1/config/"+appName)

	return &CreateUserReturn{
		Message: "Success",
		Data:    &user,
		Token:   token,
	}, nil
}

