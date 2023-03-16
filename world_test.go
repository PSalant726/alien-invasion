package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewWorld(t *testing.T) {
	testWorld := NewWorld()
	assert.NotNil(t, testWorld)
	assert.Len(t, testWorld.Cities, 0)
}

func TestWorld_Delete(t *testing.T) {
	testWorld := NewWorld()
	city := NewCity("City")
	testWorld.Store(city)

	sanity, ok := testWorld.Load(*city)
	assert.Equal(t, city, sanity)
	assert.True(t, ok)

	testWorld.Delete(city)

	deleted, ok := testWorld.Load(*city)
	assert.Nil(t, deleted)
	assert.False(t, ok)
}

func TestWorld_Load(t *testing.T) {
	testWorld := NewWorld()
	city := NewCity("City")

	before, ok := testWorld.Load(*city)
	assert.Nil(t, before)
	assert.False(t, ok)

	testWorld.Store(city)

	after, ok := testWorld.Load(*city)
	assert.Equal(t, city, after)
	assert.True(t, ok)
}

func TestWorld_LoadOrStore(t *testing.T) {
	testWorld := NewWorld()
	city1 := NewCity("City1")

	before, ok := testWorld.Load(*city1)
	assert.Nil(t, before)
	assert.False(t, ok)

	actual, ok := testWorld.LoadOrStore(city1)
	assert.Equal(t, *city1, *actual)
	assert.False(t, ok)

	city2 := NewCity("City2")
	testWorld.Store(city2)

	actual, ok = testWorld.LoadOrStore(city2)
	assert.Equal(t, city2, actual)
	assert.True(t, ok)
}

func TestWorld_Store(t *testing.T) {
	testWorld := NewWorld()
	city := NewCity("City")

	sanity, ok := testWorld.Load(*city)
	assert.Nil(t, sanity)
	assert.False(t, ok)

	testWorld.Store(city)

	stored, ok := testWorld.Load(*city)
	assert.Equal(t, city, stored)
	assert.True(t, ok)
}

func TestWorld_EstablishCity(t *testing.T) {
	testWorld := NewWorld()

	err := testWorld.EstablishCity(" ")
	assert.NotNil(t, err)

	err = testWorld.EstablishCity("NYC north")
	assert.NotNil(t, err)
	err = testWorld.EstablishCity("NYC north=Toronto south")
	assert.NotNil(t, err)
	err = testWorld.EstablishCity("NYC north=Toronto south west=Chicago")
	assert.NotNil(t, err)

	err = testWorld.EstablishCity("Central north=North south=South east=East west=West")
	assert.Nil(t, err)

	central := NewCity("Central")

	actual, ok := testWorld.Load(*central)
	assert.Equal(t, central.Name, actual.Name)
	assert.True(t, ok)

	for _, neighbor := range actual.NeighboringCities {
		for _, reverseNeighbor := range neighbor.NeighboringCities {
			if reverseNeighbor != nil {
				assert.Equal(t, actual.Name, reverseNeighbor.Name)
			}
		}
	}
}
