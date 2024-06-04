package server

import (
	v1 "accurate-webhook-prototype/internal/webhook/delivery/api/v1"
	"accurate-webhook-prototype/internal/webhook/forwarder"
	"accurate-webhook-prototype/internal/webhook/usecase"
	"context"
)

func (s *Server) Inject() {
	requestForwarder := forwarder.NewGrpcRequestForwarder(context.Background(), "localhost:9000")
	webhookUseCase := usecase.NewWebhookUseCase(requestForwarder)
	s.webhookDelivery = v1.NewWebhookDelivery(webhookUseCase)
}
