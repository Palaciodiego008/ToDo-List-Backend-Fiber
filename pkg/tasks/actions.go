package tasks

import (
	"fmt"
	"net/http"
	"todolist/pkg/models"

	"github.com/gofiber/fiber/v2"
)

func (h handler) Create(c *fiber.Ctx) error {
	task := &models.Task{}

	if err := c.BodyParser(&task); err != nil {
		return fmt.Errorf("error parsing task: %v", err)
	}

	if err := h.DB.Create(&task).Error; err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return c.Status(http.StatusOK).JSON(task)

}

func (h handler) List(c *fiber.Ctx) error {
	tasks := &models.Tasks{}

	if err := h.DB.Find(&tasks).Error; err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return c.Status(http.StatusOK).JSON(tasks)
}

func (h handler) FindByID(c *fiber.Ctx) error {
	id := c.Params("id")

	task := models.Task{}

	if err := h.DB.First(&task).Where("id = ?", id).Error; err != nil {
		return fiber.NewError(http.StatusNotFound, err.Error())
	}

	return c.Status(http.StatusOK).JSON(task)
}

func (h handler) Update(c *fiber.Ctx) error {

	task := &models.Task{}

	if err := c.BodyParser(&task); err != nil {
		return fmt.Errorf("error parsing task: %v", err)
	}

	if err := h.FindByID(c); err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	if err := h.DB.Save(&task).Error; err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return c.Status(http.StatusOK).JSON(task)
}

func (h handler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	task := &models.Task{}

	if err := h.DB.First(&task).Where("id = ?", id).Error; err != nil {
		return fiber.NewError(http.StatusNotFound, err.Error())
	}

	if err := h.DB.Delete(&task).Error; err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return c.SendStatus(http.StatusOK)
}
