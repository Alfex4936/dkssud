package dkssud

import (
	"strings"
	"unicode"
)

// splitKo disassembles Korean characters into their indexes in the QWERTY keyboard map.
func splitKo(input string) [][]interface{} {
	if input == "" {
		return nil
	}

	// Preallocate space for the separated slice to improve performance
	separated := make([][]interface{}, 0, len(input))

	for _, c := range input {
		switch {
		case c == ' ':
			separated = append(separated, []interface{}{" "})
		case c >= 44032 && c <= 55203: // Hangul Syllables range
			hexZeropoint := int(c) - 44032
			topIdx := hexZeropoint / (28 * 21)
			midIdx := (hexZeropoint / 28) % 21
			botIdx := hexZeropoint % 28
			separated = append(separated, []interface{}{topIdx, midIdx, botIdx})
		case c >= 12593 && c <= 12643: // Hangul Jamo (individual consonants and vowels)
			separated = append(separated, []interface{}{int(c)})
		default: // Non-Hangul characters
			separated = append(separated, []interface{}{string(c)})
		}
	}

	return separated
}

// splitEn remains largely the same, except it no longer converts specific letters like T, R, E, Q to lowercase
func splitEn(input string) [][]string {
	if input == "" {
		return nil
	}

	// Convert all uppercase letters to lowercase, except for specific ones like T, R, E, Q
	for _, c := range enUpper {
		if !strings.ContainsAny(string(c), "TREQ") {
			input = strings.ReplaceAll(input, string(c), strings.ToLower(string(c)))
		}
	}

	var separated [][]string
	jump := 0

	for i := 0; i < len(input); i++ {
		if jump > 0 {
			jump--
			continue
		}

		shift := 0
		combination := T
		currentIdx := i

		// Handle non-letter and non-digit characters first
		r := rune(input[i])
		if !unicode.IsLetter(r) || unicode.IsDigit(r) || input[i] == ' ' {
			separated = append(separated, []string{string(input[i])})
			continue
		}

		if currentIdx+shift+1 < len(input) && IsAttachAvailable(input[currentIdx+shift], input[currentIdx+shift+1]) == 2 {
			shift++
			combination += M

			if currentIdx+shift+1 < len(input) && IsAttachAvailable(input[currentIdx+shift], input[currentIdx+shift+1]) == 3 {
				shift++
				combination += M
			}

			if currentIdx+shift+1 < len(input) && IsAttachAvailable(input[currentIdx+shift], input[currentIdx+shift+1]) == 4 {
				shift++
				combination += B

				if currentIdx+shift+1 < len(input) {
					attachment3 := IsAttachAvailable(input[currentIdx+shift], input[currentIdx+shift+1])
					if attachment3 == 5 {
						if currentIdx+shift+2 == len(input) {
							combination += B
						} else if currentIdx+shift+2 < len(input) {
							shift++
							attachment4 := IsAttachAvailable(input[currentIdx+shift], input[currentIdx+shift+1])
							if attachment4 == 2 {
								// 자 + 자 + 모
							} else {
								combination += B
							}
						}
					} else if attachment3 == 2 {
						combination -= B // Remove 'B'
					}
				}
			}
		}

		// Based on the combination, append the appropriate slices
		switch combination {
		case T:
			separated = append(separated, []string{string(input[currentIdx])})
		case TM:
			separated = append(separated, []string{string(input[currentIdx]), string(input[currentIdx+1])})
		case TMM:
			separated = append(separated, []string{string(input[currentIdx]), input[currentIdx+1 : currentIdx+3]})
		case TMB:
			separated = append(separated, []string{string(input[currentIdx]), string(input[currentIdx+1]), string(input[currentIdx+2])})
		case TMMB:
			separated = append(separated, []string{string(input[currentIdx]), input[currentIdx+1 : currentIdx+3], string(input[currentIdx+3])})
		case TMBB:
			separated = append(separated, []string{string(input[currentIdx]), string(input[currentIdx+1]), input[currentIdx+2 : currentIdx+4]})
		case TMMBB:
			separated = append(separated, []string{string(input[currentIdx]), input[currentIdx+1 : currentIdx+3], input[currentIdx+3 : currentIdx+5]})
		}

		jump = combLen[combination] - 1
	}

	return separated
}
