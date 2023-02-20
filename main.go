package main

import (
	"fmt"
	"math/rand"
	"time"
)

type point struct {
	x, y int
}

func countRectangles(points []point) int {
	count := 0
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			for k := j + 1; k < len(points); k++ {
				if points[i].x == points[j].x && points[k].x == points[k].x && points[i].y == points[k].y && points[j].y == points[k].y {
					count++
				}
			}
		}
	}
	return count
}

func main() {
	rand.Seed(time.Now().UnixNano())
	n := 12
	points := make([]point, n)
	for i := 0; i < n; i++ {
		points[i] = point{rand.Intn(10), rand.Intn(10)}
	}
	fmt.Println("Points:", points)
	fmt.Println("Rectangles:", countRectangles(points))
}
