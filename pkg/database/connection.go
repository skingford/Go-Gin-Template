/*
 * @Author: kingford
 * @Date: 2023-03-23 09:35:31
 * @LastEditTime: 2023-03-23 09:39:15
 */
package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:255"`
}

func connectToDatabase(dbType, dsn string) (*gorm.DB, error) {
	var dialector gorm.Dialector

	switch dbType {
	case "mysql":
		dialector = mysql.Open(dsn)
	case "postgres":
		dialector = postgres.Open(dsn)
	default:
		return nil, fmt.Errorf("unsupported database type: %s", dbType)
	}

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Migrate the schema
	err = db.AutoMigrate(&User{})
	if err != nil {
		return nil, fmt.Errorf("failed to migrate schema: %w", err)
	}

	return db, nil
}

func UsePostgres() (*gorm.DB, error) {
	dbType := "postgres"
	dsn := "user=username password=password host=localhost dbname=mydb port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	return connectToDatabase(dbType, dsn)
}

func UseMySQL() (*gorm.DB, error) {
	dbType := "mysql"
	dsn := "username:password@tcp(localhost:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	return connectToDatabase(dbType, dsn)
}
