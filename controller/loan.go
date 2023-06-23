package controller

import "github.com/gofiber/fiber/v2"

// CONTROLLER

// VIEWS
func ViewReport(c *fiber.Ctx) error {
	return c.Render("reportloan", fiber.Map{
		"page":  "Report Loan",
		"title": "REPORT LOAN",
	})
}
