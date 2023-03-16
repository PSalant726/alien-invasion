package main

import (
	"errors"
)

type Alien struct {
	ID        int
	Location  *City
}

func NewAlien(id int) *Alien {
	return &Alien{ID: id}
}

func (a *Alien) InvadeRandomEmptyCity() (*City, error) {
	var startingCity *City
	for _, city := range world {
		if len(city.Residents) == 0 {
			startingCity = city
			break
		}
	}

	if startingCity == nil {
		return nil, errors.New("Nowhere left to invade, all cities are occupied")
	}

	a.Location = startingCity
	a.Location.Residents = append(a.Location.Residents, a)

	return a.Location, nil
}
