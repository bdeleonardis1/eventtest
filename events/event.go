package events

import (
	"sync"
)

type Event struct {
	Name string `json:"name"`
}

func NewEvent(name string) *Event {
	return &Event{
		Name: name,
	}
}

func (e *Event) Equals(o *Event) bool {
	return e.Name == o.Name
}

func (e *Event) String() string {
	return e.Name
}

type EventList struct {
	Events []*Event
	mutex *sync.Mutex
}

func NewEventList() *EventList {
	return &EventList{
		Events: make([]*Event, 0),
		mutex: new(sync.Mutex),
	}
}

func (el *EventList) AppendEvent(event *Event) {
	el.mutex.Lock()
	defer el.mutex.Unlock()
	el.Events = append(el.Events, event)
}

func (el *EventList) GetEvents() []*Event {
	el.mutex.Lock()
	defer el.mutex.Unlock()
	return el.Events
}

func (el *EventList) ClearEvents() {
	el.mutex.Lock()
	defer el.mutex.Unlock()
	el.Events = make([]*Event, 0)
}

func String(events []*Event) string {
	str := ""
	for _, event := range events {
		str += event.Name + ", "
	}

	if len(str) < 2 {
		return str
	}

	return str[:len(str)-2]
}
