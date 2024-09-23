package dkssud

import (
	"reflect"
	"testing"
)

// TestSplitKo checks the functionality of the splitKo function.
func TestSplitKo(t *testing.T) {
	tests := []struct {
		input string
		want  [][]interface{}
	}{
		// top, mid, final
		{
			"안녕하세요",
			[][]interface{}{
				{11, 0, 4},  // 안
				{2, 6, 21},  // 녕
				{18, 0, 0},  // 하
				{9, 5, 0},   // 세
				{11, 12, 0}, // 요
			},
		},
		{
			"가나다",
			[][]interface{}{
				{0, 0, 0}, // 가
				{2, 0, 0}, // 나
				{3, 0, 0}, // 다
			},
		},
		{
			" ",
			[][]interface{}{
				{" "}, // Space
			},
		},
		{
			"",
			nil, // Empty string
		},
		{
			"가 나",
			[][]interface{}{
				{0, 0, 0}, // 가
				{" "},     // Space
				{2, 0, 0}, // 나
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := splitKo(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitKo(%q) = %v; want %v", tt.input, got, tt.want)
			}
		})
	}
}

// TestSplitEn checks the functionality of the splitEn function.
func TestSplitEn(t *testing.T) {
	tests := []struct {
		input string
		want  [][]string
	}{
		{"rkskekfk", [][]string{{"r", "k"}, {"s", "k"}, {"e", "k"}, {"f", "k"}}}, // 가나다라
		{"Rkdlfk", [][]string{{"R", "k"}, {"d", "l"}, {"f", "k"}}},               // 까이라
		{"Z", [][]string{{"z"}}},
		{"rk zzzz ekfrqnpfr", [][]string{{"r", "k"}, {" "}, {"z"}, {"z"}, {"z"}, {"z"}, {" "}, {"e", "k", "fr"}, {"q", "np", "fr"}}}, // 가 ㅋㅋㅋㅋ 닭뷁
		{"rjRlRkwldii", [][]string{{"r", "j"}, {"R", "l"}, {"R", "k"}, {"w", "l"}, {"d", "i"}, {"i"}}},
		{"rPfydtl", [][]string{{"r", "P"}, {"f", "y", "d"}, {"t", "l"}}},
		{" ", [][]string{{" "}}},
		{"", nil},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := splitEn(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitEn(%q) = %v; want %v", tt.input, got, tt.want)
			}
		})
	}
}

// func slicesEqual(a, b [][]string) bool {
// 	if len(a) != len(b) {
// 		return false
// 	}
// 	for i := range a {
// 		if len(a[i]) != len(b[i]) {
// 			return false
// 		}
// 		for j := range a[i] {
// 			if a[i][j] != b[i][j] {
// 				return false
// 			}
// 		}
// 	}
// 	return true
// }
