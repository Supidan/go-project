package seeds

import (
	"database/sql"
	"log"

	"github.com/kristijorgji/goseeder"
	"github.com/spf13/viper"
)

func Seeds(db *sql.DB) {

	log.Println("Seeding database")
	err := goseeder.Execute(db)
	if err != nil {
		log.Fatal("Seeding test data failed\n")
	} else {
		viper.Set("seeded", true)
		viper.WriteConfig()
		log.Println("Seeding success")
	}
}
