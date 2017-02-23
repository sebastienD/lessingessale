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
	//IN  string = "me_at_the_zoo.in"
	//OUT string = "me_at_the_zoo.out"
	//IN  string = "kittens.in"
	//OUT string = "kittens.out"
	//IN  string = "trending_today.in"
	//OUT string = "trending_today.out"
	IN  string = "videos_worth_spreading.in"
	OUT string = "videos_worth_spreading.out"
)

func parseFile(in string) DataCenter {
	f, err := os.Open(in)
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
		Videos:      []*Video{},
		Endpoints:   []*Endpoint{},
		RequestDesc: nbRequestDesc,
		Requests:    []*Request{},
		Caches:      []*Cache{},
	}
	fmt.Printf("%+v\n", entry)

	grosseMap := map[int]*Cache{}
	for i := 0; i < nbCaches; i++ {
		cache := &Cache{
			Capacity: cachSize,
			Videos:   []*Video{},
		}
		entry.Caches = append(entry.Caches, cache)
		grosseMap[i] = cache
	}

	// second line
	scanner.Scan()
	for i, o := range strings.Split(scanner.Text(), " ") {
		entry.Videos = append(entry.Videos, &Video{
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

		endpoint := &Endpoint{latencyToDatacenter: convert2int(split[0])}

		for i := 0; i < nbLatencies; i++ {
			scanner.Scan()
			split = strings.Split(scanner.Text(), " ")
			cache, _ := grosseMap[i]
			cache.Latency = convert2int(split[1])
			endpoint.Caches = append(endpoint.Caches, cache)
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

		r := &Request{
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

func writeOutFile(caches []*Cache, out string) {
	f, err := os.Create(out)
	check(err)
	defer f.Close()

	f.WriteString(fmt.Sprintf("%d\n", len(caches)))
	for i, c := range caches {
		f.WriteString(fmt.Sprintf("%d", i))
		fmt.Println(len(c.Videos))
		for _, v := range c.Videos {
			f.WriteString(fmt.Sprintf(" %d", v.index))
		}
		f.WriteString("\n")
	}
	f.Sync()
}

func main() {
	files := []string{"me_at_the_zoo", "kittens", "trending_today", "videos_worth_spreading"}
	for _, f := range files {
		dc := parseFile(f + ".in")
		caches := strategyOne(dc)
		writeOutFile(caches, f+".out")
	}
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
	//Latencies           []Latency
	Requests []*Request
	Caches   []*Cache
}

type Latency map[int]int

type Request struct {
	Video    *Video
	Endpoint *Endpoint
	Nb       int
}

type Cache struct {
	Capacity int
	Videos   []*Video
	Latency  int
}

type DataCenter struct {
	Videos      []*Video
	Endpoints   []*Endpoint
	Caches      []*Cache
	RequestDesc int
	Requests    []*Request
}
