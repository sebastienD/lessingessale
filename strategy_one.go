package main

import (
	"log"
)

func strategyOne(dc DataCenter) []*Cache {
	endpoints := dc.Endpoints
	sortedEndpoints := sortEndpointsByRequestsNumber(endpoints)

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
	return dc.Caches
}
