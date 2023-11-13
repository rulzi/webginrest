package main

import (
	"context"
	"fmt"
	"webginrest/database/seeder/resources"
	"webginrest/server/config"
	"webginrest/server/services"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	fmt.Println("seeder start")

	dbHost := config.Get("database.host")
	dbPort := config.Get(`database.port`)
	dbUser := config.Get(`database.user`)
	dbPass := config.Get(`database.pass`)
	dbName := config.Get(`database.name`)
	services.OpenDBConnection(ctx, dbUser, dbPass, dbHost, dbPort, dbName)
	defer services.CloseDBConnection()

	users := resources.User()

	for _, user := range users {
		services.DB.Save(user)
	}
}
