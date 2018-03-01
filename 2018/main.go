package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type data struct {
	rows        int
	columns     int
	nbVehicules int
	nbRides     int   // entre 1 et 10000 inclus
	bonus       int   // entre 1 et 10000 inclus
	nbSteps     int64 // entre 1 et 10^9
}

type ride struct {
	startRow     int
	startColumn  int
	finishRow    int
	finishColumn int
	early        int
	late         int
}

func main() {
	//var rides = []ride
	d := parseFile("a_example.in")
	fmt.Printf("data: %v", d)
}

func parseFile(in string) data {
	f, err := os.Open(in)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// read params
	r := csv.NewReader(f)

	all, err := r.ReadAll()
	l0 := all[0]
	log.Printf("got first line %s", l0)
	log.Printf("got first line %s", l0)

	d := data{
		convert2int(l0[0]),
		convert2int(l0[1]),
		convert2int(l0[2]),
		convert2int(l0[3]),
		convert2int(l0[4]),
		conv2int(l0[5]),
	}

	return d
}

func convert2int(val string) int {
	i, _ := strconv.Atoi(val)
	return i
}

func conv2int(val string) int64 {
	i, _ := strconv.ParseInt(val, 10, 64)
	return i
}
