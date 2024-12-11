package main

import (
    "log"
    "strings"
    "math"
    "strconv"
    "container/list"
    "fmt"
    "os"
)

const data_file = "data.txt"
const test_file = "test.txt"

func parse_input(input string) []int {
    nums := make([]int, 0)
    input = strings.TrimSuffix(input, "\n")
    for _, num_str := range strings.Split(input, " ") {
        num, err := strconv.Atoi(num_str)
        if err != nil {
            fmt.Printf("Got `%s`", num_str)
            panic("Unreachable, input should be correct")
        }

        nums = append(nums, num)
    }
    return nums
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



func part_1(nums []int) int {
    l := list.New()
    for _, num := range nums {
        l.PushBack(num)
    }

    for _ = range 25 {
        for e := l.Front(); e != nil; e = e.Next() {
            val, _ := e.Value.(int)
            if val == 0 {
                e.Value = 1
            } else if a, b, even := even_digits(val); even {
                l.InsertBefore(a, e)
                new_e := l.InsertAfter(b, e)
                l.Remove(e)
                e = new_e
            } else {
                val *= 2024
                e.Value = val
            }
        }
    }


    return l.Len()
}

func part_2(nums []int) int {
    return 0;
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

