package main

import (
	"context"
	"fmt"
	"log"

	"golang_tutorial/internal/db"

	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "postgres://postgres:postgres@localhost:5432/testdb")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)

	queries := db.New(conn)

	// ✅ Create
	user, err := queries.CreateUser(ctx, db.CreateUserParams{
		Name:  "Manav",
		Email: "manav@gmail.com",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Created:", user)

	// 🔍 Search
	users, err := queries.SearchUsers(ctx, "manav")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Results:")
	for _, u := range users {
		fmt.Println(u.ID, u.Name, u.Email)
	}
}
