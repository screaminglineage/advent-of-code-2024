package main

import (
	"log"
    "regexp"
    "strings"
    "strconv"
    "fmt"
	"os"
)

const data_file = "data.txt"
const test_file = "test.txt"

func part_1(input string) int {
    mul_regex := regexp.MustCompile(`mul\(\d+,\d+\)`)
    res := mul_regex.FindAllString(input, -1)
    if (res == nil) {
        log.Fatal("Unreachable: regex is incorrect")
    }
    sum := 0
    for _, expr := range res {
        nums, _ := strings.CutPrefix(expr, "mul(")
        nums2, _ := strings.CutSuffix(nums, ")")
        nums3 := strings.Split(nums2, ",")
        a, _ := strconv.Atoi(nums3[0])
        b, _ := strconv.Atoi(nums3[1])
        sum += a*b
    }
    return sum;
}

func part_2(input string) int {
    return 0;
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


