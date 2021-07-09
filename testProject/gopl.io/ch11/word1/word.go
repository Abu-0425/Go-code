package word1

import "unicode"

// 版本1，无法解决两个特殊情况
//func IsPalindrome(s string) bool {
//	for i := range s {
//		if s[i] != s[len(s)-1-i] {
//			return false
//		}
//	}
//	return true
//}

// 升级版
func IsPalindrome(s string) bool {
	var letters []rune // 全部小写字母并且去掉空格等其他字符的集合
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}
	for i := range letters {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}