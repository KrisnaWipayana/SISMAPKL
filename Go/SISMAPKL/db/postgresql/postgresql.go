package postgresql

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Koneksi ke dalam database
func Open() {

	var err error

	const conn = "host=localhost user=postgres password=babu@123 dbname=sismapkl port=5433 sslmode=disable TimeZone=Asia/Jakarta"

	DSN := conn
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if err != nil {
		panic("Database gagal terkoneksi")
	}
	fmt.Println("Terkoneksi ke database..")
}

func GetDB() *gorm.DB {
	if DB == nil {
		log.Fatal("Database gagal terkoneksi")
	}
	return DB
}
