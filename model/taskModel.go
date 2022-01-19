package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Task struct {
	GormModel
	Title       string `gorm:"not null" json:"title" form:"title" valid:"required~Title is required"`
	Description string `gorm:"not null" json:"description" form:"description" valid:"required~Description is required"`
	Status      bool   `gorm:"not null" json:"status" form:"status"`
	UserID      uint   `json:"user_id" form:"user_id"`
	CategoryID  uint   `json:"category_id" form:"category_id"`
	User        *User
	Category    *Category
}

type GetTask struct {
	ID          uint       `json:"id"`
	Title       string     `json:"title"`
	Status      bool       `json:"status"`
	Description string     `json:"description"`
	UserID      uint       `json:"user_id"`
	CategoryID  uint       `json:"category_id"`
	CreatedAt   *time.Time `json:"created_at"`
	User        TaskUser   `json:"User"`
}

type CategoryTasks struct {
	ID          uint       `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	UserID      uint       `json:"user_id"`
	CategoryID  uint       `json:"category_id"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

func (t *Task) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(t)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (t *Task) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errUpdate := govalidator.ValidateStruct(t)

	if errUpdate != nil {
		err = errUpdate
		return
	}

	err = nil
	return
}
