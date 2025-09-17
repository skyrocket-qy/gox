package stringmatch

// BoyerMoore finds all occurrences of a pattern in a text using the Boyer-Moore algorithm.
func BoyerMoore(text, pattern string) []int {
	n := len(text)

	m := len(pattern)
	if m == 0 {
		return []int{}
	}

	if n < m {
		return []int{}
	}

	badChar := make(map[byte]int)
	for i := range m {
		badChar[pattern[i]] = i
	}

	result := []int{}

	s := 0
	for s <= (n - m) {
		j := m - 1
		for j >= 0 && pattern[j] == text[s+j] {
			j--
		}

		if j < 0 {
			result = append(result, s)
			if s+m < n {
				shift, ok := badChar[text[s+m]]
				if !ok {
					s += m
				} else {
					s += m - shift
				}
			} else {
				s++
			}
		} else {
			shift, ok := badChar[text[s+j]]
			if !ok {
				s += j + 1
			} else {
				s += max(1, j-shift)
			}
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
