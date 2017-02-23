package main

import "sort"

type RequestBySize []*Request

func (a RequestBySize) Len() int           { return len(a) }
func (a RequestBySize) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a RequestBySize) Less(i, j int) bool { return a[i].Nb < a[j].Nb }

type RequestByVideoSize []*Request

func (a RequestByVideoSize) Len() int           { return len(a) }
func (a RequestByVideoSize) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a RequestByVideoSize) Less(i, j int) bool { return a[i].Video.Size > a[j].Video.Size }

func (ep *Endpoint) GetRequestsSortedByNb() []*Request {
	sortedRequests := ep.Requests
	sort.Sort(RequestBySize(sortedRequests))
	return sortedRequests
}

func (ep *Endpoint) GetRequestsSortedByVideoSize() []*Request {
	sortedRequests := ep.Requests
	sort.Sort(RequestByVideoSize(sortedRequests))
	return sortedRequests
}

type CacheByLatency []*Cache

func (a CacheByLatency) Len() int           { return len(a) }
func (a CacheByLatency) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a CacheByLatency) Less(i, j int) bool { return a[i].Latency < a[j].Latency }

func (ep *Endpoint) GetCacheSortedByLatency() []*Cache {
	sortedCaches := ep.Caches
	sort.Sort(CacheByLatency(sortedCaches))
	return sortedCaches
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
