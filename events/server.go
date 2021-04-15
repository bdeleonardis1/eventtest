package events

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HandlerContext struct {
	eventList *EventList
}

func NewHandlerContext() *HandlerContext {
	return &HandlerContext{
		eventList: NewEventList(),
	}
}

func (hctx *HandlerContext) getEventsHandler(w http.ResponseWriter, r *http.Request) {
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

func (hctx *HandlerContext) emitEventHandler(w http.ResponseWriter, r *http.Request) {
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
		w.WriteHeader(http.StatusCreated)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Only POST requests accepted at this endpoint"))
	}
}

func (hctx *HandlerContext) clearEventHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		hctx.eventList.ClearEvents()
		w.WriteHeader(http.StatusCreated)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Only POST requests accepted at this endpoint"))
	}
}

func createServer(port string) *http.Server {
	hctx := NewHandlerContext()

	mux := http.NewServeMux()
	mux.HandleFunc("/emitevent/", hctx.emitEventHandler)
	mux.HandleFunc("/getevents/", hctx.getEventsHandler)
	mux.HandleFunc("/clearevents/", hctx.clearEventHandler)

	server := &http.Server{Addr: ":" + port, Handler: mux}

	fmt.Println("listening at 127.0.0.1:" + port)
	go func() {
		server.ListenAndServe()
	}()

	return server
}
