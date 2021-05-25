package configs

import (
	"fmt"
	"log"

	"github.com/nathakrit061103jnt/crud_fiber_go_gorm/database"
	"github.com/nathakrit061103jnt/crud_fiber_go_gorm/models"

	// _ "gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDatabase() {

	var err error
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	// dsn := "root:@tcp(127.0.0.1:3306)/crud_fiber_go_gorm?charset=utf8mb4&parseTime=True&loc=Local"
	// database.DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	dsn := "host=localhost user=postgres password=12345678 dbname=crud_fiber_go_gorm port=5432 sslmode=disable TimeZone=Asia/Kolkata"

	log.Print("Connecting to PostgreSQL DB...")
	database.DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")

	defer database.DBConn.AutoMigrate(&models.Product{})

}
