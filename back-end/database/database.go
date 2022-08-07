package database

import (
	"fmt"
	"time"

	"database/sql"

	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

const (
	host     = "localhost"
	port     = 5432
	user     = "xsky"
	password = "xsky"
	dbname   = "X-Sky"
)

func Connect() error {

	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	DB, err = sql.Open("postgres", psqlconn)

	if err != nil {
		fmt.Println(err)
	}

	go func() {
		time.Sleep(5 * time.Second)
		for {

			err = DB.Ping()
			if err != nil {
				fmt.Println("\nDatabase connection lost!")
				for {

					// open database
					DB, err = sql.Open("postgres", psqlconn)

					err = DB.Ping()
					if err != nil {
						fmt.Println("\nReconnection failed!")
						time.Sleep(5 * time.Second)
					} else {
						fmt.Println("\nReconnected!")
						time.Sleep(5 * time.Second)
						break
					}
				}
			} else {
				time.Sleep(5 * time.Second)
			}
		}
	}()

	return nil
}
