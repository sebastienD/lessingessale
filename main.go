package main

import "log"

func main() {
	log.Println("plop")
}

type Video struct {
}

type Endpoint struct {
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

type Entry2 map[int]int
