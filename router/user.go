package router

import (
	"github.com/agustfricke/fiber-jwt-auth/handlers"
	"github.com/agustfricke/fiber-jwt-auth/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
    app.Post("/signup", handlers.SignUp)
    app.Post("/signin", handlers.SignIn)
    app.Get("/logout", handlers.Logout)
    app.Get("/me",  middleware.DeserializeUser ,handlers.GetMe)
}
