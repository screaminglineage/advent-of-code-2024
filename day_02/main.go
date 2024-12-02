package main

import (
	"log"
    "strings"
    "strconv"
    "fmt"
	"os"
)

const data_file = "data.txt"
const test_file = "test.txt"

func abs(num int) int {
    if num < 0 {
        return -num;
    }
    return num
}

func part_1(input string) int {
    safe_seqs := 0
    lines := strings.Split(input, "\n")
    lines = lines[:len(lines)-1]

    for _, line := range lines {
        nums := strings.Split(line, " ")
        safe := true
        increasing := false
        for i, _ := range nums {
            if (i == len(nums) - 1) {
                break
            }

            current, err := strconv.Atoi(nums[i])
            if (err != nil) {
                break
            }
            next, err := strconv.Atoi(nums[i+1])
            if (err != nil) {
                break
            }

            diff := current - next
            if diff > 0 && increasing || diff > 0 && i == 0 {
                increasing = true
            } else if diff < 0 && !increasing || diff < 0 && i == 0 {
                increasing = false
            } else {
                safe = false
                break
            }
            diff = abs(diff)
            if diff < 1 || diff > 3 {
                safe = false
                break
            }
        }
        if safe {
            safe_seqs += 1
        }
    }
    return safe_seqs
}

func part_2(input string) int {
    safe_seqs := 0
    lines := strings.Split(input, "\n")
    lines = lines[:len(lines)-1]

    for _, line := range lines {
        nums := strings.Split(line, " ")
        safe := true
        i := 0
        j := 1
        current, err := strconv.Atoi(nums[i])
        if (err != nil) {
            panic("Unreachable")
        }
        next, err := strconv.Atoi(nums[j])
        if (err != nil) {
            panic("Unreachable")
        }

        increasing := (current - next) < 0

        for {
            diff := abs(current - next)
            if diff < 1 || diff > 3 {
                safe = false
                break
            }
            prev_increasing := increasing
            if j >= len(nums) - 1 {
                break
            }
            i += 1
            j += 1
            current, err = strconv.Atoi(nums[i])
            if (err != nil) {
                panic("Unreachable")
            }
            next, err = strconv.Atoi(nums[j])
            if (err != nil) {
                panic("Unreachable")
            }
            increasing = (current - next) < 0
            if prev_increasing != increasing {
                safe = false
                break
            }
        }
        if (safe) {
            safe_seqs += 1
        }
    }
    return safe_seqs
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


