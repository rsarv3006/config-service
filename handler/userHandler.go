package handler

import (
	"config-service/auth"
	"config-service/config"
	"config-service/ent"
	"config-service/ent/user"
	"context"
	"errors"
	"log"

	"github.com/apialerts/apialerts-go"
	"github.com/go-fuego/fuego"
)

type UserResources struct {
	DbClient        *ent.Client
	ApiAlertsClient apialerts.Client
}

type UserToCreate struct {
}

type CreateUserReturn struct {
	Message string    `json:"message"`
	Data    *ent.User `json:"data"`
	Token   string    `json:"token"`
}

type GetAllUsersReturn struct {
	UserNames []string `json:"userNames"`
}

func (h *UserResources) GetAllUsers(c fuego.ContextNoBody) (*GetAllUsersReturn, error) {
	currentUser, ok := c.Context().Value("currentUser").(*ent.User)
	if !ok {
		return nil, errors.New("current user not found in context")
	}

	if currentUser.Role != "admin" {
		return nil, fuego.ForbiddenError{
			Detail: "Fordbidden",
		}
	}

	users, err := h.DbClient.User.Query().Select(user.FieldAppName).Strings(context.Background())

	if err != nil {
		return nil, fuego.InternalServerError{
			Detail: "Failed to get app names",
			Err:    err,
		}
	}

	return &GetAllUsersReturn{
		UserNames: users,
	}, nil

}

func (h *UserResources) CreateUser(c fuego.ContextWithBody[UserToCreate]) (*CreateUserReturn, error) {
	currentUser, ok := c.Context().Value("currentUser").(*ent.User)
	if !ok {
		return nil, errors.New("current user not found in context")
	}

	if currentUser.Role != "admin" {
		return nil, fuego.ForbiddenError{
			Detail: "Fordbidden",
		}
	}

	// Get AppName from URL params instead of body
	appName := c.PathParam("AppName")
	if appName == "" {
		return nil, fuego.BadRequestError{
			Detail: "AppName parameter is required",
		}
	}

	existingUser, err := h.DbClient.User.Query().Where(user.AppName(appName)).Only(context.Background())
	if err == nil && existingUser != nil {
		return nil, fuego.ConflictError{
			Detail: "A user with the provided app name already exists.",
			Err:    err,
		}
	}

	user, err := h.DbClient.User.
		Create().
		SetAppName(appName).
		SetRole("user").
		Save(context.Background())

	if err != nil {
		log.Println(err)
		return nil, fuego.InternalServerError{
			Detail: "Error Creating User",
			Err:    err,
		}
	}

	jwtSecret := config.Config("JWT_SECRET")
	log.Println("jwtSecret: " + jwtSecret)
	token, err := auth.GenerateJWTFromSecret(user, jwtSecret)

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
		Data:    user,
		Token:   token,
	}, nil

}
