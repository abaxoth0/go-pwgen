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

type Config struct {
	// Password length
	Length  int
	// Include lower-case latin letters
	Lower   bool
	// Include upper-case latin letters
	Upper   bool
	// Include digits (0..9)
	Digits  bool
	// Include special symbols: !"#$%&'()*+,-./:;<=>?@[\]^_`{|}~
	Special bool
}

var defaultConfig = &Config{
	Length:  16,
	Lower:   true,
	Upper:   true,
	Digits:  true,
	Special: false,
}

func (c *Config) pool() (string, error) {
	pool := ""

	if c.Lower {
		pool += lowLetters
	}
	if c.Upper {
		pool += upLetters
	}
	if c.Digits {
		pool += digits
	}
	if c.Special {
		pool += special
	}
	if pool == "" {
		return "", errors.New("At least one character type must be enabled in password generation config")
	}

	return pool, nil
}

// If config is nil then will be used default config (len - 16, all symbols, except special)
func Generate(config *Config) (string, error) {
	if config == nil {
		config = defaultConfig
	}

	if config.Length <= 0 {
		return "", errors.New("password length must be greater than 0")
	}

	pool, err := config.pool()

	if err != nil {
		return "", err
	}

	buffer := make([]byte, config.Length)

	for i, cache, remain := config.Length-1, src.Int63(), iMax; i >= 0; {
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

