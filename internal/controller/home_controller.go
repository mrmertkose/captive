package controller

import (
	"captive/internal/views/pages/home"
	"github.com/gofiber/fiber/v2"
)

type HomeController struct {
}

func NewHomeController() *HomeController {
	return &HomeController{}
}

func (controller *HomeController) Index(c *fiber.Ctx) error {
	return Render(c, home.Index())
}
