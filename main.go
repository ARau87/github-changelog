package main

import (
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
)

var App *cli.App

func main() {

	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	envFile := path.Join(dir, ".env")
	if err := godotenv.Load(envFile); err != nil {
		log.Fatal("Couldn't load .env file")
	}

	app := SetupApp()

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
