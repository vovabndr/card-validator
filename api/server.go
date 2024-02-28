package api

import (
	"github.com/vovabndr/card-validator/domain"
	"github.com/vovabndr/card-validator/pb"
)

type Server struct {
	pb.UnimplementedCardValidatorServer
	validationService *domain.CardValidationService
}

func NewServer(service *domain.CardValidationService) *Server {
	server := &Server{validationService: service}
	return server
}
