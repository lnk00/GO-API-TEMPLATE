package db

import (
	"fmt"
	"go-api-template/services/config"
	"log"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var pgdb *gorm.DB
var dbOnce sync.Once

func Get() *gorm.DB {
	dbOnce.Do(func() {
		config := config.Get()
		dbConfig := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
			config.DBHost, config.DBPort, config.DBUser,
			config.DBName, config.DBPassword)

		db, err := gorm.Open(config.DBType, dbConfig)
		if err != nil {
			log.Print("DB open error: ", err)
		}

		pgdb = db
		pgdb.LogMode(true)
	})

	return pgdb
}
