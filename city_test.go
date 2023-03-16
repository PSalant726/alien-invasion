package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCity(t *testing.T) {
	foo := NewCity("Foo")
	assert.NotNil(t, foo)
	assert.Equal(t, "Foo", foo.Name)
	assert.Len(t, foo.Residents, 0)
	assert.Len(t, foo.NeighboringCities, 4)
}

func TestCity_Destroy(t *testing.T) {
	testWorld := NewWorld()

	testWorld.EstablishCity("Jail")
	jail, _ := testWorld.Load(*NewCity("Jail"))
	jail.Residents = append(jail.Residents, &Alien{ID: 0, IsTrapped: false}, &Alien{ID: 1, IsTrapped: false})

	jail.Destroy(testWorld)
	assert.NotContains(t, testWorld.Cities, "Jail", "Jail should have been destroyed")

	for _, alien := range jail.Residents {
		assert.Truef(t, alien.IsTrapped, "Alien %d should be trapped in %s", alien.ID, jail.Name)
	}

	testWorld.EstablishCity("Vanish north=North")
	vanish, _ := testWorld.Load(*NewCity("Vanish"))
	north, _ := testWorld.Load(*NewCity("North"))

	vanish.Destroy(testWorld)
	assert.NotContains(t, testWorld.Cities, "Vanish", "Vanish should have been destroyed")
	assert.Contains(t, testWorld.Cities, "North", "North should not have been destroyed")

	for _, neighbor := range north.NeighboringCities {
		assert.Nil(t, neighbor)
	}
}

func TestCity_Evict(t *testing.T) {
	foo := NewCity("Foo")
	alien0 := &Alien{ID: 0}
	alien1 := &Alien{ID: 1}
	alien2 := &Alien{ID: 2}
	foo.Residents = append(foo.Residents, alien0, alien1, alien2)

	foo.Evict(alien1)

	assert.Contains(t, foo.Residents, alien0)
	assert.NotContains(t, foo.Residents, alien1)
	assert.Contains(t, foo.Residents, alien2)
}
