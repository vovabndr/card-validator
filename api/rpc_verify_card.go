package api

import (
	"context"
	"github.com/rs/zerolog/log"
	"github.com/vovabndr/card-validator/pb"
)

func (server *Server) VerifyCard(ctx context.Context, req *pb.VerifyCardRequest) (*pb.VerifyCardResponse, error) {

	card := req.GetCard()

	log.Info().
		Any("card", card).
		Msg("received request")

	response := &pb.VerifyCardResponse{
		Valid: false,
	}

	return response, nil
}
