package main

import (
	"log"
    "strings"
    "fmt"
	"os"
)

const data_file = "data.txt"
const test_file = "test.txt"

func parse_input(input string) []string {
    parsed := make([]string, 0)
    lines := strings.Split(input, "\n")

    for _, line := range lines[:len(lines) - 1] {
        parsed = append(parsed, line)
    }
    return parsed
}

func part_1(input string) int {
    lines := parse_input(input)
    xmas := "XMAS"
    width := len(lines[0])
    height := len(lines)

    count := 0
    for y := range height {
        for x := range width {
            if lines[y][x] != 'X' {
                continue
            }
            // left to right
            if x + len(xmas) - 1 < width {
                slice := ""
                for i := range len(xmas) {
                    slice += string(lines[y][x+i])
                }
                if slice == "XMAS" {
                    count += 1
                }
            }

            // right to left
            if x - len(xmas) + 1 >= 0 {
                slice := ""
                for i := range len(xmas) {
                    slice += string(lines[y][x-i])
                }
                if slice == "XMAS" {
                    count += 1
                }
            }

            // top to bottom
            if y + len(xmas) - 1 < height {
                slice := ""
                for i := range len(xmas) {
                    slice += string(lines[y+i][x])
                }
                if slice == "XMAS" {
                    count += 1
                }
            }

            // bottom to top
            if y - len(xmas) + 1 >= 0 {
                slice := ""
                for i := range len(xmas) {
                    slice += string(lines[y-i][x])
                }
                if slice == "XMAS" {
                    count += 1
                }
            }

            // diagonals
            if y + len(xmas) - 1 < height && x + len(xmas) - 1 < width {
                slice := ""
                for i := range len(xmas) {
                    slice += string(lines[y+i][x+i])
                }
                if slice == "XMAS" {
                    count += 1
                }
            }

            if y - len(xmas) + 1 >= 0 && x - len(xmas) + 1 >= 0 {
                slice := ""
                for i := range len(xmas) {
                    slice += string(lines[y-i][x-i])
                }
                if slice == "XMAS" {
                    count += 1
                }
            }

            if y - len(xmas) + 1 >= 0 && x + len(xmas) - 1 < width {
                slice := ""
                for i := range len(xmas) {
                    slice += string(lines[y-i][x+i])
                }
                if slice == "XMAS" {
                    count += 1
                }
            }

            if y + len(xmas) - 1 < height && x - len(xmas) + 1 >= 0 {
                slice := ""
                for i := range len(xmas) {
                    slice += string(lines[y+i][x-i])
                }
                if slice == "XMAS" {
                    count += 1
                }
            }
        }
    }
    return count
}

func part_2(input string) int {
    return 0
}

func main() {
	data, err := os.ReadFile(data_file)
	if err != nil {
		log.Fatal(err)
	}
    input := string(data)
    fmt.Println("Part 1: ", part_1(input))
    fmt.Println("Part 2: ", part_2(input))
}


