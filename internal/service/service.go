package service

import (
	"github.com/Imtiaz246/Thesis-Management-System/internal/repository"
	"github.com/Imtiaz246/Thesis-Management-System/pkg/helper/sid"
	"github.com/Imtiaz246/Thesis-Management-System/pkg/log"
	"github.com/Imtiaz246/Thesis-Management-System/pkg/mailer"
	"github.com/Imtiaz246/Thesis-Management-System/pkg/token"
	"github.com/spf13/viper"
)

type Service struct {
	logger *log.Logger
	sid    *sid.Sid
	jwt    *token.JWT
	mlr    *mailer.Mailer
	tn     repository.Transaction
	conf   *viper.Viper
}

func NewService(tm repository.Transaction, logger *log.Logger, sid *sid.Sid, jwt *token.JWT, mlr *mailer.Mailer, conf *viper.Viper) *Service {
	return &Service{
		logger: logger,
		sid:    sid,
		jwt:    jwt,
		tn:     tm,
		mlr:    mlr,
		conf:   conf,
	}
}

func (s *Service) Logger() *log.Logger {
	return s.logger
}

func (s *Service) Sid() *sid.Sid {
	return s.sid
}

func (s *Service) Jwt() *token.JWT {
	return s.jwt
}

func (s *Service) Tn() repository.Transaction {
	return s.tn
}

func (s *Service) Mlr() *mailer.Mailer {
	return s.mlr
}

func (s *Service) Conf() *viper.Viper {
	return s.conf
}
