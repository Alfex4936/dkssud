package dkssud

// contains checks if a slice contains a specific string.
func contains(slice []string, item string) bool {
	for _, val := range slice {
		if val == item {
			return true
		}
	}
	return false
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

// isAttachAvailable checks if two characters can be attached based on Korean language rules.
func isAttachAvailable(i, l byte) int {
	switch {
	case contains(koTopEn, string(i)) && contains(koMidEn, string(l)):
		// 자 + 모 (Consonant + Vowel)
		return 2
	case contains(koMidEn, string(i)) && contains(koMidEn, string(l)):
		// 모 + 모 (Composite Vowel)
		if contains(koMidEn, string([]byte{i, l})) {
			return 3
		}
		return 0
	case contains(koMidEn, string(i)) && contains(koBotEn, string(l)):
		// 모 + 자 (Vowel + Consonant)
		return 4
	case contains(koBotEn, string([]byte{i, l})):
		// 자 + 자 (종) (Double Consonant)
		return 5
	default:
		// If none of the above, they are not attachable
		return 0
	}
}
