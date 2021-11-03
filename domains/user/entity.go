package user

import (
	"github.com/KotaroYamazaki/go-cleanarchtecture/models"
)

type User struct {
	*models.User
	Age int
}
