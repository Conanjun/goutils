package common

import (
	"math/rand"
	"time"
)

var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var smallletters = []rune("abcdefghijklmnopqrstuvwxyz")
var smalllettersnums = []rune("abcdefghijklmnopqrstuvwxyz0123456789")

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandInt64(min, max int64) int64 {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Int63n(max-min) + min
}

func GetRandomString(l int) string {
	b := make([]rune, l)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func GetSmallRandomString(l int) string {
	b := make([]rune, l)
	for i := range b {
		b[i] = smallletters[rand.Intn(len(smallletters))]
	}
	return string(b)
}

func GetSmallNumsRandomString(l int) string {
	b := make([]rune, l)
	for i := range b {
		b[i] = smalllettersnums[rand.Intn(len(smalllettersnums))]
	}
	return string(b)
}
