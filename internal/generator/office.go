package generator

import "math/rand"

var Offices []string

func GetRandomOffice() string {
	return Offices[rand.Intn(len(Offices))]
}
