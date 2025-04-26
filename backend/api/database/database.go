package database

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Ctx = context.Background()
var DB *pgxpool.Pool

func CreateClient() {
	var err error
	dbURL := os.Getenv("DATABSE_URL")
	DB, err = pgxpool.New(Ctx, dbURL)
	if err != nil {
		log.Fatal("unable to connect to the databse", err)
	}
	log.Println("connected to the databse")
}
