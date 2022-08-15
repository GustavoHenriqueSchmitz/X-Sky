package database

import (
	"fmt"
	"time"

	"database/sql"

	_ "github.com/lib/pq"
)

// Database
var (
	DB  *sql.DB
	err error
)

// Database configs
const (
	host     = "localhost"
	port     = 5432
	user     = "xsky"
	password = "xsky"
	dbname   = "X-Sky"
)

// Function Connect to open a connection with the database.
func Connect() error {

	// Connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// Open database connection
	DB, err = sql.Open("postgres", psqlconn)

	if err != nil {
		fmt.Println(err)
	}

	// Create goroutine function, for reconnection.
	go func() {
		time.Sleep(5 * time.Second)
		for {
			// Use ping for verify if connection is active, if the connection is inactive try to reconnect.
			err = DB.Ping()
			if err != nil {
				fmt.Println("\nDatabase connection lost!")

				for {
					// Try to reconnect.
					DB, err = sql.Open("postgres", psqlconn)

					// Use ping for verify if the reconnection is succeed.
					err = DB.Ping()
					if err != nil {
						fmt.Println("\nReconnection failed!")
						time.Sleep(10 * time.Second)
					} else {
						fmt.Println("\nReconnected!")
						time.Sleep(10 * time.Second)
						break
					}
				}
			} else {
				time.Sleep(60 * time.Second)
			}
		}
	}()

	return nil
}
