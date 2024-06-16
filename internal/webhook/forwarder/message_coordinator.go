package forwarder

import (
	"accurate-webhook-prototype/internal/webhook/entity"
	"accurate-webhook-prototype/internal/webhook/pb"
	"encoding/json"
	"log"
)

// TODO: replace the buffer using postgresql
// TODO: implement graceful quiting maybe using context
// TODO: error handling

type messageCoordinator interface {
	Send()
	Receive()
	PushToBuffer(payload entity.AccuratePayload)
}

type grpcMessageCoordinator struct {
	sendSignal      chan string
	receiveSignal   chan string
	payloadBuffer   []entity.AccuratePayload
	forwardClient   pb.WebhookForwardService_ForwardClient
	stillProcessing bool
}

func (m *grpcMessageCoordinator) PushToBuffer(payload entity.AccuratePayload) {
	prevLen := len(m.payloadBuffer)
	m.payloadBuffer = append(m.payloadBuffer, payload)
	if prevLen == 0 && !m.stillProcessing {
		m.sendSignal <- "SEND"
		log.Println("send signal sent")
	}
}

func (m *grpcMessageCoordinator) Receive() {
	for {
		m.forwardClient.Recv()
		m.receiveSignal <- "RECEIVE"
		log.Println("data received")
	}
}

func (m *grpcMessageCoordinator) Send() {
	for {
		if len(m.payloadBuffer) == 0 {
			m.stillProcessing = false
			<-m.sendSignal
			m.stillProcessing = true
		}
		var data entity.AccuratePayload
		data, m.payloadBuffer = m.payloadBuffer[0], m.payloadBuffer[1:]

		reqByte, _ := json.Marshal(data)
		m.forwardClient.Send(&pb.WebhookForwardData{
			Payload: string(reqByte),
		})
		log.Println("data sent")
		<-m.receiveSignal
	}
}

func newGrpcMessageCoordinator(forwardClient pb.WebhookForwardService_ForwardClient) messageCoordinator {
	return &grpcMessageCoordinator{
		sendSignal:      make(chan string),
		receiveSignal:   make(chan string),
		payloadBuffer:   []entity.AccuratePayload{},
		forwardClient:   forwardClient,
		stillProcessing: false,
	}
}