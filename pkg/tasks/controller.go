package tasks

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func Routers(app *fiber.App, db *gorm.DB) {
	h := handler{DB: db}

	//using cors
	app.Use(cors.New())

	app.Use(cors.New(cors.Config{
		Next:             nil,
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "",
		AllowCredentials: false,
		ExposeHeaders:    "",
		MaxAge:           0,
	}))

	var taskGroup = app.Group("/tasks")
	taskGroup.Get("/", h.List)
	taskGroup.Post("/", h.Create)
	taskGroup.Get("/:id", h.FindByID)
	taskGroup.Put("/:id", h.Update)
	taskGroup.Delete("/:id", h.Delete)

	h.DB.Debug()
}
