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


func detect_cycles(g Grid, start Point) bool {
    dirs := []Point{
        {-1, 0},    // up
        {0, 1},     // right
        {1, 0},     // down
        {0, -1},    // left
    }

    current_hare := start
    current_tort := start
    dir_hare := 0
    dir_tort := 0

    for {
        next_hare_1 := Point{current_hare.y + dirs[dir_hare].y, current_hare.x + dirs[dir_hare].x}
        if next_hare_1.x >= g.cols || next_hare_1.x < 0 || next_hare_1.y >= g.rows || next_hare_1.y < 0 {
            break
        }

        for g.grid[next_hare_1.y][next_hare_1.x] == '#' {
            dir_hare = (dir_hare + 1) % len(dirs)
            next_hare_1 = Point{current_hare.y + dirs[dir_hare].y, current_hare.x + dirs[dir_hare].x}
            if next_hare_1.x >= g.cols || next_hare_1.x < 0 || next_hare_1.y >= g.rows || next_hare_1.y < 0 {
                break
            }
        }

        next_hare := Point{next_hare_1.y + dirs[dir_hare].y, next_hare_1.x + dirs[dir_hare].x}
        if next_hare.x >= g.cols || next_hare.x < 0 || next_hare.y >= g.rows || next_hare.y < 0 {
            break
        }


        for g.grid[next_hare.y][next_hare.x] == '#' {
            dir_hare = (dir_hare + 1) % len(dirs)
            next_hare = Point{next_hare_1.y + dirs[dir_hare].y, next_hare_1.x + dirs[dir_hare].x}
            if next_hare.x >= g.cols || next_hare.x < 0 || next_hare.y >= g.rows || next_hare.y < 0 {
                break
            }
        }

        next_tort := Point{current_tort.y + dirs[dir_tort].y, current_tort.x + dirs[dir_tort].x}
        if next_tort.x >= g.cols || next_tort.x < 0 || next_tort.y >= g.rows || next_tort.y < 0 {
            break
        }

        for g.grid[next_tort.y][next_tort.x] == '#' {
            dir_tort = (dir_tort + 1) % len(dirs)
            next_tort = Point{current_tort.y + dirs[dir_tort].y, current_tort.x + dirs[dir_tort].x}
            if next_tort.x >= g.cols || next_tort.x < 0 || next_tort.y >= g.rows || next_tort.y < 0 {
                break
            }
        }

        current_hare = next_hare
        current_tort = next_tort
        if current_hare.x == current_tort.x && current_hare.y == current_tort.y && dirs[dir_hare].x == dirs[dir_tort].x && dirs[dir_hare].y == dirs[dir_tort].y { 
            return true
        }
    }
    return false
}

func print_grid(g Grid, start Point) int {
    for y := range g.grid {
        for x := range g.grid[y] {
            if start.x == x && start.y == y {
                fmt.Print("^")
            } else {
                fmt.Print(string(g.grid[y][x]))
            }
        }
        fmt.Println()
    }
    fmt.Println("###############################")
    return 0
}

func part_2(g Grid, start Point) int {
    count := 0
    for y := range g.grid {
        for x := range g.grid[y] {
            if (start.x == x && start.y == y) || g.grid[y][x] == '#' {
                continue
            }
            old := g.grid[y][x]
            g.grid[y][x] = '#'
            if detect_cycles(g, start) {
                count += 1
            }
            g.grid[y][x] = old
        }
    }
    return count
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


