package usecase

import (
	"accurate-webhook-prototype/internal/webhook"
	"accurate-webhook-prototype/internal/webhook/entity"
	"context"
)

type webhookUseCase struct {
	requestForwarder webhook.RequestForwarder
}

func NewWebhookUseCase(requestForwarder webhook.RequestForwarder) webhook.WebhookUseCase {
	return &webhookUseCase{
		requestForwarder: requestForwarder,
	}
}

func (uc *webhookUseCase) ForwardRequest(ctx context.Context, payload []entity.AccuratePayload) error {
	uc.requestForwarder.ForwardRequest(ctx, payload)

	return nil
}
