package lead

import (
	"github.com/Nishith-Savla/go-fiber-crm-basic/database"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Lead struct {
	gorm.Model
	Name    string `json:"name,omitempty"`
	Company string `json:"company,omitempty"`
	Email   string `json:"email,omitempty"`
	Phone   int    `json:"phone,omitempty"`
}

func GetLeads(c *fiber.Ctx) error {
	var leads []Lead
	database.DBConn.Find(&leads)
	return c.JSON(leads)
}

func GetLead(c *fiber.Ctx) error {
	id := c.Params("id")
	var lead Lead
	database.DBConn.Find(&lead, id)
	return c.JSON(lead)
}

func NewLead(c *fiber.Ctx) error {
	lead := new(Lead)
	if err := c.BodyParser(lead); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	database.DBConn.Create(&lead)
	return c.Status(201).JSON(lead)
}

func DeleteLead(c *fiber.Ctx) error {
	id := c.Params("id")
	var lead Lead
	database.DBConn.First(&lead, id)
	if lead.Name == "" {
		return c.Status(400).SendString("No lead found with given ID")
	}
	database.DBConn.Delete(&lead)
	return c.SendString("Lead successfully deleted")
}
