package tests

import (
	"context"

	"github.com/bimaagung/laboratcall-microservices/checkup/src/Domains/entities"
	database "github.com/bimaagung/laboratcall-microservices/checkup/src/Infrastructures/database/postgres"
	"gorm.io/gorm"
)

type CheckupTableTestHelper struct {
	DB *gorm.DB
}

func  NewCheckupTableTestHelper() *CheckupTableTestHelper {
	db := database.NewPostgresDB().WithContext(context.TODO())
	return &CheckupTableTestHelper{
		DB: db,
	}
}

func (c *CheckupTableTestHelper) Add(payload *entities.Checkup)(string, error){
	err := c.DB.Create(&payload).Error
	if err != nil {
		return "", err
	}
	return payload.Id, nil
}

func (c *CheckupTableTestHelper) GetById(id string)(payload *entities.Checkup, err error){
	err = c.DB.First(&payload, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return payload, nil
}

func (c *CheckupTableTestHelper) Clear() error {
	query := "TRUNCATE TABLE checkups"
	err := c.DB.Raw(query).Error
	if err != nil {
		return err
	}

	return nil
} 


