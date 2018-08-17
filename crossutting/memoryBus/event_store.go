package memoryBus

import (
	"errors"
	"sync"

	dddcore "../../dddcore"
	"github.com/google/uuid"
)

var ErrEventNotFound = errors.New("Event not found")

//内存里的ES，需要持久化的实现。
type eventStore struct {
	mtx    sync.RWMutex
	events map[string]*dddcore.Event
}

func (s *eventStore) Store(events []*dddcore.Event) error {
	if len(events) == 0 {
		return nil
	}

	s.mtx.Lock()
	defer s.mtx.Unlock()

	// todo check event version
	for _, e := range events {
		s.events[e.ID.String()] = e
	}

	return nil
}

func (s *eventStore) Get(id uuid.UUID) (*dddcore.Event, error) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	if val, ok := s.events[id.String()]; ok {
		return val, nil
	}
	return nil, ErrEventNotFound
}

func (s *eventStore) FindAll() []*dddcore.Event {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	es := make([]*dddcore.Event, 0, len(s.events))
	for _, val := range s.events {
		es = append(es, val)
	}
	return es
}

func (s *eventStore) GetStream(streamID uuid.UUID, streamName string) []*dddcore.Event {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	e := make([]*dddcore.Event, 0, 0)
	for _, val := range s.events {
		if val.Metadata.StreamName == streamName && val.Metadata.StreamID == streamID {
			e = append(e, val)
		}
	}
	return e
}

// New creates in memory event store
func NewEventStore() dddcore.EventStore {
	return &eventStore{
		events: make(map[string]*dddcore.Event),
	}
}
