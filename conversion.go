package dkssud

import (
	"strings"
	"unicode"
)

// QwertyToHangul converts a string of QWERTY keyboard inputs to their corresponding Hangul characters.
// It supports the mapping of Korean characters typed using the QWERTY keyboard layout.
//
// Example usage:
//
//	converted := QwertyToHangul("rkskekfk")
//	fmt.Println(converted) // Output: "가나다라"
//
//	converted = QwertyToHangul("gk zzzz apfhd")
//	fmt.Println(converted) // Output: "하 ㅋㅋㅋㅋ 메롱"
//
//	converted = QwertyToHangul("123")
//	fmt.Println(converted) // Output: "123" (Numbers remain unchanged)
//
// Parameters:
//   - input: A string containing QWERTY keyboard inputs.
//
// Returns:
//   - A string containing the corresponding Hangul characters.
func QwertyToHangul(input string) string {
	charGroups := splitEn(input)
	var convertedString strings.Builder

	for _, charGroup := range charGroups {
		if len(charGroup) == 1 {
			if idx, ok := rawMapperMap[charGroup[0]]; ok {
				convertedString.WriteRune(rune(idx + 12593))
			} else {
				convertedString.WriteString(charGroup[0])
			}
			continue
		}

		topIdx, midIdx, botIdx := -1, -1, 0

		for j, char := range charGroup {
			r := rune(char[0])
			if char == " " || unicode.IsDigit(r) || !unicode.IsLetter(r) {
				convertedString.WriteString(char)
				break
			}

			switch j {
			case 0:
				topIdx = koTopEnMap[char] // Use map lookup
			case 1:
				midIdx = koMidEnMap[char]
			case 2:
				botIdx = koBotEnMap[char]
			}
		}

		if topIdx != -1 && midIdx != -1 {
			combinedChar := rune((topIdx*21*28 + midIdx*28 + botIdx) + 44032)
			convertedString.WriteRune(combinedChar)
		}
	}

	return convertedString.String()
}

// HangulToQwerty converts a string of Hangul characters into their corresponding QWERTY keyboard inputs.
// This function reverses the process performed by QwertyToHangul.
//
// Example usage:
//
//	converted := HangulToQwerty("가나다라")
//	fmt.Println(converted) // Output: "rkskekfk"
//
//	converted = HangulToQwerty("하 ㅋㅋㅋㅋ 메롱")
//	fmt.Println(converted) // Output: "gk zzzz apfhd"
//
// Parameters:
//   - input: A string containing Hangul characters.
//
// Returns:
//   - A string representing the corresponding QWERTY keyboard inputs.
func HangulToQwerty(input string) string {
	idxGroups := splitKo(input)
	var convertedString strings.Builder

	for _, idxGroup := range idxGroups {
		for j, idxCapsule := range idxGroup {
			switch v := idxCapsule.(type) {
			case string:
				// Append spaces or non-integer characters directly
				convertedString.WriteString(v)
			case int:
				switch {
				case v >= 12593 && v <= 12643:
					// Handle individual Jamo characters using `rawMapper`
					convertedString.WriteString(rawMapper[v-12593])
				default:
					// Handle composite Hangul syllables
					switch j {
					case 0:
						convertedString.WriteString(koTopEn[v])
					case 1:
						convertedString.WriteString(koMidEn[v])
					case 2:
						convertedString.WriteString(koBotEn[v])
					}
				}
			}
		}
	}

	return convertedString.String()
}
