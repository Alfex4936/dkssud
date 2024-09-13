package dkssud

import (
	"fmt"
	"testing"
)

// TestContainsByte checks the functionality of the containsByte function.
func TestContainsByte(t *testing.T) {
	tests := []struct {
		set  map[byte]struct{}
		item byte
		want bool
	}{
		// Test case where the item is present in the set
		{map[byte]struct{}{'a': {}, 'b': {}, 'c': {}}, 'b', true},
		// Test case where the item is not present in the set
		{map[byte]struct{}{'a': {}, 'b': {}, 'c': {}}, 'd', false},
		// Test case with an empty set
		{map[byte]struct{}{}, 'a', false},
		// Test case with similar characters to check case sensitivity
		{map[byte]struct{}{'a': {}, 'A': {}}, 'A', true},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Item: %q", tt.item), func(t *testing.T) {
			got := containsByte(tt.set, tt.item)
			if got != tt.want {
				t.Errorf("containsByte(%v, %q) = %v; want %v", tt.set, tt.item, got, tt.want)
			}
		})
	}
}

// TestContainsString checks the functionality of the containsString function.
func TestContainsString(t *testing.T) {
	tests := []struct {
		set  map[string]struct{}
		item string
		want bool
	}{
		// Test case where the item is present in the set
		{map[string]struct{}{"a": {}, "b": {}, "c": {}}, "b", true},
		// Test case where the item is not present in the set
		{map[string]struct{}{"a": {}, "b": {}, "c": {}}, "d", false},
		// Test case with an empty set
		{map[string]struct{}{}, "a", false},
		// Test case with similar strings to check case sensitivity
		{map[string]struct{}{"a": {}, "A": {}}, "A", true},
		// Test case with multi-character strings
		{map[string]struct{}{"hk": {}, "ho": {}, "hl": {}}, "ho", true},
		{map[string]struct{}{"hk": {}, "ho": {}, "hl": {}}, "ha", false},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Item: %q", tt.item), func(t *testing.T) {
			got := containsString(tt.set, tt.item)
			if got != tt.want {
				t.Errorf("containsString(%v, %q) = %v; want %v", tt.set, tt.item, got, tt.want)
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
			got := IsAttachAvailable(tt.i, tt.l)
			if got != tt.want {
				t.Errorf("IsAttachAvailable(%q, %q) = %v; want %v", tt.i, tt.l, got, tt.want)
			}
		})
	}
}

func TestIsQwertyHangul(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		// Test cases covering all possible scenarios
		{"경기도", false},                       // Hangul characters
		{"rudrleh", true},                    // QWERTY Hangul
		{"123", false},                       // Only digits
		{"1경기", false},                       // Mixed digits and Hangul
		{"123abbabb", true},                  // Mixed digits and letters (QWERTY Hangul)
		{"가낟ㄱㄴㅁㄱ닥ㄴㄷ", false},                 // Hangul and Hangul Jamo
		{"", false},                          // Empty string
		{"!@#", false},                       // Special characters
		{"qwerty", true},                     // All letters, valid QWERTY Hangul input
		{"abcd1234", true},                   // Mixed letters and digits, valid QWERTY Hangul input
		{"abcd!@#efg", false},                // Mixed letters and special characters
		{"공감하다", false},                      // Pure Hangul phrase
		{"학문rudrleh", false},                 // Mixed Hangul and QWERTY Hangul
		{"한글123", false},                     // Hangul with numbers
		{"123abc한글", false},                  // Mixed digits, letters, and Hangul
		{"abcd1234efg", true},                // Continuous QWERTY Hangul input
		{"1234567890", false},                // Only numbers
		{"", false},                          // Empty string (edge case)
		{" ", false},                         // Single space (edge case)
		{"abcdefghijklmnopqrstuvwxyz", true}, // Full alphabet
		{"dkssud gktpdy", true},              // Full alphabet
		{"hello there", true},                // Full alphabet
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := IsQwertyHangul(tt.input)
			if result != tt.expected {
				t.Errorf("IsQwertyHangul(%q) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

// sliceToMap converts a slice of strings into a map[string]struct{} for testing.
func sliceToMap(slice []string) map[string]struct{} {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}
	return set
}
