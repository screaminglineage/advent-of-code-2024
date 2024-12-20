package main

import (
	"log"
    "slices"
    "fmt"
	"os"
)

const data_file = "data.txt"
const test_file = "test.txt"
const test_file_2 = "test2.txt"

func parse_input(input string) []int {
    file_blocks := make([]int, 0)
    free_space := false
    id := 0
    for _, ch := range input {
        if ch < '0' || ch > '9' {
            if ch == '\n' {
                break
            }
            panic("Unreachable, input should be correct")
        }
        num := int(ch - '0')
        if free_space {
            for _ = range num {
                file_blocks = append(file_blocks, -1)
            }
        } else {
            for _ = range num {
                file_blocks = append(file_blocks, id)
            }
            id += 1
        }

        free_space = !free_space
    }
    return file_blocks
}

func part_1(file_blocks []int) int {
    empty := slices.Index(file_blocks, -1)
    last := len(file_blocks) - 1

    for empty < last {
        file_blocks[empty], file_blocks[last] = file_blocks[last], file_blocks[empty]
        last -= 1
        empty = empty + slices.Index(file_blocks[empty:], -1)
    }

    for file_blocks[last] == -1 {
        last -= 1
    }
    file_blocks = file_blocks[:last+1]

    sum := 0
    for i, num := range file_blocks {
        sum += i*num
    }

    return sum;
}

func part_2(file_blocks []int) int {
    last := len(file_blocks) - 1

    start := slices.Index(file_blocks, -1)
    for start < last {
        for last > -1 && file_blocks[last] == -1 {
            last -= 1
        }
        file_id := file_blocks[last]
        count := 0
        for last > -1 && file_blocks[last] == file_id {
            last -= 1
            count += 1
        }

        found := false
        for start < last {
            free_space_count := 0
            next := start
            for start < last && file_blocks[next] == -1 {
                next += 1
                free_space_count += 1
            }
            if free_space_count >= count {
                found = true
                break
            }
            start = next + slices.Index(file_blocks[next:], -1)
        }

        if found {
            for i := range(count) {
                file_blocks[start + i], file_blocks[last + i + 1] = file_blocks[last + i + 1], file_blocks[start + i]
            }
        } else {
            start = slices.Index(file_blocks, -1)
        }
    }
    // fmt.Println(file_blocks)

    sum := 0
    for i, num := range file_blocks {
        if num != -1 {
            sum += i*num
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
    file_blocks := parse_input(input)
    fmt.Println("Part 1: ", part_1(slices.Clone(file_blocks)))
    fmt.Println("Part 2: ", part_2(file_blocks))
}


