package mysql

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Database() {
	var err error
	// dsn := "root@tcp(127.0.0.1:3306)/waysfood?charset=utf8mb4&parseTime=True&loc=Local"
	// // dsn := "{user}:{password}@tcp({Host}:{Port})/{Database name}?charset=utf8mb4&parseTime=True&loc=Local"
	// DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	var DB_PASSWORD = os.Getenv("DB_PASSWORD")
	var DB_HOST = os.Getenv("DB_HOST")
	var DB_PORT = os.Getenv("DB_PORT")
	var DB_NAME = os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	
	if err != nil {
		panic(err)
	}

	fmt.Println("Database Has Connected")

}
