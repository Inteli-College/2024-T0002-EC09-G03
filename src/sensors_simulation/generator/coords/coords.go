package coords

import (
	"math/rand"
	"sync"
)

var ExistingCoords = make(map[string]map[float64]bool)
var MutexMap = sync.RWMutex{}

func GenerateCoords(typeName *string) (*float64, *float64) {
	baseX := -46.633308
	baseY := -23.550520

	maxOffsetX, minOffsetX := 0.3, -0.2
	maxOffsetY, minOffsetY := 0.1, -0.2

	defer MutexMap.Unlock()

	for {
		offsetX := minOffsetX + (rand.Float64() * (maxOffsetX - minOffsetX))
		offsetY := minOffsetY + (rand.Float64() * (maxOffsetY - minOffsetY))

		coordsX := baseX + offsetX
		coordsY := baseY + offsetY

		MutexMap.Lock()
		if ExistingCoords[*typeName] == nil {
			ExistingCoords[*typeName] = make(map[float64]bool)
		}
		MutexMap.Unlock()

		MutexMap.RLock()
		if ExistingCoords[*typeName][coordsX+coordsY] {
			MutexMap.RUnlock()
			continue
		} else {
			MutexMap.RUnlock()
			MutexMap.Lock()
			ExistingCoords[*typeName][coordsX+coordsY] = true
			return &coordsX, &coordsY
		}
	}
}
