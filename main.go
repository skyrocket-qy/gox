package main

import (
	"errors"
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID   uint
	Name string
}

type Order struct {
	ID     uint
	Amount float64
}

func main() {
	// Setup the first database connection (dbA)
	dbA, err := gorm.Open(sqlite.Open("dbA.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to dbA: %v", err)
	}

	// Setup the second database connection (dbB)
	dbB, err := gorm.Open(sqlite.Open("dbB.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to dbB: %v", err)
	}

	// Auto-migrate the schema
	if err := dbA.AutoMigrate(&User{}); err != nil {
		log.Fatalf("failed to migrate schema in dbA: %v", err)
	}
	if err := dbB.AutoMigrate(&Order{}); err != nil {
		log.Fatalf("failed to migrate schema in dbB: %v", err)
	}

	if err := dbA.Transaction(func(txA *gorm.DB) error {
		if err := txA.Create(&User{Name: "John Doe"}).Error; err != nil {
			return err
		}

		if err := dbB.Transaction(func(txB *gorm.DB) error {
			if err := txB.Create(&Order{Amount: 100.0}).Error; err != nil {
				return err
			}
			return nil
		}); err != nil {
			return err
		}

		// suppose this is failed
		return errors.New("mock failed ")
	}); err != nil {
		log.Fatalf("failed to do transaction in dbA: %v", err)
	}

	fmt.Println("both transactions committed successfully")
}
