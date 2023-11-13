package services

import (
	"context"
	"fmt"

	"webginrest/server/config"

	"github.com/jinzhu/gorm"

	// Mysql gorm
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB variable global db
var DB *gorm.DB

// OpenDBConnection Open connection Database
func OpenDBConnection(ctx context.Context, dbUser string, dbPass string, dbHost string, dbPort string, dbName string) {

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4,utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)

	var err error
	DB, err = gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}

	DB.DB().SetMaxIdleConns(0)

	if config.GetBoolean(`debug`) {
		DB.LogMode(true)
	}
}

// CloseDBConnection Closing DB Connection
func CloseDBConnection() error {
	return DB.Close()
}
