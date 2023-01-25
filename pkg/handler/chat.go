package handler

import (
	"doduykhang/hermes-chat/pkg/service"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Chat struct{
	service service.Chat	
}

func NewChat(s service.Chat) *Chat {
	return &Chat{
		service: s,
	}
}

func (c *Chat) HandleChat(w http.ResponseWriter, r *http.Request) {
	roomID := chi.URLParam(r, "roomID")
	ctx := r.Context()
	userID := ctx.Value("userID").(string)
	err := c.service.ConnectToRoom(w, r, roomID, userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
}

func (c *Chat) HandleMessage() {
	c.service.HandleMessage()
}

func (c *Chat) HandleWaitForMessage() {
	c.service.HandleWaitForMessage()
}
