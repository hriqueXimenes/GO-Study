package main

import "fmt"

func main() {
	fmt.Println(checkPermutation("ab", "asdczxcwba"))
}

// ## -- Not great solution - BIG O(n^2)
func checkPermutation(s1 string, s2 string) bool {
	var s1nums [26]int
	var s2nums [26]int

	if len(s1) > len(s2) {
		return false
	}

	checkEqual := func(arr1, arr2 [26]int) bool {
		for i := 0; i < 26; i++ {
			if arr1[i] != arr2[i] {
				return false
			}
		}

		return true
	}

	for i := 0; i < len(s2)-len(s1)+1; i++ {
		for l := 0; l < len(s1); l++ {
			s1nums[s1[l]-'a']++
			s2nums[s2[l+i]-'a']++
		}
		if checkEqual(s1nums, s2nums) {
			return true
		}

		s1nums = [26]int{}
		s2nums = [26]int{}
	}

	return false

}
