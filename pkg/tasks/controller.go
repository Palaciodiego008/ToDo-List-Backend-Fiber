package tasks

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func Routers(app *fiber.App, db *gorm.DB) {
	h := handler{DB: db}

	var taskGroup = app.Group("/tasks")
	taskGroup.Get("/", h.List)
	taskGroup.Post("/", h.Create)
	taskGroup.Get("/:id", h.FindByID)
	taskGroup.Put("/:id", h.Update)
	taskGroup.Delete("/:id", h.Delete)

	h.DB.Debug()
}
