package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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
	index        int
	startRow     int
	startColumn  int
	finishRow    int
	finishColumn int
	early        int
	late         int
}

const fileIn = "a_example.in"
const fileOut = "a_axample.out"

func main() {

	d, rides := parseFile(fileIn)
	fmt.Printf("data: %v, %v", d, rides)

	final := make([][]ride, d.nbVehicules)

	for i, r := range rides {
		j := i % d.nbVehicules
		final[j] = append(final[j], r)

	}

	writeOutFile(final, fileOut)
}

func writeOutFile(final [][]ride, out string) {
	f, _ := os.Create(out)
	defer f.Close()

	for _, rides := range final {
		var t string
		for _, r := range rides {
			t += fmt.Sprintf(" %d", r.index)
		}

		f.WriteString(fmt.Sprintf("%d %s\n", len(rides), strings.Trim(t, " ")))
	}

	f.Sync()
}

func parseFile(in string) (data, []ride) {
	f, err := os.Open(in)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// read params
	re := csv.NewReader(f)
	re.Comma = ' '
	all, err := re.ReadAll()
	l0 := all[0]
	d := data{
		convert2int(l0[0]),
		convert2int(l0[1]),
		convert2int(l0[2]),
		convert2int(l0[3]),
		convert2int(l0[4]),
		conv2int(l0[5]),
	}
	l := all[1:]
	rides := make([]ride, 0)
	for i, r := range l {
		rides = append(rides, ride{
			i,
			convert2int(r[0]),
			convert2int(r[1]),
			convert2int(r[2]),
			convert2int(r[3]),
			convert2int(r[4]),
			convert2int(r[5]),
		})
	}
	return d, rides
}

func convert2int(val string) int {
	i, _ := strconv.Atoi(val)
	return i
}

func conv2int(val string) int64 {
	i, _ := strconv.ParseInt(val, 10, 64)
	return i
}
