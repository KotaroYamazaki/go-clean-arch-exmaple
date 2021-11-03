package repository

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/KotaroYamazaki/go-clean-arch-example/models"
	"github.com/stretchr/testify/require"
)

const (
	getQuery   = "select * from `users` where `id`=?"
	storeQuery = "INSERT INTO `users` (`firebase_uid`,`name`,`birthday`,`created_at`,`updated_at`) VALUES (?,?,?,?,?)"
)

func Test_repository_Get(t *testing.T) {

	db, mock, err := sqlmock.New(
		sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual),
	)
	require.NoError(t, err)
	defer db.Close()

	repo := New(db)

	t.Run("ID=1が渡された場合はID=1のユーザーが返る", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "firebase_uid", "name"}).
			AddRow(1, "hoge", "fuga")

		u := &models.User{
			ID:          1,
			FirebaseUID: "hoge",
			Name:        "fuga",
		}
		mock.ExpectQuery(getQuery).WithArgs(1).WillReturnRows(rows)

		getU, err := repo.Get(context.Background(), 1)
		require.NoError(t, err)
		require.Equal(t, u, getU)
	})

	t.Run("存在しないIDが渡された時、エラーが返る", func(t *testing.T) {
		getErr := sql.ErrNoRows
		mock.ExpectQuery(getQuery).WithArgs(2).WillReturnError(getErr)

		getU, err := repo.Get(context.Background(), 2)
		require.Nil(t, getU)
		require.NotNil(t, err)
	})

}

func Test_repository_Store(t *testing.T) {
	db, mock, err := sqlmock.New(
		sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual),
	)
	require.NoError(t, err)
	defer db.Close()

	repo := New(db)

	t.Run("正しい*models.Userが渡された場合は正常にStoreされる", func(t *testing.T) {
		u := &models.User{
			FirebaseUID: "hoge",
			Name:        "fuga",
			Birthday:    time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}
		mock.ExpectExec(storeQuery).
			WithArgs(u.FirebaseUID, u.Name, u.Birthday, u.CreatedAt, u.UpdatedAt).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := repo.Store(context.Background(), u)
		require.NoError(t, err)
	})
}
