package dkssud

import (
	"strings"
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

	// Convert uppercase letters to lowercase, except for letters in uppercaseLettersToKeep
	var sb strings.Builder
	sb.Grow(len(input))
	for i := 0; i < len(input); i++ {
		c := input[i]
		if c >= 'A' && c <= 'Z' {
			if _, keep := uppercaseLettersToKeep[c]; !keep {
				c += 'a' - 'A'
			}
		}
		sb.WriteByte(c)
	}
	processedInput := sb.String()

	separated := make([][]string, 0, len(processedInput))
	inputLen := len(processedInput)
	jump := 0

	for i := 0; i < inputLen; i++ {
		if jump > 0 {
			jump--
			continue
		}

		shift := 0
		combination := T
		currentIdx := i
		c := processedInput[i]

		if !isLetter(c) || isDigit(c) || c == ' ' {
			separated = append(separated, []string{string(c)})
			continue
		}

		// Attachment checks with caching
		idx := currentIdx + shift
		idxNext := idx + 1

		if idxNext < inputLen {
			attachType := IsAttachAvailable(processedInput[idx], processedInput[idxNext])
			if attachType == 2 {
				shift++
				combination += M

				idx = currentIdx + shift
				idxNext = idx + 1

				if idxNext < inputLen {
					attachType = IsAttachAvailable(processedInput[idx], processedInput[idxNext])
					if attachType == 3 {
						shift++
						combination += M
						idx = currentIdx + shift
						idxNext = idx + 1
					}
				}

				if idxNext < inputLen {
					attachType = IsAttachAvailable(processedInput[idx], processedInput[idxNext])
					if attachType == 4 {
						shift++
						combination += B
						idx = currentIdx + shift
						idxNext = idx + 1

						if idxNext < inputLen {
							attachment3 := IsAttachAvailable(processedInput[idx], processedInput[idxNext])
							if attachment3 == 5 {
								if idxNext+1 == inputLen {
									combination += B
								} else if idxNext+1 < inputLen {
									shift++
									idx = currentIdx + shift
									idxNext = idx + 1
									attachment4 := IsAttachAvailable(processedInput[idx], processedInput[idxNext])
									if attachment4 != 2 {
										combination += B
									}
								}
							} else if attachment3 == 2 {
								combination -= B
							}
						}
					}
				}
			}
		}

		// Append slices based on combination
		separated = append(separated, createSliceByCombination(processedInput, currentIdx, combination))

		jump = combLen[combination] - 1
	}

	return separated
}

// Helper function to create the slice based on the combination and range
func createSlice(input string, startIdx int, sliceLengths ...int) []string {
	result := make([]string, len(sliceLengths))
	idx := startIdx
	for i, length := range sliceLengths {
		result[i] = input[idx : idx+length]
		idx += length
	}
	return result
}

func isLetter(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func createSliceByCombination(input string, startIdx int, combination int) []string {
	switch combination {
	case T:
		return createSlice(input, startIdx, 1)
	case TM:
		return createSlice(input, startIdx, 1, 1)
	case TMM:
		return createSlice(input, startIdx, 1, 2)
	case TMB:
		return createSlice(input, startIdx, 1, 1, 1)
	case TMMB:
		return createSlice(input, startIdx, 1, 2, 1)
	case TMBB:
		return createSlice(input, startIdx, 1, 1, 2)
	case TMMBB:
		return createSlice(input, startIdx, 1, 2, 2)
	default:
		// Handle unexpected combination
		return []string{string(input[startIdx])}
	}
}
