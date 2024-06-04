package forwarder

import (
	"accurate-webhook-prototype/internal/webhook"
	"accurate-webhook-prototype/internal/webhook/entity"
	"accurate-webhook-prototype/internal/webhook/pb"
	"context"
	"encoding/json"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type grpcRequestForwarder struct {
	forwardClient pb.WebhookForwardService_ForwardClient
}

func (g *grpcRequestForwarder) ForwardRequest(ctx context.Context, payload []entity.AccuratePayload) error {
	log.Println("forwarded request")

	bytePayload, err := json.Marshal(payload)
	if err != nil {
		log.Println("error marshaling")
		return err
	}

	err = g.forwardClient.Send(&pb.WebhookForwardData{
		Payload: string(bytePayload),
	})
	if err != nil {
		log.Println("error forwarding request")
		return err
	}

	return nil
}

func NewGrpcRequestForwarder(ctx context.Context, targetAUri string) webhook.RequestForwarder {
	conn, err := grpc.NewClient(targetAUri, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	client := pb.NewWebhookForwardServiceClient(conn)

	forwardClient, err := client.Forward(ctx)
	if err != nil {
		panic(err)
	}

	return &grpcRequestForwarder{
		forwardClient: forwardClient,
	}
}
