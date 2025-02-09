package batch

import (
	"context"
	"errors"
	"github.com/Imtiaz246/Thesis-Management-System/internal/apis/v1"
	"github.com/Imtiaz246/Thesis-Management-System/internal/model"
	"github.com/Imtiaz246/Thesis-Management-System/internal/repository"
	"github.com/Imtiaz246/Thesis-Management-System/internal/service"
	"github.com/Imtiaz246/Thesis-Management-System/internal/service/utils"
)

type Service interface {
	CreateBatch(ctx context.Context, requesterUniId string, req *v1.CreateBatchRequest) error
	ListBatches(ctx context.Context) ([]*v1.BatchInfo, error)
	GetBatchById(ctx context.Context, id uint) (*v1.BatchInfo, error)
	UpdateBatch(ctx context.Context, requesterUniId string, id uint, req *v1.UpdateBatchRequest) error
	DeleteBatch(ctx context.Context, requesterUniId string, id uint) error
}

func NewBatchService(service *service.Service, userRepo repository.UserRepository, batchRepo repository.BatchRepository) Service {
	return &batchService{
		batchRepo: batchRepo,
		userRepo:  userRepo,
		Service:   service,
	}
}

type batchService struct {
	batchRepo repository.BatchRepository
	userRepo  repository.UserRepository
	*service.Service
}

func (s *batchService) CreateBatch(ctx context.Context, requesterUniId string, req *v1.CreateBatchRequest) error {
	requester, err := s.userRepo.GetByUniversityId(ctx, requesterUniId)
	if err != nil {
		if errors.Is(err, v1.ErrNotFound) {
			return v1.ErrUnauthorized
		}
		return v1.ErrInternalServerError
	}
	// TODO: currently only admin can create a batch
	// This will be updated when roles are implemented
	// Like convenor can create a batch
	if !requester.IsAdmin {
		return v1.ErrUnauthorized
	}

	existingBatch, err := s.batchRepo.GetByName(ctx, req.Name)
	if err != nil && !errors.Is(err, v1.ErrBatchNotFound) {
		return err
	}
	if existingBatch != nil {
		return v1.ErrBatchAlreadyExists
	}

	batch := &model.Batch{
		Name:  req.Name,
		Quota: req.Quota,

		TeamRegDeadline: req.TeamRegDeadline,
		MaxTeamMember:   req.MaxTeamMember,
		MaxTeacherPref:  req.MaxTeacherPref,
		PreDefenceAt:    req.PreDefenceAt,
		DefenceAt:       req.DefenceAt,
		CreatedBy:       requester,
	}

	err = s.batchRepo.Create(ctx, batch)
	if err != nil {
		return err
	}

	return nil
}

func (s *batchService) ListBatches(ctx context.Context) ([]*v1.BatchInfo, error) {
	batches, err := s.batchRepo.GetAllWithCreatorInfo(ctx)
	if err != nil {
		return nil, err
	}
	batchesApiFormat := make([]*v1.BatchInfo, len(batches))
	for i, batch := range batches {
		batchesApiFormat[i] = batch.ConvertToApiFormat()
	}

	return batchesApiFormat, nil
}

func (s *batchService) GetBatchById(ctx context.Context, id uint) (*v1.BatchInfo, error) {
	batch, err := s.batchRepo.GetById(ctx, id)
	if err != nil {
		if errors.Is(err, v1.ErrBatchNotFound) {
			return nil, v1.ErrBatchNotFound
		}
		return nil, v1.ErrInternalServerError
	}

	return batch.ConvertToApiFormat(), nil
}

func (s *batchService) UpdateBatch(ctx context.Context, requesterUniId string, id uint, req *v1.UpdateBatchRequest) error {
	requester, err := s.userRepo.GetByUniversityId(ctx, requesterUniId)
	if err != nil {
		if errors.Is(err, v1.ErrNotFound) {
			return v1.ErrUnauthorized
		}
		return v1.ErrInternalServerError
	}
	// TODO: currently only admin can update a batch
	// This will be updated when roles are implemented
	// Like convenor can update a batch
	if !requester.IsAdmin {
		return v1.ErrUnauthorized
	}

	batch, err := s.batchRepo.GetById(ctx, id)
	if err != nil {
		if errors.Is(err, v1.ErrNotFound) {
			return v1.ErrBatchNotFound
		}
		return v1.ErrInternalServerError
	}

	utils.SetIfNonDefault(&req.Name, &batch.Name)
	utils.SetIfNonDefault(&req.Quota, &batch.Quota)
	utils.SetIfNonDefault(&req.MaxTeamMember, &batch.MaxTeamMember)
	utils.SetIfNonDefault(&req.MaxTeacherPref, &batch.MaxTeacherPref)
	utils.SetIfNonDefault(req.TeamRegDeadline, &batch.TeamRegDeadline)
	utils.SetIfNonDefault(req.PreDefenceAt, &batch.PreDefenceAt)
	utils.SetIfNonDefault(req.DefenceAt, &batch.DefenceAt)

	err = s.batchRepo.Update(ctx, batch)
	if err != nil {
		return err
	}

	return nil
}

func (s *batchService) DeleteBatch(ctx context.Context, requesterUniId string, id uint) error {
	requester, err := s.userRepo.GetByUniversityId(ctx, requesterUniId)
	if err != nil {
		if errors.Is(err, v1.ErrNotFound) {
			return v1.ErrUnauthorized
		}
		return v1.ErrInternalServerError
	}
	// TODO: currently only admin can update a batch
	// This will be updated when roles are implemented
	// Like convenor can update a batch
	if !requester.IsAdmin {
		return v1.ErrUnauthorized
	}

	batch, err := s.batchRepo.GetById(ctx, id)
	if err != nil {
		if errors.Is(err, v1.ErrBatchNotFound) {
			return v1.ErrBatchNotFound
		}
		return v1.ErrInternalServerError
	}

	err = s.batchRepo.Delete(ctx, batch.ID)
	if err != nil {
		return err
	}

	return nil
}
