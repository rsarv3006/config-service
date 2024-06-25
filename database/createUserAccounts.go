package database

import (
	"RjsConfigService/auth"
	"RjsConfigService/ent"
	"RjsConfigService/ent/user"
	"context"
	"log"
)

func CreateUserAccounts(dbClient *ent.Client, jwtSecret string) {
	log.Println("Creating admin account")
	adminAccountToken, err := createAdminAccount(dbClient, jwtSecret)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(adminAccountToken)
	log.Println("Admin account created")

	appAccounts, err := CreateAppAccounts(dbClient, jwtSecret)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(appAccounts)
	log.Println("App accounts created")

}

func createAdminAccount(dbClient *ent.Client, jwtSecret string) (string, error) {
	doesUserExist, err := dbClient.User.
		Query().
		Where(user.Role("admin")).
		Exist(context.Background())

	if err != nil {
		return "", err
	}

	if doesUserExist {
		return "", nil
	}

	adminAccount, err := dbClient.User.
		Create().
		SetRole("admin").
		SetAppName("").
		Save(context.Background())

	if err != nil {
		return "", err
	}

	return auth.GenerateJWTFromSecret(adminAccount, jwtSecret)
}

var appAccounts = []string{"basketbuddy", "wishlistwrangler"}

type CreatedAppAccount struct {
	Name  string
	Token string
}

func CreateAppAccounts(dbClient *ent.Client, jwtSecret string) ([]CreatedAppAccount, error) {
	log.Println("Creating app accounts")
	var errorFromProcess error

	createdAccounts := []CreatedAppAccount{}

	for _, appName := range appAccounts {
		doesUserExist, err := dbClient.User.
			Query().
			Where(user.AppName(appName)).
			Exist(context.Background())

		if err != nil {
			errorFromProcess = err
		}

		if doesUserExist {
			continue
		}

		appAccount, err := dbClient.User.
			Create().
			SetRole("app").
			SetAppName(appName).
			Save(context.Background())

		if err != nil {
			errorFromProcess = err
		}

		token, err := auth.GenerateJWTFromSecret(appAccount, jwtSecret)

		if err != nil {
			errorFromProcess = err
		}

		createdAccounts = append(createdAccounts, CreatedAppAccount{
			Name:  appName,
			Token: token,
		})

	}
	log.Println("App accounts created")
	return createdAccounts, errorFromProcess
}
