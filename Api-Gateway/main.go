package main

import (
	"fmt"
	"log"

	"api-gateway/api"
	"api-gateway/api/handler"
	pbl "api-gateway/genproto/learning"
	pbu "api-gateway/genproto/user"
	"api-gateway/kafka"
	

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "api-gateway/genproto/game"
)

func main() {
	LearningConn, err := grpc.NewClient(fmt.Sprintf("learning_service%s", ":8070"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error while Newclient: ", err.Error())
	}
	defer LearningConn.Close()

	GameCon, err := grpc.NewClient(fmt.Sprintf("game_service%s", ":8060"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error while Newclient: ", err.Error())
	}
	defer GameCon.Close()

	UsrCon, err := grpc.NewClient(fmt.Sprintf("auth_service%s", ":8088"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error while Newclient: ", err.Error())
	}
	defer GameCon.Close()

	kaf, err := kafka.NewKafkaProducer([]string{"kafka:9092"})
	if err != nil {
		log.Fatal("Error while NewKafkaProducer: ", err.Error())
	}
	defer kaf.Close()

	us := pbl.NewLearningServiceClient(LearningConn)
	cs := pb.NewGameServiceClient(GameCon)
	usr := pbu.NewUserServiceClient(UsrCon)


	h := handler.NewHandler(us, cs, usr, kaf)
	r := api.NewGin(h)

	fmt.Println("Server started on port:8077")
	err = r.Run(":8077")
	if err != nil {
		log.Fatal("Error while running server: ", err.Error())
	}
}
