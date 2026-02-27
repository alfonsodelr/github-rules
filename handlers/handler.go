package handlers

import (
	"fmt"
	"net/http"

	"github.com/alfonsodelr/sample-server/clients"
)

type Handler struct {
	name string
}

func NewHandler(name string) *Handler {
	return &Handler{
		name: name,
	}
}

// Swaggo annotations for API documentation
// @Summary Handle version request
// @Description Processes the version request and returns a response.
// @Tags version
// @Accept json
// @Produce json
// @Success 200 {string} string "Handler response"
// @Router /version [get]
func (h *Handler) HandleVersionRequest() string {
	// Create a client instance to demonstrate usage
	client := clients.NewClient("1.0")
	fmt.Printf("Client version: %s\n", client.GetVersion())
	return fmt.Sprintf("Handler %s is processing the version request with version %s", h.name, client.GetVersion())
}

func HandlerVersionRequest(w http.ResponseWriter, r *http.Request) {
	h := NewHandler("default")
	_, _ = w.Write([]byte(h.HandleVersionRequest()))
}
