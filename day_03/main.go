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
const test_file_part2 = "test_2.txt"

func eval_mul(expr string) int {
    nums, _ := strings.CutPrefix(expr, "mul(")
    nums2, _ := strings.CutSuffix(nums, ")")
    nums3 := strings.Split(nums2, ",")
    a, _ := strconv.Atoi(nums3[0])
    b, _ := strconv.Atoi(nums3[1])
    return a * b
}

func part_1(input string) int {
    mul_regex := regexp.MustCompile(`mul\(\d+,\d+\)`)
    res := mul_regex.FindAllString(input, -1)
    if (res == nil) {
        log.Fatal("Unreachable: regex is incorrect")
    }
    sum := 0
    for _, expr := range res {
        sum += eval_mul(expr)
    }
    return sum;
}

func part_2(input string) int {
    mul_regex := regexp.MustCompile(`mul\(\d+,\d+\)`)
    do_regex := regexp.MustCompile(`do(n't)?\(\)`)

    sum := 0
    bounds := mul_regex.FindStringIndex(input)
    if (bounds == nil) {
        log.Fatal("Unreachable: regex is incorrect")
    }
    sum += eval_mul(input[bounds[0]:bounds[1]])
    input = input[bounds[1]:]

    do := true
    for len(input) > 0 {
        mul_bounds := mul_regex.FindStringIndex(input)
        if (mul_bounds == nil) {
            break
        }

        do_bounds := do_regex.FindStringIndex(input[:mul_bounds[0]])
        if do_bounds != nil {
            res := input[do_bounds[0]:do_bounds[1]]
            if res == "don't()" {
                do = false
            } else {
                do = true
            }
        }
        if do {
            sum += eval_mul(input[mul_bounds[0]:mul_bounds[1]])
        }
        input = input[mul_bounds[1]:]
    }
    return sum;
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


