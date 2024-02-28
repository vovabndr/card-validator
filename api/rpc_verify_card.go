package api

import (
	"context"
	"github.com/vovabndr/card-validator/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) VerifyCard(ctx context.Context, req *pb.VerifyCardRequest) (*pb.VerifyCardResponse, error) {
	card := toDomain(req.GetCard())

	isValid, err := server.validationService.Validate(card)

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "message: %s", err)
	}

	response := &pb.VerifyCardResponse{
		Valid: isValid,
	}

	return response, nil
}
