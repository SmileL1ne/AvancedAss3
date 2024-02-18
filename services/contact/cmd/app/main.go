package main

import (
	"architecture_go/pkg/store/postgres"
	"fmt"
)

func main() {
	fmt.Println("Hello World!")

	dbInfo := &postgres.Database{
		Host:     "localhost",
		Port:     "7777",
		User:     "postgres",
		Password: "postgres",
		Name:     "ap_ass3",
	}

	db, err := postgres.OpenDB(dbInfo)
	if err != nil {
		panic("nah" + err.Error())
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INT PRIMARY KEY,
			name VARCHAR(255) UNIQUE NOT NULL 
		)
	`)
	if err != nil {
		panic(err)
	}
}
