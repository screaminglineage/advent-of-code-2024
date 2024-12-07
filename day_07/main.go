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



type Equation struct {
    nums []int
    target int
}

func parse_input(input string) []Equation {
    equations := make([]Equation, 0)
    lines := strings.Split(input, "\n")
    for i := range len(lines) - 1 {
        var equation Equation
        target := strings.Split(lines[i], ":")[0]

        var err error
        equation.target, err = strconv.Atoi(target)
        if err != nil {
            panic("Unreachable, input should be correct")
        }

        nums := strings.Split(lines[i], " ")[1:]
        equation.nums = make([]int, 0)
        for _, num := range nums {
            n, err := strconv.Atoi(num)
            if err != nil {
                panic("Unreachable, input should be correct")
            }
            equation.nums = append(equation.nums, n)
        }
        equations = append(equations, equation)
    }
    return equations
}

func calc(nums []int, op string, operators []string, acc int, target int) bool {
    if len(nums) == 1 {
        return acc == target
    }
    var new_value int
    switch op {
        case "+": new_value = acc + nums[1]
        case "*": new_value = acc * nums[1]
    }
    for _, op := range operators {
        if calc(nums[1:], op, operators, new_value, target) {
            return true
        }
    }
    return false
}

func part_1(equations []Equation) int {
    operators := []string{"+", "*"}
    sum := 0
    for _, equation := range equations {
        for _, op := range operators {
            if calc(equation.nums, op, operators, equation.nums[0], equation.target) {
                sum += equation.target
                break
            }
        }
    }
    return sum
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
    equations := parse_input(input)
    fmt.Println("Part 1: ", part_1(equations))
    fmt.Println("Part 2: ", part_2(input))
}


