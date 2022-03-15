package util

import (
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func RandomNameOfCity() string {
	return RandomString(6)
}

func RandomPosition() float64 {
	return RandomFloat(-1000, 1000)
}

func RandomString(n int) string {
	k := len(alphabet)
	var sb strings.Builder
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min+1)
}
