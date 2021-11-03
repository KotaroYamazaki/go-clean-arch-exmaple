package utils

import (
	"time"
)

func GetAge(bd time.Time) int {
	now := time.Now()
	age := now.Year() - bd.Year()
	if now.YearDay() < bd.YearDay() {
		age--
	}
	return age
}
