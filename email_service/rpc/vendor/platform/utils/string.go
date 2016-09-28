package utils

import (
	"fmt"
	"strings"
	"time"
	"unicode"
)

func GetRandomName() string {
	t := time.Now()
	filename := fmt.Sprintf("%04d%02d%02d%02d%02d%02d%03d%06d", t.Year(), int(t.Month()), t.Day(), t.Hour(), t.Minute(), t.Second(), int(t.Nanosecond()/1000000), GetRand().Intn(1000000))
	return filename

}

// 从指定位置截取string，当start ＋ length 超出末尾后，返回长度根据实际长度返回，如Substr（“chinarun”， 7， 3），返回为n
// start为负值时反向（从最末向前）截取.比如Substr（“chinarun”， -2， 3），返回为nar
func SubStr(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)

	if start >= rl || length < 1 {
		return ""
	}

	end := 0

	if start < 0 {
		end = rl + start
		if end < 0 {
			return ""
		}

		start = end - length
		if start < 0 {
			start = 0
		}
	} else {
		end = start + length
		if end > rl {
			end = rl
		}
	}

	return string(rs[start:end])
}

// 驼峰命名转成下划线或短横等命名
// src - 要转换的驼峰字符串
// sep - 分隔符，传空为下划线，可指定其他字符
// 返回转换后的字符串
func SepSplitName(src, sep string) (dst string) {
	if sep == "" {
		sep = "_"
	}

	for _, s := range src {
		if unicode.IsUpper(s) {
			dst += sep + strings.ToLower(string(s))
		} else {
			dst += strings.ToLower(string(s))
		}
	}

	if string(dst[0]) == sep {
		return dst[1:]
	}

	return
}

func CamelToSnake(s string) string {
	var result string
	var words []string
	var lastPos int
	rs := []rune(s)

	for i := 0; i < len(rs); i++ {
		if i > 0 && unicode.IsUpper(rs[i]) {
			if initialism := startsWithInitialism(s[lastPos:]); initialism != "" {
				words = append(words, initialism)

				i += len(initialism) - 1
				lastPos = i
				continue
			}

			words = append(words, s[lastPos:i])
			lastPos = i
		}
	}

	// append the last word
	if s[lastPos:] != "" {
		words = append(words, s[lastPos:])
	}

	for k, word := range words {
		if k > 0 {
			result += "_"
		}

		result += strings.ToLower(word)
	}

	return result
}

// SnakeToCamel returns a string converted from snake case to uppercase
func SnakeToCamel(s string) string {
	var result string

	words := strings.Split(s, "_")

	for _, word := range words {
		if upper := strings.ToUpper(word); commonInitialisms[upper] {
			result += upper
			continue
		}

		if len(word) > 0 {
			w := []rune(word)
			w[0] = unicode.ToUpper(w[0])
			result += string(w)
		}
	}

	return result
}

// startsWithInitialism returns the initialism if the given string begins with it
func startsWithInitialism(s string) string {
	var initialism string
	// the longest initialism is 5 char, the shortest 2
	for i := 1; i <= 5; i++ {
		if len(s) > i-1 && commonInitialisms[s[:i]] {
			initialism = s[:i]
		}
	}
	return initialism
}

// commonInitialisms, taken from
// https://github.com/golang/lint/blob/32a87160691b3c96046c0c678fe57c5bef761456/lint.go#L702
var commonInitialisms = map[string]bool{
	"API":   true,
	"ASCII": true,
	"CPU":   true,
	"CSS":   true,
	"DNS":   true,
	"EOF":   true,
	"GUID":  true,
	"HTML":  true,
	"HTTP":  true,
	"HTTPS": true,
	"ID":    true,
	"IP":    true,
	"JSON":  true,
	"LHS":   true,
	"QPS":   true,
	"RAM":   true,
	"RHS":   true,
	"RPC":   true,
	"SLA":   true,
	"SMTP":  true,
	"SQL":   true,
	"SSH":   true,
	"TCP":   true,
	"TLS":   true,
	"TTL":   true,
	"UDP":   true,
	"UI":    true,
	"UID":   true,
	"UUID":  true,
	"URI":   true,
	"URL":   true,
	"UTF8":  true,
	"VM":    true,
	"XML":   true,
	"XSRF":  true,
	"XSS":   true,
}
