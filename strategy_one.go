package main

func startegyOne(dc DataCenter) []Cache {
	endpoints := []Endpoint{}
	sortedEndpoints := sortEndpointsByRequestsNumber(endpoints)

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
}
