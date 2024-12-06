package main

import (
	"log"
    "strings"
    "fmt"
	"os"
)

const data_file = "data.txt"
const test_file = "test.txt"

type Point struct {
    y, x int
}

type Grid struct {
    rows, cols int
    grid [][]rune
}


func parse_input(input string) (g Grid, start Point) {
    grid := make([][]rune, 0)
    lines := strings.Split(input, "\n")
    for y, line := range lines {
        if line == "" {
            break
        }
        row := make([]rune, 0)
        for x, ch := range line {
            if ch == '^' {
                start.x = x
                start.y = y
                row = append(row, '.')
            } else {
                row = append(row, rune(ch))
            }
        }
        grid = append(grid, row)
    }
    g = Grid{len(lines) - 1, len(lines[0]), grid}
    return g, start
}

var Empty struct{}

func part_1(g Grid, start Point) int {
    dirs := []Point{
        {-1, 0},    // up
        {0, 1},     // right
        {1, 0},     // down
        {0, -1},    // left
    }

    current := start
    dir := 0
    visited := make(map[Point]struct{})

    count := 1
    for {
        next := Point{current.y + dirs[dir].y, current.x + dirs[dir].x}
        if next.x >= g.cols || next.x < 0 || next.y >= g.rows || next.y < 0 {
            return count
        }
        for g.grid[next.y][next.x] == '#' {
            dir = (dir + 1) % len(dirs)
            next = Point{current.y + dirs[dir].y, current.x + dirs[dir].x}
            if next.x >= g.cols || next.x < 0 || next.y >= g.rows || next.y < 0 {
                return count
            }
        }
        current = next
        if _, found := visited[current]; !found {
            visited[current] = Empty
            count += 1
        }
    }
    return count
}


func part_2(g Grid, start Point) int {
    return 0
}

func main() {
	data, err := os.ReadFile(data_file)
	if err != nil {
		log.Fatal(err)
	}
    input := string(data)
    grid, start := parse_input(input)
    fmt.Println("Part 1: ", part_1(grid, start))
    fmt.Println("Part 2: ", part_2(grid, start))
}


