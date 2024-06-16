package server

import (
	"sync"

	"github.com/gofiber/fiber/v2"
)

var i = 0

func (s *Server) MapRoutes() {
	var mtx sync.Mutex

	// Request counter
	// s.app.Use(func(c *fiber.Ctx) error {
	// 	mtx.Lock()
	// 	i++
	// 	mtx.Unlock()
	// 	return c.Next()
	// })

	s.app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	s.app.Post("/webhook", s.webhookDelivery.Receive)
}
