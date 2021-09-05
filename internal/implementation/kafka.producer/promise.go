package kafkaproducer

import (
	"context"
	"time"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	jsoniter "github.com/json-iterator/go"
	"github.com/segmentio/kafka-go"

	"github.com/ozonva/ova-promise-api/internal/domain"
)

func NewEvent() *cloudevents.Event {
	v := cloudevents.NewEvent()
	return &v
}

func (w w) Send(ctx context.Context, event *cloudevents.Event) error {
	bytes, err := jsoniter.Marshal(event)
	if err != nil {
		return err
	}

	if err := w.writer.WriteMessages(ctx, kafka.Message{Value: bytes, Key: []byte(event.Type())}); err != nil {
		return err
	}

	return nil
}

func (w w) NewEventPromiseCreated(ctx context.Context, promise *domain.Promise) error {
	ev := NewEvent()
	ev.SetID(domain.GenerateID().String())
	ev.SetType("PromiseCreated")

	if err := ev.SetData(cloudevents.ApplicationJSON, map[string]interface{}{
		"created_at": promise.CreatedAt.Format(time.RFC3339Nano),
		"promise_id": promise.ID.String(),
	}); err != nil {
		return err
	}

	return w.Send(ctx, ev)
}

func (w w) NewEventPromiseRemoved(ctx context.Context, id domain.ID) error {
	ev := NewEvent()
	ev.SetID(domain.GenerateID().String())
	ev.SetType("PromiseDeleted")

	if err := ev.SetData(cloudevents.ApplicationJSON, map[string]interface{}{
		"created_at": time.Now().UTC().Format(time.RFC3339Nano),
		"promise_id": id.String(),
	}); err != nil {
		return err
	}

	return w.Send(ctx, ev)
}

func (w w) NewEventPromiseUpdated(ctx context.Context, id domain.ID) error {
	ev := NewEvent()
	ev.SetID(domain.GenerateID().String())
	ev.SetType("PromiseUpdated")

	if err := ev.SetData(cloudevents.ApplicationJSON, map[string]interface{}{
		"created_at": time.Now().UTC().Format(time.RFC3339Nano),
		"promise_id": id.String(),
	}); err != nil {
		return err
	}

	return w.Send(ctx, ev)
}
