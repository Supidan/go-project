package main

import (
	"log"
	"net/http"

	"example/go-project/config"
	"example/go-project/db/migration"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG Mode")
	}
}

func main() {
	//Setup Database

	db := config.NewMysqlGormDatabase()

	//Start Migration
	migration.Migrate(db)

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World test!")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
