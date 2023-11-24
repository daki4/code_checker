package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ConnectionOptions struct {
	Host     string
	User     string
	Database string
	Password string
	Port     int
}

var db *gorm.DB
var dbErr error

func SpecificConnect(options ConnectionOptions) error {
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Europe/Amsterdam",
		options.Host,
		options.User,
		options.Password,
		options.Database,
		options.Port,
	)
	db, dbErr = gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	return dbErr
}

func Connect() error {
	if host, err := os.LookupEnv("DB_HOST"); !err {
		return fmt.Errorf("DB_HOST not found")
	} else if user, err := os.LookupEnv("DB_USER"); !err {
		return fmt.Errorf("DB_USER not found")
	} else if password, err := os.LookupEnv("DB_PASSWORD"); !err {
		return fmt.Errorf("DB_PASSWORD not found")
	} else if database, err := os.LookupEnv("DB_NAME"); !err {
		return fmt.Errorf("DB_NAME not found")
	} else if port, err := os.LookupEnv("DB_PORT"); !err {
		return fmt.Errorf("DB_PORT not found")
	} else {
		connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Amsterdam",
			host,
			user,
			password,
			database,
			port,
		)
		db, dbErr = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
		return dbErr
	}
}

func Migrate(db *gorm.DB, migratable interface{}) error {
	return db.AutoMigrate(migratable)
}

func GetDb() *gorm.DB {
	return db
}
