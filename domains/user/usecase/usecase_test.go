package usecase

import (
	"context"
	"database/sql"
	"reflect"
	"testing"
	"time"

	"github.com/KotaroYamazaki/go-clean-arch-sample/domains/user"
	"github.com/KotaroYamazaki/go-clean-arch-sample/domains/user/repository"
	"github.com/KotaroYamazaki/go-clean-arch-sample/domains/user/repository/mock_repository"
	"github.com/KotaroYamazaki/go-clean-arch-sample/models"

	"github.com/golang/mock/gomock"
)

var (
	bd        = time.Now().AddDate(-20, 0, 0)
	dummyUser = &models.User{
		ID:          1,
		FirebaseUID: "2",
		Name:        "3",
		Birthday:    bd,
		CreatedAt:   time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt:   time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
	}
)

func TestNew(t *testing.T) {
	type args struct {
		repo repository.Repository
	}
	tests := []struct {
		name string
		args args
		want Usecase
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_usecase_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	mockRepository := mock_repository.NewMockRepository(ctrl)
	mockRepository.EXPECT().Get(ctx, 1).Return(dummyUser, nil)
	mockRepository.EXPECT().Get(ctx, 2).Return(nil, sql.ErrNoRows)

	uc := New(mockRepository)

	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name    string
		args    args
		want    *user.User
		wantErr bool
	}{
		{
			name: "ユーザーID=1が渡された場合はID=1のユーザーエンティティを返す",
			args: args{
				ctx: ctx,
				id:  1,
			},
			want:    &user.User{User: dummyUser, Age: 20},
			wantErr: false,
		},
		{
			name: "`存在しないユーザーIDが渡された場合はエラーを返す`",
			args: args{
				ctx: ctx,
				id:  2,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := uc.Get(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("usecase.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_usecase_Signup(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	mockRepository := mock_repository.NewMockRepository(ctrl)
	mockRepository.EXPECT().Store(ctx, &models.User{FirebaseUID: "1", Name: "2", Birthday: bd}).Return(nil)

	uc := New(mockRepository)

	type args struct {
		ctx    context.Context
		params *user.SignupParams
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "適切なパラメータが渡された場合は正常に処理される",
			args: args{
				ctx: ctx,
				params: &user.SignupParams{
					FirebaseUID: "1",
					Name:        "2",
					Birthday:    &bd,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := uc.Signup(tt.args.ctx, tt.args.params); (err != nil) != tt.wantErr {
				t.Errorf("usecase.Signup() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
