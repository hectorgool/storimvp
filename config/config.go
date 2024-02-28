package config

import (
	"fmt"
	"os"
	"storimvp/schema"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func init() {

	dbParams := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=%v&loc=%v",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"),
		os.Getenv("MYSQL_CHARSET"),
		os.Getenv("MYSQL_PARSETIME"),
		os.Getenv("MYSQL_LOC"))
	db, err = gorm.Open(mysql.Open(dbParams), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&schema.DBDocument{})
}

// GetDB gets mysql conection
func GetDB() *gorm.DB {
	return db
}
