package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Category struct {
	GormModel
	Type  string `gorm:"not null" json:"type" form:"type" valid:"required~Type is required"`
	Tasks []Task `gorm:"constraint:OnUpdate:CASCADE, OnDelete:CASCADE;" json:"tasks"`
}

type GetCategory struct {
	ID        uint            `json:"id"`
	Type      string          `json:"type"`
	UpdatedAt *time.Time      `json:"updated_at"`
	CreatedAt *time.Time      `json:"created_at"`
	Tasks     []CategoryTasks `json:"Tasks"`
}

func (c *Category) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(c)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (c *Category) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errUpdate := govalidator.ValidateStruct(c)

	if errUpdate != nil {
		err = errUpdate
		return
	}

	err = nil
	return
}
