package main

import (
	"fmt"
	"os"
	"scaffold/config"
	"scaffold/shared/log"
	"scaffold/src"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	ctx := "main"

	defer func() {
		if r := recover(); r != nil {
			log.Log(log.PanicLevel, fmt.Sprint(r), ctx, "main_process")
			time.Sleep(5 * time.Second)
		}
	}()

	// load config
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(".env is not loaded properly")
		os.Exit(2)
	}

	// run the app
	src.Application(config.NewConfig()).Run()
}
