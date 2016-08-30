// Match provides a simple pattern matcher with unicode support.
package match

import "unicode/utf8"

// Match returns true if str matches pattern. This is a very
// simple wildcard match where '*' matches on any number characters
// and '?' matches on any one character.

// pattern:
// 	{ term }
// term:
// 	'*'         matches any sequence of non-Separator characters
// 	'?'         matches any single non-Separator character
// 	c           matches character c (c != '*', '?', '\\')
// 	'\\' c      matches character c
//
func Match(str, pattern string) bool {
	if pattern == "*" {
		return true
	}
	return deepMatch(str, pattern)
}
func deepMatch(str, pattern string) bool {
	for len(pattern) > 0 {
		if pattern[0] > 0x7f {
			return deepMatchRune(str, pattern)
		}
		switch pattern[0] {
		default:
			if len(str) == 0 {
				return false
			}
			if str[0] > 0x7f {
				return deepMatchRune(str, pattern)
			}
			if str[0] != pattern[0] {
				return false
			}
		case '?':
			if len(str) == 0 {
				return false
			}
		case '*':
			return deepMatch(str, pattern[1:]) ||
				(len(str) > 0 && deepMatch(str[1:], pattern))
		}
		str = str[1:]
		pattern = pattern[1:]
	}
	return len(str) == 0 && len(pattern) == 0
}

func deepMatchRune(str, pattern string) bool {
	var sr, pr rune
	var srsz, prsz int

	// read the first rune ahead of time
	if len(str) > 0 {
		if str[0] > 0x7f {
			sr, srsz = utf8.DecodeRuneInString(str)
		} else {
			sr, srsz = rune(str[0]), 1
		}
	} else {
		sr, srsz = utf8.RuneError, 0
	}
	if len(pattern) > 0 {
		if pattern[0] > 0x7f {
			pr, prsz = utf8.DecodeRuneInString(pattern)
		} else {
			pr, prsz = rune(pattern[0]), 1
		}
	} else {
		pr, prsz = utf8.RuneError, 0
	}
	// done reading
	for pr != utf8.RuneError {
		switch pr {
		default:
			if srsz == utf8.RuneError {
				return false
			}
			if sr != pr {
				return false
			}
		case '?':
			if srsz == utf8.RuneError {
				return false
			}
		case '*':
			return deepMatchRune(str, pattern[prsz:]) ||
				(srsz > 0 && deepMatchRune(str[srsz:], pattern))
		}
		str = str[srsz:]
		pattern = pattern[prsz:]
		// read the next runes
		if len(str) > 0 {
			if str[0] > 0x7f {
				sr, srsz = utf8.DecodeRuneInString(str)
			} else {
				sr, srsz = rune(str[0]), 1
			}
		} else {
			sr, srsz = utf8.RuneError, 0
		}
		if len(pattern) > 0 {
			if pattern[0] > 0x7f {
				pr, prsz = utf8.DecodeRuneInString(pattern)
			} else {
				pr, prsz = rune(pattern[0]), 1
			}
		} else {
			pr, prsz = utf8.RuneError, 0
		}
		// done reading
	}

	return srsz == 0 && prsz == 0
}
