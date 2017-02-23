package main

func StartegyOne(dc DataCenter) {
	endpoints := []Endpoint{}
	sortedEndpoints := sortEndpointsByRequestsNumber(endpoints)

	for _, e := range endpoints {
		requests := e.GetRequestsSortedByNb() // TODO: or by size of video ?
		for _, r := range requests {
			inserted :=  false
			caches := e.GetCacheSortedByLatency()
			video := r.Video()

			for i:= 0; i < len(caches) && !inserted; i++ {
				// Insert the best video in cache
				cache := caches[0]
				if /*cache.Capacity >  video.Size */true {
					cache.Insert(video)
					inserted = true;
				}
			}
		}
	}
}
