package main

import (
	"fmt"
	"leetcode/utils"
)

var (
	memStartIdx int
	memEndIdx   int
)

func main() {
	testCases := []struct {
		s        string
		expected []string
	}{
		{"babad", []string{"bab", "aba"}},
		{"cbbd", []string{"bb"}},
		{"forgeeksskeegfor", []string{"geeksskeeg"}},
		{"Geeks", []string{"ee"}},
		{"abcde", []string{"a", "b", "c", "d", "e"}},
	}

	for _, tc := range testCases {
		fmt.Printf("* Case: %s, expected: %+v\n", tc.s, tc.expected)
		fmt.Printf("  - Recursive : %s\n", longestPalindrome(tc.s))
		fmt.Printf("  - DP        : %s\n", LPbyDP(tc.s))
	}
}

/** Recurrence Relations

	f(startIdx, endIdx, s) = is this  palindrome string

	- 1: startIdx == endIdx or start > endIdx,
	- f(startIdx, endIdx, s)
		- s[startIdx] != s[endIdx] => Max(f(startIdx, endIdx-1, s), f(startIdx+1, endIdx, s))
		- s[startIdx] == s[endIdx]
			- if f(startIdx+1, endIdx-1, s) == (endIdx-1 - startIdx+1 + 1) : sub-string is palindrome string
			  => f(startIdx+1, endIdx-1, s) + 2
			- else
			  => f(startIdx+1, endIdx-1, s)
**/

func longestPalindrome(s string) string {
	var memIdx []int = []int{0, 0}
	LPbyRecursion(0, len(s)-1, s, memIdx)
	return s[memIdx[0] : memIdx[1]+1]
}

func LPbyRecursion(start, end int, s string, memIdx []int) int {
	if start == end {
		return 1
	} else if start > end {
		return 0
	}

	if s[start] != s[end] {
		return utils.Max(LPbyRecursion(start+1, end, s, memIdx), LPbyRecursion(start, end-1, s, memIdx))
	}

	subStart := start + 1
	subEnd := end - 1
	sub := LPbyRecursion(subStart, subEnd, s, memIdx)
	if sub == (subEnd - subStart + 1) {
		memLongestPalindromeStringIdx(start, end, memIdx)
		return sub + 2
	}
	return sub
}

func memLongestPalindromeStringIdx(start, end int, mem []int) {
	if (mem[0] == 0 && mem[1] == 0) || ((end - start + 1) > (mem[1] - mem[0] + 1)) {
		mem[0] = start
		mem[1] = end
	}
}

/* s/e
        e[0] e[1] e[2] e[3] e[4]
   s[0]  1    ?    ?    ?    ?
   s[1]  X    1    ?    ?    ?
   s[2]  X    X    1    ?    ?
   s[3]  X    X    X    1    ?
   s[4]  X    X    X    X    1
*/
func LPbyDP(s string) string {
	res := make([][]int, len(s)) // first: startIdx, second: endIdx
	for i := 0; i < len(s); i++ {
		res[i] = make([]int, len(s))
	}

	var memStartIdx int = 0
	var maxLength int = 1

	/* looping 도는 순서 (startIdx/endIdx)
		00 11 22 33 44
	 	01 12 23 34
	 	02 13 24
	 	03 14
	 	04
	*/

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s)-i; j++ {
			start := j
			end := start + i

			if start == end {
				res[start][end] = 1
				continue
			}

			if s[start] != s[end] {
				res[start][end] = utils.Max(res[start+1][end], res[start][end-1])
			} else {
				subStart := start + 1
				subEnd := end - 1
				subLength := subEnd - subStart + 1
				if subStart > subEnd {
					res[start][end] = 2
				} else if res[subStart][subEnd] == subLength {
					res[start][end] = res[subStart][subEnd] + 2
				} else {
					res[start][end] = res[subStart][subEnd]
				}

				if res[start][end] > maxLength {
					memStartIdx = start
					maxLength = res[start][end]
				}
			}
		}
	}
	return s[memStartIdx : memStartIdx+maxLength]
}
