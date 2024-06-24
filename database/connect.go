package database

import (
	"RjsConfigService/config"
	"RjsConfigService/ent"
	"context"
	"log"

	_ "github.com/lib/pq"
)

func Connect() *ent.Client {
	host := config.Config("DB_HOST")
	port := config.Config("DB_PORT")
	user := config.Config("DB_USER")
	database := config.Config("DB_NAME")
	pass := config.Config("DB_PASS")
	sslMode := config.Config("DB_SSL_MODE")

	log.Println("Connecting to database...")

	connString := "host=" + host + " port=" + port + " user=" + user + " dbname=" + database + " password=" + pass + " sslmode=" + sslMode

	client, err := ent.Open("postgres", connString)

	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client
}
