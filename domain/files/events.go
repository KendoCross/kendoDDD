package files

import (
	"github.com/KendoCross/kendoDDD/interfaces"
	eh "github.com/looplab/eventhorizon"
)

const (
	FileAddedEvent eh.EventType = "FileAddedEvent"
)

func init() {
	// Only the event for creating an invite has custom data.
	eh.RegisterEventData(FileAddedEvent, func() eh.EventData {
		return &interfaces.FileInfo{}
	})

}
