package services

import (
	"fmt"

	"github.com/datake914/lonbutu/models"
	"github.com/jinzhu/gorm"
)

type WordGetService struct {
	baseService
}

func (s *WordGetService) Execute(tx *gorm.DB) {
	tx.AutoMigrate(&models.WordModel{})
	users := []models.WordModel{}
	tx.Find(&users)
	fmt.Println(users)
}
