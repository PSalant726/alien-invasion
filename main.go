package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
)

const MaxAlienMoves = 10000

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

	world = make(map[string]*City)
)

func main() {
	flag.Parse()
	if err := validateFlags(); err != nil {
		log.Fatalf("Invalid option value provided: %s", err)
	}

	if err := buildWorld(); err != nil {
		log.Fatalf("Invalid city map provided: %s", err)
	}

	invaders, err := invadeCities()
	if err != nil {
		log.Fatalf("Alien invasion failed: %s", err)
	}

	trappedAliens := 1 // If only a single Alien remains, then they have nobody to fight!
	for trappedAliens < *aliens {
		for _, city := range world {
			if len(city.Residents) > 1 {
				city.Destroy()
				log.Printf(
					"%s has been destroyed by Alien %d and Alien %d!",
					city.Name,
					city.Residents[0].ID,
					city.Residents[1].ID,
				)
			}
		}

		for _, alien := range invaders {
			if alien.IsTrapped {
				trappedAliens++
				continue
			}

			if err := alien.Move(); err != nil {
				log.Printf("Failed to move Alien: %s", err)
			}
		}
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

func buildWorld() error {
	input, err := os.Open(*worldMap)
	if err != nil {
		return fmt.Errorf("Failed to open input file: %s", err)
	}
	defer input.Close()

	mapScanner := bufio.NewScanner(input)
	mapScanner.Split(bufio.ScanLines)

	for mapScanner.Scan() {
		err := EstablishCity(world, mapScanner.Text())
		if err != nil {
			return fmt.Errorf("Failed to parse city details: %s", err)
		}
	}
	if err := mapScanner.Err(); err != nil {
		return fmt.Errorf("Failed to parse world map: %s", err)
	}

	return nil
}

func invadeCities() ([]*Alien, error) {
	invaders := make([]*Alien, *aliens)
	for i := 0; i < *aliens; i++ {
		invader := NewAlien(i)
		if city, err := invader.InvadeRandomEmptyCity(); err != nil {
			return nil, fmt.Errorf("Alien %d failed to invade a city: %s", i, err)
		} else {
			log.Printf("Alien %d has invaded %s!", invader.ID, city.Name)
		}

		invaders[i] = invader
	}

	return invaders, nil
}
