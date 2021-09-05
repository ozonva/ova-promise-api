package grpcserver

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/ozonva/ova-promise-api/internal/domain"

	"go.uber.org/zap"

	pb "github.com/ozonva/ova-promise-api/internal/implementation/grpc.server/protocol"
	"github.com/ozonva/ova-promise-api/internal/usecase"
)

const chunkSize = 100

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
	span, ctx := opentracing.StartSpanFromContext(ctx, "grpc_handler:create_promise")
	defer span.Finish()

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
	span, ctx := opentracing.StartSpanFromContext(ctx, "grpc_handler:describe_update")
	defer span.Finish()

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

func (s *PromiseService) ListPromises(ctx context.Context, in *pb.ListPromisesRequest) (*pb.ListPromisesRequestResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "grpc_handler:list_promise")
	defer span.Finish()

	promises, err := s.ucHandler.PromiseGetList(ctx, in.Limit, in.Offset)
	if err != nil {
		return nil, err
	}

	res := &pb.ListPromisesRequestResponse{
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
	span, ctx := opentracing.StartSpanFromContext(ctx, "grpc_handler:remove_promise")
	defer span.Finish()

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

func (s *PromiseService) SavePromiseList(ctx context.Context, in *pb.ListPromisesRequestResponse) (*pb.SuccessMessage, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "grpc_handler:save_promise_list")
	defer span.Finish()

	res := pb.SuccessMessage{
		Message: "",
	}

	protoPromises := in.GetPromises()

	promises := make([]domain.Promise, len(protoPromises))

	for idx, promise := range protoPromises {
		id, err := uuid.Parse(promise.ID)
		if err != nil {
			return nil, err
		}

		promises[idx] = domain.Promise{
			ID:          id,
			UserID:      promise.UserID,
			Description: promise.Description,
			Status:      promise.Status,
			CreatedAt:   promise.CreatedAt.AsTime(),
			UpdatedAt:   promise.UpdatedAt.AsTime(),
		}

		if promise.DateDeadline != nil {
			t := promise.DateDeadline.AsTime()
			promises[idx].DateDeadline = &t
		}
	}

	if err := s.ucHandler.PromiseSaveListChunks(ctx, promises, chunkSize); err != nil {
		return &res, err
	}

	res.Message = "list successfully saved"

	return &res, nil
}

func (s *PromiseService) UpdatePromise(ctx context.Context, in *pb.UpdatePromiseRequest) (*pb.Promise, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "grpc_handler:update_promise")
	defer span.Finish()

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, err
	}

	in.GetDateDeadline()

	updateFileds := map[domain.PromiseUpdateProperty]interface{}{
		domain.PromiseStatus:       &in.Status,
		domain.PromiseDescription:  &in.Description,
		domain.PromiseDateDeadline: &in.DateDeadline,
	}

	if in.GetDateDeadline() != nil {
		updateFileds[domain.PromiseDateDeadline] = &in.DateDeadline
	}

	promise, err := s.ucHandler.PromiseUpdate(ctx, id, updateFileds)
	if err != nil {
		return nil, err
	}

	res := pb.Promise{
		ID:          promise.ID.String(),
		UserID:      promise.UserID,
		Description: promise.Description,
		Status:      promise.Status,
		CreatedAt:   timestamppb.New(promise.CreatedAt),
		UpdatedAt:   timestamppb.New(promise.UpdatedAt),
	}

	if promise.DateDeadline != nil {
		res.DateDeadline = timestamppb.New(*promise.DateDeadline)
	}

	return &res, nil
}
