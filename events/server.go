package events

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type handlerContext struct {
	eventList *EventList
}

const (
	defaultPort = "1111"
)

func newHandlerContext() *handlerContext {
	return &handlerContext{
		eventList: NewEventList(),
	}
}

func (hctx *handlerContext) getEventsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getEvents handler", r.Method)

	switch r.Method {
	case http.MethodGet:
		fmt.Println("events in getEvents", hctx.eventList.GetEvents())

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		json.NewEncoder(w).Encode(hctx.eventList.GetEvents())
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Only GET requests accepted at this endpoint"))
	}
}

func (hctx *handlerContext) emitEventHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside emitEvent", r.Method)
	switch r.Method {
	case http.MethodPost:
		var event Event
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&event)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("There was an error unmarshalling your event: %v", err)))
		}
		hctx.eventList.AppendEvent(&event)
		fmt.Println("Events after emitting:", hctx.eventList.GetEvents())
		w.WriteHeader(http.StatusCreated)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Only POST requests accepted at this endpoint"))
	}
}

func (hctx *handlerContext) clearEventHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		hctx.eventList.ClearEvents()
		w.WriteHeader(http.StatusCreated)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Only POST requests accepted at this endpoint"))
	}
}

func createServer() *http.Server {
	hctx := newHandlerContext()

	mux := http.NewServeMux()
	mux.HandleFunc("/emitevent/", hctx.emitEventHandler)
	mux.HandleFunc("/getevents/", hctx.getEventsHandler)
	mux.HandleFunc("/clearevents/", hctx.clearEventHandler)

	server := &http.Server{Addr: ":" + defaultPort, Handler: mux}

	go func() {
		server.ListenAndServe()
	}()

	return server
}
