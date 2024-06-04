package webhook

import (
	"accurate-webhook-prototype/internal/webhook/entity"
	"context"
)

type WebhookUseCase interface {
	ForwardRequest(ctx context.Context, payload []entity.AccuratePayload) error
}
