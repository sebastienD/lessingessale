package main

import "sort"

type RequestBySize []*Request

func (a RequestBySize) Len() int           { return len(a) }
func (a RequestBySize) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a RequestBySize) Less(i, j int) bool { return a[i].Nb < a[j].Nb }


func (ep *Endpoint) GetRequestsSortedByNb() []*Request {
	sortedRequests := ep.Requests
	sort.Sort(RequestBySize(sortedRequests))
	return sortedRequests
}

func (ep *Endpoint) GetCacheSortedByLatency() []*Cache {
	return ep.Caches
}

type EndpointsByNbRequests []*Endpoint

func (a EndpointsByNbRequests) Len() int           { return len(a) }
func (a EndpointsByNbRequests) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a EndpointsByNbRequests) Less(i, j int) bool { return len(a[i].Requests) < len(a[j].Requests) }


func sortEndpointsByRequestsNumber(endpoints []*Endpoint) []*Endpoint {
	sortedEndpoints := endpoints
	sort.Sort(EndpointsByNbRequests(sortedEndpoints))
	return sortedEndpoints
}