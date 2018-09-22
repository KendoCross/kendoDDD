package presentation

import (
	"container/list"
)

//WsEventType 类型
type WsEventType int

//WsEvent 数据交互结构
type WsEvent struct {
	Type      WsEventType // JOIN, LEAVE, MESSAGE
	User      string
	Timestamp int // Unix timestamp (secs)
	Content   string
}

const archiveSize = 20

// Event archives.
var archive = list.New()

// NewArchive saves new event to archive list.
func NewArchive(wsEvent WsEvent) {
	if archive.Len() >= archiveSize {
		archive.Remove(archive.Front())
	}
	archive.PushBack(wsEvent)
}

// GetWsEvents returns all events after lastReceived.
func GetWsEvents(lastReceived int) []WsEvent {
	wsEvents := make([]WsEvent, 0, archive.Len())
	for wsEvent := archive.Front(); wsEvent != nil; wsEvent = wsEvent.Next() {
		e := wsEvent.Value.(WsEvent)
		if e.Timestamp > int(lastReceived) {
			wsEvents = append(wsEvents, e)
		}
	}
	return wsEvents
}
