package events

type Event struct {
	Name string `json:"name"`
}

func NewEvent(name string) *Event {
	return &Event{
		Name: name,
	}
}

type EventList struct {
	Events []*Event
}

func NewEventList() *EventList {
	return &EventList{
		Events: make([]*Event, 0),
	}
}

func (el *EventList) AppendEvent(event *Event) {
	el.Events = append(el.Events, event)
}

func (el *EventList) GetEvents() []*Event {
	return el.Events
}

func (el *EventList) ClearEvents() {
	el.Events = make([]*Event, 0)
}
