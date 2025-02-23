package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var tasks = []Task{}
var nextID = 1

func main() {
	app := fiber.New()

	app.Post("/tasks", createTask)
	app.Get("/tasks", getTasks)
	app.Get("/tasks/:id", getTaskByID)
	app.Put("/tasks/:id", updateTask)
	app.Delete("/tasks/:id", deleteTask)

	app.Listen(":8080")
}

func createTask(c *fiber.Ctx) error {
	newTask := new(Task)
	if err := c.BodyParser(newTask); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	newTask.ID = nextID
	nextID++
	tasks = append(tasks, *newTask)
	return c.Status(fiber.StatusCreated).JSON(newTask)
}

func getTasks(c *fiber.Ctx) error {
	return c.JSON(tasks)
}

func getTaskByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}
	for _, task := range tasks {
		if task.ID == id {
			return c.JSON(task)
		}
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found"})
}

func updateTask(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}
	updatedTask := new(Task)
	if err := c.BodyParser(updatedTask); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Title = updatedTask.Title
			tasks[i].Done = updatedTask.Done
			return c.JSON(tasks[i])
		}
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found"})
}

func deleteTask(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return c.JSON(fiber.Map{"message": "Task deleted"})
		}
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found"})
}
