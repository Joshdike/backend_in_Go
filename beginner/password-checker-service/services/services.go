package services

import (
	"errors"
	"strings"
)

func StrengthChecker(password string) (string, []string, error) {
	suggestions := []string{}
	if len(password) < 4 || len(password) > 32 {
		return "", suggestions, errors.New("invalid password, length must be from 4 to 32 characters")
	}
	passByte := []byte(password)
	common := []string{"password", "admin", "welcome", "123", "321", "qwerty", "abcd", "p@ssword", "adm1n", "welcome1"}
	lowercase := "abcdefghijklmnopqrstuvwxyz"
	uppercase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers := "0123456789"
	special := "!@#$%^&*()-_=+[]{}|;:,.<>?/"
	strength := 0
	if len(password) >= 8 {
		strength += 2
	} else {
		suggestions = append(suggestions, "Add more characters")
	}
	if containsAny(passByte, lowercase) {
		strength += 1
	} else {
		suggestions = append(suggestions, "Include lowercase characters")
	}
	if containsAny(passByte, uppercase) {
		strength += 1
	} else {
		suggestions = append(suggestions, "Include Uppercase characters")
	}
	if containsAny(passByte, numbers) {
		strength += 1
	} else {
		suggestions = append(suggestions, "Include numbers")
	}
	if containsAny(passByte, special) {
		strength += 1
	} else {
		suggestions = append(suggestions, "Include special characters")
	}

	if hasCommon(strings.ToLower(password), common) || HasRepetitions(password) {
		strength = max(strength-2, 0)
		suggestions = append(suggestions, "Avoid common or repetitive patterns")
	}

	if strength == 6 {
		return "very strong", suggestions, nil
	} else if strength == 5 {
		return "strong", suggestions, nil
	} else if strength >= 3 {
		return "medium", suggestions, nil
	} else {
		return "weak", suggestions, nil
	}

}

func containsAny(password []byte, word string) bool {
	for _, char := range password {
		if strings.ContainsRune(word, rune(char)) {
			return true
		}
	}
	return false
}

func hasRepeatedChars(password string) bool {
	for i := 0; i < len(password)-2; i++ {
		if password[i] == password[i+1] && password[i] == password[i+2] {
			return true
		}
	}
	return false
}

func hasRepeatedSubstrings(password string) bool {
	for l := 2; l <= len(password)/2; l++ {
		for i := 0; i <= len(password)-2*l; i++ {
			sub1 := password[i : i+l]
			sub2 := password[i+l : i+2*l]
			if sub1 == sub2 {
				return true
			}
		}
	}
	return false
}

func HasRepetitions(password string) bool {
	return hasRepeatedChars(password) || hasRepeatedSubstrings(password)
}

func hasCommon(password string, words []string) bool {
	for _, i := range words {
		if strings.Contains(password, i) {
			return true
		}
	}
	return false
}
