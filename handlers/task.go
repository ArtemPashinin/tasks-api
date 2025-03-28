package handlers

import (
	"tasks-api/db"
	"tasks-api/models"

	"github.com/gofiber/fiber/v2"
)

func RegisterTaskHandlers(app *fiber.App, repo db.Repository) {
	app.Get("/tasks", getTasks(repo))
	app.Post("/tasks", createTask(repo))
	app.Put("/tasks/:id", updateTask(repo))
	app.Delete("/tasks/:id", deleteTask(repo))
}

func getTasks(repo db.Repository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tasks, err := repo.FindAll()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(tasks)
	}
}

func createTask(repo db.Repository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var task models.Task
		if err := c.BodyParser(&task); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}

		if task.Status == "" {
			task.Status = "new"
		}

		id, err := repo.CreateOne(task)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"id": id,
		})
	}
}

func updateTask(repo db.Repository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid ID",
			})
		}

		var task models.Task
		if err := c.BodyParser(&task); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}

		if err := repo.UpdateOne(id, task); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.SendStatus(fiber.StatusOK)
	}
}

func deleteTask(repo db.Repository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid ID",
			})
		}

		if err := repo.DeleteOne(id); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.SendStatus(fiber.StatusOK)
	}
}
