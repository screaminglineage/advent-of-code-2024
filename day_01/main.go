package main

import (
	"log"
    "strings"
    "strconv"
    "fmt"
    "slices"
	"os"
)

const data_file = "data.txt"
const test_file = "test.txt"

func part_1(input string) int {
    left := make([]int, 0)
    right := make([]int, 0)

    for _, line := range strings.Split(input, "\n") {
        elems := strings.Split(line, "   ")
        a, err := strconv.Atoi(elems[0])
        if (err != nil) { 
            break
        }
        b, err := strconv.Atoi(elems[1])
        if (err != nil) {
            break
        }
        fmt.Println(a, b)
        left = append(left, a)
        right = append(right, b)
    }
    slices.Sort(left)
    slices.Sort(right)

    count := 0
    for i := range len(left) {
        var distance int
        if (left[i] > right[i]) {
            distance = left[i] - right[i]
        } else {
            distance = right[i] - left[i]
        }
        count += distance
    }

    return count
}

func main() {
	data, err := os.ReadFile(data_file)
	if err != nil {
		log.Fatal(err)
	}
    input := string(data)
    fmt.Println(input)
    fmt.Println("Part 1: ", part_1(input))
}


