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
        SetName("John Doe").
        SetAge(30).
        SetEmail("test@example.com").
        SetDescription("A short description about John Doe."). // Descriptionフィールドの設定
        Save(context.Background())
    if err != nil {
        log.Fatalf("failed creating user: %v", err)
    }
    log.Printf("user was created: %v", user)
}