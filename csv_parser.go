package scheduler

import (
	"os"
	"log"
	"errors"
	"strconv"
	"bufio"
	"strings"
)
/*
	Takes in 3 command line arguments. A file name for the csv that should be
	parsed, the number of slots each person is expected to fill, and the name
	of the file you want to create.
*/

const Yes = "Yes"
const Name = "name"
const Starter = "Events Start Here"


func main() {
	arguments := os.Args[1:]
	if len(arguments) != 3 {
		err1 := errors.New("Wrong number of command line arguments")
		log.Fatal(err1)
	}
	file, err2 := os.Open(arguments[0])
	if err2 != nil {
		log.Fatal(err2)
	}
	capacity, err3 := strconv.ParseUint(arguments[1], 10, 64)
	if err3 != nil {
		log.Fatal(err3)
	}
	events, candidates, maxEventCapacity, err4 := read_from_sheet(file, capacity)
	if err4 != nil {
		log.Fatal(err4)
	}
	err5 := file.Close()
	if err5 != nil {
		log.Fatal(err5)
	}
	g, src, dst := graphConstructor(events, candidates)
	g.maxFlow(src, dst)
	err6 := create_sheet(g, src, arguments[2], maxEventCapacity)
	if err6 != nil {
		log.Fatal(err6)
	}
	os.Exit(0)
}

func read_from_sheet(file *os.File, capacity uint64) (events []*Event, candidates []*Candidate , maxEventCapacity uint64, err error) {
	var i int
	var eventStart int
	var nameCol int
	var firstLine string
	var line string
	var linesplit[] string
	var newCap uint64
	var name string
	var potentialEvents[] string
	events = make([] *Event, 0)
	candidates = make([] *Candidate, 0) 
	reader := bufio.NewReader(file)
	firstLine, err = reader.ReadString('\n')
	if err != nil {
		return
	}
	firstRow := strings.Split(firstLine, ",")
	for i = 0; i < len(firstRow); i++ {
		if strings.ToLower(firstRow[i]) == name {
			nameCol = i
		}
		if strings.ToLower(firstRow[i]) == Starter {
			eventStart = i + 1
		}
	}
	line, err = reader.ReadString('\n')
	if err != nil {
		return
	}
	linesplit = strings.Split(line, ",")
	for j := eventStart; j < len(firstRow); j++ {
		newCap, err = strconv.ParseUint(linesplit[j], 10, 64)

		if err != nil {
			return
		}
		if newCap > maxEventCapacity {
			maxEventCapacity = newCap
		}
		events = append(events, &(Event{name: firstRow[j], capacity: uint(newCap)}))
	}
	for  ;reader.Buffered() != 0; {
		line, err = reader.ReadString('\n')
		if err != nil {
			return
		}
		linesplit = strings.Split(line, ",")
		potentialEvents = make([]string, 0)
		name = linesplit[nameCol]
		for j := eventStart; j < len(firstRow); j++ {
			if linesplit[j] == Yes {
				potentialEvents = append(potentialEvents, firstRow[j])
			}
		}
		candidates = append(candidates, &Candidate{name: name, capacity: uint(capacity), potentialEvents: potentialEvents})
	}
	return
}

func create_sheet(g *Graph, src *Vertex, name string, maxEventCapacity uint64) (err error) {
	var rowNum uint
	var csv *os.File
	size := uint(maxEventCapacity + 1)
	rows := make([]string, maxEventCapacity + 1)
	csv, err = os.Create(name + ".csv")
	if err != nil {
		return
	}
	for _, element := range src.edges {
		rowNum = 0
		for _, item := range element.dst.edges {
			if item.dst != src {
				if item.cost == 0 {
					rows[rowNum] = rows[rowNum] + item.dst.name + ","
				}
			}
		}
		for ; rowNum < size; rowNum += 1 {
			rows[rowNum] = rows[rowNum] + ","
		}
	}
	for _, item := range rows {
		_, err = csv.WriteString(item + "\n")
		if err != nil {
			return
		}
	}
	err = csv.Close()
	return
}