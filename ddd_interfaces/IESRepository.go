package ddd_interfaces

import (
	"context"

	"../dddcore"
	//"github.com/google/uuid"
)

type IESRepository interface {
	Save(ctx context.Context, evetnMsg dddcore.IEventMsg) error
	//Get(id uuid.UUID) *User
}
