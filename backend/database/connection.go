package database

import (
	"fmt"

	"github.com/oyen-bright/MyTradePad/backend/models"
	"github.com/oyen-bright/MyTradePad/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Init(config *config.DatabaseConfig) (*gorm.DB, error) {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.Host,
		config.User,
		config.Password,
		config.DBName,
		config.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(
		&models.User{},
		&models.TempUser{},
	)

	return db, err
}
