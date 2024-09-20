package model

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error

	dsn := "root:@tcp(127.0.0.1:3306)/go_resapi_mux"

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("koneksi database gagal: %v", err)
	}

	if err := DB.AutoMigrate(&Product{}); err != nil {
		log.Fatalf("migrate database gagal: %v", err)
	}
}
