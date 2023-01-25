package route

import (
	"doduykhang/hermes-chat/pkg/handler"
	"doduykhang/hermes-chat/pkg/middleware"

	"github.com/go-chi/chi/v5"
)

type Chat struct {
	handler *handler.Chat
} 

func NewChat(h *handler.Chat) *Chat {
	return &Chat{
		handler: h,
	}	
}

func (c *Chat) Register(r chi.Router) {
	c.handler.HandleMessage()
	go c.handler.HandleWaitForMessage()
	r.With(middleware.GetUserId).Get("/connect/{roomID}", c.handler.HandleChat)	
}
