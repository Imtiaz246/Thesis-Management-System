package repository

import (
	"context"
	"errors"
	v1 "github.com/Imtiaz246/Thesis-Management-System/internal/apis/v1"
	"github.com/Imtiaz246/Thesis-Management-System/internal/model"

	"gorm.io/gorm"
)

type BatchRepository interface {
	Create(ctx context.Context, batch *model.Batch) error
	Update(ctx context.Context, batch *model.Batch) error
	Delete(ctx context.Context, batchID uint) error
	GetById(ctx context.Context, batchID uint) (*model.Batch, error)
	GetByName(ctx context.Context, name string) (*model.Batch, error)
	RegisterStudent(ctx context.Context, batchID, studentID uint) error
	GetRegisteredStudents(ctx context.Context, batchID uint) ([]*model.Student, error)
	IsBatchRegisterer(ctx context.Context, batchID, studentID uint) (bool, error)
	ListAllBatches(ctx context.Context) ([]*model.Batch, error)
	ListOpenBatches(ctx context.Context) ([]*model.Batch, error)
	CheckBatchExistence(ctx context.Context, batchName string) (bool, error)
}

func NewBatchRepository(r *Repository) BatchRepository {
	return &batchRepository{
		Repository: r,
	}
}

type batchRepository struct {
	*Repository
}

func (r *batchRepository) Create(ctx context.Context, batch *model.Batch) error {
	if err := r.DB(ctx).Create(batch).Error; err != nil {
		return err
	}
	return nil
}

func (r *batchRepository) Update(ctx context.Context, batch *model.Batch) error {
	if err := r.DB(ctx).Save(batch).Error; err != nil {
		return err
	}
	return nil
}

func (r *batchRepository) Delete(ctx context.Context, batchID uint) error {
	if err := r.DB(ctx).Delete(&model.Batch{}, batchID).Error; err != nil {
		return err
	}
	return nil
}

func (r *batchRepository) GetById(ctx context.Context, batchID uint) (*model.Batch, error) {
	var batch model.Batch
	if err := r.DB(ctx).Preload("CreatedBy").Where("id = ?", batchID).First(&batch).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, v1.ErrBatchNotFound
		}
		return nil, err
	}
	return &batch, nil
}

func (r *batchRepository) GetByName(ctx context.Context, name string) (*model.Batch, error) {
	var batch model.Batch
	if err := r.db.WithContext(ctx).Preload("CreatedBy").Where("name = ?", name).First(&batch).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, v1.ErrBatchNotFound
		}
		return nil, err
	}
	return &batch, nil
}

func (r *batchRepository) ListAllBatches(ctx context.Context) ([]*model.Batch, error) {
	var batches []*model.Batch
	if err := r.DB(ctx).Find(&batches).Error; err != nil {
		return nil, err
	}

	return batches, nil
}

func (r *batchRepository) ListOpenBatches(ctx context.Context) ([]*model.Batch, error) {
	var batches []*model.Batch
	if err := r.DB(ctx).Find(&batches, "closed = ?", false).Error; err != nil {
		return nil, err
	}

	return batches, nil
}

func (r *batchRepository) CheckBatchExistence(ctx context.Context, batchName string) (bool, error) {
	_, err := r.GetByName(ctx, batchName)
	if err != nil {
		if errors.Is(err, v1.ErrBatchNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (r *batchRepository) RegisterStudent(ctx context.Context, batchID, studentID uint) error {
	batchRegistration := &model.BatchRegistration{
		BatchID:   batchID,
		StudentID: studentID,
	}

	if err := r.DB(ctx).First(batchRegistration, *batchRegistration).Error; err == nil {
		return v1.ErrAlreadyRegForBatch
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	return r.DB(ctx).Create(batchRegistration).Error
}

func (r *batchRepository) GetRegisteredStudents(ctx context.Context, batchID uint) ([]*model.Student, error) {
	var batchRegistrations []*model.BatchRegistration
	if err := r.DB(ctx).Where("batch_id = ?", batchID).Preload("Student").Find(&batchRegistrations).Error; err != nil {
		return nil, err
	}
	var students []*model.Student
	for _, batchRegistration := range batchRegistrations {
		students = append(students, batchRegistration.Student)
	}

	return students, nil
}

func (r *batchRepository) IsBatchRegisterer(ctx context.Context, batchID, studentID uint) (bool, error) {
	batchRegistration := &model.BatchRegistration{
		BatchID:   batchID,
		StudentID: studentID,
	}
	if err := r.DB(ctx).First(batchRegistration).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
