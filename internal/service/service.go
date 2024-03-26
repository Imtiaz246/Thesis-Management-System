package service

import (
	"github.com/Imtiaz246/Thesis-Management-System/internal/repository"
	"github.com/Imtiaz246/Thesis-Management-System/pkg/helper/sid"
	"github.com/Imtiaz246/Thesis-Management-System/pkg/jwt"
	"github.com/Imtiaz246/Thesis-Management-System/pkg/log"
	"github.com/Imtiaz246/Thesis-Management-System/pkg/mailer"
	"github.com/spf13/viper"
)

type Service struct {
	logger *log.Logger
	sid    *sid.Sid
	jwt    *jwt.JWT
	mlr    *mailer.Mailer
	tm     repository.Transaction
	conf   *viper.Viper
}

func NewService(tm repository.Transaction, logger *log.Logger, sid *sid.Sid, jwt *jwt.JWT, mlr *mailer.Mailer, conf *viper.Viper) *Service {
	return &Service{
		logger: logger,
		sid:    sid,
		jwt:    jwt,
		tm:     tm,
		mlr:    mlr,
		conf:   conf,
	}
}
