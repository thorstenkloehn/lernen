package main

import "testing"

func Test_haus(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := haus(); got != tt.want {
				t.Errorf("haus() = %v, want %v", got, tt.want)
			}
		})
	}
}