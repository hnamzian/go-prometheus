package utils

import (
	"math/rand"
	"time"
)

func Sleep(ms int) {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	now := time.Now()
	n := rand.Intn(ms + now.Second())
	time.Sleep(time.Duration(n) * time.Millisecond)
}
