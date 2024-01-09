package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
	Body  string `json:"body"`
}

type UpdateDone struct {
	Done bool `json:"done"`
}

func main() {
	fmt.Print("Hello word")

	app := fiber.New()

	todos := []Todo{}

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	app.Post("/api/todos/", func(c *fiber.Ctx) error {
		todo := &Todo{}

		if err := c.BodyParser(todo); err != nil {
			return err
		}

		todo.ID = len(todos) + 1

		todos = append(todos, *todo)

		return c.JSON(todos)
	})

	app.Get("/api/todos/", func(c *fiber.Ctx) error {
		return c.JSON(todos)
	})

	app.Patch("/api/todos/:id/done", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		todo := &UpdateDone{}

		// Log the request body
		fmt.Println("Request Body:", string(c.Request().Body()))

		if err := c.BodyParser(todo); err != nil {
			log.Println("Error parsing request body:", err)
			return err
		}

		if err != nil {
			return c.Status(401).SendString("invalid id")
		}

		for i, t := range todos {
			if t.ID == id {
				todos[i].Done = todo.Done
				break
			}
		}

		return c.JSON(todos)
	})

	log.Fatal(app.Listen(":4000"))
}
