package usecase

import (
	"context"

	"gitlab.com/laboct2021/pnl-game/internal/models"
	"gitlab.com/laboct2021/pnl-game/pkg/logwrapper"
)

type UserService struct {
	logger logwrapper.LogWrapper
	repo   models.UserRepo
}

func NewUserService(log logwrapper.LogWrapper, r models.UserRepo) UserService {
	return UserService{
		logger: log,
		repo:   r,
	}
}

func (u UserService) Create(ctx context.Context, user *models.User) error {
	err := u.repo.Create(ctx, user)
	return err
}

func (u UserService) GetAll(ctx context.Context) ([]models.User, error) {
	getAll, err := u.repo.GetAll(ctx)
	return getAll, err
}

func (u UserService) GetByID(ctx context.Context, id uint64) (models.User, error) {
	getByID, err := u.repo.GetByID(ctx, id)
	return getByID, err
}

func (u UserService) Update(ctx context.Context, users *models.User) error {
	err := u.repo.Update(ctx, users)
	return err
}

func (u UserService) Delete(ctx context.Context, id uint64) error {
	err := u.repo.Delete(ctx, id)
	return err
}
func (u UserService) FindByLogin(ctx context.Context, email string) (models.User, error) {
	user, err := u.repo.FindByLogin(ctx, email)
	return user, err
}
