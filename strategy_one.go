package main

import (
	"log"
)

type CacheWithEndpoint struct {
	Endpoint *Endpoint
	Cache    *Cache
}

func strategyOne(dc DataCenter) []*Cache {
	endpoints := dc.Endpoints
	sortedEndpoints := sortEndpointsByRequestsNumber(endpoints)

	emptyCaches := map[*Cache]CacheWithEndpoint{}

	for _, e := range sortedEndpoints {
		log.Println("Endpoint", e.index)
		requests := e.GetRequestsSortedByNb() // TODO: or by size of video ?
		for _, r := range requests {
			log.Println("Requests", r.Nb)
			inserted := false
			caches := e.GetCacheSortedByLatency()
			video := r.Video
			log.Println("Video", video.index)

			for i := 0; i < len(caches) && !inserted; i++ {
				// Insert the best video in cache
				cache := caches[i]
				inserted = cache.Insert(video)
				if inserted {
					log.Println("Inserted", video.index)
				}
			}
		}
	}

	for _, e := range sortedEndpoints {
		for _, c := range e.Caches {
			log.Println("Video", len(c.Videos))
			if len(c.Videos) == 0 {
				emptyCaches[c] = CacheWithEndpoint{
					Endpoint: e,
					Cache:    c,
				}
			}
		}
	}

	for cache, assoc := range emptyCaches {
		requests := assoc.Endpoint.GetRequestsSortedByVideoSize()
		for _, r := range requests {
			cache.Insert(r.Video)
		}
	}

	return dc.Caches
}
