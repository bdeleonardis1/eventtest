package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Event struct {
	Name string
}

type HandlerContext struct {
	events []Event
}

func NewHandlerContext() *HandlerContext {
	events := make([]Event, 0, 10)

	return &HandlerContext{
		events: events,
	}
}


func(hctx *HandlerContext) getEventsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(hctx.events)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Only GET requests accepted at this endpoint"))
	}
}

func(hctx *HandlerContext) emitEventHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var event Event
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&event)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("There was an error unmarshalling your event: %v", err)))
		}
		hctx.events = append(hctx.events, event)
		w.WriteHeader(http.StatusCreated)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Only POST requests accepted at this endpoint"))
	}
}

func main() {
	hctx := NewHandlerContext()

	http.HandleFunc("/emitevent/", hctx.emitEventHandler)
	http.HandleFunc("/getevents/", hctx.getEventsHandler)

	fmt.Println("listening at 127.0.0.1:1111")
	log.Fatal(http.ListenAndServe(":1111", nil))
}