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
    grid [][]int
}

var DIRS = [4]Point{
    {-1, 0},    // up
    {0, 1},     // right
    {1, 0},     // down
    {0, -1},    // left
}

func parse_input(input string) (g Grid, starts []Point, ends []Point) {
    starts = make([]Point, 0)
    ends = make([]Point, 0)
    grid := make([][]int, 0)
    lines := strings.Split(input, "\n")
    for y, line := range lines {
        if line == "" {
            break
        }
        row := make([]int, 0)
        for x, ch := range line {
            if ch < '0' || ch > '9' {
                fmt.Printf("Got `%c`\n", ch)
                panic("Unreachable, input must be correct")
            }
            num := int(ch - '0')
            row = append(row, num)
            if num == 0 {
                starts = append(starts, Point{y,x})
            } else if num == 9 {
                ends = append(ends, Point{y,x})
            }
        }
        grid = append(grid, row)
    }
    g = Grid{len(lines) - 1, len(lines[0]), grid}
    return g, starts, ends
}


var Empty struct{}

func (p1 Point) add(p2 Point) Point {
    return Point{p1.y + p2.y, p1.x + p2.x}
}

func traverse_trail(g Grid, start Point) int {
    visited := make(map[Point]struct{})
    queue := make([]Point, 1)
    queue[0] = start

    count := 0
    for len(queue) > 0 {
        current := queue[0]
        queue = queue[1:]

        if _, f := visited[current]; f {
            continue
        }
        visited[current] = Empty

        if g.grid[current.y][current.x] == 9 {
            count += 1
            continue
        }

        for _, dir := range DIRS {
            next := current.add(dir)
            if next.x >= g.cols || next.x < 0 || next.y >= g.rows || next.y < 0 {
                continue
            }
            if _, found := visited[next]; found {
                continue
            }
            if g.grid[next.y][next.x] - g.grid[current.y][current.x] != 1 {
                continue
            }
            queue = append(queue, next)
        }
    }
    return count
}

func part_1(g Grid, starts []Point) int {
    sum := 0
    for _, start := range starts {
        sum += traverse_trail(g, start)
    }
    return sum
}

func traverse_trail_2(g Grid, start Point, end Point) int {
    rating := 0
    queue := make([]Point, 1)
    queue[0] = start

    for len(queue) > 0 {
        current := queue[0]
        queue = queue[1:]

        if current == end {
            rating += 1
            continue
        }

        for _, dir := range DIRS {
            next := current.add(dir)
            if next.x >= g.cols || next.x < 0 || next.y >= g.rows || next.y < 0 {
                continue
            }
            if g.grid[next.y][next.x] - g.grid[current.y][current.x] != 1 {
                continue
            }
            queue = append(queue, next)
        }
    }
    return rating
}


func part_2(g Grid, starts []Point, ends []Point) int {
    sum := 0
    for _, start := range starts {
        ratings := 0
        for _, end := range ends {
            ratings += traverse_trail_2(g, start, end)
        }
        sum += ratings
    }
    return sum
}

func main() {
    data, err := os.ReadFile(data_file)
    if err != nil {
        log.Fatal(err)
    }
    input := string(data)
    grid, starts, ends := parse_input(input)
    fmt.Println("Part 1: ", part_1(grid, starts))
    fmt.Println("Part 2: ", part_2(grid, starts, ends))
}


