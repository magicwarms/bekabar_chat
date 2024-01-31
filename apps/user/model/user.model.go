package model

import (
	"bekabar_chat/apps/utils"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// UserModel Constructs your UserModel under entities
type UserModel struct {
	ID        string `gorm:"primaryKey" json:"id"`
	Username  string `gorm:"not null;uniqueIndex;size:70" json:"username"`
	Email     string `gorm:"not null;uniqueIndex;size:50" json:"email"`
	Password  string `gorm:"not null" json:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// Set table name (GORM)
func (UserModel) TableName() string {
	return "users"
}

// DEFINE HOOKS
func (user *UserModel) BeforeCreate(tx *gorm.DB) (err error) {
	hashedPassword, err := utils.HashPassword(user.Password)
	// if hash password error, return the error
	if err != nil {
		return err
	}
	// set password before create
	user.Password = string(hashedPassword)
	generateUuid, errGenerateUuid := uuid.NewV7()
	if errGenerateUuid != nil {
		fmt.Println("Error generate uuid", errGenerateUuid)
		return errGenerateUuid
	}
	user.ID = generateUuid.String()
	return
}

func (user *UserModel) BeforeUpdate(tx *gorm.DB) (err error) {
	if tx.Statement.Changed() {
		tx.Statement.SetColumn("UpdatedAt", time.Now())
	}
	return
}
