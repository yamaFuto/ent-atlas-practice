package main

import (
	"context"
	"database/sql"
	"log"

	"ent-atlas-test/ent"
	"ent-atlas-test/ent/migrate"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/lib/pq"

	"github.com/bxcodec/faker/v3"
)

// Open new connection
func Open(databaseUrl string) *ent.Client {
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv))
}

func main() {
	dns := "host=localhost port=5432 user=postgres password=postgres dbname=ent-atlas-test sslmode=disable timezone=UTC connect_timeout=5"
	ctx := context.Background()

	// refresh migration
	ent_client := Open(dns)

	defer ent_client.Close()

	// Run migration
	if err := ent_client.Schema.Create(ctx, migrate.WithDropIndex(true)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	log.Println("migration successful")

	// Create a new user
	user, err := ent_client.User.
		Create().
		SetName(faker.Name()).
		SetAge(30).
		SetEmail(faker.Email()).
		SetDescription(faker.Sentence()). // Descriptionフィールドの設定
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating user: %v", err)
	}
	log.Printf("user was created: %v", user)

	// Create a new post for the user
	post, err := ent_client.Post.
		Create().
		SetTitle("First Post").
		SetContent("This is a post written by the user.").
		SetUser(user). // Userに紐づける（user_posts フィールドに）
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating post: %v", err)
	}
	log.Printf("post was created: %v", post)

	// Create a new book
	book, err := ent_client.Book.
		Create().
		SetTitle("Harry Potter").
		SetBody("it's a wizard world!!").
		SetPrice(300).
		SetThoughts("nice").
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating user: %v", err)
	}
	log.Printf("book was created: %v", book)
}
