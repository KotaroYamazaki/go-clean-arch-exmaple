package usecase

import (
	"context"
	"log"

	"github.com/KotaroYamazaki/go-clean-arch-example/domains/user"
	"github.com/KotaroYamazaki/go-clean-arch-example/domains/user/repository"
	"github.com/KotaroYamazaki/go-clean-arch-example/models"
	"github.com/KotaroYamazaki/go-clean-arch-example/utils"
)

type Usecase interface {
	Get(ctx context.Context, id int) (*user.User, error)
	Signup(ctx context.Context, params *user.SignupParams) error
}

type usecase struct {
	repo repository.Repository
}

func New(repo repository.Repository) Usecase {
	return &usecase{
		repo: repo,
	}
}

func (uc *usecase) Get(ctx context.Context, id int) (*user.User, error) {
	u, err := uc.repo.Get(ctx, id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &user.User{
		User: u,
		Age:  utils.GetAge(u.Birthday),
	}, nil
}

func (uc *usecase) Signup(ctx context.Context, params *user.SignupParams) error {
	u := &models.User{
		FirebaseUID: params.FirebaseUID,
		Name:        params.Name,
		Birthday:    *params.Birthday,
	}
	if err := uc.repo.Store(ctx, u); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
