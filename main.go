package main

import (
	"fmt"
	"strings"
)

type searchEngine struct {
	query     string
	frequency int
}

var searchEngineList []searchEngine

func addQuery(q string, f int) {
	s := searchEngine{
		query:     q,
		frequency: f,
	}
	searchEngineList = append(searchEngineList, s)
}
func filterOutTop3(prefix string) []searchEngine {
	var top3Searches []searchEngine
	for _, s := range searchEngineList {

		if strings.HasPrefix(s.query, prefix) {
			top3Searches = append(top3Searches, s)
		}
	}
	return top3Searches
}
func main() {
	//go SampleRoutine()
	//	fmt.Println("Hello World!...")
	// fmt.Print("Is Madam a palindrome ", isPalindrome("shiva"))
	//fmt.Println("Longest Palindrome :" + longestPalindromeSubStr("shivaUpTownFUNKMADAmsabbcbbashiva"))
	//time.Sleep(1 * time.Second)
	//fmt.Println("Main thread done")
	addQuery("coffee late", 20)
	addQuery("coffee Tapri", 10)
	addQuery("continental", 5)
	addQuery("Hazelnut", 30)
	fmt.Print(searchEngineList)
	fmt.Println(filterOutTop3("co"))

}
func SampleRoutine() {
	fmt.Println("In Sample Routine")
}

// func longestPalindromeSubStr(s string) string {
// 	if len(s) == 0 {
// 		return ""
// 	}
// 	res := ""
// 	subStr := ""
// 	for i := 0; i < len(s); i++ {
// 		for j := i; j < len(s); j++ {
// 			if isPalindrome(s[i:j]) {
// 				subStr = s[i:j]
// 				if len(subStr) > len(res) {
// 					res = subStr
// 				}
// 			}
// 		}
// 	}
// 	return res
// }

// func isPalindrome(s string) bool {
// 	if len(s) == 0 {
// 		return false
// 	}
// 	// madam
// 	for i, j := 0, len(s)-1; i < len(s); i++ {
// 		if s[i] != s[j] {
// 			return false
// 		}
// 		j--
// 	}
// 	return true
// }

func longestPalindromeSubStr(s string) string {
	if len(s) == 0 {
		return ""
	}
	T := "^#" + strings.Join(strings.Split(s, ""), "#") + "#$"
	n := len(T)
	p := make([]int, n)
	c, r := 0, 0
	for i := 0; i < n; i++ {
		if r > i {
			p[i] = min(r-i, p[2*c-1])
		}
		for T[i+1+p[i]] == T[i-1-p[i]] {
			p[i]++
		}
		if i+p[i] > r {
			c, r = i, i+p[i]
		}
	}
	maxLen := 0
	maxCenter := 0
	for i, v := range p {
		if v > maxLen {
			maxLen = v
			maxCenter = i
		}
	}

	return s[(maxCenter-maxLen)/2 : (maxCenter+maxLen)/2]
}
