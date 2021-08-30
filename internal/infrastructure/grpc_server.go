package infrastructure

import (
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcZap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpcRecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpcCtxTags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpcOpentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	grpcPrometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	grpcServer "github.com/ozonva/ova-promise-api/internal/implementation/grpc.server"
	pb "github.com/ozonva/ova-promise-api/internal/implementation/grpc.server/protocol"
	"github.com/ozonva/ova-promise-api/internal/usecase"
)

var grpcMetrics = grpcPrometheus.NewServerMetrics()

func InitGRPCServer(uc usecase.Handler, logger *zap.Logger) *grpc.Server {
	s := grpc.NewServer(
		grpc.StreamInterceptor(grpcMiddleware.ChainStreamServer(
			grpcRecovery.StreamServerInterceptor(),
			grpcCtxTags.StreamServerInterceptor(),
			grpcOpentracing.StreamServerInterceptor(),
			grpcMetrics.StreamServerInterceptor(),
			grpcZap.StreamServerInterceptor(logger),
		)),
		grpc.UnaryInterceptor(grpcMiddleware.ChainUnaryServer(
			grpcRecovery.UnaryServerInterceptor(),
			grpcCtxTags.UnaryServerInterceptor(),
			grpcOpentracing.UnaryServerInterceptor(),
			grpcMetrics.UnaryServerInterceptor(),
			grpcZap.UnaryServerInterceptor(logger),
		)),
	)
	server := grpcServer.NewPromiseService(uc, logger)
	pb.RegisterPromiseHandlerServer(s, server)

	return s
}
