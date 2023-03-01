// Package natsort implements natural strings sorting
package natsort

import (
	"sort"
	"strings"
)

type stringSlice []string

func (s stringSlice) Len() int {
	return len(s)
}

func (s stringSlice) Less(a, b int) bool {
	return Compare(s[a], s[b])
}

func (s stringSlice) Swap(a, b int) {
	s[a], s[b] = s[b], s[a]
}

// Sort sorts a list of strings in a natural order
func Sort(l []string) {
	sort.Sort(stringSlice(l))
}

// Compare returns true if the first string precedes the second one according to natural order
func Compare(a, b string) bool {
	lnA := len(a)
	lnB := len(b)
	posA := 0
	posB := 0

	for {
		if lnA <= posA {
			if lnB <= posB {
				// eof on both at the same time (equal)
				return false
			}
			return true
		} else if lnB <= posB {
			// eof on b
			return false
		}

		av, bv := a[posA], b[posB]

		if av >= '0' && av <= '9' && bv >= '0' && bv <= '9' {
			// go into numeric mode
			intLnA := 1
			intLnB := 1
			for {
				if posA+intLnA >= lnA {
					break
				}
				x := a[posA+intLnA]
				if av == '0' {
					posA += 1
					av = x
					continue
				}
				if x >= '0' && x <= '9' {
					intLnA += 1
				} else {
					break
				}
			}
			for {
				if posB+intLnB >= lnB {
					break
				}
				x := b[posB+intLnB]
				if bv == '0' {
					posB += 1
					bv = x
					continue
				}
				if x >= '0' && x <= '9' {
					intLnB += 1
				} else {
					break
				}
			}
			if intLnB > intLnA {
				// length of a value is longer, means it's a bigger number
				return true
			} else if intLnA > intLnB {
				return false
			}
			// both have same length, let's compare as string
			v := strings.Compare(a[posA:posA+intLnA], b[posB:posB+intLnB])
			if v < 0 {
				return true
			} else if v > 0 {
				return false
			}
			// equal
			posA += intLnA
			posB += intLnB
			continue
		}

		if av == bv {
			posA += 1
			posB += 1
			continue
		}

		return av < bv
	}
}
