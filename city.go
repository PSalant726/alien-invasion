package main

import (
	"errors"
	"fmt"
	"strings"
)

type City struct {
	Name      string
	Residents []*Alien

	// A city may have at most four neighboring cities. They
	// are stored here in order of North, South, East, West.
	NeighboringCities []*City
}

func NewCity(name string) *City {
	return &City{
		Name:              name,
		Residents:         []*Alien{},
		NeighboringCities: make([]*City, 4),
	}
}

// EstablishCity adds a new City to the world, and updates any
// neighboring Cities to become aware of the newly established City.
// It expects `line` to be formatted as specified by the `world-map`
// parameter's help text.
func EstablishCity(world map[string]*City, line string) error {
	cityDetails := strings.Split(line, " ")
	if len(cityDetails) == 0 {
		return errors.New("No city details provided")
	}

	city := NewCity(cityDetails[0])
	for _, directionAndDestination := range cityDetails[1:] {
		road := strings.Split(directionAndDestination, "=")
		if len(road) != 2 {
			return fmt.Errorf("Invalid directional key/value pair provided for city %s: %s", city.Name, road)
		}

		destination := world[road[1]]
		if destination == nil {
			destination = NewCity(road[1])
			world[destination.Name] = destination
		}

		switch road[0] {
		case "north":
			if destination.NeighboringCities[1] != nil && destination.NeighboringCities[1].Name != city.Name {
				return fmt.Errorf(
					"City %s routes North to %s, which already has a road heading South to %s",
					city.Name,
					destination.Name,
					destination.NeighboringCities[1].Name,
				)
			}

			city.NeighboringCities[0] = destination
			destination.NeighboringCities[1] = city

		case "south":
			if destination.NeighboringCities[0] != nil && destination.NeighboringCities[0].Name != city.Name {
				return fmt.Errorf(
					"City %s routes South to %s, which already has a road heading North to %s",
					city.Name,
					destination.Name,
					destination.NeighboringCities[0].Name,
				)
			}

			city.NeighboringCities[1] = destination
			destination.NeighboringCities[0] = city

		case "east":
			if destination.NeighboringCities[3] != nil && destination.NeighboringCities[3].Name != city.Name {
				return fmt.Errorf(
					"City %s routes East to %s, which already has a road heading West to %s",
					city.Name,
					destination.Name,
					destination.NeighboringCities[3].Name,
				)
			}

			city.NeighboringCities[2] = destination
			destination.NeighboringCities[3] = city

		case "west":
			if destination.NeighboringCities[2] != nil && destination.NeighboringCities[2].Name != city.Name {
				return fmt.Errorf(
					"City %s routes West to %s, which already has a road heading East to %s",
					city.Name,
					destination.Name,
					destination.NeighboringCities[2].Name,
				)
			}

			city.NeighboringCities[3] = destination
			destination.NeighboringCities[2] = city

		default:
			return fmt.Errorf("Invalid direction provided from city %s: %s", city.Name, road[0])
		}
	}

	world[city.Name] = city
	return nil
}
