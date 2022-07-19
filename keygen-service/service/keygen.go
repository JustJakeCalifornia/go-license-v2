package service

import (
	"math/rand"
	"strings"
	"time"
)

const (
	length = 16
)

func generateRandomLicenseKey() string {
	rand.Seed(time.Now().UnixNano())
	key := make([]rune, length)
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	for i := range key {
		key[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	keyRune := []rune(strings.ToUpper(string(key)))

	return splitKey(keyRune)
}

func splitKey(key []rune) string {
	return strings.Join([]string{string(key[0:4]), string(key[4:8]), string(key[8:12]), string(key[12:16])}, "-")
}
