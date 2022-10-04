package main

import (
	"math"
	"os"
	"strconv"
	"time"

	"golang.org/x/text/language"
	"golang.org/x/text/message"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.RFC3339,
		NoColor:    false,
	})

	if len(os.Args) != 3 {
		log.Fatal().Msg("must specify two integer values")
	}

	// n is the number of possible characters, such as 26 lowercase letters,
	// or 10 numeric digits
	n, err := strconv.ParseUint(os.Args[1], 10, 64)
	if err != nil {
		log.Fatal().Err(err).Msgf("could not parse %q as integer", os.Args[1])
	}

	// l is the length of the password
	l, err := strconv.ParseUint(os.Args[2], 10, 64)
	if err != nil {
		log.Fatal().Err(err).Msgf("could not parse %q as integer", os.Args[2])
	}

	// cache the value of n^l
	value := math.Pow(float64(n), float64(l))

	// tp will add commas to indicate thousands
	tp := message.NewPrinter(language.English)

	// vs is the value string, e.g., 1,234,567 instead of 1234567
	vs := tp.Sprintf("%d", int64(value))

	// e is the entropy bits
	e := math.Log2(value)

	// debugging info for checking the math
	log.Debug().Msgf("%d ^ %d = %s", n, l, vs)
	log.Debug().Msgf("log2( %s ) = %f", vs, e)

	// print the entropy bits
	log.Info().Msgf("%f", e)
}
