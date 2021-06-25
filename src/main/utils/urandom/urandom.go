package urandom

import (
	"math/rand"
	"time"
)

const (
	LOWER   string = "qwertyuiopasdfghjklzxcvbnm"
	UPPER   string = "QWERTYUIOPASDFGHJKLZXCVBNM"
	NUMBER  string = "0123456789"
	SPECIAL string = "~!@#$%^&*()_+-=[]{}|;:,.<>?"
)

func GenString(seeds string, length int) string {
	rand.Seed(time.Now().UnixNano())
	byteSeed := []byte(seeds)
	var result []byte
	for i := 0; i < length; i++ {
		result = append(result, byteSeed[rand.Intn(len(byteSeed))])
	}
	return string(result)
}

func GenStringTimes(seeds string, length int, times int) []string {
	var result []string
	if times <= 0 {
		result = append(result, GenString(seeds, length))
	} else {
		for i := 0; i < times; i++ {
			result = append(result, GenString(seeds, length))
			time.Sleep(time.Millisecond)
		}
	}
	return result
}
