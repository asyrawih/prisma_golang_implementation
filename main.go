package main

import (
	"fmt"
	"log"

	"github.com/hananloser/prismago/prisma/db"
	"github.com/hananloser/prismago/router"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	fmt.Println("starting")

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	echo := echo.New()
	echo.Use(middleware.Logger())

	client := db.NewClient()

	if err := client.Prisma.Connect(); err != nil {
		log.Fatal("Error", err)
	}

	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	router.TodoRouter(echo, client)

	log.Println(echo.Router())

	echo.Logger.Fatal(echo.Start(":8000"))

}
