package finding

import (
	"fmt"
	"testing"
	"time"
)

func Test_toEbayDateTime(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "2021/11/27T00/28/30.123",
			want:  "2021-11-27T00:28:30.123Z",
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			dt, _ := time.Parse("2006/01/02T15/04/05.000", tt.input)
			fmt.Println(dt)
			fmt.Println(toEbayDateTime(dt))
			if got := toEbayDateTime(dt); got != tt.want {
				t.Errorf("toEbayDateTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fromEbayDateTime(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "2021-11-27T00:28:30.123Z",
			want:  "2021/11/27T00/28/30.123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			dt, _ := fromEbayDateTime(tt.input)
			datetime := dt.Format("2006/01/02T15/04/05.000")
			if got := datetime; got != tt.want {
				t.Errorf("fromEbayDateTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
