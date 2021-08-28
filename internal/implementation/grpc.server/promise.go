package grpcserver

import (
	"context"

	"go.uber.org/zap"

	pb "github.com/ozonva/ova-promise-api/internal/implementation/grpc.server/protocol"
	"github.com/ozonva/ova-promise-api/internal/usecase"
)

type PromiseService struct {
	pb.UnimplementedPromiseHandlerServer
	ucHandler usecase.Handler
	logger    *zap.Logger
}

func NewPromiseService(uc usecase.Handler, logger *zap.Logger) *PromiseService {
	return &PromiseService{
		ucHandler: uc,
		logger:    logger,
	}
}

func (s *PromiseService) CreatePromise(ctx context.Context, in *pb.CreateRequest) (*pb.Promise, error) {
	s.logger.Info(
		"incoming request",
		zap.Int64("user-id", in.UserID),
	)

	return nil, nil
}

func (s *PromiseService) DescribePromise(ctx context.Context, in *pb.UUID) (*pb.Promise, error) {
	s.logger.Info(
		"incoming request",
		zap.Any("promise-id", in.Id),
	)

	return nil, nil
}

func (s *PromiseService) ListPromises(ctx context.Context, in *pb.ListPromisesRequest) (*pb.ListPromisesResponse, error) {
	s.logger.Info(
		"incoming request",
		zap.Int64("limit", in.Limit),
		zap.Int64("offset", in.Offset),
	)

	return nil, nil
}

func (s *PromiseService) RemovePromise(ctx context.Context, in *pb.UUID) (*pb.SuccessMessage, error) {
	s.logger.Info(
		"incoming request",
		zap.Any("promise-id", in.Id),
	)

	return nil, nil
}
