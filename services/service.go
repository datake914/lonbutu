package services

import "github.com/jinzhu/gorm"

type Service interface {
	Execute(tx *gorm.DB)
}
