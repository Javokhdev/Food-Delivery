package main

import (
	"log"
	"net"

	pb "learning-service/genproto/learning"
	"learning-service/service"
	postgres "learning-service/storage/postgres"
	"google.golang.org/grpc"
)

func main() {
	db, err := postgres.NewpostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}
	liss, err := net.Listen("tcp", ":8070")
	if err != nil {
		log.Fatal("Error while connection on tcp: ", err.Error())
	}

	s := grpc.NewServer()
	pb.RegisterLearningServiceServer(s, service.NewLearningService(db))

	log.Printf("Server listening at %v", liss.Addr())
	if err := s.Serve(liss); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
