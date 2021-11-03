package utils

import (
	"time"
)

func GetAge(bd time.Time) int {
	now := time.Now().In(time.UTC)
	age := now.Year() - bd.Year()
	if now.YearDay() < bd.YearDay() {
		age--
	}
	return age
}
