package events

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

func String(events []*Event) string {
	str := ""
	for _, event := range events {
		str += event.Name + ", "
	}
	return str[:len(str)-2]
}
