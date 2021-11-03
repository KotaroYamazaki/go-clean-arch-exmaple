package user

import (
	"github.com/KotaroYamazaki/go-clean-arch-example/models"
)

type User struct {
	*models.User
	Age int
}
