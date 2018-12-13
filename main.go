package main

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/jimmysawczuk/try"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	// postgres driver
	_ "github.com/lib/pq"
)

var timeout = 60 * time.Second
var interval = 1 * time.Second

func init() {
	flag.DurationVar(&timeout, "timeout", 60*time.Second, "total amount of time to try")
	flag.DurationVar(&interval, "interval", 1*time.Second, "amount of time to wait between tries")

	flag.Parse()
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("missing required argument: database connection string")
		os.Exit(1)
	}

	connectionString := os.Args[1]
	start := time.Now()

	log.Printf("attempting to connect to postgresql (will try for %s, %s between attempts)", timeout, interval)

	if err := try.Try(connectToPSQL(connectionString), timeout, interval); err != nil {
		log.Fatal(err)
	}

	log.Printf("connected in %s", time.Now().Sub(start).Truncate(time.Millisecond))

}

// connectToPSQL returns a function which attempts to connect to a Postgres server and ping it using the connection string provided.
func connectToPSQL(connectionString string) func() error {
	return func() error {
		db, err := sqlx.Open("postgres", connectionString)
		if err != nil {
			return errors.Wrap(err, "failed to open database")
		}

		defer db.Close()

		err = db.Ping()
		if err != nil {
			return errors.Wrap(err, "couldn't ping")
		}

		return nil
	}
}
