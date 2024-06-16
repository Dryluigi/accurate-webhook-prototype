package forwarder

import (
	"accurate-webhook-prototype/internal/webhook"
	"accurate-webhook-prototype/internal/webhook/entity"
	"accurate-webhook-prototype/internal/webhook/pb"
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type grpcRequestForwarder struct {
	forwardClient pb.WebhookForwardService_ForwardClient
	mc            messageCoordinator
}

func (g *grpcRequestForwarder) ForwardRequest(ctx context.Context, payload []entity.AccuratePayload) error {
	for i := range payload {
		g.mc.PushToBuffer(payload[i])
	}

	return nil
}

func NewGrpcRequestForwarder(ctx context.Context, targetAUri string) webhook.RequestForwarder {
	conn, err := grpc.NewClient(targetAUri, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	client := pb.NewWebhookForwardServiceClient(conn)

	retryMax := 3
	retryCount := 0
	retrySecond := 1
	var forwardClient pb.WebhookForwardService_ForwardClient
	for forwardClient, err = client.Forward(ctx); err != nil; {
		retryCount++
		if retryCount > retryMax {
			break
		}
		log.Printf("Retrying in %d seconds", retrySecond)
		time.Sleep(time.Duration(retrySecond) * time.Second)
		retrySecond *= 2
	}
	if err != nil {
		log.Panic(err)
	}

	mc := newGrpcMessageCoordinator(forwardClient)

	go mc.Send()
	go mc.Receive()

	return &grpcRequestForwarder{
		forwardClient: forwardClient,
		mc:            mc,
	}
}
