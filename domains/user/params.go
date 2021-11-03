package user

import (
	"time"
)

type SignupParams struct {
	FirebaseUID string     `json:"firebase_uid"`
	Name        string     `json:"name" binding:"required"`
	Birthday    *time.Time `json:"birthday" binding:"required"`
}

func (p *SignupParams) Validate() error {
	if p.Name == "" {
		return ErrNameRequired
	}
	if p.Birthday == nil {
		return ErrBirthdayRequired
	}
	return nil
}
