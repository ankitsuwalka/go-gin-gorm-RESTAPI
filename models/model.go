package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	Id        uint   `gorm:"primaryKey" json:"id"`
	Name      string `json:"name"`
	DOB       string `json:"dob"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func CreateEmp(db *gorm.DB, Emp *Employee) error {

	if err1 := db.Create(Emp).Error; err1 != nil {
		return errors.New("can't create User")
	}
	return nil
}

func GetEmps(db *gorm.DB, Emps *[]Employee) error {
	if err1 := db.Find(Emps).Error; err1 != nil {
		return errors.New("can't get employees")
	}
	return nil
}

func GetEmp(db *gorm.DB, Emps *Employee, id uint) error {
	if err1 := db.Where("id = ?", id).First(Emps).Error; err1 != nil {
		return errors.New("can't get employees")
	}
	return nil
}

func DeleteEmp(db *gorm.DB, Emp *Employee, id uint) error {
	if err1 := db.Where("id = ?", id).Delete(Emp).Error; err1 != nil {
		return errors.New("can't delete employee")
	}
	return nil
}

func UpdateEmp(Db *gorm.DB, Emp *Employee) error {
	if err1 := Db.Save(&Emp).Error; err1 != nil {
		return err1
	}
	return nil
}
