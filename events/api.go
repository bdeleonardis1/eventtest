package events

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

type IsOrdered int
const (
	Ordered IsOrdered = iota
	Unordered
)

func EmitEvent(event *Event) error {
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

func GetEvents() ([]*Event, error) {
	res, err := http.Get("http://127.0.0.1:1111/getevents")
	if err != nil {
		return nil, err
	}
	var events []*Event
	err = json.NewDecoder(res.Body).Decode(&events)
	if err != nil {
		return nil, err
	}
	return events, nil
}

func ClearEvents() error {
	res, err := http.Post("http://127.0.0.1:1111/clearevents/", "application/json", bytes.NewBuffer([]byte("{}")))
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusAccepted {
		return fmt.Errorf("received a %v status code when trying to clear the events", res.StatusCode)
	}
	return nil
}

func ExpectExactEvents(t *testing.T, expectedEvents []*Event) {
	t.Helper()

	actualEvents, err := GetEvents()
	if err != nil {
		t.Error(err)
	}

	if len(actualEvents) != len(expectedEvents) {
		t.Fatalf("actual events: %v, not the same length as expected events: %v", String(actualEvents), String(expectedEvents))
	}

	for i, actualEvent := range actualEvents {
		if !actualEvent.Equals(expectedEvents[i]) {
			t.Errorf("the %vth actual event: %v, does not equal the expected event: %v", actualEvent, expectedEvents[i])
		}
	}
}

func ExpectEvents(t *testing.T, expectedEvents []*Event, ordered IsOrdered) {
	t.Helper()

	actualEvents, err := GetEvents()
	fmt.Println("actualEvents in ExpectEvents: ", String(actualEvents))

	if err != nil {
		t.Fatalf("error getting events: %v", err)
	}

	if ordered == Ordered {
		expectEventsOrdered(t, expectedEvents, actualEvents)
	} else {
		expectEventsUnordered(t, expectedEvents, actualEvents)
	}
}

func expectEventsOrdered(t *testing.T, expectedEvents, actualEvents []*Event) {
	t.Helper()

	actualIdx := 0
	for _, expectedEvent := range expectedEvents {
		found := false
		for actualIdx < len(actualEvents) {
			if expectedEvent.Equals(actualEvents[actualIdx]) {
				found = true
				break
			}
			actualIdx += 1
		}
		if !found {
			t.Fatalf("could not find expected event: %v", expectedEvent)
		}
	}
}

func expectEventsUnordered(t *testing.T, expectedEvents, actualEvents []*Event) {
	t.Helper()

	for _, expectedEvent := range expectedEvents {
		found := false
		for _, actualEvent := range actualEvents {
			if expectedEvent.Equals(actualEvent) {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("could not find expected event: %v", expectedEvent)
		}
	}
}

func UnexpectedEvents(t *testing.T, unexpectedEvents []*Event) {
	t.Helper()

	actualEvents, err := GetEvents()
	if err != nil {
		t.Fatal(err)
	}

	for _, unexpected := range unexpectedEvents {
		for _, actualEvent := range actualEvents {
			if unexpected.Equals(actualEvent) {
				t.Errorf("event: %v occurred even though it should not have", unexpected)
			}
		}
	}
}