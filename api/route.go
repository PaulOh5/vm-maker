package api

import (
	"vm-maker/api/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupVmRouter(apiRoute fiber.Router) {
	apiRoute.Post("/vm", handler.PostVm)
}
