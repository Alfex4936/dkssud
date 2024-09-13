package dkssud

import "unicode"

func containsByte(set map[byte]struct{}, item byte) bool {
	_, exists := set[item]
	return exists
}

func containsString(set map[string]struct{}, item string) bool {
	_, exists := set[item]
	return exists
}

// Helper function to find the index of a string in a slice
func indexInSlice(slice []string, item string) int {
	for idx, val := range slice {
		if val == item {
			return idx
		}
	}
	return -1
}

// IsAttachAvailable 함수는 두 개의 문자(i, l)가 한국어 조합 규칙에 따라 붙일 수 있는지를 확인합니다.
// 이 함수는 주로 자음과 모음, 또는 자음과 자음 간의 결합 가능 여부를 판단하는 데 사용됩니다.
//
// 이 함수는 다음과 같은 경우에 특정 값을 반환합니다:
// - 2: 첫 번째 문자가 자음(초성)이고, 두 번째 문자가 모음(중성)인 경우
// - 3: 첫 번째와 두 번째 문자가 모두 모음(중성)이며, 이들이 합쳐져 이중 모음(복모음)을 구성할 수 있는 경우
// - 4: 첫 번째 문자가 모음(중성)이고, 두 번째 문자가 자음(종성)인 경우
// - 5: 두 문자가 결합하여 이중 자음(겹받침)을 구성할 수 있는 경우
// - 0: 위의 경우에 해당하지 않는 경우 (결합 불가)
//
// 예시:
//   result := isAttachAvailable('r', 'k')
//   fmt.Println(result) // 출력: 2 (자 + 모)
//
//   result = isAttachAvailable('k', 'o')
//   fmt.Println(result) // 출력: 0 (모 + 모, 결합 불가)
//
//   result = isAttachAvailable('k', 'r')
//   fmt.Println(result) // 출력: 4 (모 + 자)
//
// Parameters:
// - i: 첫 번째 문자 (byte 형식)
// - l: 두 번째 문자 (byte 형식)
//
// Returns:
// - int: 문자의 결합 가능 여부를 나타내는 값. 0은 결합 불가, 2는 자 + 모, 3은 복모음, 4는 모 + 자, 5는 겹받침을 의미합니다.
func IsAttachAvailable(i, l byte) int {
	// Check for Consonant + Vowel (자 + 모)
	if containsByte(koTopEnSet, i) {
		if containsByte(koMidEnSetSingleChar, l) {
			// Both are single-byte characters
			return 2
		} else if containsString(koMidEnSetMultiChar, string(l)) {
			// 'l' is a multi-character vowel
			return 2
		}
	}

	// Check for Vowel + Vowel (Composite Vowel) (모 + 모)
	if containsByte(koMidEnSetSingleChar, i) {
		if containsByte(koMidEnSetSingleChar, l) {
			combined := string([]byte{i, l})
			if containsString(koMidEnSetMultiChar, combined) {
				return 3
			}
		} else if containsString(koMidEnSetMultiChar, string(l)) {
			combined := string(i) + string(l)
			if containsString(koMidEnSetMultiChar, combined) {
				return 3
			}
		}
	} else if containsString(koMidEnSetMultiChar, string(i)) {
		if containsByte(koMidEnSetSingleChar, l) || containsString(koMidEnSetMultiChar, string(l)) {
			combined := string(i) + string(l)
			if containsString(koMidEnSetMultiChar, combined) {
				return 3
			}
		}
	}

	// Check for Vowel + Consonant (모 + 자)
	if containsByte(koMidEnSetSingleChar, i) || containsString(koMidEnSetMultiChar, string(i)) {
		if containsByte(koBotEnSetSingleChar, l) {
			return 4
		} else if containsString(koBotEnSetMultiChar, string(l)) {
			return 4
		}
	}

	// Check for Double Consonant (자 + 자) in the final position
	if containsByte(koBotEnSetSingleChar, i) || containsString(koBotEnSetMultiChar, string(i)) {
		combined := string([]byte{i, l})
		if containsString(koBotEnSetMultiChar, combined) {
			return 5
		}
	}

	// If none of the above, they are not attachable
	return 0
}

// IsQwertyHangul 함수는 입력된 문자열이 QWERTY 키보드에서 한글을 입력한 것으로 보이는지를 확인합니다.
//
// !주의사항:
// 이 함수는 ASCII 문자로 이루어진 영어 문장도 QWERTY 한글로 잘못 인식할 수 있습니다.
// 예를 들어, "hello there" 같은 영어 문장은 QWERTY 한글로 인식됩니다.
// 이 함수는 주로 QWERTY 한글 입력이 예상되는 상황에서 사용되는 것이 좋습니다.
//
// 이 함수는 다음과 같은 경우에 true를 반환합니다:
// - 문자열이 한글이 아닌 ASCII 문자(영문자 또는 숫자)로만 구성되어 있을 때
// - 문자열이 영문자와 숫자의 조합으로 이루어진 경우
//
// 이 함수는 다음과 같은 경우에 false를 반환합니다:
// - 문자열에 한글 문자가 포함된 경우
// - 문자열에 ASCII 범위를 벗어나는 문자가 포함된 경우 (예: 특수 문자, 다른 언어의 문자 등)
// - 문자열이 오직 숫자로만 이루어진 경우
//
// 예시:
//  fmt.Println(IsQwertyHangul("경기도"))  // false - 한글 포함
//  fmt.Println(IsQwertyHangul("rudrleh")) // true - QWERTY로 입력된 한글
//  fmt.Println(IsQwertyHangul("123"))     // false - 숫자만 포함
//  fmt.Println(IsQwertyHangul("1경기"))   // false - 한글 포함
//  fmt.Println(IsQwertyHangul("123abbabb")) // true - QWERTY로 입력된 한글과 숫자 조합
//  fmt.Println(IsQwertyHangul("가낟ㄱㄴㅁㄱ닥ㄴㄷ")) // false - 한글 및 한글 자모 포함
//
// Parameters:
// - input: 검사할 문자열
//
// Returns:
// - 문자열이 QWERTY 키보드로 입력된 한글일 가능성이 있으면 true, 그렇지 않으면 false를 반환합니다.
func IsQwertyHangul(input string) bool {
	hasLetters := false

	for _, r := range input {
		// If the character is within the Hangul Unicode block, it's not QWERTY Hangul
		if unicode.Is(unicode.Hangul, r) {
			return false
		}
		// If the character is not an ASCII letter or digit, it's also not QWERTY Hangul
		if r > unicode.MaxASCII || (!unicode.IsLetter(r) && !unicode.IsDigit(r) && r != ' ') {
			return false
		}
		// Mark that we've seen at least one letter
		if unicode.IsLetter(r) {
			hasLetters = true
		}
	}

	// If no letters were found, return false (e.g., input is all digits)
	return hasLetters
}
