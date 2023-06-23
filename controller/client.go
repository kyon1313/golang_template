package controller

import "github.com/gofiber/fiber/v2"

//CONTROLLER

// VIEWS
func ViewClient(c *fiber.Ctx) error {
	return c.Render("clientlist", fiber.Map{
		"page":  "Client List",
		"title": "CLIENT LIST",
	})
}

func ViewCheckClient(c *fiber.Ctx) error {
	return c.Render("checkclient", fiber.Map{
		"page":  "Check Client",
		"title": "CHECK CLIENT",
	})
}
