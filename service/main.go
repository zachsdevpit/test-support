package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django/v3"
)

func main() {
	app, err := initApp()
	if err != nil {
		log.Fatal(err)
	}
	app.Get("/", handleGetHome)
	log.Fatal(app.Listen(":8080"))
}

func initApp() (*fiber.App, error) {
	//	if err := godotenv.Load(); err != nil {
	//		return nil, err
	//	}
	engine := django.New("./views", ".html")
	engine.Reload(true)
	app := fiber.New(fiber.Config{
		DisableStartupMessage: false,
		PassLocalsToViews:     true,
		Views:                 engine,
	})
	return app, nil
}

func handleGetHome(c *fiber.Ctx) error {
	return c.Render("home/index", fiber.Map{})
}
