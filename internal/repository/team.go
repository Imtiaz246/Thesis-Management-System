package repository

type TeamRepository interface {
}

func NewTeamRepository(r *Repository) TeamRepository {
	return &teamRepository{
		Repository: r,
	}
}

type teamRepository struct {
	*Repository
}
