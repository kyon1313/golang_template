package controller

import "github.com/gofiber/fiber/v2"

// CONTROLLER

// VIEWS

func ViewDelete(c *fiber.Ctx) error {
	return c.Render("deletefile", fiber.Map{
		"page":  "Delete File",
		"title": "DELETE FILE",
	})
}
