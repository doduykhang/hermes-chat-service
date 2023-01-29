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

var port = "8082"

func NewRest() {
	mux := chi.NewMux()
	mux.Use(middleware.Heartbeat("/ping"))
	mux.Use(middleware.Logger)

	//config
	conf := config.LoadEnv(".")
	melody := melody.New()
	redis := config.NewRedisClient(conf)

	gorm, err := config.NewGormConnection(conf)
	if err != nil {
		log.Panicf("Error connecting to database, %v\n", err)
	}
	rabbitMQ, err := config.NewRabbitMQConnection(conf)
	if err != nil {
		log.Panicf("Error connecting to rabbit mq, %v\n", err)
	}
	defer rabbitMQ.Close()

	//repository
	userRepository := repository.NewUserRepository(gorm)

	//services
	pubSubService := service.NewPubSubService(redis)
	queueService := service.NewQueue(rabbitMQ)

	chatService := service.NewChat(melody, pubSubService, userRepository, queueService)
	userService := service.NewUserService(userRepository)

	//handlers
	chatHandler := handler.NewChat(chatService)
	userHandler := handler.NewUser(userService, queueService)

	//routes
	mux.Route("/api/chat", route.NewChat(chatHandler).Register)

	//other
	go userHandler.HandleAddUser()
	go userHandler.HandleRemoveUser()

	log.Printf("Chat service listening at port %s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), mux))
	
}
