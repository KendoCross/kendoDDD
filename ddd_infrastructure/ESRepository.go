package ddd_infrastructure

import (
	"context"

	"../ddd_interfaces"
	"../dddcore"
)

type eventSourcedRepository struct {
	streamName string
	eventStore dddcore.EventStore
	eventBus   dddcore.EventBus
}

// Save current user changes to event store and publish each event with an event bus
func (r *eventSourcedRepository) Save(ctx context.Context, u dddcore.IEventMsg) error {
	err := r.eventStore.Store(u.Changes())
	if err != nil {
		return err
	}

	for _, event := range u.Changes() {
		r.eventBus.Publish(ctx, event.Metadata.Type, *event)
	}

	return nil
}

// Get user with current state applied
// 需要实现从数据库里获取
// func (r *eventSourcedRepository) Get(id uuid.UUID) *dddcore.IEventMsg {
// 	events := r.eventStore.GetStream(id, r.streamName)

// 	u := user.New()
// 	u.FromHistory(events)

// 	return u
// }

// NewUserEventSourcedRepository creates new user event sourced repository
func NewESRepository(streamName string, store dddcore.EventStore, bus dddcore.EventBus) ddd_interfaces.IESRepository {
	return &eventSourcedRepository{streamName, store, bus}
}
