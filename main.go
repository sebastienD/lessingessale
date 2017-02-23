package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

const (
	IN  string = "me_at_the_zoo_example.in"
	OUT string = "me_at_the_zoo_example.out"
)

func parseFile() DataCenter {
	f, err := os.Open(IN)
	check(err)
	defer f.Close()

	// read params
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	line := strings.Split(scanner.Text(), " ")

	// first line
	//nbVideos := convert2int(line[0])
	nbEndpoints := convert2int(line[1])
	nbRequestDesc := convert2int(line[2])
	nbCaches := convert2int(line[3])
	cachSize := convert2int(line[4])

	entry := DataCenter{
		Videos:      []Video{},
		Endpoints:   []Endpoint{},
		RequestDesc: nbRequestDesc,
		Requests:    []Request{},
		Caches:      []Cache{},
	}
	fmt.Printf("%+v\n", entry)

	for i := 0; i < nbCaches; i++ {
		entry.Caches = append(entry.Caches, Cache{
			Capacity: cachSize,
		})
	}

	// second line
	scanner.Scan()
	for i, o := range strings.Split(scanner.Text(), " ") {
		entry.Videos = append(entry.Videos, Video{
			index: i,
			Size:  convert2int(o),
		})
	}

	// second line
	//  Parse endpoints
	for n := 0; n < nbEndpoints; n++ {
		scanner.Scan()
		split := strings.Split(scanner.Text(), " ")
		nbLatencies := convert2int(split[1])

		endpoint := Endpoint{latencyToDatacenter: convert2int(split[0])}
		for i := 0; i < nbLatencies; i++ {
			scanner.Scan()
			split = strings.Split(scanner.Text(), " ")
			latence := Latency{}
			latence[convert2int(split[0])] = convert2int(split[1])
			endpoint.Latencies = append(endpoint.Latencies, latence)
		}

		entry.Endpoints = append(entry.Endpoints, endpoint)
	}

	fmt.Println(entry.Endpoints)

	for k := 0; k < nbRequestDesc; k++ {
		scanner.Scan()
		split := strings.Split(scanner.Text(), " ")

		endpoint := entry.Endpoints[convert2int(split[1])]
		video := entry.Videos[convert2int(split[0])]
		nbRequests := convert2int(split[2])

		r := Request{
			Endpoint: endpoint,
			Nb:       nbRequests,
			Video:    video,
		}

		entry.Requests = append(entry.Requests, r)
		endpoint.Requests = append(endpoint.Requests, r)
	}

	fmt.Println(entry)

	return entry
}

func writeOutFile(caches []Cache) {
	f, err := os.Create(OUT)
	check(err)
	defer f.Close()

	f.WriteString(fmt.Sprintf("%d\n", len(caches)))
	for i, c := range caches {
		f.WriteString(fmt.Sprintf("%d", i))
		for j, v := range c.Videos {
			f.WriteString(fmt.Sprintf(" %d", v.index))
		}
		f.WriteString("\n")
	}
	f.Sync()
}

func main() {
	dc := parseFile()
	caches := startegyOne(dc)
	writeOutFile(caches)
}

func convert2int(val string) int {
	i, err := strconv.Atoi(val)
	check(err)
	return i
}

type Video struct {
	index int
	Size  int
}

// REPETABLE
type Endpoint struct {
	index               int
	latencyToDatacenter int
	Latencies           []Latency
	Requests            []Request
}

type Latency map[int]int

type Request struct {
	Video    Video
	Endpoint Endpoint
	Nb       int
}

type Cache struct {
	Capacity int
	Videos   []Video
}

type DataCenter struct {
	Videos      []Video
	Endpoints   []Endpoint
	Caches      []Cache
	RequestDesc int
	Requests    []Request
}
