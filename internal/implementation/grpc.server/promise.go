package grpcserver

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/ozonva/ova-promise-api/internal/domain"

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
	res := &pb.Promise{
		UserID:       in.UserID,
		Description:  in.Description,
		DateDeadline: in.DateDeadline,
	}

	dateDeadline := in.DateDeadline.AsTime()

	p, err := domain.NewPromise(
		domain.GenerateID(),
		in.UserID,
		in.Description,
		&dateDeadline,
	)

	if err != nil {
		return res, err
	}

	if err := s.ucHandler.PromiseSave(ctx, p); err != nil {
		return nil, err
	}

	res.ID = p.ID.String()
	res.Status = p.Status
	res.CreatedAt = timestamppb.New(p.CreatedAt)
	res.UpdatedAt = timestamppb.New(p.UpdatedAt)

	if p.DateDeadline != nil {
		res.DateDeadline = timestamppb.New(*p.DateDeadline)
	}

	return res, nil
}

func (s *PromiseService) DescribePromise(ctx context.Context, in *pb.UUID) (*pb.Promise, error) {
	id, err := uuid.Parse(in.Id)
	if err != nil {
		return nil, err
	}

	p, err := s.ucHandler.PromiseGetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	res := pb.Promise{
		ID:          p.ID.String(),
		UserID:      p.UserID,
		Description: p.Description,
		Status:      p.Status,
		CreatedAt:   timestamppb.New(p.CreatedAt),
		UpdatedAt:   timestamppb.New(p.UpdatedAt),
	}

	if p.DateDeadline != nil {
		res.DateDeadline = timestamppb.New(*p.DateDeadline)
	}

	return &res, nil
}

func (s *PromiseService) ListPromises(ctx context.Context, in *pb.ListPromisesRequest) (*pb.ListPromisesResponse, error) {
	promises, err := s.ucHandler.PromiseGetList(ctx, in.Limit, in.Offset)
	if err != nil {
		return nil, err
	}

	res := &pb.ListPromisesResponse{
		Promises: nil,
	}

	for _, p := range promises {
		r := pb.Promise{
			ID:          p.ID.String(),
			UserID:      p.UserID,
			Description: p.Description,
			Status:      p.Status,
			CreatedAt:   timestamppb.New(p.CreatedAt),
			UpdatedAt:   timestamppb.New(p.UpdatedAt),
		}

		if p.DateDeadline != nil {
			r.DateDeadline = timestamppb.New(*p.DateDeadline)
		}

		res.Promises = append(res.Promises, &r)
	}

	return res, nil
}

func (s *PromiseService) RemovePromise(ctx context.Context, in *pb.UUID) (*pb.SuccessMessage, error) {
	res := pb.SuccessMessage{
		Message: "",
	}

	id, err := uuid.Parse(in.Id)
	if err != nil {
		return &res, err
	}

	if err := s.ucHandler.PromiseRemove(ctx, id); err != nil {
		return &res, err
	}

	res.Message = fmt.Sprintf("promise with id=%s successfully deleted", in.Id)

	return &res, nil
}
