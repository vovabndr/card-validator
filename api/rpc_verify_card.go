package api

import (
	"context"
	"errors"
	"github.com/vovabndr/card-validator/domain"
	"github.com/vovabndr/card-validator/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) VerifyCard(ctx context.Context, req *pb.VerifyCardRequest) (*pb.VerifyCardResponse, error) {
	card := toDomain(req)

	isValid, err := server.validationService.Validate(card)

	if errors.Is(err, domain.ErrNoPaymentSystem) {
		return nil, status.Errorf(codes.Unimplemented, "message: %s", err)
	}

	response := &pb.VerifyCardResponse{Valid: isValid}

	if err != nil {
		response.Error = &pb.VerifyCardErrorResponse{
			Code:    int64(codes.InvalidArgument),
			Message: err.Error(),
		}
	}

	return response, nil
}
