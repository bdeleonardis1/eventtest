package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bdeleonardis1/eventtest/events"
)

func EmitEvent(event *events.Event) error {
	marshaledEvent, err := json.Marshal(event)
	if err != nil {
		return err
	}

	// TODO: get this URL from a configuration
	resp , err := http.Post("http://127.0.0.1:1111/emitevent/", "application/json", bytes.NewBuffer(marshaledEvent))
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("received a %v status code when trying to emit an event", resp.StatusCode)
	}

	return nil
}

func GetEvents(event *events.Event) ([]*events.Event, error) {
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