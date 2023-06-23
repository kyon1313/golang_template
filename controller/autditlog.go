package controller

import "github.com/gofiber/fiber/v2"

// CONTROLLER

// VIEWS
func ViewAuditLog(c *fiber.Ctx) error {
	return c.Render("auditlog", fiber.Map{
		"page":  "Audit Trails / Log",
		"title": "AUDIT TRAILS / LOG",
	})
}
