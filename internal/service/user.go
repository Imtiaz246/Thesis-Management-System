package service

import (
	"bytes"
	"context"
	"fmt"
	v1 "github.com/Imtiaz246/Thesis-Management-System/api/v1"
	"github.com/Imtiaz246/Thesis-Management-System/internal/model"
	"github.com/Imtiaz246/Thesis-Management-System/internal/repository"
	"github.com/Imtiaz246/Thesis-Management-System/internal/templates"
	"github.com/Imtiaz246/Thesis-Management-System/pkg/mailer"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"time"
)

type UserService interface {
	Register(ctx context.Context, req *v1.RegisterRequest, token string) error
	ReqRegister(ctx context.Context, req *v1.ReqRegister) (*v1.StudentInfo, error)
	Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginResponseData, error)
	GetProfile(ctx context.Context, userId string) (*v1.GetProfileResponseData, error)
	UpdateProfile(ctx context.Context, userId string, req *v1.UpdateProfileRequest) error
	VerifyEmail(ctx context.Context, token string) (*v1.StudentInfo, error)
}

func NewUserService(service *Service, userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
		Service:  service,
	}
}

type userService struct {
	userRepo repository.UserRepository
	*Service
}

func (s *userService) ReqRegister(ctx context.Context, req *v1.ReqRegister) (*v1.StudentInfo, error) {
	found, err := s.userRepo.CheckUserExistence(ctx, req.UniversityId)
	if err != nil {
		return nil, err
	}
	if found {
		return nil, v1.ErrUserAlreadyExists
	}

	studentInfo, err := s.mockStudentInfoApi(req.UniversityId)
	if err != nil {
		return nil, err
	}
	token, err := s.sid.GenRandomToken(20)
	if err != nil {
		return nil, err
	}

	err = s.userRepo.ReqRegisterCache(context.TODO(), token, studentInfo)
	if err != nil {
		return nil, err
	}
	expAt := time.Now().Add(time.Minute * 20)
	htd, err := s.renderEmailVerifyTemplate(studentInfo.UniversityId, token, expAt)
	if err != nil {
		return nil, err
	}
	err = s.mlr.SendMail(mailer.WithReceiver(studentInfo.Email),
		mailer.WithHTMLTemplate(htd), mailer.WithSubject("Email verification"))
	if err != nil {
		return nil, err
	}

	return studentInfo, nil
}

func (s *userService) VerifyEmail(ctx context.Context, token string) (*v1.StudentInfo, error) {
	return s.userRepo.ReqRegisterCacheGet(ctx, token)
}

func (s *userService) Register(ctx context.Context, req *v1.RegisterRequest, token string) error {
	studentInfo, err := s.userRepo.ReqRegisterCacheGet(ctx, token)
	if err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &model.User{
		UniversityId: studentInfo.UniversityId,
		Email:        studentInfo.Email,
		Password:     string(hashedPassword),
		IsAdmin:      false,
		Role:         model.RoleStudent,
		IsVerified:   true,
		ChangePass:   false,
		Student: &model.Student{
			Name:            req.Name,
			Mobile:          req.Mobile,
			AlternateMobile: req.AlternateMobile,
			Section:         req.Section,

			Department: studentInfo.Department,
			CGPA:       studentInfo.CGPA,
			Country:    studentInfo.Country,
			Batch:      studentInfo.Batch,
		},
	}

	err = s.tm.Transaction(ctx, func(ctx context.Context) error {
		if err = s.userRepo.Create(ctx, user); err != nil {
			return err
		}
		if err = s.userRepo.ReqRegisterCacheClear(ctx, token); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *userService) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginResponseData, error) {
	user, err := s.userRepo.GetByUniversityId(ctx, req.UniversityId)
	if err != nil || user == nil {
		return nil, v1.ErrUnauthorized
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, err
	}
	if !user.IsVerified {
		return nil, v1.ErrEmailNotVerified
	}
	accessToken, err := s.jwt.GenToken(user.UniversityId, time.Now().Add(time.Minute*15))
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.jwt.GenToken(user.UniversityId, time.Now().Add(time.Hour*24*15))
	if err != nil {
		return nil, err
	}

	return &v1.LoginResponseData{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *userService) GetProfile(ctx context.Context, userId string) (*v1.GetProfileResponseData, error) {
	//user, err := s.userRepo.GetByID(ctx, userId)
	//if err != nil {
	//	return nil, err
	//}
	//
	//return &v1.GetProfileResponseData{
	//	UserId:   user.UserId,
	//	Nickname: user.Nickname,
	//}, nil
	return nil, nil
}

func (s *userService) UpdateProfile(ctx context.Context, userId string, req *v1.UpdateProfileRequest) error {
	//user, err := s.userRepo.GetByID(ctx, userId)
	//if err != nil {
	//	return err
	//}
	//
	//user.Email = req.Email
	//user.Nickname = req.Nickname
	//
	//if err = s.userRepo.Update(ctx, user); err != nil {
	//	return err
	//}

	return nil
}

func (s *userService) renderEmailVerifyTemplate(universityId, token string, expirationTime time.Time) ([]byte, error) {
	tmpl, err := template.New("email_verification").Parse(templates.EmailVerifyTmpl)
	if err != nil {
		return nil, err
	}
	data := map[string]interface{}{
		"Link":    fmt.Sprintf("%s?universityId=%s,token=%s", s.conf.GetString("urls.front_end.reg_confirmation_path"), universityId, token),
		"ExpTime": expirationTime.Format(time.DateTime),
	}
	var out bytes.Buffer
	if err = tmpl.Execute(&out, data); err != nil {
		return nil, err
	}

	return out.Bytes(), nil
}

// todo: replace that with actual api call
func (s *userService) mockStudentInfoApi(universityId string) (*v1.StudentInfo, error) {
	return &v1.StudentInfo{
		UniversityId:     universityId,
		Email:            fmt.Sprintf("%s@ugrad.iiuc.ac.bd", "c183074"),
		Department:       "CSE",
		CGPA:             3.50,
		Batch:            47,
		CompletedCredits: 160,
		Nationality:      "Bangladeshi",
		Mobile:           "01773820154",
		AlternateMobile:  "01773820154",
	}, nil
}
