package main

import (
	"captive/internal/container"
	"captive/internal/router"
	"captive/pkg/config"
	"embed"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"net/http"
)

//go:embed all:internal/assets
var embedDirStatic embed.FS

func main() {
	app := fiber.New()
	config.Set()
	app.Use(recover.New(), helmet.New(), cors.New())

	// db settings

	// ---> session
	//csrfConfig := csrf.Config{
	//	ContextKey:     "csrf",
	//	KeyLookup:      "header:X-CSRF-Token",
	//	Expiration:     time.Duration(config.Get().Server.SessionExpire) * time.Hour,
	//	Session:        sessions.SessionCon(),
	//	KeyGenerator:   utils.UUIDv4,
	//	CookieHTTPOnly: true,
	//	CookieSameSite: fiber.CookieSameSiteLaxMode,
	//}
	//if config.Get().App.Ssl {
	//	csrfConfig.CookieName = "__Host-csrf"
	//	csrfConfig.CookieSecure = true
	//} else {
	//	csrfConfig.CookieName = "csrf_"
	//	csrfConfig.CookieSecure = false
	//}
	//app.Use(csrf.New(csrfConfig))

	app.Use("/assets", filesystem.New(filesystem.Config{
		Root:       http.FS(embedDirStatic),
		PathPrefix: "assets",
		Browse:     true,
		MaxAge:     2592000,
	}))

	newContainer := container.NewContainer()
	router.Setup(app, newContainer)

	// ---> not found
	//app.Use(func(c *fiber.Ctx) error {
	//	c.Status(http.StatusNotFound)
	//	return controller.Render(c, pages.Page404())
	//})

	log.Fatal(app.Listen(fmt.Sprintf("%s:%s", config.Get().Server.Host, config.Get().Server.Port)))
}
