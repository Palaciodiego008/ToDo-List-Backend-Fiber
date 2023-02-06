package main

import (
	"log"
	"todolist/pkg/common/config"
	"todolist/pkg/db"
	"todolist/pkg/tasks"

	"github.com/gofiber/fiber/v2"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln(err.Error())
	}

	app := fiber.New()

	db := db.Init(c.DBUrl)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	tasks.Routers(app, db)

	app.Listen(c.Port)
}
