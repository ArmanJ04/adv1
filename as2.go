package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"html/template"
	"log"
	"net/http"
	"strconv"
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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("register.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
	})

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		name := r.FormValue("name")
		email := r.FormValue("email")

		newUser := User{Name: name, Email: email}
		db.Create(&newUser)

		fmt.Fprintf(w, "User registered successfully:\n%+v", newUser)
	})

	http.HandleFunc("/update", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		userIDStr := r.FormValue("userIdUpdate")
		newName := r.FormValue("newName")
		newEmail := r.FormValue("newEmail")

		userID, err := strconv.ParseUint(userIDStr, 10, 64)
		if err != nil {
			http.Error(w, "Invalid User ID", http.StatusBadRequest)
			return
		}

		var updateUser User
		result := db.First(&updateUser, userID)
		if result.Error != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		if newName != "" {
			updateUser.Name = newName
		}
		if newEmail != "" {
			updateUser.Email = newEmail
		}
		db.Save(&updateUser)

		fmt.Fprintf(w, "User updated successfully:\n%+v", updateUser)
	})

	http.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		userIDStr := r.FormValue("userIdDelete")

		userID, err := strconv.ParseUint(userIDStr, 10, 64)
		if err != nil {
			http.Error(w, "Invalid User ID", http.StatusBadRequest)
			return
		}

		var deleteUser User
		result := db.First(&deleteUser, userID)
		if result.Error != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		db.Delete(&deleteUser, userID)

		fmt.Fprintf(w, "User deleted successfully:\n%+v", deleteUser)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
