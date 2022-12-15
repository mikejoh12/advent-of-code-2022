package main

import (
	"bufio"
	"fmt"
	"image"
	"os"
)

type sensor struct {
	position	image.Point
	beacon		image.Point
	mhRange		int
}

func absDiffInt(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func getManhattanDist(pos1, pos2 image.Point) int {
	return absDiffInt(pos1.X, pos2.X) + absDiffInt(pos1.Y, pos2.Y)
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var sensors []sensor
	beaconPts := make(map[image.Point]bool)
	var maxX int

	for scanner.Scan() {
		t := scanner.Text()
		var sx, sy, bx, by int
		fmt.Sscanf(t, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sx, &sy, &bx, &by)
		
		sensors = append(sensors, sensor{
			position: image.Point{sx, sy},
			beacon: image.Point{bx, by},
			mhRange: getManhattanDist(image.Point{sx, sy}, image.Point{bx, by}),
		})
		beaconPts[image.Point{bx, by}] = true
		maxX = max(maxX, sx)
		maxX = max(maxX, bx)
	}
	
	y,  noBeaconPoints := 2000000, 0
	for x := -2 * maxX; x <= 2 * maxX; x++ {
		for _, sensor := range sensors {
			mhDist := getManhattanDist(image.Point{x, y}, sensor.position)
			if _, ok := beaconPts[image.Point{x, y}]; !ok && mhDist <= sensor.mhRange {
				noBeaconPoints++
				break
			}
		}
	}

	fmt.Println("part 1:", noBeaconPoints)
}