package controller

import (
	"captive/internal/views/pages/auth"
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
}

func NewAuthController() *AuthController {
	return &AuthController{}
}

func (controller *AuthController) Login(c *fiber.Ctx) error {
	return Render(c, auth.Login())
}

func (controller *AuthController) Register(c *fiber.Ctx) error {
	return Render(c, auth.Register())
}
