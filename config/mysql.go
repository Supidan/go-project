package config

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysqlDatabase() *sql.DB {
	var dbHost = viper.GetString(`database.host`)
	var dbPort = viper.GetString(`database.port`)
	var dbName = viper.GetString(`database.name`)
	var dbUser = viper.GetString(`database.user`)
	var dbPass = viper.GetString(`database.pass`)
	var connection = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	count := 0
	dbConn, err := sql.Open(dbHost, connection)
	//Reconnect to database until 3 minute timeout
	if err != nil {
		for {
			if err == nil {
				fmt.Println("")
				break
			}

			fmt.Print("Reconnecting...")
			time.Sleep(time.Second)
			count++

			if count >= 180 {
				fmt.Println("")
				fmt.Println("DB Connection Failure")
				panic(err)
			}

			dbConn, err = sql.Open(dbHost, connection)
		}
	}

	fmt.Println("DB connection successful")

	return dbConn
}

func NewMysqlGormDatabase() *gorm.DB {
	var dbHost = viper.GetString(`database.host`)
	var dbPort = viper.GetString(`database.port`)
	var dbName = viper.GetString(`database.name`)
	var dbUser = viper.GetString(`database.user`)
	var dbPass = viper.GetString(`database.pass`)
	var connection = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	count := 0
	dbConn, err := gorm.Open(mysql.Open(connection), &gorm.Config{})
	//Reconnect to database until 3 minute timeout
	if err != nil {
		for {
			if err == nil {
				fmt.Println("")
				break
			}

			fmt.Print("Reconnecting...")
			time.Sleep(time.Second)
			count++

			if count >= 180 {
				fmt.Println("")
				fmt.Println("DB Connection Failure")
				panic(err)
			}

			dbConn, err = gorm.Open(mysql.Open(connection), &gorm.Config{})
		}
	}

	fmt.Println("DB connection successful")

	return dbConn
}
