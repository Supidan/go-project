package migration

import (
	"example/go-project/db/seeds"
	"example/go-project/entity"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func Migrate(dbGorm *gorm.DB) {

	sqlDB, err := dbGorm.DB()
	isSeeded := viper.GetBool("seeded")

	if err != nil {
		panic(err)
	}

	err = dbGorm.AutoMigrate(
		entity.Buyer{},
		entity.Seller{},
		entity.Product{},
		entity.Order{},
		entity.OrderDetail{},
	)

	if err == nil && !isSeeded {
		seeds.Seeds(sqlDB)
	}
}
