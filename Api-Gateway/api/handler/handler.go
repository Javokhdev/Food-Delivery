package handler

import (
	pb "auth/genproto/auth"
	pbu "auth/genproto/user"
)

type Handler struct {
	Auth  pb.AuthServiceClient
	User  pbu.UserServiceClient
	Redis InMemoryStorageI
}

func NewHandler(auth pb.AuthServiceClient, user pbu.UserServiceClient, redis InMemoryStorageI) *Handler {
	return &Handler{
		Auth:  auth,
		User:  user,
		Redis: redis,
	}

}
