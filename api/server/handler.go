package server

import (
	"context"
	"log"
	"net/http"

	"github.com/acmcsufoss/api.acmcsuf.com/api"
	"github.com/go-chi/chi/v5"
)

// HandlerOptions is the options for creating a new Handler instance.
type HandlerOptions struct {
	Ctx   context.Context
	Port  string
	Store api.Store
}

// Handler is the API server handler.
type Handler struct {
	Client
	Port   string
	router chi.Router
	store  api.Store
}

// Serve starts the server.
func (h *Handler) Serve() error {
	log.Printf("Listening on http://127.0.0.1%s...", h.Port)
	return http.ListenAndServe(h.Port, h.router)
}

func (h *Handler) getEvents(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))

	/*
		result, err := s.CreateEvent(api.NewCreateEventRequest(
			"Test Event",
			"Test Event Content",
			"https://placekitten.com/200/300",
			"Test Location",
			// New Year's 2023
			time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			// 2 hours in milliseconds.
			2*60*60*1000,
			false,
			"Test Host",
			api.VisibilityPublic,
		))
		if err != nil {
			log.Fatalf("Error creating event: %v", err)
		}

		log.Printf("Created event with ID %s and created at %d", result.ID, result.CreatedAt)
	*/
}

// NewHandler creates a new Handler instance.
func NewHandler(o HandlerOptions) Handler {
	h := Handler{
		Client: *NewClient(o.Ctx),
		Port:   o.Port,
		store:  o.Store,
		router: chi.NewRouter(),
	}

	h.router.Get("/events", h.getEvents)

	return h
}

// Client wraps around state.State for some common functionalities.
type Client struct {
	ctx context.Context
}

// NewClient creates a new Client instance.
func NewClient(ctx context.Context) *Client {
	return &Client{ctx}
}

// Context returns the internal context.
func (c *Client) Context() context.Context {
	return c.ctx
}
