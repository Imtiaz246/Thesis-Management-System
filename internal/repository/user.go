package repository

import (
	"context"
	"encoding/json"
	"errors"
	apisv1 "github.com/Imtiaz246/Thesis-Management-System/internal/apis/v1"
	"github.com/Imtiaz246/Thesis-Management-System/internal/model"
	"gorm.io/gorm"
	"time"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
	GetByUniversityId(ctx context.Context, universityId string) (*model.User, error)
	CheckUserExistence(ctx context.Context, universityId string) (bool, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	ReqRegisterCache(ctx context.Context, token string, studentInfo *apisv1.StudentInfo) error
	ReqRegisterCacheGet(ctx context.Context, token string) (*apisv1.StudentInfo, error)
	ReqRegisterCacheClear(ctx context.Context, token string) error
}

func NewUserRepository(r *Repository) UserRepository {
	return &userRepository{
		Repository: r,
	}
}

type userRepository struct {
	*Repository
}

func (r *userRepository) Create(ctx context.Context, user *model.User) error {
	if err := r.DB(ctx).Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) Update(ctx context.Context, user *model.User) error {
	if err := r.DB(ctx).Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) GetByUniversityId(ctx context.Context, universityId string) (*model.User, error) {
	var user model.User
	if err := r.DB(ctx).Where("university_id = ?", universityId).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apisv1.ErrNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) CheckUserExistence(ctx context.Context, universityId string) (bool, error) {
	user, err := r.GetByUniversityId(ctx, universityId)
	if err != nil && !errors.Is(err, apisv1.ErrNotFound) {
		return false, err
	}
	if user != nil {
		return true, nil
	}

	return false, nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	if err := r.DB(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) ReqRegisterCache(ctx context.Context, token string, studentInfo *apisv1.StudentInfo) error {
	data, err := json.Marshal(studentInfo)
	if err != nil {
		return err
	}
	err = r.rdb.Set(ctx, token, data, time.Minute*20).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) ReqRegisterCacheGet(ctx context.Context, token string) (*apisv1.StudentInfo, error) {
	data, err := r.rdb.Get(ctx, token).Bytes()
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, apisv1.ErrNotFound
	}
	studentInfo := new(apisv1.StudentInfo)
	if err = json.Unmarshal(data, studentInfo); err != nil {
		return nil, err
	}

	return studentInfo, nil
}

func (r *userRepository) ReqRegisterCacheClear(ctx context.Context, token string) error {
	return r.rdb.Del(ctx, token).Err()
}
