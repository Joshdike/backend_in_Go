package service

import (
	"crypto/rand"
	"errors"
	"math/big"
	"strings"
)

func GeneratePassword(length int, includeUppercase, includeNumbers, includeSpecial bool) (string, string, error) {
	if length < 4 || length > 32 {
		return "", "", errors.New("requested password length must be from 4 to 32 characters")
	}
	lowercase := "abcdefghijklmnopqrstuvwxyz"
	uppercase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers := "0123456789"
	special := "!@#$%^&*()-_=+[]{}|;:,.<>?/"
	pool := lowercase
	password := make([]byte, length)
	requiredPools := [][]byte{[]byte(lowercase)}

	if includeUppercase {
		pool += uppercase
		requiredPools = append(requiredPools, []byte(uppercase))
	}
	if includeNumbers {
		pool += numbers
		requiredPools = append(requiredPools, []byte(numbers))
	}
	if includeSpecial {
		pool += special
		requiredPools = append(requiredPools, []byte(special))
	}

	for i, charPool := range requiredPools {
		n, err := randomElement(charPool)
		if err != nil {
			return "", "", errors.New(err.Error())
		}
		password[i] = n
	}

	poolBytes := []byte(pool)
	for i := len(requiredPools); i < length; i++ {
		n, err := randomElement(poolBytes)
		if err != nil {
			return "", "", errors.New(err.Error())
		}
		password[i] = n
	}

	err := shuffle(password)
	if err != nil {
		return "", "", errors.New(err.Error())
	}

	return string(password), passwordStrength(password), nil
}

func randomElement(pool []byte) (byte, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(pool))))
	if err != nil {
		return 0, err
	}
	return pool[n.Int64()], nil
}

func shuffle(a []byte) error {
	for i := len(a) - 1; i > 0; i-- {
		j, err := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
		if err != nil {
			return err
		}
		a[i], a[j.Int64()] = a[j.Int64()], a[i]
	}
	return nil
}

func passwordStrength(password []byte) string {
	strength := 0
	lowercase := "abcdefghijklmnopqrstuvwxyz"
	uppercase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers := "0123456789"
	special := "!@#$%^&*()-_=+[]{}|;:,.<>?/"
	if len(password) >= 8 {
		strength += 2
	}
	if containsAny(password, lowercase) {
		strength += 1
	}
	if containsAny(password, uppercase) {
		strength += 1
	}
	if containsAny(password, numbers) {
		strength += 1
	}
	if containsAny(password, special) {
		strength += 1
	}
	if strength == 6 {
		return "very strong"
	} else if strength == 5 {
		return "strong"
	} else if strength >= 3 {
		return "medium"
	} else {
		return "weak"
	}
}

func containsAny(char []byte, word string) bool {
	for _, i := range char {
		if strings.ContainsRune(word, rune(i)) {
			return true
		}
	}
	return false
}
