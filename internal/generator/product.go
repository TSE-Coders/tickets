package generator

import "math/rand"

var Products SupportedProducts

type SupportedProducts struct {
	productList []*string
	Monitors    string
	SynRUM      string
	APM         string
	DBM         string
}

func (s *SupportedProducts) addProduct(loc *string, region string) {
	*loc = region
	s.productList = append(s.productList, loc)
}

func (s SupportedProducts) GetRandom() string {
	return *s.productList[rand.Intn(len(s.productList))]
}
