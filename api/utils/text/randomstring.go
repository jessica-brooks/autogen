package text

import (
	"math/rand"
	"strconv"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

//UniqueString returns unique string from letters and timestamp
func UniqueString(n int) string {
	t := time.Now().Unix()
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return strconv.FormatInt(t, 10) + string(b)
}

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
