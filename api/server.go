package api

import "github.com/vovabndr/card-validator/pb"

type Server struct {
	pb.UnimplementedCardValidatorServer
}

func NewServer() (*Server, error) {
	server := &Server{}
	return server, nil
}
