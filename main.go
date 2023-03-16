package main

import (
	"errors"
	"flag"
	"log"
	"os"
)

var (
	aliens   = flag.Int("aliens", 2, "The amount of violent Alien invaders to unleash upon the world.")
	worldMap = flag.String("world-map", "", "The path to the input file.\n\n"+
		"The input file must include one City per line, connect any two Cities with at most a single route, "+
		"and provide at least as many Cities as there are Aliens. The City's name must come first, "+
		"followed by up to four directional key/value pairs, all separated by a single space. "+
		"Inter-City routes may be provided in any order, but the only valid keys are the four directions shown in the following example file:\n\n"+
		"\tCity1 north=City2 south=City3 east=City4 west=City5\n"+
		"\tCity2 south=City1\n"+
		"\tCity3 north=City1\n"+
		"\tCity4 west=City1\n"+
		"\tCity5 east=City1",
	)

)

func main() {
	flag.Parse()
	if err := validateFlags(); err != nil {
		log.Fatalf("Invalid option value provided: %s", err)
	}

}

func validateFlags() error {
	if *worldMap == "" {
		return errors.New("world-map is required")
	}

	if _, err := os.Stat(*worldMap); err != nil {
		return fmt.Errorf("Invalid value provided for world-map: %s", err)
	}

	return nil
}
