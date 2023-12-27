# Golang PostgreSQL CRUD Example

This Golang project demonstrates the basic setup for working with a PostgreSQL database, handling migrations using `golang-migrate`, and performing CRUD operations using the Gorm ORM library.

## Table of Contents

- [Database Preparation](#database-preparation)
  - [Install PostgreSQL](#install-postgresql)
  - [Create Database and User](#create-database-and-user)
  - [Create Users Table](#create-users-table)
- [Connecting to the Database in Golang](#connecting-to-the-database-in-golang)
- [Creating Migrations](#creating-migrations)
  - [Install golang-migrate](#install-golang-migrate)
  - [Create Migrations](#create-migrations)
- [Working with CRUD Operations](#working-with-crud-operations)
  - [Install Gorm](#install-gorm)
  - [Create User Model](#create-user-model)
  - [CRUD Operations](#crud-operations)

## Database Preparation

### Install PostgreSQL

Follow the official instructions to [install PostgreSQL](https://www.postgresql.org/download/) on your system.

### Create Database and User

Create a PostgreSQL database and user for the project. Replace `[Your Database Name]` and `[Your User]` with your preferred names.

```sql
CREATE DATABASE [Your Database Name];
CREATE USER [Your User] WITH PASSWORD '[Your Password]';
ALTER ROLE [Your User] SET client_encoding TO 'utf8';
ALTER ROLE [Your User] SET default_transaction_isolation TO 'read committed';
ALTER ROLE [Your User] SET timezone TO 'UTC';
```
## Connecting to the Database in Golang
Use the provided Golang script to connect to your PostgreSQL database. Ensure you have the required dependencies installed.
package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=[Your User] dbname=[Your Database Name] sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected to the database")
	defer db.Close()
}
## Creating Migrations
Install golang-migrate
Install the golang-migrate tool for managing migrations.

go get -tags 'postgres' -u github.com/golang-migrate/migrate/cmd/migrate

Create Migrations
Create two migrations using the provided commands:

# Create new migration
migrate create -ext sql -dir db/migrations -seq create_users_table

# Create another migration
migrate create -ext sql -dir db/migrations -seq add_age_column
Working with CRUD Operations
Install Gorm
Install the Gorm ORM library for Golang.

go get -u gorm.io/gorm
Create User Model
Define a Gorm model for the users table in the main.go file.

package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Name  string
	Email string
}
CRUD Operations
Implement CRUD operations using Gorm. Update the main.go file with the following:

// ... (Previous code)

func main() {
	dsn := "user=[Your User] dbname=[Your Database Name] sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	// Auto Migrate
	db.AutoMigrate(&User{})

	// Create
	db.Create(&User{Name: "John Doe", Email: "john@example.com"})

	// Read
	var user User
	db.First(&user, 1)
	log.Println(user)

	// Update
	db.Model(&user).Update("Name", "Doe John")

	// Delete
	db.Delete(&user, 1)
}
