package main

import (
	"fmt"

	"github.com/ivanovladyslav/waves-backend/m/config"
)

func main() {
	cfg, err := config.LoadConfig()

	if err != nil {
		fmt.Println("Unable to load env")
	}

	fmt.Println(cfg.Amqp.URL)

	fmt.Println("Waves")
}
