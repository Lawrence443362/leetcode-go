package main
import (
	"strings"
	"unicode"
)
/*
Фраза называется палиндромом , если после преобразования всех заглавных букв в строчные и удаления всех небуквенно-цифровых символов она читается одинаково как в прямом, так и в обратном порядке. К буквенно-цифровым символам относятся буквы и цифры.

Если задана строка s, вернуть значение, true если она является палиндромом , или false значение в противном случае .
*/
func isPalindrome(s string) bool {
	i, j := 0, len(s) - 1

	for i < j {
		for i < j && !isAlnum(rune(s[i])) {i++}
		for i < j && !isAlnum(rune(s[j])) {j--}

		if unicode.ToLower(rune(s[i])) != unicode.ToLower(rune(s[j])) {
			return false
		}
		i++
		j--
	}

	return true
}

func isAlnum(c rune) bool {
	return unicode.IsLetter(c) || unicode.IsDigit(c)
}


func isPalindromeOld(s string) bool {
	cleaned := strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
				return unicode.ToLower(r)
		}
		return -1 // -1 = удалить символ
	}, s)

	if len(cleaned) == 0 || len(cleaned) == 1 {
		return true
	}

	for i, j := 0, len(cleaned) - 1; i < j; i, j = i + 1, j - 1 {
		if cleaned[i] != cleaned[j] {
			return false
		}
	}

	return true
}