package models

import (
	"github.com/andikanugraha11/rest-api-jwt/helpers"
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	FullName string    `gorm:"not null" json:"full_name" form:"full_name" valid:"required~full_name is required"`
	Email    string    `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required,email~email invalid"`
	Password string    `gorm:"not null" json:"password" form:"password" valid:"required,minstringlength(8)~Password minimal 8 character"`
	Products []Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"products"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(u)
	if err != nil {
		return err
	}
	u.Password = helpers.HashPassword(u.Password)
	return nil
}