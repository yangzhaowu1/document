package main

import "fmt"

/*
242. 有效的字母异位词
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
注意：若 s 和 t 中每个字符出现的次数都相同，则称 s 和 t 互为字母异位词。
*/

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	byteS := make(map[byte]int)
	byteT := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		byteS[s[i]]++
		byteT[t[i]]++
	}

	for k, v := range byteS {
		if v1, ok := byteT[k]; !ok || v != v1 {
			return false
		}
	}

	return true
}

func main()  {
	fmt.Println(isAnagram("anagram", "nagaram"))
}