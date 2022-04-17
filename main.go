package main

import (
	"log"

	"github.com/Nishith-Savla/go-fiber-crm-basic/database"
	"github.com/Nishith-Savla/go-fiber-crm-basic/lead"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func initDatabase() {
	var err error
	if database.DBConn, err = gorm.Open(sqlite.Open("leads.db"), &gorm.Config{}); err != nil {
		log.Fatalln(err)
	}
	database.DBConn.AutoMigrate(&lead.Lead{})
}

func main() {
	app := fiber.New()
	initDatabase()
	defer func() {
		db, err := database.DBConn.DB()
		if err != nil {
			log.Fatalln(err)
		}
		db.Close()
	}()
	setupRoutes(app)
	app.Listen(":3000")
}
