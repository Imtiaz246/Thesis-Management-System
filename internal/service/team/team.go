package team

import (
	"context"
	"errors"
	v1 "github.com/Imtiaz246/Thesis-Management-System/internal/apis/v1"
	"github.com/Imtiaz246/Thesis-Management-System/internal/model"
	"github.com/Imtiaz246/Thesis-Management-System/internal/repository"
	"github.com/Imtiaz246/Thesis-Management-System/internal/service"
)

type Service interface {
	CreateTeam(ctx context.Context, batchId uint, req *v1.CreateTeamRequest, requesterUniId string) (*v1.TeamInfo, error)
	GetJoinedTeam(ctx context.Context, batchId uint, requesterUniId string) (*v1.TeamInfo, error)
	LeaveTeam(ctx context.Context, teamId uint, requesterUniId string) error
	SendInvitation(ctx context.Context, batchId, teamId uint, requesterUniId, targetUserUniId string) error
	RejectInvitation(ctx context.Context, teamId uint, requesterUniId string) error
	AcceptInvitation(ctx context.Context, batchId, teamId uint, requesterUniId string) error
	ListInvitations(ctx context.Context, teamId uint, requesterUniId string) ([]*v1.TeamInvitationInfo, error)
}

func NewTeamService(r repository.TeamRepository, s *service.Service) Service {
	return &teamService{
		teamRepo: r,
		Service:  s,
	}
}

type teamService struct {
	userRepo  repository.UserRepository
	batchRepo repository.BatchRepository
	teamRepo  repository.TeamRepository
	*service.Service
}

func (s *teamService) CreateTeam(ctx context.Context, batchId uint, req *v1.CreateTeamRequest, requesterUniId string) (*v1.TeamInfo, error) {
	requester, err := s.userRepo.GetByUniversityId(ctx, requesterUniId)
	if err != nil {
		if errors.Is(err, v1.ErrNotFound) {
			return nil, v1.ErrUserNotExists
		}
		return nil, v1.ErrInternalServerError
	}

	isRegisterer, err := s.batchRepo.IsBatchRegisterer(ctx, batchId, requester.ID)
	if err != nil {
		return nil, err
	}
	if !isRegisterer {
		return nil, v1.ErrNotBatchRegisterer
	}

	team := &model.Team{
		Name:    req.Name,
		Subject: req.Subject,
		BatchID: batchId,
	}
	if err = s.teamRepo.CreateTeam(ctx, team); err != nil {
		return nil, err
	}

	return team.ConvertToApiFormat(), nil
}

func (s *teamService) GetJoinedTeam(ctx context.Context, batchId uint, requesterUniId string) (*v1.TeamInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (s *teamService) LeaveTeam(ctx context.Context, teamId uint, requesterUniId string) error {
	requester, err := s.userRepo.GetByUniversityId(ctx, requesterUniId)
	if err != nil {
		if errors.Is(err, v1.ErrNotFound) {
			return v1.ErrUserNotExists
		}
		return v1.ErrInternalServerError
	}

	if err = s.teamRepo.DeleteTeamMember(ctx, teamId, requester.ID); err != nil {
		return err
	}

	return nil
}

func (s *teamService) SendInvitation(ctx context.Context, batchId, teamId uint, requesterUniId, targetStudentUniId string) error {
	requester, err := s.userRepo.GetByUniversityId(ctx, requesterUniId)
	if err != nil {
		if errors.Is(err, v1.ErrNotFound) {
			return v1.ErrUserNotExists
		}
		return v1.ErrInternalServerError
	}
	targetStudent, err := s.userRepo.GetByUniversityId(ctx, targetStudentUniId)
	if err != nil {
		if errors.Is(err, v1.ErrNotFound) {
			return v1.ErrUserNotExists
		}
		return v1.ErrInternalServerError
	}

	isRequesterRegisterer, err := s.batchRepo.IsBatchRegisterer(ctx, batchId, requester.ID)
	if err != nil {
		return err
	}
	if !isRequesterRegisterer {
		return v1.ErrNotBatchRegisterer
	}
	isTargetStudentRegisterer, err := s.batchRepo.IsBatchRegisterer(ctx, batchId, targetStudent.ID)
	if err != nil {
		return err
	}
	if !isTargetStudentRegisterer {
		return v1.ErrTeamInvitationToNonRegisterer
	}

	teamInvitation := &model.TeamInvitation{
		TeamID:    teamId,
		StudentID: targetStudent.ID,
	}
	if err = s.teamRepo.CreateTeamInvitation(ctx, teamInvitation); err != nil {
		return err
	}

	return nil
}

func (s *teamService) RejectInvitation(ctx context.Context, teamId uint, requesterUniId string) error {
	requester, err := s.userRepo.GetByUniversityId(ctx, requesterUniId)
	if err != nil {
		if errors.Is(err, v1.ErrNotFound) {
			return v1.ErrUserNotExists
		}
		return v1.ErrInternalServerError
	}

	if err = s.teamRepo.RemoveTeamInvitation(ctx, teamId, requester.ID); err != nil {
		return err
	}

	return nil
}

func (s *teamService) AcceptInvitation(ctx context.Context, batchId, teamId uint, requesterUniId string) error {
	//TODO implement me
	panic("implement me")
}

func (s *teamService) ListInvitations(ctx context.Context, teamId uint, requesterUniId string) ([]*v1.TeamInvitationInfo, error) {
	requester, err := s.userRepo.GetByUniversityId(ctx, requesterUniId)
	if err != nil {
		if errors.Is(err, v1.ErrNotFound) {
			return nil, v1.ErrUserNotExists
		}
		return nil, v1.ErrInternalServerError
	}

	teamInvitations, err := s.teamRepo.GetTeamInvitations(ctx, teamId, requester.ID)
	if err != nil {
		return nil, err
	}

	result := make([]*v1.TeamInvitationInfo, len(teamInvitations))
	for _, teamInvitation := range teamInvitations {
		result = append(result, &v1.TeamInvitationInfo{
			TeamInfo:  teamInvitation.Team.ConvertToApiFormat(),
			InvitedAt: teamInvitation.CreatedAt,
		})
	}

	return result, nil
}
