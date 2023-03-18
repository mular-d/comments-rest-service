package main

import (
	"context"
	"fmt"

	"github.com/targerian1999/comments/internal/comment"
	"github.com/targerian1999/comments/internal/db"
)

func Run() error {
	fmt.Println("Setting up the api")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to the database")
		return err
	}

	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate database")
		return err
	}

	cmtService := comment.NewService(db)
	fmt.Println(cmtService.GetComment(context.Background(), "9121bf83-28dc-4b8d-bf70-7d347a24ff2e"))

	cmtService.PostComment(
		context.Background(),
		comment.Comment{
			ID:     "9121bf83-28dc-4b8d-bf70-7d347a24ff2e",
			Slug:   "manual-test",
			Author: "mular",
			Body:   "Hello World",
		},
	)

	return nil
}

func main() {
	fmt.Println("Entry point of the application")

	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
