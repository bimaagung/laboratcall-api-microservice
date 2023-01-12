package entities

import (
	"time"

	"gorm.io/gorm"
)

type Checkup struct {
	Id        string `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	Unit  	  string `gorm:"not null"`
	Price  	  int `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
