package wal

import (
	"simplewal/internal/wal"
	"testing"
)

// TestAdd verifies the Add function for various input cases.
func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{name: "positive numbers", a: 2, b: 3, expected: 5},
		{name: "negative numbers", a: -1, b: -4, expected: -5},
		{name: "zero and positive", a: 0, b: 7, expected: 7},
		{name: "zero and zero", a: 0, b: 0, expected: 0},
		{name: "positive and negative", a: 10, b: -3, expected: 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := wal.Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%d, %d): got %d, want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}
