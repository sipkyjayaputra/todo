package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
	"todo/model/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect() (*gorm.DB, *sql.DB, error) {
	// load configuration
	connectionString := fmt.Sprintf(`%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local`,
		CONFIG["DBUSER"],
		CONFIG["DBPWD"],
		CONFIG["DBHOST"],
		CONFIG["DBPORT"],
		CONFIG["DBNAME"])
	log.Println(connectionString)

	// create logger
	logger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // define log writer
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: false,
			Colorful:                  false,
		},
	)

	// make a connection
	connection, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{Logger: logger})
	if err != nil {
		return nil, nil, err
	}

	sqlDB, err := connection.DB()
	if err != nil {
		return nil, nil, err
	} else {
		sqlDB.SetMaxIdleConns(2)
		sqlDB.SetMaxOpenConns(1000)
	}

	// migrate if we have a table schema here
	connection.AutoMigrate(&entity.Task{})

	log.Println("MySQL connection success!")
	return connection, sqlDB, nil
}
