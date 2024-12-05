package main

import (
	"log"
    "strings"
    "slices"
    "strconv"
    "fmt"
	"os"
)

const data_file = "data.txt"
const test_file = "test.txt"

type PageOrdering struct {
    before, after int
}

func parse_input(input string) ([]PageOrdering, [][]int) {
    rules := make([]PageOrdering, 0)
    lines := strings.Split(input, "\n")
    var j int
    for i, line := range lines {
        if line == "" {
            j = i
            break
        }
        elems := strings.Split(line, "|")
        before, err := strconv.Atoi(elems[0])
        if (err != nil) { 
            break
        }
        after, err := strconv.Atoi(elems[1])
        if (err != nil) {
            break
        }
        rules = append(rules, PageOrdering{before, after})
    }

    pages := make([][]int, 0)
    for _, line := range lines[j+1:len(lines)-1] {
        nums := strings.Split(line, ",")
        p := make([]int, 0)
        for i := range len(nums) {
            n, err := strconv.Atoi(nums[i])
            if err != nil {
                break
            }
            p = append(p, n)
        }
        pages = append(pages, p)
    }
    return rules, pages
}

func get_afters(rules []PageOrdering, page int) []int {
    afters := make([]int, 0)
    for _, rule := range rules {
        if rule.before == page {
            afters = append(afters, rule.after)
        }
    }
    return afters
}

func is_correct(rules []PageOrdering, update []int) bool {
    is_correct := true
    for i, page := range update {
        afters := get_afters(rules, page)
        for _, after := range afters {
            index := slices.Index(update, after)
            if index != -1 && index < i {
                is_correct = false
                break
            }
        }
        if !is_correct {
            break
        }
    }
    return is_correct
}

func part_1(rules []PageOrdering, pages [][]int) int {
    sum := 0
    for _, update := range pages {
        if is_correct(rules, update) {
            a := update[len(update)/2]
            fmt.Println(a)
            sum += a
        }
    }
    return sum
}

func part_2(rules []PageOrdering, pages [][]int) int {
    return 0;
}

func main() {
	data, err := os.ReadFile(data_file)
	if err != nil {
		log.Fatal(err)
	}
    input := string(data)
    rules, pages := parse_input(input)
    fmt.Println("Part 1: ", part_1(rules, pages))
    fmt.Println("Part 2: ", part_2(rules, pages))
}


