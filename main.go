package main

import (
	"context"
	"errors"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/vovabndr/card-validator/api"
	"github.com/vovabndr/card-validator/domain"
	"github.com/vovabndr/card-validator/pb"
	"github.com/vovabndr/card-validator/utils"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var interruptSignals = []os.Signal{
	os.Interrupt,
	syscall.SIGTERM,
	syscall.SIGINT,
}

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal().Msg("Couldn't load config")
	}

	if config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	ctx, stop := signal.NotifyContext(context.Background(), interruptSignals...)
	defer stop()

	waitGroup, ctx := errgroup.WithContext(ctx)

	paymentSystems := []domain.PaymentSystem{
		domain.NewMastercardPaymentSystem(),
		domain.NewVisaPaymentSystem(),
	}
	validationService := domain.NewCardValidationService(paymentSystems)

	runGrpcServer(ctx, waitGroup, config, validationService)
	runGatewayServer(ctx, waitGroup, config, validationService)

	err = waitGroup.Wait()
	if err != nil {
		log.Fatal().Err(err).Msg("error from wait group")
	}
}

func runGrpcServer(
	ctx context.Context,
	waitGroup *errgroup.Group,
	config utils.Config,
	service *domain.CardValidationService,
) {
	server := api.NewServer(service)

	logger := grpc.UnaryInterceptor(api.GrpcLogger)
	grpcServer := grpc.NewServer(logger)
	pb.RegisterCardValidatorServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal().Msg("Couldn't create listener")
	}

	waitGroup.Go(func() error {
		log.Info().Msgf("start gRPC at %s", listener.Addr().String())
		err = grpcServer.Serve(listener)
		if err != nil {
			if errors.Is(err, grpc.ErrServerStopped) {
				return nil
			}
			log.Error().Msg("gRPC server failed to server")
			return err
		}
		return nil
	})

	waitGroup.Go(func() error {
		<-ctx.Done()
		log.Info().Msg("graceful shutdown gRPC server")
		grpcServer.GracefulStop()
		log.Info().Msg("gRPC server is stopped")
		return nil
	})
}

func runGatewayServer(
	ctx context.Context,
	waitGroup *errgroup.Group,
	config utils.Config,
	service *domain.CardValidationService,
) {
	server := api.NewServer(service)

	jsonOption := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames:     true,
			EmitDefaultValues: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})

	grpcMux := runtime.NewServeMux(jsonOption)

	err := pb.RegisterCardValidatorHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatal().Msg("Couldn't register handler server")
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	httpServer := http.Server{
		Handler: api.HttpLogger(mux),
		Addr:    config.HTTPServerAddress,
	}

	waitGroup.Go(func() error {
		log.Info().Msgf("start HTTP gateway at %s", httpServer.Addr)
		err := httpServer.ListenAndServe()

		if err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				return nil
			}

			log.Error().Msg("HTTP Gateway server failed to serve")
			return err
		}
		return nil
	})

	waitGroup.Go(func() error {
		<-ctx.Done()
		log.Info().Msg("graceful shutdown HTTP gateway server")
		err := httpServer.Shutdown(context.Background())
		if err != nil {
			log.Error().Err(err).Msg("failed to shutdown HTTP gateway server")
		}
		log.Info().Msg("HTTP gateway server is stopped")
		return nil
	})
}
