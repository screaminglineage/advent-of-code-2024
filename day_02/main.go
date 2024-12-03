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

func is_safe(nums []string) bool {
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
            return false
        }
        diff = abs(diff)
        if diff < 1 || diff > 3 {
            return false
        }
    }
    return true
}

func part_1(input string) int {
    safe_seqs := 0
    lines := strings.Split(input, "\n")
    lines = lines[:len(lines)-1]

    for _, line := range lines {
        nums := strings.Split(line, " ")
        if is_safe(nums) {
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
        if is_safe(nums) {
            fmt.Println(nums)
            safe_seqs += 1
            continue
        }
        for i := range len(nums) {
            sliced := append([]string{}, nums[:i]...)
            sliced = append(sliced, nums[i+1:]...)
            if is_safe(sliced) {
                fmt.Println(nums)
                safe_seqs += 1
                break
            }
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


