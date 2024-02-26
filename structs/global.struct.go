package structs

import "gorm.io/gorm"

type MyConnectionService struct {
	DB *gorm.DB
}
