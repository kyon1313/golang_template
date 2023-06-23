package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/xuri/excelize/v2"
)

// CONTROLLER
func ReadExcelFile(c *fiber.Ctx) error {
	f, err := excelize.OpenFile("users.xlsx")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return err
	}

	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
	return nil
}

// VIEWS
func ViewUpload(c *fiber.Ctx) error {
	return c.Render("uploadfile", fiber.Map{
		"page":  "Upload File",
		"title": "UPLOAD FILE",
	})
}
