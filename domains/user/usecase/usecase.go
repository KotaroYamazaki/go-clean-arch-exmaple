package usecase

import (
	"context"

	"github.com/KotaroYamazaki/go-cleanarchtecture/domains/user"
	"github.com/KotaroYamazaki/go-cleanarchtecture/domains/user/repository"
	"github.com/KotaroYamazaki/go-cleanarchtecture/models"
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
		return nil, err
	}
	return &user.User{
		User: u,
		Age:  18,
	}, nil
}

func (uc *usecase) Signup(ctx context.Context, params *user.SignupParams) error {
	u := &models.User{
		FirebaseUID: params.FirebaseUID,
		Name:        params.Name,
		Birthday:    *params.Birthday,
	}
	return uc.repo.Store(ctx, u)
}
