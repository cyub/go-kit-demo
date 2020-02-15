package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// Account struct
type Account struct {
	Name   string `gorm:"type:varchar(30);unique" json:"name"`
	Gender int    `gorm:"type:tinyint(1);default:0;" json:"gender"`
	gorm.Model
}

// Create user record
func (m *Account) Create(name string, gender int) error {
	notFound := db.Where("name = ?", name).First(m).RecordNotFound()
	if !notFound {
		return fmt.Errorf("%s has exist, please change anthor", name)
	}

	m.Name = name
	m.Gender = gender
	err := db.Create(m).Error

	fmt.Print(m)
	return err
}

// Find user record
func (m *Account) Find(ID int) error {
	err := db.First(m, ID).Error
	if err != nil {
		return err
	}
	return nil
}
