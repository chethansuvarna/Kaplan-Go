package utils

import (
	"testing"
	"time"
)

func TestGetTimeInFormat(t *testing.T) {
	currentTime := time.Now()
	type args struct {
		format string
		time   time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success",
			args: args{
				format: "MM-DD-YYYY",
				time:   currentTime,
			},
			want: currentTime.Format("01-02-2006"),
		},
		{
			name: "success",
			args: args{
				format: "YYYY-MM-DD",
				time:   currentTime,
			},
			want: currentTime.Format("2006-01-02"),
		},
		{
			name: "success",
			args: args{
				format: "YYYY.MM.DD",
				time:   currentTime,
			},
			want: currentTime.Format("2006.12.31"),
		},
		{
			name: "success",
			args: args{
				format: "YYYY",
				time:   currentTime,
			},
			want: currentTime.Format("2006-01-02 15:04:05"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTimeInFormat(tt.args.format, tt.args.time); got != tt.want {
				t.Errorf("GetTimeInFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}
