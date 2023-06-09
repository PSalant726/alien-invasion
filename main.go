package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
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
)

func main() {
	flag.Parse()
	if err := validateFlags(); err != nil {
		log.Fatalf("Invalid option value provided: %s", err)
	}

	world, err := buildWorld()
	if err != nil {
		log.Fatalf("Invalid city map provided: %s", err)
	}

	invaders, err := invadeCities(world)
	if err != nil {
		log.Fatalf("Alien invasion failed: %s", err)
	}

	trappedAliens := 1 // If only a single Alien remains, then they have nobody to fight!
	for trappedAliens < *aliens {
		destroyCities(world)
		trappedAliens += moveAliens(invaders)
	}

	log.Println("All Aliens have been trapped or destroyed")

	printWorld(world)
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

func buildWorld() (*World, error) {
	input, err := os.Open(*worldMap)
	if err != nil {
		return nil, fmt.Errorf("Failed to open input file: %s", err)
	}
	defer input.Close()

	mapScanner := bufio.NewScanner(input)
	mapScanner.Split(bufio.ScanLines)

	world := NewWorld()
	for mapScanner.Scan() {
		err := world.EstablishCity(mapScanner.Text())
		if err != nil {
			return nil, fmt.Errorf("Failed to parse city details: %s", err)
		}
	}
	if err := mapScanner.Err(); err != nil {
		return nil, fmt.Errorf("Failed to parse world map: %s", err)
	}

	return world, nil
}

func invadeCities(inWorld *World) ([]*Alien, error) {
	var wg sync.WaitGroup
	var errs int

	invaders := make([]*Alien, *aliens)
	for i := 0; i < *aliens; i++ {
		wg.Add(1)
		go func(alienID int) {
			defer wg.Done()

			invader := NewAlien(alienID)
			if city, err := invader.InvadeRandomEmptyCity(inWorld); err != nil {
				log.Printf("Alien %d failed to invade a city: %s", alienID, err)
				errs++
			} else {
				invaders[alienID] = invader
				log.Printf("Alien %d has invaded %s!", invader.ID, city.Name)
			}
		}(i)
	}

	wg.Wait()

	if errs > 0 {
		return nil, errors.New("There are not enough Cities for all Aliens to invade")
	}

	return invaders, nil
}

func destroyCities(inWorld *World) {
	var wg sync.WaitGroup
	inWorld.Range(func(_ string, city *City) bool {
		wg.Add(1)
		go func(city *City) {
			defer wg.Done()

			if len(city.Residents) > 1 {
				city.Destroy(inWorld)
				log.Printf(
					"%s has been destroyed by Alien %d and Alien %d!",
					city.Name,
					city.Residents[0].ID,
					city.Residents[1].ID,
				)
			}
		}(city)

		return true
	})

	wg.Wait()
}

func moveAliens(invaders []*Alien) int {
	var trappedAliens int
	var wg sync.WaitGroup

	for _, alien := range invaders {
		wg.Add(1)
		go func(alien *Alien) {
			defer wg.Done()

			if alien.IsTrapped {
				trappedAliens++
				return
			}

			if err := alien.Move(); err != nil {
				log.Printf("Failed to move Alien: %s", err)
			}
		}(alien)
	}

	wg.Wait()

	return trappedAliens
}

func printWorld(world *World) {
	log.Println("The current state of the world is:")

	world.Range(func(cityName string, city *City) bool {
		cityOut := cityName + " "

		for i, neighbor := range city.NeighboringCities {
			if neighbor != nil {
				var direction string
				switch i {
				case 0:
					direction = "north"
				case 1:
					direction = "south"
				case 2:
					direction = "east"
				case 3:
					direction = "west"
				}

				cityOut += fmt.Sprintf("%s=%s ", direction, neighbor.Name)
			}
		}

		fmt.Println(strings.TrimSpace(cityOut))
		return true
	})
}
