package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

// App holds the App configuration
type App struct {
	Host        string
	Port        int
	ReadTimeout time.Duration
}

// LoadApp loads App configuration
func LoadApp() App {

	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Some error occured. Err: %s", err)
	}

	app := App{}

	app.Host = os.Getenv("APP_HOST")
	app.Port, _ = strconv.Atoi(os.Getenv("APP_PORT"))
	timeOut, _ := strconv.Atoi(os.Getenv("APP_READ_TIMEOUT"))
	app.ReadTimeout = time.Duration(timeOut) * time.Second

	return app
}
