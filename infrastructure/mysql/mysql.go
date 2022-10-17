package infrastructure

import (
	models "101-hexagonal-golang/infrastructure/mysql/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	//"gorm.io/gorm/clause"
)

func ConnMysql(config map[string]string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config["USER"], config["PASS"], config["HOST"], config["PORT"], config["DBNAME"])
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}

func Migration(db *gorm.DB) {
	db.Table("members").AutoMigrate(
		&models.Members{},
	)
}
