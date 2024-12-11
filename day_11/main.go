package main

import (
    "log"
    "strings"
    "math"
    "strconv"
    "fmt"
    "os"
)

const data_file = "data.txt"
const test_file = "test.txt"

func insert_or_increment(data map[int]int, num int, count int) {
    if _, found := data[num]; found {
        data[num] += count
    } else {
        data[num] = count
    }
}


func parse_input(input string) map[int]int {
    data := make(map[int]int)
    input = strings.TrimSuffix(input, "\n")
    for _, num_str := range strings.Split(input, " ") {
        num, err := strconv.Atoi(num_str)
        if err != nil {
            fmt.Printf("Got `%s`", num_str)
            panic("Unreachable, input should be correct")
        }
        insert_or_increment(data, num, 1)
    }
    return data
}


func even_digits(num int) (first int, second int, has_even bool)  {
    num_len := int(math.Ceil(math.Log10(float64(num + 1))))
    has_even = num_len % 2 == 0
    if !has_even {
        return
    }
    num_str := strconv.Itoa(num)
    first, _ = strconv.Atoi(num_str[:num_len/2])
    second, _ = strconv.Atoi(num_str[num_len/2:])
    return
}

func count_stones(blinks int, data map[int]int) int {
    next := make(map[int]int)
    for range blinks {
        for val, count := range data {
            if val == 0 {
                insert_or_increment(next, 1, count)
            } else if a, b, even := even_digits(val); even {
                insert_or_increment(next, a, count)
                insert_or_increment(next, b, count)
            } else {
                insert_or_increment(next, val*2024, count)
            }
        }
        next, data = data, next
        next = make(map[int]int)
    }

    sum := 0
    for _, v := range data {
        sum += v
    }
    return sum
}

func part_1(data map[int]int) int {
    return count_stones(25, data)
}

func part_2(data map[int]int) int {
    return count_stones(75, data)
}


func main() {
    data, err := os.ReadFile(data_file)
    if err != nil {
        log.Fatal(err)
    }
    input := string(data)
    nums := parse_input(input)
    fmt.Println("Part 1: ", part_1(nums))
    fmt.Println("Part 2: ", part_2(nums))
}


