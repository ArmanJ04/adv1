package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Name  string
	Email string
}

func main() {
	dsn := "user=postgres password=jansatov04 dbname=postgres sslmode=disable host=localhost port=3000"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&User{})

	newUser := User{Name: "John Doe", Email: "john@example.com"}
	db.Create(&newUser)
	fmt.Printf("Created user: %+v\n", newUser)

	var retrievedUser User
	db.First(&retrievedUser, newUser.ID)
	fmt.Printf("Retrieved user by ID: %+v\n", retrievedUser)

	db.Model(&retrievedUser).Update("Name", "Arman John")
	fmt.Printf("Updated user: %+v\n", retrievedUser)

	db.Delete(&retrievedUser, retrievedUser.ID)
	fmt.Printf("Deleted user with ID %d\n", retrievedUser.ID)

	var users []User
	db.Find(&users)
	fmt.Println("List of all users:")
	for _, u := range users {
		fmt.Printf("%+v\n", u)
	}
}
