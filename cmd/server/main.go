package main

import (
	"context"
	"fmt"

	"github.com/souviks72/go-rest-api/internal/comment"
	"github.com/souviks72/go-rest-api/internal/db"
)

// Run - is going to be responsible for
// the instatiation and startup of our
// go application
func Run() error {
	fmt.Println("starting up our application")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to the database")
		return err
	}
	if err := db.Ping(context.Background()); err != nil {
		return err
	}
	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate database")
		return err
	}

	cmtService := comment.NewService(db)
	fmt.Println(cmtService.PostComment(context.Background(), comment.Comment{
		ID:     "some id",
		Slug:   "manual-test",
		Author: "Elliot",
		Body:   "Hello World",
	}))

	return nil
}

func main() {
	fmt.Println("Go REST API Course")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
