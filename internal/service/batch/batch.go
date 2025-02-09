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
	ListAllBatches(ctx context.Context) ([]*v1.BatchInfo, error)
	ListOpenBatches(ctx context.Context) ([]*v1.BatchInfo, error)
	GetBatchById(ctx context.Context, id uint) (*v1.BatchInfo, error)
	UpdateBatch(ctx context.Context, requesterUniId string, id uint, req *v1.UpdateBatchRequest) error
	CloseBatch(ctx context.Context, requesterUniId string, id uint) error
	DeleteBatch(ctx context.Context, requesterUniId string, id uint) error
	Register(ctx context.Context, requesterUniId string, batchId uint) error
	ListBatchRegisters(ctx context.Context, batchId uint) ([]*v1.StudentInfo, error)
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
			return v1.ErrUserNotExists
		}
		return v1.ErrInternalServerError
	}
	// TODO: currently only admin can create a batch
	// This will be updated when roles are implemented
	// Like convenor can create a batch
	if !requester.IsAdmin {
		return v1.ErrForbiddenAction
	}

	existingBatch, err := s.batchRepo.GetByName(ctx, req.Name)
	if err != nil && !errors.Is(err, v1.ErrBatchNotFound) {
		return err
	}
	if existingBatch != nil {
		return v1.ErrBatchAlreadyExists
	}

	batch := &model.Batch{
		Name:    req.Name,
		Quota:   req.Quota,
		MinCGPA: req.MinCGPARequired,
		MinCH:   req.MinCHRequired,

		TeamRegDeadline: req.TeamRegDeadline,
		MaxTeamMember:   req.MaxTeamMember,
		MaxTeacherPref:  req.MaxTeacherPref,
		PreDefenceAt:    req.PreDefenceAt,
		DefenceAt:       req.DefenceAt,
		CreatedBy:       requester,
	}

	if err = batch.VerifyBeforeUpsert(); err != nil {
		return err
	}

	err = s.batchRepo.Create(ctx, batch)
	if err != nil {
		return err
	}

	return nil
}

func (s *batchService) CloseBatch(ctx context.Context, requesterUniId string, batchId uint) error {
	requester, err := s.userRepo.GetByUniversityId(ctx, requesterUniId)
	if err != nil {
		if errors.Is(err, v1.ErrNotFound) {
			return v1.ErrUserNotExists
		}
		return v1.ErrInternalServerError
	}

	// TODO: currently only admin can close a batch
	// This will be updated when roles are implemented
	// Like convenor can close a batch
	if !requester.IsAdmin {
		return v1.ErrForbiddenAction
	}

	batch, err := s.batchRepo.GetById(ctx, batchId)
	if err != nil {
		return err
	}
	if batch.GetCurrentStage() < model.StageResult {
		return v1.ErrBatchCanNotBeClosed
	}

	batch.Closed = true
	if err = s.batchRepo.Update(ctx, batch); err != nil {
		return err
	}

	return nil
}

func (s *batchService) Register(ctx context.Context, requesterUniId string, batchId uint) error {
	requester, err := s.userRepo.GetByUniversityId(ctx, requesterUniId)
	if err != nil {
		if errors.Is(err, v1.ErrNotFound) {
			return v1.ErrUnauthorized
		}
		return v1.ErrInternalServerError
	}
	if requester.Role != model.RoleStudent {
		return v1.ErrForbiddenAction
	}

	batch, err := s.batchRepo.GetById(ctx, batchId)
	if err != nil {
		return err
	}
	if requester.Student.CGPA < batch.MinCGPA || requester.Student.CompletedCredits < batch.MinCH {
		return v1.ErrStudentNotEligible
	}

	if batch.GetCurrentStage() > model.StageTeamRegistration {
		return v1.ErrTeamRegDeadlinePassed
	}

	err = s.batchRepo.RegisterStudent(ctx, batchId, requester.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *batchService) ListBatchRegisters(ctx context.Context, batchId uint) ([]*v1.StudentInfo, error) {
	students, err := s.batchRepo.GetRegisteredStudents(ctx, batchId)
	if err != nil {
		return nil, err
	}
	studentsApiFormat := make([]*v1.StudentInfo, len(students))
	for i, student := range students {
		studentsApiFormat[i] = student.ConvertToMinimalApiFormat()
	}

	return studentsApiFormat, nil
}

func (s *batchService) ListAllBatches(ctx context.Context) ([]*v1.BatchInfo, error) {
	batches, err := s.batchRepo.ListAllBatches(ctx)
	if err != nil {
		return nil, err
	}
	batchesApiFormat := make([]*v1.BatchInfo, len(batches))
	for i, batch := range batches {
		batchesApiFormat[i] = batch.ConvertToApiFormat()
	}

	return batchesApiFormat, nil
}

func (s *batchService) ListOpenBatches(ctx context.Context) ([]*v1.BatchInfo, error) {
	batches, err := s.batchRepo.ListOpenBatches(ctx)
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
			return v1.ErrUserNotExists
		}
		return v1.ErrInternalServerError
	}
	// TODO: currently only admin can update a batch
	// This will be updated when roles are implemented
	// Like convenor can update a batch
	if !requester.IsAdmin {
		return v1.ErrForbiddenAction
	}

	batch, err := s.batchRepo.GetById(ctx, id)
	if err != nil {
		return err
	}

	utils.SetIfNonDefault(&req.Name, &batch.Name)
	utils.SetIfNonDefault(&req.Quota, &batch.Quota)
	utils.SetIfNonDefault(&req.MinCHRequired, &batch.MinCH)
	utils.SetIfNonDefault(&req.MinCGPARequired, &batch.MinCGPA)
	utils.SetIfNonDefault(&req.MaxTeamMember, &batch.MaxTeamMember)
	utils.SetIfNonDefault(&req.MaxTeacherPref, &batch.MaxTeacherPref)
	utils.SetIfNonDefault(req.TeamRegDeadline, &batch.TeamRegDeadline)
	utils.SetIfNonDefault(req.PreDefenceAt, &batch.PreDefenceAt)
	utils.SetIfNonDefault(req.DefenceAt, &batch.DefenceAt)

	if err = batch.VerifyBeforeUpsert(); err != nil {
		return err
	}

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
			return v1.ErrUserNotExists
		}
		return v1.ErrInternalServerError
	}
	// TODO: currently only admin can update a batch
	// This will be updated when roles are implemented
	// Like convenor can update a batch
	if !requester.IsAdmin {
		return v1.ErrForbiddenAction
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
