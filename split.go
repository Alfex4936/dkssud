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

	// Use strings.Builder for efficient string manipulation
	var sb strings.Builder
	sb.Grow(len(input)) // Preallocate space

	// Iterate through input and convert uppercase letters to lowercase, except for T, R, E, Q
	for _, c := range input {
		if !strings.ContainsRune("TREQ", c) && unicode.IsUpper(c) {
			sb.WriteRune(unicode.ToLower(c))
		} else {
			sb.WriteRune(c)
		}
	}
	processedInput := sb.String()

	separated := make([][]string, 0, len(processedInput)) // Preallocate memory

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
							if attachment4 != 2 {
								combination += B
							} // non 자 + 자 + 모
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
			separated = append(separated, createSlice(input, currentIdx, 1)) // T
		case TM:
			separated = append(separated, createSlice(input, currentIdx, 1, 1)) // TM
		case TMM:
			separated = append(separated, createSlice(input, currentIdx, 1, 2)) // TMM
		case TMB:
			separated = append(separated, createSlice(input, currentIdx, 1, 1, 1)) // TMB
		case TMMB:
			separated = append(separated, createSlice(input, currentIdx, 1, 2, 1)) // TMMB
		case TMBB:
			separated = append(separated, createSlice(input, currentIdx, 1, 1, 2)) // TMBB
		case TMMBB:
			separated = append(separated, createSlice(input, currentIdx, 1, 2, 2)) // TMMBB
		}

		jump = combLen[combination] - 1
	}

	return separated
}

// Helper function to create the slice based on the combination and range
func createSlice(input string, currentIdx int, sliceLengths ...int) []string {
	result := make([]string, len(sliceLengths))
	start := currentIdx
	for i, length := range sliceLengths {
		if length == 1 {
			result[i] = string(input[start]) // Single character
		} else {
			result[i] = input[start : start+length] // Sub-slice for multi-character strings
		}
		start += length
	}
	return result
}
