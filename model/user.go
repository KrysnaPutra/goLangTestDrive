package model

import (
	"belajargo/database"
	"database/sql"
	"fmt"
	"time"
)

type User struct {
	Id           uint      `gorm:"primaryKey"`
	Name         string    `gorm:"not null"`
	Email        string    `gorm:"not null"`
	Age          uint8     `gorm:"not null"`
	Dob          time.Time `gorm:"not null"`
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time `gorm:"not null"`
	UpdatedAt    time.Time `gorm:"not null"`
}

func CreateUser(name string, email string, age uint8, dob time.Time, memberNumber sql.NullString, activatedAt sql.NullTime) {
	db := database.ConnectDatabase()
	user := User{
		Name:         name,
		Email:        email,
		Age:          age,
		Dob:          dob,
		MemberNumber: memberNumber,
		ActivatedAt:  activatedAt,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	result := db.Create(&user)

	fmt.Println(result.Error)
	fmt.Println(result.RowsAffected)
	fmt.Println("User berhasil dibuat")
}

func SelectUser() map[string]interface{} {
	db := EmDb()
	result := map[string]interface{}{}
	db.Model(&User{}).Select("*").Find(&result)
	return result
}

func UpdateUser(id uint, name string, email string, age uint8, dob time.Time, memberNumber sql.NullString, activatedAt sql.NullTime) {
	db := EmDb()

	user := User{}
	db.Where("id", id).First(&user)

	user.Name = name
	user.Email = email
	user.Age = age
	user.Dob = dob
	user.MemberNumber = memberNumber
	user.ActivatedAt = activatedAt
	user.UpdatedAt = time.Now()

	db.Save(&user)

	fmt.Printf("Sukeses mengupdate user id %v", id)
}

func DeleteUser(id uint) {
	db := EmDb()
	db.Delete(&User{}, id)

	fmt.Printf("Sukeses mengdelete user id %v", id)
}
