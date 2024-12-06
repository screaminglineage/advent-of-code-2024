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

func parse_input(input string) (map[int][]int, [][]int) {
    page_orderings := make(map[int][]int)
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
        if _, found := page_orderings[before]; !found {
            page_orderings[before] = make([]int, 1)
            page_orderings[before][0] = after
        } else {
            page_orderings[before] = append(page_orderings[before], after)
        }
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
    return page_orderings, pages
}

// returns the indices which fail the check
func is_correct(page_orderings map[int][]int, update []int) (a int, b int, is_correct bool) {
    for i, page := range update {
        afters := page_orderings[page]
        for _, after := range afters {
            index := slices.Index(update, after)
            if index != -1 && index < i {
                a = index
                b = i
                is_correct = false
                return
            }
        }
    }
    return -1, -1, true
}

func part_1(page_orderings map[int][]int, pages [][]int) int {
    sum := 0
    for _, update := range pages {
        if _, _, correct := is_correct(page_orderings, update); correct {
            a := update[len(update)/2]
            sum += a
        }
    }
    return sum
}


func part_2(page_orderings map[int][]int, pages [][]int) int {
    sum := 0
    for _, update := range pages {
        incorrect := false
        for {
            if i, j, correct := is_correct(page_orderings, update); !correct {
                incorrect = true
                update[i], update[j] = update[j], update[i]
            } else {
                break
            }
        }
        if incorrect {
            a := update[len(update)/2]
            sum += a
        }
    }
    return sum
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


