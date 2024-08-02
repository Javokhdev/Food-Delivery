package handler

import (
	pb "api-gateway/genproto/game"
	pbl "api-gateway/genproto/learning"
	pbu "api-gateway/genproto/user"
	"api-gateway/kafka"
)

type Handler struct {
	Learning pbl.LearningServiceClient
	Game     pb.GameServiceClient
	User     pbu.UserServiceClient
	Kaf      kafka.KafkaProducer
}

func NewHandler(learn pbl.LearningServiceClient, game pb.GameServiceClient, user pbu.UserServiceClient, kaff kafka.KafkaProducer) *Handler {
	return &Handler{
		Learning: learn,
		Game:     game,
		User:     user,
        Kaf:      kaff,
	}
}
