package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	dbURL := os.Getenv("DIRECT_URL")
	if dbURL == "" {
		dbURL = os.Getenv("DATABASE_URL")
	}
	if dbURL == "" {
		log.Fatal("DATABASE_URL or DIRECT_URL must be set")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	fmt.Println("Connecting to Supabase PostgreSQL...")
	config, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		log.Fatalf("ParseConfig failed: %v", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Fatalf("Connect failed: %v", err)
	}
	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("Ping failed: %v", err)
	}
	fmt.Println("Successfully connected to Supabase!")

	files := []string{
		"../supabase/migrations/001_initial_schema.sql",
		"../supabase/migrations/002_seed_initial_data.sql",
	}

	for _, file := range files {
		fmt.Printf("Applying %s...\n", file)
		content, err := os.ReadFile(file)
		if err != nil {
			log.Fatalf("Failed to read file %s: %v", file, err)
		}

		_, err = pool.Exec(ctx, string(content))
		if err != nil {
			log.Fatalf("Failed to execute %s: %v", file, err)
		}
		fmt.Printf("Successfully applied %s!\n", file)
	}

	fmt.Println("ALL MIGRATIONS AND SEED DATA APPLIED SUCCESSFULLY!")
}
