package main


type CacheWithEndpoint struct {
	Endpoint *Endpoint
	Cache *Cache
}



func strategyOne(dc DataCenter) []*Cache {
	endpoints := dc.Endpoints
	sortedEndpoints := sortEndpointsByRequestsNumber(endpoints)

	emptyCaches := map[*Cache]CacheWithEndpoint{}

	for _, e := range sortedEndpoints {
		requests := e.GetRequestsSortedByNb() // TODO: or by size of video ?
		for _, r := range requests {
			inserted := false
			caches := e.GetCacheSortedByLatency()
			video := r.Video

			for i := 0; i < len(caches) && !inserted; i++ {
				// Insert the best video in cache
				cache := caches[0]
				inserted = cache.Insert(video)
			}
		}
	}

	for _, e := range sortedEndpoints {
		for _, c := range e.Caches {
			if len(c.Videos) {
				emptyCaches[c]=CacheWithEndpoint{
					Endpoint: e,
					Cache: c,
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
