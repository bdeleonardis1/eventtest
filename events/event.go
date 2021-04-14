package events

import "fmt"

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
	fmt.Println("we're appending the event", event.Name)
	el.Events = append(el.Events, event)
	fmt.Println("events after appending:", String(el.Events))
}

func (el *EventList) GetEvents() []*Event {
	fmt.Println("getting the events:", el.Events)
	return el.Events
}

func (el *EventList) ClearEvents() {
	fmt.Println("clearing events")
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
