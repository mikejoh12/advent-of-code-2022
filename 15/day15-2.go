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

func getOutsideDiamondPts(s sensor) (points []image.Point) {
	radius := s.mhRange + 1
	points = append(points, image.Point{s.position.X, s.position.Y-radius}) // top
	points = append(points, image.Point{s.position.X, s.position.Y+radius}) // bottom
	points = append(points, image.Point{s.position.X-radius, s.position.Y}) // left
	points = append(points, image.Point{s.position.X+radius, s.position.Y}) // right
	for diff := 1; diff <= radius; diff++ {
		points = append(points, image.Point{s.position.X+radius-diff, s.position.Y-diff}) // top right
		points = append(points, image.Point{s.position.X-radius+diff, s.position.Y-diff}) // top left
		points = append(points, image.Point{s.position.X+radius-diff, s.position.Y+diff}) // bottom right
		points = append(points, image.Point{s.position.X-radius+diff, s.position.Y+diff}) // bottom left
	}
	return points
}

func getAllOutsideDiamondPts(sensors []sensor) (points []image.Point) {
	for _, s := range sensors {
		points = append(points, getOutsideDiamondPts(s)...)
	}
	return points
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var sensors []sensor
	beaconPts := make(map[image.Point]bool)

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
	}

	maxXY := 4000000
	outsidePts := getAllOutsideDiamondPts(sensors)

	for _, point := range outsidePts {

		if point.X < 0 || point.X > maxXY || point.Y < 0 || point.Y > maxXY {
			continue
		}

		possible := true

		for _, sensor := range sensors {
			mhDist := getManhattanDist(point, sensor.position)
			if mhDist <= sensor.mhRange {
				possible = false
				break
			}
		}

		if possible {
			fmt.Println("part 2:", point.X * 4000000 + point.Y)
			break
		}
	}
}