package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type PhoneNumber struct {
	gorm.Model
	Value string
}

func main() {
	db, err := gorm.Open(sqlite.Open("phone_numbers.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(fmt.Errorf("failed to connect database: %w", err))
	}

	// Drop table if exists
	if err = db.Migrator().DropTable(&PhoneNumber{}); err != nil {
		panic(fmt.Errorf("failed to drop table: %w", err))
	}

	// Migrate the schema
	if err = db.AutoMigrate(&PhoneNumber{}); err != nil {
		panic(fmt.Errorf("failed to migrate database: %w", err))
	}

	// Insert initial data
	if err = seed(db); err != nil {
		panic(fmt.Errorf("failed to seed database: %w", err))
	}

	// Show initial data
	fmt.Println("Before:")
	showPhoneNumbers(db)

	// Fetch all phone numbers
	var phoneNumbers []PhoneNumber
	result := db.Find(&phoneNumbers)
	if result.Error != nil {
		panic(fmt.Errorf("failed to fetch phone numbers: %w", result.Error))
	}

	// Normalize and save phone numbers
	for i := range phoneNumbers {
		phoneNumbers[i].Value = normalize(phoneNumbers[i].Value)

		// If the normalized phone number already exists, just delete the current duplicate one
		if db.Take(&PhoneNumber{}, "value = ?", phoneNumbers[i].Value).RowsAffected > 0 {
			db.Delete(&phoneNumbers[i])
			continue
		}

		// Update the phone number
		db.Model(&PhoneNumber{}).Where("id = ?", phoneNumbers[i].ID).Updates(phoneNumbers[i])
	}

	fmt.Println("After:")
	showPhoneNumbers(db)
}

func seed(db *gorm.DB) error {
	r, err := os.OpenFile("data.txt", os.O_RDONLY, 0)
	if err != nil {
		return fmt.Errorf("failed to open initial data file: %w", err)
	}

	data, err := io.ReadAll(r)
	if err != nil {
		return fmt.Errorf("failed to read initial data file: %w", err)
	}
	values := strings.Split(string(data), "\n")

	phoneNumbers := make([]PhoneNumber, 0, len(values))
	for _, val := range values {
		phoneNumbers = append(phoneNumbers, PhoneNumber{
			Value: val,
		})
	}
	db.CreateInBatches(&phoneNumbers, len(phoneNumbers))

	return nil
}

func normalize(phoneNumber string) string {
	digits := make([]rune, 0, 10)
	for _, digit := range phoneNumber {
		if digit < '0' || digit > '9' {
			continue
		}
		digits = append(digits, digit)
	}
	return string(digits)
}

func showPhoneNumbers(db *gorm.DB) {
	var phoneNumbers []PhoneNumber
	result := db.Find(&phoneNumbers)
	if result.Error != nil {
		panic(fmt.Errorf("failed to fetch phone numbers: %w", result.Error))
	}

	for _, phoneNumber := range phoneNumbers {
		fmt.Println(phoneNumber.Value)
	}
}
