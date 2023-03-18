package main

import (
	"context"
	"fmt"

	"github.com/targerian1999/comments/internal/db"
)

func Run() error {
	fmt.Println("Setting up the api")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to the database")
		return err
	}

	if err := db.Ping(context.Background()); err != nil {
		return err
	}

	fmt.Println("Successfully connected and pinged database")

	return nil
}

func main() {
	fmt.Println("Entry point of the application")

	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
