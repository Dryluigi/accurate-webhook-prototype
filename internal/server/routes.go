package server

import "github.com/gofiber/fiber/v2"

func (s *Server) MapRoutes() {
	s.app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	s.app.Post("/webhook", s.webhookDelivery.Receive)
}
