package main

import (
	"belajargo/model"
	"database/sql"
	"fmt"
	"time"
)

func main() {
	model.MigrateAll()

	model.CreateProduct(
		"Nasi Padang",
		900,
	)
	model.CreateProduct(
		"Be Guling",
		100,
	)
	fmt.Println(model.SelectProduct())

	model.UpdateProduct(1, "Nasi Campur", 400)
	model.DeleteProduct(2)

	fmt.Println(model.SelectProduct())

	//

	model.CreateUser(
		"Made Sigma",
		"madebigmasigma@example.com",
		22,
		time.Now(),
		sql.NullString{String: "T-992288", Valid: true},
		sql.NullTime{Time: time.Now(), Valid: true},
	)

	model.CreateUser(
		"Ketut Baller",
		"ketutballer@example.com",
		32,
		time.Now(),
		sql.NullString{Valid: false},
		sql.NullTime{Valid: false},
	)

	fmt.Println(model.SelectUser())

	model.UpdateUser(
		1,
		"Made Sigma Ligma",
		"madesigmaligma@example.com",
		23,
		time.Now(),
		sql.NullString{Valid: false},
		sql.NullTime{Valid: false},
	)

	model.DeleteUser(2)

	fmt.Println(model.SelectUser())
}
