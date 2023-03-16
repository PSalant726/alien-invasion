package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAlien(t *testing.T) {
	alien := NewAlien(5)
	assert.NotNil(t, alien)
	assert.Equal(t, 5, alien.ID)
}

func TestAlien_InvadeRandomEmptyCity(t *testing.T) {
	testWorld := NewWorld()
	testWorld.EstablishCity("City1")
	testWorld.EstablishCity("City2")
	testWorld.EstablishCity("City3")

	testWorld.Range(func(_ string, city *City) bool {
		assert.Len(t, city.Residents, 0)
		return true
	})

	alien0 := NewAlien(0)

	invaded, err := alien0.InvadeRandomEmptyCity(testWorld)
	assert.NotNil(t, invaded)
	assert.Nil(t, err)
	assert.NotNil(t, alien0.Location)
	assert.Len(t, alien0.Location.Residents, 1)
	assert.Equal(t, alien0, alien0.Location.Residents[0])

	aliens := []*Alien{NewAlien(1), NewAlien(2)}
	for _, alien := range aliens {
		_, err := alien.InvadeRandomEmptyCity(testWorld)
		assert.Nil(t, err)
		assert.NotNil(t, alien.Location)
		assert.Len(t, alien.Location.Residents, 1)
		assert.Equal(t, alien, alien.Location.Residents[0])
	}

	alien3 := NewAlien(3)
	failed, err := alien3.InvadeRandomEmptyCity(testWorld)
	assert.Nil(t, failed)
	assert.NotNil(t, err)
	assert.Nil(t, alien3.Location)
}

func TestAlien_Move(t *testing.T) {
	testWorld := NewWorld()
	testWorld.EstablishCity("North south=South")
	testWorld.EstablishCity("South north=North")
	north, _ := testWorld.Load(*NewCity("North"))
	south, _ := testWorld.Load(*NewCity("South"))
	alien := NewAlien(0)
	alien.Location = north
	north.Residents = append(north.Residents, alien)

	assert.Zero(t, alien.Moves)
	err := alien.Move()

	assert.Nil(t, err)
	assert.Equal(t, 1, alien.Moves)
	assert.NotContains(t, north.Residents, alien)
	assert.Contains(t, south.Residents, alien)

	testWorld.EstablishCity("Island")
	island, _ := testWorld.Load(*NewCity("Island"))
	alien1 := NewAlien(1)
	alien1.Location = island
	island.Residents = append(island.Residents, alien1)

	assert.Zero(t, alien1.Moves)
	err = alien1.Move()

	assert.NotNil(t, err)
	assert.Zero(t, alien1.Moves)
	assert.Contains(t, island.Residents, alien1)
	assert.True(t, alien1.IsTrapped)
}
