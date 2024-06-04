package webhook

import "github.com/gofiber/fiber/v2"

type WebhookDelivery interface {
	Receive(c *fiber.Ctx) error
}
