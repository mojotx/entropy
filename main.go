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

	N, err := strconv.ParseUint(os.Args[1], 10, 64)
	if err != nil {
		log.Fatal().Err(err).Msgf("could not parse %q as integer", os.Args[1])
	}

	L, err := strconv.ParseUint(os.Args[2], 10, 64)
	if err != nil {
		log.Fatal().Err(err).Msgf("could not parse %q as integer", os.Args[2])
	}

	product := math.Pow(float64(N), float64(L))
	// p := message.NewPrinter(language.English)
	// s := p.Sprintf("%d\n", 1000)

	printer := message.NewPrinter(language.English)
	prodString := printer.Sprintf("%d", int64(product))

	log.Info().Msgf("%d ^ %d = %s", N, L, prodString)
	log.Info().Msgf("log2( %s ) = %f", prodString, math.Log2(product))
}
