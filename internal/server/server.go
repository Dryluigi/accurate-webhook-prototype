package server

import (
	"accurate-webhook-prototype/internal/webhook"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	app *fiber.App

	webhookDelivery webhook.WebhookDelivery
}

func NewServer() *Server {
	app := fiber.New()

	serv := Server{
		app: app,
	}

	serv.Inject()
	serv.MapRoutes()

	return &serv
}

func (s *Server) Run() error {
	return s.app.Listen(":3000")
}
