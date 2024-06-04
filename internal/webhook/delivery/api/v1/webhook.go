package v1

import (
	"accurate-webhook-prototype/internal/webhook"
	"accurate-webhook-prototype/internal/webhook/entity"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type webhookDelivery struct {
	webhookUseCase webhook.WebhookUseCase
}

func (w *webhookDelivery) Receive(c *fiber.Ctx) error {
	req := []entity.AccuratePayload{}
	err := c.BodyParser(&req)
	if err != nil {
		return err
	}

	err = w.webhookUseCase.ForwardRequest(c.UserContext(), req)
	if err != nil {
		return err
	}

	return c.SendStatus(http.StatusOK)
}

func NewWebhookDelivery(webhookUseCase webhook.WebhookUseCase) webhook.WebhookDelivery {
	return &webhookDelivery{
		webhookUseCase: webhookUseCase,
	}
}
