package main

import (
	"errors"
	"fmt"
	"math/rand"
)

type Alien struct {
	ID        int
	Moves     int
	Location  *City
	IsTrapped bool
}

func NewAlien(id int) *Alien {
	return &Alien{ID: id}
}

func (a *Alien) InvadeRandomEmptyCity(inWorld *World) (*City, error) {
	var startingCity *City
	inWorld.Range(func(_ string, city *City) bool {
		if len(city.Residents) == 0 {
			startingCity = city
			return false
		}

		return true
	})

	if startingCity == nil {
		return nil, errors.New("Nowhere left to invade, all cities are occupied")
	}

	a.Location = startingCity
	a.Location.Residents = append(a.Location.Residents, a)

	return a.Location, nil
}

func (a *Alien) Move() error {
	var possibleDestinations []*City
	for _, city := range a.Location.NeighboringCities {
		if city != nil {
			possibleDestinations = append(possibleDestinations, city)
		}
	}

	if len(possibleDestinations) == 0 {
		a.IsTrapped = true
		return fmt.Errorf("Alien %d is trapped in %s", a.ID, a.Location.Name)
	}

	destination := possibleDestinations[rand.Intn(len(possibleDestinations))]
	a.Location.Evict(a)
	a.Location = destination
	a.Location.Residents = append(a.Location.Residents, a)

	a.Moves++
	if a.Moves >= MaxAlienMoves {
		a.IsTrapped = true
	}

	return nil
}
