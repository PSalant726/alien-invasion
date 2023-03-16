package main

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

func (c *City) Destroy() {
	for _, alien := range c.Residents {
		alien.IsTrapped = true
	}

	for _, neighbor := range c.NeighboringCities {
		if neighbor != nil {
			for i, city := range neighbor.NeighboringCities {
				if city != nil && city.Name == c.Name {
					neighbor.NeighboringCities[i] = nil
				}
			}
		}
	}

	world.Delete(c)
}

// Evict removes the specified Alien from the list of the City's residents.
func (c *City) Evict(alien *Alien) {
	for i, resident := range c.Residents {
		if resident == alien {
			c.Residents = append(c.Residents[:i], c.Residents[i+1:]...)
		}
	}
}
