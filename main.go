package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/svartvit93/crud_go/bootstrap"
	"github.com/svartvit93/crud_go/repository"
)

type Repository repository.Repository

func main() {
	app := fiber.New()
	boostrap.InitializeApp(app)
}