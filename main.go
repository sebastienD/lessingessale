package main

import "log"

func main() {
	log.Println("plop")
}

type Entry struct {
	Videos      int
	Endpoints   int
	RequestDesc int
	NbCache     int
	CacheSize   int
	VideoSizes  []int
}
