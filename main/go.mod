module belajargo/main

replace belajargo/database => ../database

replace belajargo/model => ../model

go 1.23.1

require belajargo/model v0.0.0-00010101000000-000000000000

require (
	belajargo/database v0.0.0-00010101000000-000000000000 // indirect
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/go-sql-driver/mysql v1.8.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/text v0.18.0 // indirect
	gorm.io/driver/mysql v1.5.7 // indirect
	gorm.io/gorm v1.25.12 // indirect
)
