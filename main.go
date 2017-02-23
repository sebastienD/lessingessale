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

func main() {
	f, err := os.Open("me_at_the_zoo.in")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	// read params
	scanner.Scan()
	line := strings.Split(scanner.Text(), " ")

	// convert params to []int
	entry := Entry{
		Videos:      convert2int(line[0]),
		Endpoints:   convert2int(line[1]),
		RequestDesc: convert2int(line[2]),
		NbCache:     convert2int(line[3]),
		CacheSize:   convert2int(line[4]),
	}
	fmt.Println(entry)
}

func convert2int(val string) int {
	i, err := strconv.Atoi(val)
	check(err)
	return i
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
