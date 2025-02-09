package team

import (
	"github.com/Imtiaz246/Thesis-Management-System/internal/repository"
	"github.com/Imtiaz246/Thesis-Management-System/internal/service"
)

type Service interface {
}

func NewTeamService(r repository.TeamRepository, s *service.Service) Service {
	return &teamService{
		teamRepo: r,
		Service:  s,
	}
}

type teamService struct {
	teamRepo repository.TeamRepository
	*service.Service
}
