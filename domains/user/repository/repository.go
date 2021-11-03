package repository

import (
	"context"
	"database/sql"

	"github.com/KotaroYamazaki/go-cleanarchtecture/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Repository interface {
	Get(ctx context.Context, id int) (*models.User, error)
	Store(ctx context.Context, user *models.User) error
}

type repository struct {
	db *sql.DB
}

func New(
	db *sql.DB,

) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Get(ctx context.Context, id int) (*models.User, error) {
	return nil, nil
}

func (repo *repository) Store(ctx context.Context, user *models.User) error {
	return user.Insert(ctx, repo.db, boil.Infer())
}
