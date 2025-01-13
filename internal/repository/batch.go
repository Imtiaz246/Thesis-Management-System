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
	GetAllWithCreatorInfo(ctx context.Context) ([]*model.Batch, error)
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
	if err := r.DB(ctx).First(&batch, batchID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, v1.ErrBatchNotFound
		}
		return nil, err
	}
	return &batch, nil
}

func (r *batchRepository) GetByName(ctx context.Context, name string) (*model.Batch, error) {
	var batch model.Batch
	if err := r.db.WithContext(ctx).Where("name = ?", name).First(&batch).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, v1.ErrBatchNotFound
		}
		return nil, err
	}
	return &batch, nil
}

func (r *batchRepository) GetAllWithCreatorInfo(ctx context.Context) ([]*model.Batch, error) {
	var batches []*model.Batch
	if err := r.DB(ctx).Preload("CreatedBy").Find(&batches).Error; err != nil {
		return nil, err
	}
	return batches, nil
}

func (r *batchRepository) CheckBatchExistence(ctx context.Context, batchName string) (bool, error) {
	var count int64
	if err := r.DB(ctx).Model(&model.Batch{}).Where("name = ?", batchName).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
