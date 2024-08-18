package generator

import "math/rand"

var Regions []string

func GetRandomRegion() string {
	return Regions[rand.Intn(len(Regions))]
}
