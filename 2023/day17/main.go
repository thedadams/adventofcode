package main

import (
	"container/heap"
	"embed"
	"fmt"
	"math"
	"slices"
	"strings"

	"github.com/thedadams/adventofcode/2023/util"
)

//go:embed input.txt
var f embed.FS

func main() {
	// The answers that are outputted here are correct, but I am sure there are optimizations to be made.
	// This currently takes about 5 minutes to calculate both answers on my M1 MacBook Pro.
	partOne()
	partTwo()
}

func partOne() {
	s, err := util.ReadInputFile(f)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = s.Close()
	}()

	grid := make([][]int, 0)
	for s.Scan() {
		grid = append(grid, make([]int, 0))
		for _, spot := range strings.Split(s.Text(), "") {
			grid[len(grid)-1] = append(grid[len(grid)-1], util.MustAtoi(spot))
		}
	}

	fmt.Printf("Answer Day Seventeen, Part One: %v\n", dijkstra(grid, position{x: 0, y: 0, direction: "E", streak: 0}, 0, 3))
}

func partTwo() {
	s, err := util.ReadInputFile(f)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = s.Close()
	}()

	grid := make([][]int, 0)
	for s.Scan() {
		grid = append(grid, make([]int, 0))
		for _, spot := range strings.Split(s.Text(), "") {
			grid[len(grid)-1] = append(grid[len(grid)-1], util.MustAtoi(spot))
		}
	}

	fmt.Printf("Answer Day Seventeen, Part Two: %v\n", dijkstra(grid, position{x: 0, y: 0, direction: "START", streak: 0}, 4, 10))
}

func dijkstra(grid [][]int, source position, streakMin, streakMax int) int {
	queue := make([]position, 0, len(grid)*len(grid[0])*4*streakMax+1)
	dist := make(map[position]int, len(grid)*len(grid[0])*4*streakMax+1)
	processed := make(map[position]struct{}, len(grid)*len(grid[0])*4*streakMax)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if i+j > 0 {
				for _, dir := range []string{"N", "S", "W", "E"} {
					for k := 1; k <= streakMax; k++ {
						// My first guess at an optimization would be to not have to include the streak information.
						pos := position{x: j, y: i, direction: dir, streak: k}
						dist[pos] = math.MaxInt
						queue = append(queue, pos)
					}
				}
			}
		}
	}

	queue = append(queue, source)
	dist[source] = 0

	priorityQueue := &pq{
		queue: queue,
		dist:  dist,
	}
	heap.Init(priorityQueue)

	for len(queue) != 0 {
		current := heap.Pop(priorityQueue).(position)
		processed[current] = struct{}{}

		if current.y == len(grid)-1 && current.x == len(grid[current.y])-1 {
			if current.streak >= streakMin {
				return dist[current]
			}
			continue
		}

		for _, next := range neighbors(grid, current, streakMin, streakMax) {
			if _, ok := processed[next]; !ok {
				alt := dist[current] + grid[next.y][next.x]
				if alt < dist[next] {
					dist[next] = alt
					heap.Fix(priorityQueue, slices.Index(queue, next))
				}
			}
		}
	}

	return -1
}

type pq struct {
	queue []position
	dist  map[position]int
}

func (p *pq) Len() int {
	return len(p.queue)
}

func (p *pq) Less(i, j int) bool {
	if p.dist[p.queue[i]] != p.dist[p.queue[j]] {
		return p.dist[p.queue[i]] < p.dist[p.queue[j]]
	}

	if p.queue[i].y != p.queue[j].y {
		return p.queue[i].y < p.queue[j].y
	}

	if p.queue[i].x != p.queue[j].x {
		return p.queue[i].x < p.queue[j].x
	}

	return p.queue[i].streak < p.queue[j].streak
}

func (p *pq) Swap(i, j int) {
	p.queue[i], p.queue[j] = p.queue[j], p.queue[i]
}

func (p *pq) Push(x interface{}) {
	p.queue = append(p.queue, x.(position))
}

func (p *pq) Pop() interface{} {
	x := p.queue[len(p.queue)-1]
	p.queue = p.queue[:len(p.queue)-1]
	return x
}

type position struct {
	x, y, streak int
	direction    string
}

func neighbors(grid [][]int, current position, streakMin, streakMax int) []position {
	var next, possible []position
	switch current.direction {
	case "N":
		possible = []position{
			{x: current.x, y: current.y - 1, direction: "N", streak: current.streak + 1},
			{x: current.x - 1, y: current.y, direction: "W", streak: 1},
			{x: current.x + 1, y: current.y, direction: "E", streak: 1},
		}

	case "S":
		possible = []position{
			{x: current.x, y: current.y + 1, direction: "S", streak: current.streak + 1},
			{x: current.x - 1, y: current.y, direction: "W", streak: 1},
			{x: current.x + 1, y: current.y, direction: "E", streak: 1},
		}
	case "W":
		possible = []position{
			{x: current.x - 1, y: current.y, direction: "W", streak: current.streak + 1},
			{x: current.x, y: current.y - 1, direction: "N", streak: 1},
			{x: current.x, y: current.y + 1, direction: "S", streak: 1},
		}
	case "E":
		possible = []position{
			{x: current.x + 1, y: current.y, direction: "E", streak: current.streak + 1},
			{x: current.x, y: current.y - 1, direction: "N", streak: 1},
			{x: current.x, y: current.y + 1, direction: "S", streak: 1},
		}
	case "START":
		possible = []position{
			{x: current.x + 1, y: current.y, direction: "E", streak: 1},
			{x: current.x, y: current.y + 1, direction: "S", streak: 1},
		}
	}

	for _, pos := range possible {
		if pos.x >= 0 && pos.y >= 0 && pos.y < len(grid) && pos.x < len(grid[pos.y]) && pos.streak <= streakMax {
			if current.direction == "START" || current.direction == pos.direction || current.streak >= streakMin {
				next = append(next, pos)
			}
		}
	}

	return next
}
