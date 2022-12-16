package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type interval struct {
	min, max int
}

type position struct {
	row int
	col int
}

type sensor struct {
	position position
	beacon   position
	distance int
}

func toInt(raw string) int {
	res, err := strconv.Atoi(raw)

	if err != nil {
		panic(err)
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func loadPosition(raw string) position {
	coordsSplit := strings.Split(raw, ", ")
	x := toInt(strings.Split(coordsSplit[0], "=")[1])
	y := toInt(strings.Split(coordsSplit[1], "=")[1])

	return position{x, y}
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

func distance(a, b position) int {
	return abs(a.row-b.row) + abs(a.col-b.col)
}

func loadSensors(filename string) []sensor {
	sensors := []sensor{}

	readFile, err := os.Open(filename)
	defer readFile.Close()
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	for fileScanner.Scan() {
		line := fileScanner.Text()

		sensorRawPos := line[strings.Index(line, "x="):strings.Index(line, ":")]
		beaconRawPos := line[strings.LastIndex(line, "x="):]
		sensorPos := loadPosition(sensorRawPos)
		beaconPos := loadPosition(beaconRawPos)
		distance := distance(sensorPos, beaconPos)

		sensors = append(sensors, sensor{position: sensorPos, beacon: beaconPos, distance: distance})
	}

	return sensors
}

func getXLimits(sensors []sensor) (int, int) {
	minXS, maxXS, maxD := sensors[0].position.row, sensors[0].position.row, sensors[0].distance

	for i := 1; i < len(sensors); i++ {
		if sensors[i].position.row < minXS {
			minXS = sensors[i].position.row
		}
		if sensors[i].position.row > maxXS {
			maxXS = sensors[i].position.row
		}
		if sensors[i].distance > maxD {
			maxD = sensors[i].distance
		}
	}

	return minXS - maxD, maxXS + maxD
}

func getYLimits(sensors []sensor) (int, int) {
	minYS, maxYS, maxD := sensors[0].position.col, sensors[0].position.col, sensors[0].distance

	for i := 1; i < len(sensors); i++ {
		if sensors[i].position.col < minYS {
			minYS = sensors[i].position.row
		}
		if sensors[i].position.col > maxYS {
			maxYS = sensors[i].position.row
		}
		if sensors[i].distance > maxD {
			maxD = sensors[i].distance
		}
	}

	return minYS - maxD, maxYS + maxD
}

func contains(arr []int, n int) bool {
	for _, el := range arr {
		if el == n {
			return true
		}
	}
	return false
}

func getBeaconsAtColN(sensors []sensor, col int) int {
	b := 0
	x := []int{}

	for _, s := range sensors {
		if s.beacon.col == col && !contains(x, s.beacon.row) {
			b++
			x = append(x, s.beacon.row)
		}
	}

	return b
}

// ordered list of disjoint intervals
type FreeSpace struct {
	intervals []interval
}

func (fs FreeSpace) String() string {
	res := ""
	//res += fmt.Sprintf("[%d] ", len(fs.intervals))
	for _, i := range fs.intervals {
		res += fmt.Sprintf("%v ", i)
	}
	return res
}

func (fs *FreeSpace) Add(item interval) {
	// https://coderbyte.com/algorithm/insert-interval-into-list-of-sorted-disjoint-intervals
	newSet := make([]interval, 0)
	endSet := make([]interval, 0)
	i := 0
	// add intervals that come before the new interval
	for i < len(fs.intervals) && fs.intervals[i].max < item.min {
		newSet = append(newSet, fs.intervals[i])
		i++
	}

	// add our new interval to this final list
	newSet = append(newSet, item)

	// check each interval that comes after the new interval to determine if we can merge
	// if no merges are required then populate a list of the remaining intervals
	for i < len(fs.intervals) {
		var last = newSet[len(newSet)-1]
		if fs.intervals[i].min < last.max {
			newInterval := interval{min(last.min, fs.intervals[i].min), max(last.max, fs.intervals[i].max)}
			newSet[len(newSet)-1] = newInterval
		} else {
			endSet = append(endSet, fs.intervals[i])
		}
		i++
	}
	fs.intervals = append(newSet, endSet...)
}

func (fs *FreeSpace) Merge() {
	if len(fs.intervals) == 0 {
		return
	}
	newSet := make([]interval, 0)
	var last = fs.intervals[0]
	i := 1

	for i < len(fs.intervals) {
		if last.max+1 >= fs.intervals[i].min {
			last.max = fs.intervals[i].max
		} else {
			newSet = append(newSet, last)
			last = fs.intervals[i]
		}
		i++
	}
	newSet = append(newSet, last)
	fs.intervals = newSet
}

func (fs *FreeSpace) Intersect(item interval) {
	newSet := make([]interval, 0)
	for _, i := range fs.intervals {
		if i.max < item.min {
			continue
		}
		if i.min > item.max {
			break
		}
		newSet = append(newSet, interval{max(i.min, item.min), min(i.max, item.max)})
	}
	fs.intervals = newSet
}

func (fs *FreeSpace) Cardinality() int {
	res := 0
	for _, i := range fs.intervals {
		res += i.max - i.min + 1
	}
	return res
}

func getImpossibleBeaconPositionsAtColN(sensors []sensor, minX, maxX, col int) (int, FreeSpace) {
	line := FreeSpace{}

	for _, s := range sensors {
		r := s.distance - abs(s.position.col-col)
		if r > 0 {

			x1, x2 := s.position.row-r, s.position.row+r

			if x1 < minX {
				x1 = minX
			}
			if x2 > maxX {
				x2 = maxX
			}

			line.Add(interval{x1, x2})
		}
	}

	line.Merge()
	n := line.Cardinality()
	b := getBeaconsAtColN(sensors, col)
	return n - b, line
}

func getDistressBeacon(sensors []sensor, minX, maxX, minY, maxY int) position {
	pos := interval{minX, maxX}

	for i := minY; i <= maxY; i++ {
		_, line := getImpossibleBeaconPositionsAtColN(sensors, minX, maxX, i)

		line.Intersect(pos)

		if len(line.intervals) > 1 {
			return position{line.intervals[0].max + 1, i}
		}
	}

	panic("Signal not found")
}

func getTuningFrequency(sensors []sensor, minX, maxX, minY, maxY int) int {
	beacon := getDistressBeacon(sensors, minX, maxX, minY, maxY)

	return beacon.row*4000000 + beacon.col
}

func main() {
	input := os.Args[1]
	col := toInt(os.Args[2])
	min := toInt(os.Args[3])
	max := toInt(os.Args[4])

	sensors := loadSensors(input)
	minX, maxX := getXLimits(sensors)
	nPositions, _ := getImpossibleBeaconPositionsAtColN(sensors, minX, maxX, col)
	fmt.Println("1:", nPositions)

	frequency := getTuningFrequency(sensors, min, max, min, max)
	fmt.Println("2:", frequency)
}
