package dkssud

import (
	"testing"
)

// TestQwertyToHangul tests the QwertyToHangul function.
func TestQwertyToHangul(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"rkskekfk", "가나다라"},           // Simple conversion
		{"anlfrRkfrRkR", "뮑깕깎"},        // Mixed case, double consonants
		{"gk zzzz apfhd", "하 ㅋㅋㅋㅋ 메롱"}, // Complex sentence
		{"", ""},                  // Empty string
		{"rjRlRkwldii", "거끼까지야ㅑ"}, // Combination with uppercase and lowercase
		{"abcd", "뮻ㅇ"},            // Simple consonants
		{"123", "123"},            // Numbers should remain unchanged
		{"tla", "심"},              // Testing with complex consonant combinations
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := QwertyToHangul(tt.input)
			if got != tt.want {
				t.Errorf("QwertyToHangul(%q) = %q; want %q", tt.input, got, tt.want)
			}
		})
	}
}

// Test쿼티 tests the QwertyToHangul function.
func Test쿼티(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"rkskekfk", "가나다라"},           // Simple conversion
		{"anlfrRkfrRkR", "뮑깕깎"},        // Mixed case, double consonants
		{"gk zzzz apfhd", "하 ㅋㅋㅋㅋ 메롱"}, // Complex sentence
		{"", ""},                  // Empty string
		{"rjRlRkwldii", "거끼까지야ㅑ"}, // Combination with uppercase and lowercase
		{"abcd", "뮻ㅇ"},            // Simple consonants
		{"123", "123"},            // Numbers should remain unchanged
		{"tla", "심"},              // Testing with complex consonant combinations
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := 쿼티(tt.input)
			if got != tt.want {
				t.Errorf("QwertyToHangul(%q) = %q; want %q", tt.input, got, tt.want)
			}
		})
	}
}

// TestHangulToQwerty tests the HangulToQwerty function.
func TestHangulToQwerty(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"가나다라", "rkskekfk"},           // Simple conversion
		{"뮑깕깎", "anlfrRkfrRkR"},        // Mixed case, double consonants
		{"하 ㅋㅋㅋㅋ 메롱", "gk zzzz apfhd"}, // Complex sentence
		{"", ""},                  // Empty string
		{"거끼까지야ㅑ", "rjRlRkwldii"}, // Combination with uppercase and lowercase
		{"뮻ㅇ", "abcd"},            // Simple consonants
		{"123", "123"},            // Numbers should remain unchanged
		{"심", "tla"},              // Testing with complex consonant combinations
		{"안녕", "dkssud"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := HangulToQwerty(tt.input)
			if got != tt.want {
				t.Errorf("HangulToQwerty(%q) = %q; want %q", tt.input, got, tt.want)
			}
		})
	}
}

// Test한글 tests the HangulToQwerty function.
func Test한글(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"가나다라", "rkskekfk"},           // Simple conversion
		{"뮑깕깎", "anlfrRkfrRkR"},        // Mixed case, double consonants
		{"하 ㅋㅋㅋㅋ 메롱", "gk zzzz apfhd"}, // Complex sentence
		{"", ""},                  // Empty string
		{"거끼까지야ㅑ", "rjRlRkwldii"}, // Combination with uppercase and lowercase
		{"뮻ㅇ", "abcd"},            // Simple consonants
		{"123", "123"},            // Numbers should remain unchanged
		{"심", "tla"},              // Testing with complex consonant combinations
		{"안녕", "dkssud"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := 한글(tt.input)
			if got != tt.want {
				t.Errorf("HangulToQwerty(%q) = %q; want %q", tt.input, got, tt.want)
			}
		})
	}
}

func BenchmarkQwertyToHangul(b *testing.B) {
	input := "rkSkEkfkQkTkdkzk" // Example Qwerty input that maps to Hangul
	for i := 0; i < b.N; i++ {
		QwertyToHangul(input)
	}
}

func BenchmarkHangulToQwerty(b *testing.B) {
	input := "경기 수원시 123 한강 레디 카페" // Example Hangul input that maps to Qwerty
	for i := 0; i < b.N; i++ {
		HangulToQwerty(input)
	}
}

func BenchmarkQwertyToHangulLongInput(b *testing.B) {
	input := "rkSkEkfkQkTkdkzkrkSkEkfkQkTkdkzk" // Longer Qwerty input for testing
	for i := 0; i < b.N; i++ {
		QwertyToHangul(input)
	}
}

func BenchmarkHangulToQwertyLongInput(b *testing.B) {
	input := "가나다라마바사아자차카타파하가나다라마바사아자차카타파하" // Longer Hangul input for testing
	for i := 0; i < b.N; i++ {
		HangulToQwerty(input)
	}
}
