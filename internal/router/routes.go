package router

import (
	"captive/internal/container"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App, c *container.Container) {
	cntrl := c.Controllers

	app.Get("/", cntrl.HomeCntrl.Index)

	app.Get("/auth/login", cntrl.AuthCntrl.Login)
	app.Get("/auth/register", cntrl.AuthCntrl.Register)
}
