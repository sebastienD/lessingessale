package main

import "log"

func main() {
	log.Println("plop")
}

type Video struct {
}

// REPETABLE
type Endpoint struct {
	index          int
	latency        int
	connectedCache int
}

type DataCenter struct {
}

type Request struct {
}

type Cache struct {
}

type Entry struct {
	Videos      int
	Endpoints   int
	RequestDesc int
	NbCache     int
	CacheSize   int
}

type VideosSize map[int]int
