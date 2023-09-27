package main

import (
	_ "fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Open a log file
	file, err := os.OpenFile("server.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Set log output to the file
	log.SetOutput(file)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	app := fiber.New(fiber.Config{})

	app.Static("/", "./public")

	app.Get("/api", func(c *fiber.Ctx) error {
		log.Println("GET /api from", c.IP())
		return c.SendString("I'm a GET request!")
	})

	app.Route("/test", func(api fiber.Router) {
		api.Get("/foo", handler1).Name("foo")
		api.Get("/bar", handler2).Name("ber")
	}, "test.")

	log.Fatal(app.Listen(":80"))
}

func handler1(c *fiber.Ctx) error {
	log.Println("GET /test/foo from", c.IP())
	return c.SendString("I'm a GET request /test/foo!")
}

func handler2(c *fiber.Ctx) error {
	log.Println("GET /test/bar from", c.IP())
	return c.SendString("I'm a GET request /test/bar!")
}
