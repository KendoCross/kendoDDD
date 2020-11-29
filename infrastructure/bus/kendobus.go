package bus

import (
	"context"
	"errors"
	"sync"

	eh "github.com/looplab/eventhorizon"
)

var ErrHandlerAlreadySet = errors.New("handler is already set")

var ErrHandlerNotFound = errors.New("no handlers for command")

var dealers map[eh.CommandType]CommandDealer = make(map[eh.CommandType]CommandDealer)
var dealersMu sync.RWMutex

type CommandDealer interface {
	DealCommand(context.Context, eh.Command) (interface{}, error)
}

func DealerCommand(ctx context.Context, cmd eh.Command) (interface{}, error) {
	dealersMu.RLock()
	defer dealersMu.RUnlock()

	if dealer, ok := dealers[cmd.CommandType()]; ok {
		return dealer.DealCommand(ctx, cmd)
	}

	return nil, ErrHandlerNotFound
}

func SetDealer(dealer CommandDealer, cmdType eh.CommandType) error {
	dealersMu.Lock()
	defer dealersMu.Unlock()

	if _, ok := dealers[cmdType]; ok {
		return ErrHandlerAlreadySet
	}

	dealers[cmdType] = dealer
	return nil
}
