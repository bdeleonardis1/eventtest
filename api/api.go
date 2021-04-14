package eventtestapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/bdeleonardis1/eventtest/events"
)

func EmitEvent(event *events.Event) error {
	marshaledEvent, err := json.Marshal(event)
	if err != nil {
		return err
	}

	// TODO: get this URL from a configuration
	res, err := http.Post("http://127.0.0.1:1111/emitevent/", "application/json", bytes.NewBuffer(marshaledEvent))
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusCreated {
		return fmt.Errorf("received a %v status code when trying to emit an event", res.StatusCode)
	}

	return nil
}

func GetEvents() ([]*events.Event, error) {
	res, err := http.Get("http://127.0.0.1:1111/getevents")
	if err != nil {
		return nil, err
	}
	var events []*events.Event
	err = json.NewDecoder(res.Body).Decode(&events)
	if err != nil {
		return nil, err
	}
	return events, nil
}

func ClearEvents() error {
	fmt.Println("Attempting to clear the events")

	res, err := http.Post("http://127.0.0.1:1111/clearevents/", "application/json", bytes.NewBuffer([]byte("{}")))
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusAccepted {
		return fmt.Errorf("received a %v status code when trying to clear the events", res.StatusCode)
	}
	return nil
}

func ExpectEvents(t *testing.T, expectedEvents []*events.Event) {
	actualEvents, err := GetEvents()
	if err != nil {
		t.Error(err)
	}

	if len(actualEvents) != len(expectedEvents) {
		t.Errorf("actual events: %v, not the same length as expected events: %v", events.String(actualEvents), events.String(expectedEvents))
	}

	for i, actualEvent := range actualEvents {
		if !actualEvent.Equals(expectedEvents[i]) {
			t.Errorf("the %vth actual event: %v, does not equal the expected event: %v", actualEvent, expectedEvents[i])
		}
	}
}