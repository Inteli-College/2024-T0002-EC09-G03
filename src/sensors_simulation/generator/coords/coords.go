package coords

import (
	"math/rand"
	"sync"
)

var existingCoords = make(map[string]map[float64]bool)
var mMap = sync.RWMutex{}

func GenerateCoords(typeName *string) (*float64, *float64) {
	baseX := -46.633308
	baseY := -23.550520

	maxOffsetX, minOffsetX := 0.3, -0.2
	maxOffsetY, minOffsetY := 0.1, -0.2

	defer mMap.Unlock()

	for {
		offsetX := minOffsetX + (rand.Float64() * (maxOffsetX - minOffsetX))
		offsetY := minOffsetY + (rand.Float64() * (maxOffsetY - minOffsetY))

		coordsX := baseX + offsetX
		coordsY := baseY + offsetY

		mMap.Lock()
		if existingCoords[*typeName] == nil {
			existingCoords[*typeName] = make(map[float64]bool)
		}
		mMap.Unlock()

		mMap.RLock()
		if existingCoords[*typeName][coordsX+coordsY] {
			mMap.RUnlock()
			continue
		} else {
			mMap.RUnlock()
			mMap.Lock()
			existingCoords[*typeName][coordsX+coordsY] = true
			return &coordsX, &coordsY
		}
	}
}
