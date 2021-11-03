package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/KotaroYamazaki/go-clean-arch-example/domains/user"
	"github.com/KotaroYamazaki/go-clean-arch-example/domains/user/usecase"
	"github.com/KotaroYamazaki/go-clean-arch-example/domains/user/usecase/mock_usecase"
	"github.com/KotaroYamazaki/go-clean-arch-example/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"
)

var (
	bd        = time.Now().AddDate(-20, 0, 0)
	dummyUser = &user.User{
		User: &models.User{
			ID:          1,
			FirebaseUID: "2",
			Name:        "3",
			Birthday:    bd,
			CreatedAt:   time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
			UpdatedAt:   time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		Age: 20,
	}
)

func TestNew(t *testing.T) {
	type args struct {
		uc usecase.Usecase
	}
	tests := []struct {
		name string
		args args
		want *Handler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.uc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandler_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	param := gin.Param{
		Key:   "id",
		Value: "1",
	}
	params := gin.Params{param}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("GET", "/users/1", nil)
	c.Params = params

	mockUC := mock_usecase.NewMockUsecase(ctrl)
	mockUC.EXPECT().Get(c, 1).Return(dummyUser, nil)

	h := New(mockUC)

	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
		wantRes  *user.User
	}{
		{
			name:     "users/1 のパスにGETリクエスト場合はユーザーID=1のエンティティが返る",
			args:     args{c: c},
			wantCode: http.StatusOK,
			wantRes:  dummyUser,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h.Get(tt.args.c)
			assert.Equal(t, tt.wantCode, w.Code)
		})
	}
}

func TestHandler_Signup(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	body := bytes.NewBufferString("{\"firebase_uid\":\"hoge\",\"name\":\"foo\",\"birthday\":\"2014-08-25T00:00:00Z\"}")
	c.Request, _ = http.NewRequest("POST", "/singup", body)
	c.Request.Header.Add("Content-Type", binding.MIMEJSON)

	c1, _ := gin.CreateTestContext(w)
	body1 := bytes.NewBufferString("{\"firebase_uid\":\"hoge\",\"birthday\":\"2014-08-25T00:00:00Z\"}")
	c1.Request, _ = http.NewRequest("POST", "/singup", body1)
	c1.Request.Header.Add("Content-Type", binding.MIMEJSON)

	c2, _ := gin.CreateTestContext(w)
	body2 := bytes.NewBufferString("{\"firebase_uid\":\"duplicated_id\",\"name\":\"foo\",\"birthday\":\"2014-08-25T00:00:00Z\"}")
	c2.Request, _ = http.NewRequest("POST", "/singup", body2)
	c2.Request.Header.Add("Content-Type", binding.MIMEJSON)

	bd := time.Date(2014, 8, 25, 0, 0, 0, 0, time.UTC)
	params := &user.SignupParams{
		FirebaseUID: "hoge",
		Name:        "foo",
		Birthday:    &bd,
	}
	params2 := &user.SignupParams{
		FirebaseUID: "duplicated_id",
		Name:        "foo",
		Birthday:    &bd,
	}
	mockUC := mock_usecase.NewMockUsecase(ctrl)
	mockUC.EXPECT().Signup(c, params).Return(nil)
	mockUC.EXPECT().Signup(c2, params2).Return(errors.Wrap(&mysql.MySQLError{Number: 1062, Message: "Duplicate entry 'duplicated_id' for key 'firebase_uid'"}, "models: unable to insert into"))

	h := New(mockUC)

	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "正常なパラメータを受け取ってPOSTされた場合はステータスコード200を返す",
			args: args{
				c: c,
			},
			want: http.StatusOK,
		},
		{
			name: "不正なパラメータを受け取ってPOSTされた場合はステータスコード400を返す",
			args: args{
				c: c1,
			},
			want: http.StatusBadRequest,
		},
		{
			name: "Sinup関数内でのエラーがあった場合はステータスコード500を返す",
			args: args{
				c: c2,
			},
			want: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h.Signup(tt.args.c)
		})
	}
}
