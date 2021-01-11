package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v4"
)

func main() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	tic := time.Now()
	err = conn.Ping(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "ping failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("we pinged in %v\n", time.Now().Sub(tic))
}