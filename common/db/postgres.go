package postgres


import (
	"fmt"
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPostgresDB() (*gorm.DB, error) {
	//get database connection string from env
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
	)

	//connect to DB
	gormDB, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("cannot open database: %v", err)
	}
	//get DB from gorm.DB
	sqlDB, err := gormDB.DB()
	if err != nil {
		return nil, err
	}

	//Check is connectio available
	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	return gormDB, nil
}
