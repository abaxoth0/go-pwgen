package pwgen

import (
	"errors"
	"math/rand"
	"time"
	"unsafe"
)

var lowLetters = "abcdefghijklmnopqrstuvwxyz"
var upLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
var digits = "0123456789"
var special = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"

const iBits = 6
const iMask = 1<<iBits - 1
const iMax = 63 / iBits

var s = rand.NewSource(time.Now().UnixNano())

type Options struct {
	Length  int
	Lower   bool
	Upper   bool
	Digits  bool
	Special bool
}

var opt = &Options{
	Length:  16,
	Lower:   true,
	Upper:   true,
	Digits:  true,
	Special: false,
}

func (o *Options) pool() (string, error) {
	pool := ""

	if o.Lower {
		pool += lowLetters
	}
	if o.Upper {
		pool += upLetters
	}
	if o.Digits {
		pool += digits
	}
	if o.Special {
		pool += special
	}
	if pool == "" {
		return "", errors.New("at least one character type must be selected")
	}

	return pool, nil
}

func Generate(options *Options) (string, error) {
	if options == nil {
		options = opt
	}

	if options.Length <= 0 {
		return "", errors.New("password length must be greater than 0")
	}

	pool, err := options.pool()

	if err != nil {
		return "", err
	}

	b := make([]byte, options.Length)

	for i, cache, remain := options.Length-1, s.Int63(), iMax; i >= 0; {
		if remain == 0 {
			cache, remain = s.Int63(), iMax
		}

		if idx := int(cache & iMask); idx < len(pool) {
			b[i] = pool[idx]
			i--
		}

		cache >>= iBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b)), nil
}
