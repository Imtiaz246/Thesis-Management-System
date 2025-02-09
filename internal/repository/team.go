package repository

import (
	"context"
	"errors"
	"github.com/Imtiaz246/Thesis-Management-System/internal/model"
	"gorm.io/gorm"
)

type TeamRepository interface {
	CreateTeam(ctx context.Context, team *model.Team) error
	UpdateTeam(ctx context.Context, team *model.Team) error
	DeleteTeamIfZeroMember(ctx context.Context, teamID uint) error
	GetTeamByBatchAndTeamID(ctx context.Context, batchID, teamID uint) (*model.Team, error)

	IsTeamMember(ctx context.Context, teamID, studentID uint) (bool, error)
	DeleteTeamMember(ctx context.Context, teamID, studentID uint) error
	CountTeamMember(ctx context.Context, teamID uint) (int, error)

	CreateTeamInvitation(ctx context.Context, teamInvitation *model.TeamInvitation) error
	RemoveTeamInvitation(ctx context.Context, teamID, studentID uint) error
	GetTeamInvitations(ctx context.Context, teamID, studentID uint) ([]*model.TeamInvitation, error)
}

func NewTeamRepository(r *Repository) TeamRepository {
	return &teamRepository{
		Repository: r,
	}
}

type teamRepository struct {
	*Repository
}

func (r *teamRepository) CreateTeam(ctx context.Context, team *model.Team) error {
	if err := r.DB(ctx).Create(team).Error; err != nil {
		return err
	}
	return nil
}

func (r *teamRepository) UpdateTeam(ctx context.Context, team *model.Team) error {
	if err := r.DB(ctx).Save(team).Error; err != nil {
		return err
	}
	return nil
}

func (r *teamRepository) DeleteTeamIfZeroMember(ctx context.Context, teamID uint) error {
	count, err := r.CountTeamMember(ctx, teamID)
	if err != nil {
		return err
	}
	if count > 0 {
		return nil
	}
	if err = r.DB(ctx).Delete(&model.Team{}, teamID).Error; err != nil {
		return err
	}

	return nil
}

func (r *teamRepository) GetTeamByBatchAndTeamID(ctx context.Context, batchID, teamID uint) (*model.Team, error) {
	team := &model.Team{
		Model: gorm.Model{
			ID: teamID,
		},
		BatchID: batchID,
	}
	if err := r.DB(ctx).First(team).Error; err != nil {
		return nil, err
	}

	return team, nil
}

func (r *teamRepository) IsTeamMember(ctx context.Context, teamID uint, studentID uint) (bool, error) {
	tm := &model.TeamMember{
		TeamID:    teamID,
		StudentID: studentID,
	}
	if err := r.DB(ctx).First(tm).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (r *teamRepository) DeleteTeamMember(ctx context.Context, teamID, studentID uint) error {
	teamMember := &model.TeamMember{
		TeamID:    teamID,
		StudentID: studentID,
	}
	if err := r.DB(ctx).Delete(teamMember).Error; err != nil {
		return err
	}

	return nil
}

func (r *teamRepository) CountTeamMember(ctx context.Context, teamID uint) (int, error) {
	var teamMembers []model.TeamMember
	if err := r.DB(ctx).Where("team_id = ?", teamID).Find(&teamMembers).Error; err != nil {
		return 0, err
	}

	return len(teamMembers), nil
}

func (r *teamRepository) CreateTeamInvitation(ctx context.Context, teamInvitation *model.TeamInvitation) error {
	if err := r.DB(ctx).Create(teamInvitation).Error; err != nil {
		return err
	}

	return nil

}

func (r *teamRepository) RemoveTeamInvitation(ctx context.Context, teamID, studentID uint) error {
	teamInvitation := &model.TeamInvitation{
		TeamID:    teamID,
		StudentID: studentID,
	}
	if err := r.DB(ctx).Delete(teamInvitation).Error; err != nil {
		return err
	}

	return nil
}

func (r *teamRepository) GetTeamInvitations(ctx context.Context, teamID, studentID uint) ([]*model.TeamInvitation, error) {
	var teamInvitations []*model.TeamInvitation
	// TODO: Preload team with students info
	if err := r.DB(ctx).Where("team_id = ? AND student_id = ?", teamID, studentID).Preload("Team").Find(&teamInvitations).Error; err != nil {
		return nil, err
	}

	return teamInvitations, nil
}
