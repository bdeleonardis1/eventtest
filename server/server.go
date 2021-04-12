package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/bdeleonardis1/eventtest/events"
)

type HandlerContext struct {
	eventList *events.EventList
}

func NewHandlerContext() *HandlerContext {
	return &HandlerContext{
		eventList: events.NewEventList(),
	}
}

func(hctx *HandlerContext) getEventsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(hctx.eventList.GetEvents())
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Only GET requests accepted at this endpoint"))
	}
}

func(hctx *HandlerContext) emitEventHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var event *events.Event
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(event)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("There was an error unmarshalling your event: %v", err)))
		}
		hctx.eventList.AppendEvent(event)
		w.WriteHeader(http.StatusCreated)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Only POST requests accepted at this endpoint"))
	}
}

func Serve() {
	hctx := NewHandlerContext()

	http.HandleFunc("/emitevent/", hctx.emitEventHandler)
	http.HandleFunc("/getevents/", hctx.getEventsHandler)

	fmt.Println("listening at 127.0.0.1:1111")
	log.Fatal(http.ListenAndServe(":1111", nil))
}