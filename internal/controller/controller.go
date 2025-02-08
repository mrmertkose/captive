package controller

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"html"
	"strings"
)

func Render(c *fiber.Ctx, component templ.Component) error {
	c.Set("Content-Type", "text/html")
	return component.Render(c.Context(), c.Response().BodyWriter())
}

func HxRedirect(c *fiber.Ctx, to string) error {
	if len(c.Get("HX-Request")) > 0 {
		c.Set("HX-Redirect", to)
		return c.Status(fiber.StatusSeeOther).SendString("")
	}

	return c.Redirect(to, fiber.StatusSeeOther)
}

func clearInput(value string) string {
	return html.EscapeString(strings.TrimSpace(value))
}

func requestInput(c *fiber.Ctx, key string) string {
	return clearInput(strings.TrimSpace(c.FormValue(key)))
}
