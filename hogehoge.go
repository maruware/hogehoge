package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var vowel_letters = []rune("bcdfghjklmnpqrstvwxyz_")
var cons_letters = []rune("aiueo")

func randVowel() rune {
	return vowel_letters[rand.Intn(len(vowel_letters))]
}

func randConsonant() rune {
	return cons_letters[rand.Intn(len(cons_letters))]
}

func randKana() string {
	v := randVowel()
	if v == '_' {
		return string(randConsonant())
	} else {
		return string([]rune{v, randConsonant()})
	}
}

func main() {
	flag.Parse()

	n := int64(1)
	if flag.NArg() > 0 {
		n, _ = strconv.ParseInt(flag.Arg(0), 10, 64)
	}

	rand.Seed(time.Now().Unix())

	for i := 0; i < int(n); i++ {

		k1 := randKana()
		k2 := randKana()

		s := k1 + k2 + k1 + k2
		fmt.Println(s)
	}
}
