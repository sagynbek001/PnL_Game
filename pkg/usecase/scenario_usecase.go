package usecase

import (
	"context"

	"gitlab.com/laboct2021/pnl-game/internal/models"
	"gitlab.com/laboct2021/pnl-game/pkg/logwrapper"
)

type ScenarioService struct {
	logger logwrapper.LogWrapper
	repo   models.ScenarioRepo
}

func NewScenarioService(log logwrapper.LogWrapper, r models.ScenarioRepo) ScenarioService {
	return ScenarioService{
		logger: log,
		repo:   r,
	}
}

func (u ScenarioService) Create(ctx context.Context, Scenario *models.Scenario) error {
	err := u.repo.Create(ctx, Scenario)
	return err
}

func (u ScenarioService) GetAll(ctx context.Context) ([]models.Scenario, error) {
	getAll, err := u.repo.GetAll(ctx)
	return getAll, err
}

func (u ScenarioService) GetByID(ctx context.Context, id uint64) (models.Scenario, error) {
	getByID, err := u.repo.GetByID(ctx, id)
	return getByID, err
}

func (u ScenarioService) Update(ctx context.Context, Scenarios *models.Scenario) error {
	err := u.repo.Update(ctx, Scenarios)
	return err
}

func (u ScenarioService) Delete(ctx context.Context, id uint64) error {
	err := u.repo.Delete(ctx, id)
	return err
}
