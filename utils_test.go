package dkssud

import (
	"testing"
)

// TestContains checks the functionality of the contains function.
func TestContains(t *testing.T) {
	tests := []struct {
		slice []string
		item  string
		want  bool
	}{
		{[]string{"a", "b", "c"}, "b", true},
		{[]string{"a", "b", "c"}, "d", false},
		{[]string{}, "a", false},
		{[]string{"a", "A"}, "A", true},
	}

	for _, tt := range tests {
		t.Run(tt.item, func(t *testing.T) {
			got := contains(tt.slice, tt.item)
			if got != tt.want {
				t.Errorf("contains(%v, %q) = %v; want %v", tt.slice, tt.item, got, tt.want)
			}
		})
	}
}

// TestIndexInSlice checks the functionality of the indexInSlice function.
func TestIndexInSlice(t *testing.T) {
	tests := []struct {
		slice []string
		item  string
		want  int
	}{
		{[]string{"a", "b", "c"}, "b", 1},
		{[]string{"a", "b", "c"}, "d", -1},
		{[]string{"a", "A", "a"}, "a", 0},
		{[]string{"a", "b", "c"}, "c", 2},
	}

	for _, tt := range tests {
		t.Run(tt.item, func(t *testing.T) {
			got := indexInSlice(tt.slice, tt.item)
			if got != tt.want {
				t.Errorf("indexInSlice(%v, %q) = %v; want %v", tt.slice, tt.item, got, tt.want)
			}
		})
	}
}

// TestIsAttachAvailable checks the functionality of the isAttachAvailable function.
func TestIsAttachAvailable(t *testing.T) {
	tests := []struct {
		i, l byte
		want int
	}{
		{'r', 'k', 2}, // 자 + 모
		{'k', 'o', 0}, // 모 + 모
		{'k', 'r', 4}, // 모 + 자
		{'r', 'r', 0}, // 자 + 자 (not attachable)
		{'R', 'R', 0}, // 자 + 자 (double consonant)
		{'k', 'z', 4}, // 모 + other (not attachable)
	}

	for _, tt := range tests {
		t.Run(string([]byte{tt.i, tt.l}), func(t *testing.T) {
			got := isAttachAvailable(tt.i, tt.l)
			if got != tt.want {
				t.Errorf("isAttachAvailable(%q, %q) = %v; want %v", tt.i, tt.l, got, tt.want)
			}
		})
	}
}
