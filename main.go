package main

import (
	"log"
	"time"
	"webApp/controller/util"
	"webApp/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html"
)

func main() {

	util.CreateConnection()

	app := fiber.New(fiber.Config{
		Views: html.New("./template/views", ".html"),
	})

	// Dito iloload ung path ng CSS,JS,IMAGES etc...
	app.Static("/", "./template/css")
	app.Static("/", "./template/js")
	app.Static("/", "./template/images")

	// Setup log file
	logFileName := "./logs/insta-" + time.Now().Format("2006-01-02") + ".log"

	file, err := util.OpenLogFile(logFileName)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)
	log.Println("Start of Entry After Re-run")
	// set default log output to the 'new' file
	defer file.Close()

	// Configure application CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	// Declare & initialize logger
	routes.AppRoutes(app)

	log.Fatal(app.Listen(":5020"))
}
