package main

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/template/html/v2"
)

type Message struct{
	MessageContent string `json:"message_content" form:"message_content"`
}

func main(){
	engine := html.New("./templates", ".html")
	app := fiber.New(fiber.Config{
        Views: engine,
    })

	app.Static("/static/", "./static")

	app.Get("/", func(c fiber.Ctx) error {
		return c.Status(http.StatusOK).Render("chat", nil)
	})

	app.Post("/", func(c fiber.Ctx) error {
		// c.Request().

		inputValue := c.FormValue("chat-input")

		fmt.Println("input:", inputValue)

		return nil
	})
	
	fmt.Println("running on port 8080")
	app.Listen(":8080")
}