package model

import (
	"fmt"
	"time"
)

type Product struct {
	Id        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	Quatitiy  uint32 `gorm:"not null;default:100"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func CreateProduct(name string, quantitiy uint32) {
	db := EmDb()
	product := Product{
		Name:      name,
		Quatitiy:  quantitiy,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	result := db.Create(&product)

	fmt.Println(result.Error)
	fmt.Println(result.RowsAffected)
	fmt.Println("Sukses membuat produk")
}

func SelectProduct() map[string]interface{} {
	db := EmDb()
	result := map[string]interface{}{}
	db.Model(&Product{}).Select("*").Find(&result)
	return result
}

func UpdateProduct(id uint, name string, quatitiy uint32) {
	db := EmDb()

	product := Product{}
	db.Where("id", id).First(&product)

	product.Name = name
	product.Quatitiy = quatitiy
	product.UpdatedAt = time.Now()

	db.Save(&product)

	fmt.Printf("Sukeses Mengupdate Produk id %v", id)
}

func DeleteProduct(id uint) {
	db := EmDb()
	db.Delete(&Product{}, id)

	fmt.Printf("Sukeses Mengdelete Produk id %v", id)
}
