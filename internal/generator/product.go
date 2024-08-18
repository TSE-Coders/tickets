package generator

import "math/rand"

var Products []string

func GetRandomProduct() string {
	r := rand.Intn(len(Products))
	return Products[r]
}
