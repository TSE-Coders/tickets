package generator

import "math/rand"

var Regions SupportedRegions

type SupportedRegions struct {
	regionList []*string
	NYC        string
	Boston     string
	Denver     string
}

func (r *SupportedRegions) addRegion(loc *string, region string) {
	*loc = region
	r.regionList = append(r.regionList, loc)
}

func (r SupportedRegions) GetRandom() string {
	return *r.regionList[rand.Intn(len(r.regionList))]
}
