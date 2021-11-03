package utils

import (
	"testing"
	"time"
)

func TestGetAge(t *testing.T) {
	now := time.Now()
	twentyYearsAgo := now.AddDate(-20, 0, 0)
	type args struct {
		bd time.Time
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "TestGetAge",
			args: args{
				bd: twentyYearsAgo,
			},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAge(tt.args.bd); got != tt.want {
				t.Errorf("GetAge() = %v, want %v", got, tt.want)
			}
		})
	}
}
