package model

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"my_kanban_board/helper"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type roleType string

const (
	ADMIN  roleType = "admin"
	MEMBER roleType = "member"
)

//Validate enum when set to database
func (r roleType) Value() (driver.Value, error) {
	switch r {
	case ADMIN, MEMBER:
		return string(r), nil
	}
	return nil, errors.New("Invalid role type value")
}

//validate enum where read from database
func (r *roleType) Scan(value interface{}) error {
	var rt roleType
	if value == nil {
		*r = ""
		return nil
	}

	st, ok := value.(string)
	if !ok {
		return errors.New("Invalid data from role type")
	}
	rt = roleType(st)
	switch rt {
	case ADMIN, MEMBER:
		*r = rt
		return nil
	}
	return fmt.Errorf("Invalid role type value : %s", rt)
}

type User struct {
	GormModel
	FullName string   `gorm:"not null" json:"full_name" form:"full_name" valid:"required~Your Full Name is required"`
	Email    string   `gorm:"not null; uniqueIndex" json:"email" form:"email" valid:"required~Your Email is required, email~Invalid Email format"`
	Password string   `gorm:"not null" json:"password" form:"password" valid:"required~Your Password is required, minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Role     roleType `gorm:"not null;type:roleType" json:"role" form:"role" valid:"required~Your Role is required"`
	Tasks    []Task   `gorm:"constraint:OnUpdate:CASCADE, OnDelete:CASCADE;" json:"tasks"`
}

type TaskUser struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}
	fmt.Println("Create password model.User", u.Password)
	u.Password = helper.HashPass(u.Password)
	err = nil
	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errUpdate := govalidator.ValidateStruct(u)

	if errUpdate != nil {
		err = errUpdate
		return
	}

	err = nil
	return
}
