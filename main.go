package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	// Connection parameters
	connStr := "user=postgres password=postgres dbname=vidkrix sslmode=disable"

	// Open PostgreSQL database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
						id SERIAL PRIMARY KEY,
						name TEXT NOT NULL,
						age INTEGER NOT NULL
					)`)
	if err != nil {
		log.Fatal(err)
	}

	// Insert data into the table
	insertStmt, err := db.Prepare("INSERT INTO users(name, age) VALUES($1, $2)")
	if err != nil {
		log.Fatal(err)
	}
	defer insertStmt.Close()

	_, err = insertStmt.Exec("Alice", 30)
	if err != nil {
		log.Fatal(err)
	}

	_, err = insertStmt.Exec("Bob", 35)
	if err != nil {
		log.Fatal(err)
	}

	// Query data from the table
	rows, err := db.Query("SELECT id, name, age FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Users:")
	for rows.Next() {
		var id, age int
		var name string
		err := rows.Scan(&id, &name, &age)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", id, name, age)
	}
}
