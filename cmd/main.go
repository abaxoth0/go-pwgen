package main

import (
	"fmt"
	"os"

	"github.com/StepanAnanin/go-pwgen/packages/pwgen"
	"github.com/akamensky/argparse"
)

func main() {
	parser := argparse.NewParser("pwgen", "Generate a random password")

	length := parser.Int("L", "length", &argparse.Options{Default: 16, Help: "Password length"})
	upper := parser.Flag("u", "upper", &argparse.Options{Help: "Include uppercase letters"})
	lower := parser.Flag("l", "lower", &argparse.Options{Help: "Include lowercase letters"})
	digits := parser.Flag("d", "digits", &argparse.Options{Help: "Include digits"})
	special := parser.Flag("s", "special", &argparse.Options{Help: "Include special characters"})

	if err := parser.Parse(os.Args); err != nil {
		fmt.Println(parser.Usage(err))
		return
	}

	if !*upper && !*lower && !*digits && !*special {
		*upper = true
		*lower = true
		*digits = true
	}

	p, err := pwgen.Generate(&pwgen.Options{
		Length:  *length,
		Upper:   *upper,
		Lower:   *lower,
		Digits:  *digits,
		Special: *special,
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	println(p)
}
