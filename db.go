package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:root@tcp(127.0.0.1:3306)/mydb?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to MariaDB:", err)
	}
	log.Println("Connected to MariaDB successfully")

	// Panggil migrasi tabel User dan Todo
	Migrate()
}

func Migrate() {
	// Migrasikan tabel User dan Todo ke dalam database
	//DB.AutoMigrate(&User{}, &Todo{})
}
