package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv(env string) {
	if env == "dev" {
		err := godotenv.Load("./configs/.dev.env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	} else {
		err := godotenv.Load("./configs/.test.env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println(os.Getenv("MONGOURI"))
	return
}
