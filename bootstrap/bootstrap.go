package bootstrap

import (
	"log"
	"vm-maker/api"
	"vm-maker/config"
	"vm-maker/utils"

	"github.com/gofiber/fiber/v2"
)

func NewApplication() *fiber.App {
	if check, err := utils.CheckGroup("kvm"); err != nil {
		log.Fatalf("Error checking group kvm: %v", err)
	} else if !check {
		log.Fatal("User is not in group kvm")
	}

	settings := config.SetupSettings()

	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("settings", settings)
		return c.Next()
	})

	apiRoutes := app.Group("/api")
	api.SetupVmRouter(apiRoutes)

	return app
}
