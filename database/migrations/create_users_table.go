package migrations

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID      uint      `gorm:"primary key;autoIncrement" json: "id"`
	Name    *string   `json:"name"`
	Email   *string   `json:"email"`
	Date    time.Time `json:"date"`
	City    *string   `json:"city"`
	Country *string   `json:"country"`
}

func MigrateUsers(db *gorm.DB) error {
	err := db.AutoMigrate(&User{})
	return err
}
