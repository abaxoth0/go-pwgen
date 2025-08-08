package pwgen

import (
	"errors"
	"math/rand"
	"time"
	"unsafe"
)

const (
	lowLetters = "abcdefghijklmnopqrstuvwxyz"
	upLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits = "0123456789"
	special = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
)

const iBits = 6
const iMask = 1<<iBits - 1
const iMax = 63 / iBits

var src = rand.NewSource(time.Now().UnixNano())

const (
	// Lower-case latin letters
	LOWER uint8 = 1 << iota
	// Upper-case latin letters
	UPPER
	// Digits (0..9)
	DIGITS
	// Special symbols: !"#$%&'()*+,-./:;<=>?@[\]^_`{|}~
	SPECIAL
)

func pool(charset uint8) (string, error) {
	pool := ""

	if charset&LOWER != 0 {
		pool += lowLetters
	}
	if charset&UPPER != 0 {
		pool += upLetters
	}
	if charset&DIGITS != 0 {
		pool += digits
	}
	if charset&SPECIAL != 0 {
		pool += special
	}
	if pool == "" {
		return "", errors.New("Invalid character set: at least one character type must selected")
	}

	return pool, nil
}

// If config is nil then will be used default config (len - 16, all symbols, except special)
func Generate(length int, charset uint8) (string, error) {
	if length <= 0 {
		return "", errors.New("password length must be greater than 0")
	}

	pool, err := pool(charset)
	if err != nil {
		return "", err
	}

	buffer := make([]byte, length)

	for i, cache, remain := length-1, src.Int63(), iMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), iMax
		}

		if idx := int(cache & iMask); idx < len(pool) {
			buffer[i] = pool[idx]
			i--
		}

		cache >>= iBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&buffer)), nil
}

