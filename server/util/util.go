package util

import (
	"math/rand"
	"time"
)

func RandomUsername(n int) string {
	var letters = []byte("qweryuiopdfghjvbnmvzaDBGHIUJGQWIPIVGASBHUZ")
	var res string
	rand.Seed(time.Now().Unix())
	for i := 0; i < n; i++ {
		res += string(letters[rand.Int()%len(letters)])
	}
	return res
}
