package dkssud

// Constants for Hangul components and their English mappings
var (
	koTopEn   = []string{"r", "R", "s", "e", "E", "f", "a", "q", "Q", "t", "T", "d", "w", "W", "c", "z", "x", "v", "g"}
	koMidEn   = []string{"k", "o", "i", "O", "j", "p", "u", "P", "h", "hk", "ho", "hl", "y", "n", "nj", "np", "nl", "b", "m", "ml", "l"}
	koBotEn   = []string{"", "r", "R", "rt", "s", "sw", "sg", "e", "f", "fr", "fa", "fq", "ft", "fx", "fv", "fg", "a", "q", "qt", "t", "T", "d", "w", "c", "z", "x", "v", "g"}
	enUpper   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	rawMapper = []string{"r", "R", "rt", "s", "sw", "sg", "e", "E", "f", "fr", "fa", "fq", "ft", "fx", "fv", "fg", "a", "q", "Q", "qt", "t", "T", "d", "w", "W", "c", "z", "x", "v", "g", "k", "o", "i", "O", "j", "p", "u", "P", "h", "hk", "ho", "hl", "y", "n", "nj", "np", "nl", "b", "m", "ml", "l"}

	// combLen returns the Korean character's length in the view of English characters (e.g., T -> 1, TMMBB -> 5)
	combLen = map[int]int{
		T:     1,
		M:     1,
		B:     1,
		TM:    2,
		TMM:   3,
		TMB:   3,
		TMMB:  4,
		TMBB:  4,
		TMMBB: 5,
	}
)

// Constants for bitmasks
const (
	T     = 0b00010000
	M     = 0b00000100
	B     = 0b00000001
	TM    = T + M
	MM    = M + M
	MB    = M + B
	BB    = B + B
	TMM   = T + M + M
	TMB   = T + M + B
	TMMB  = T + M + M + B
	TMBB  = T + M + B + B
	TMMBB = T + M + M + B + B
)

var (
	koTopEnMap = map[string]int{
		"E": 4,
		"Q": 8,
		"R": 1,
		"T": 10,
		"W": 13,
		"a": 6,
		"c": 14,
		"d": 11,
		"e": 3,
		"f": 5,
		"g": 18,
		"q": 7,
		"r": 0,
		"s": 2,
		"t": 9,
		"v": 17,
		"w": 12,
		"x": 16,
		"z": 15,
	}

	koMidEnMap = map[string]int{
		"O":  3,
		"P":  7,
		"b":  17,
		"h":  8,
		"hk": 9,
		"hl": 11,
		"ho": 10,
		"i":  2,
		"j":  4,
		"k":  0,
		"l":  20,
		"m":  18,
		"ml": 19,
		"n":  13,
		"nj": 14,
		"nl": 16,
		"np": 15,
		"o":  1,
		"p":  5,
		"u":  6,
		"y":  12,
	}

	koBotEnMap = map[string]int{
		"":   0,
		"R":  2,
		"T":  20,
		"a":  16,
		"c":  23,
		"d":  21,
		"e":  7,
		"f":  8,
		"fa": 10,
		"fg": 15,
		"fq": 11,
		"fr": 9,
		"ft": 12,
		"fv": 14,
		"fx": 13,
		"g":  27,
		"q":  17,
		"qt": 18,
		"r":  1,
		"rt": 3,
		"s":  4,
		"sg": 6,
		"sw": 5,
		"t":  19,
		"v":  26,
		"w":  22,
		"x":  25,
		"z":  24,
	}
	rawMapperMap = map[string]int{
		"E":  7,
		"O":  33,
		"P":  37,
		"Q":  18,
		"R":  1,
		"T":  21,
		"W":  24,
		"a":  16,
		"b":  47,
		"c":  25,
		"d":  22,
		"e":  6,
		"f":  8,
		"fa": 10,
		"fg": 15,
		"fq": 11,
		"fr": 9,
		"ft": 12,
		"fv": 14,
		"fx": 13,
		"g":  29,
		"h":  38,
		"hk": 39,
		"hl": 41,
		"ho": 40,
		"i":  32,
		"j":  34,
		"k":  30,
		"l":  50,
		"m":  48,
		"ml": 49,
		"n":  43,
		"nj": 44,
		"nl": 46,
		"np": 45,
		"o":  31,
		"p":  35,
		"q":  17,
		"qt": 19,
		"r":  0,
		"rt": 2,
		"s":  3,
		"sg": 5,
		"sw": 4,
		"t":  20,
		"u":  36,
		"v":  28,
		"w":  23,
		"x":  27,
		"y":  42,
		"z":  26,
	}
)
