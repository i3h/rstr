package rstr

import (
	"math/rand"
	"time"
)

const (
	DefaultSet string = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	LetterSet  string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	NumberSet  string = "0123456789"
)

func New(num int64, charSet string) string {
	runeSet := []rune(charSet)
	rand.Seed(time.Now().UTC().UnixNano())
	b := make([]rune, num)
	for i := range b {
		b[i] = runeSet[rand.Intn(len(runeSet))]
	}
	return string(b)
}
