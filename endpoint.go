package main

import "sort"

type ByVideoSize []Request

func (a ByVideoSize) Len() int           { return len(a) }
func (a ByVideoSize) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByVideoSize) Less(i, j int) bool { return a[i].Nb < a[j].Nb }


func (ep *Endpoint) GetRequestsSortedByNb() []Request {
	return sort.Sort(ep.Requests)
}

func (ep *Endpoint) GetCacheSortedByLatency() []Cache {
	return  []Cache{}
}


func sortEndpointsByRequestsNumber(endpoints []Endpoint) []Endpoint {
	return endpoints
}