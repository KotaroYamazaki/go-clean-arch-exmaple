package user

import (
	"github.com/KotaroYamazaki/go-clean-arch-sample/models"
)

type User struct {
	*models.User
	Age int
}
