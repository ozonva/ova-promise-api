package infrastructure

import (
	"go.uber.org/zap"
	"google.golang.org/grpc"

	grpcserver "github.com/ozonva/ova-promise-api/internal/implementation/grpc.server"
	pb "github.com/ozonva/ova-promise-api/internal/implementation/grpc.server/protocol"
	"github.com/ozonva/ova-promise-api/internal/usecase"
)

func InitGRPCServer(uc usecase.Handler, logger *zap.Logger) *grpc.Server {
	s := grpc.NewServer()
	server := grpcserver.NewPromiseService(uc, logger)
	pb.RegisterPromiseHandlerServer(s, server)

	return s
}
