package api

import (
	"doduykhang/hermes-chat/pkg/config"
	"doduykhang/hermes-chat/pkg/handler"
	"doduykhang/hermes-chat/pkg/repository"
	"doduykhang/hermes-chat/pkg/route"
	"doduykhang/hermes-chat/pkg/service"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/olahol/melody"
)

var port = "8080"

func NewRest() {
	mux := chi.NewMux()
	mux.Use(middleware.Heartbeat("/ping"))

	//config
	melody := melody.New()
	redis := config.NewRedisClient()

	_, err := config.NewGormConnection()
	if err != nil {
		log.Panicf("Error connecting to database, %v\n", err)
	}

	gorm, err := config.NewGormConnection()
	if err != nil {
		log.Panicf("Error connecting to database, %v\n", err)
	}

	//repository
	userRepository := repository.NewUserRepository(gorm)

	//services
	pubSubService := service.NewPubSubService(redis)
	chatService := service.NewChat(melody, pubSubService, userRepository)

	//handlers
	chatHandler := handler.NewChat(chatService)

	//routes
	mux.Route("/chat", route.NewChat(chatHandler).Register)

	log.Printf("Chat service listening at port %s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), mux))
	
}
