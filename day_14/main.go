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

type Vector2 struct {
    x, y int
}

type Robot struct {
    pos Vector2
    vel Vector2
}

func parse_input(input string) []Robot {
    robots := make([]Robot, 0)
    lines := strings.Split(input, "\n")
    for i := range len(lines) - 1 {
        values := strings.Split(lines[i], " ")

        pos_str := strings.TrimPrefix(values[0], "p=")
        position := strings.Split(pos_str, ",")
        vel_str := strings.TrimPrefix(values[1], "v=")
        velocity := strings.Split(vel_str, ",")

        p_x, _ := strconv.Atoi(position[0])
        p_y, _ := strconv.Atoi(position[1])
        v_x, _ := strconv.Atoi(velocity[0])
        v_y, _ := strconv.Atoi(velocity[1])
        robots = append(robots, Robot{Vector2{p_x, p_y}, Vector2{v_x, v_y}})
    }
    return robots
}

// modulo function with wrap-around for negative numbers
func modulo(a, b int) int {
    return (a % b + b) % b
}

func (v1 *Vector2) add(v2 Vector2, width int, height int) {
    v1.x = modulo(v1.x + v2.x, width)
    v1.y = modulo(v1.y + v2.y, height)
}



func part_1(robots []Robot) int {
    const width = 101
    const height = 103
    for range 100 {
        for i := range robots {
            robots[i].pos.add(robots[i].vel, width, height)
        }
    }
    var quad_counts [4]int
    for _, robot := range robots {
        if robot.pos.x < width/2 && robot.pos.y < height/2 {
            quad_counts[0] += 1
        } else if robot.pos.x > width/2 && robot.pos.y < height/2 {
            quad_counts[1] += 1
        } else if robot.pos.x < width/2 && robot.pos.y > height/2 {
            quad_counts[2] += 1
        } else if robot.pos.x > width/2 && robot.pos.y > height/2 {
            quad_counts[3] += 1
        }
    }
    return quad_counts[0] * quad_counts[1] * quad_counts[2] * quad_counts[3]
}

func has_robot(robots []Robot, x int, y int) bool {
    for _, robot := range robots {
        if robot.pos.x == x && robot.pos.y == y {
            return true
        }
    }
    return false
}

func part_2(robots []Robot) {
    const width = 101
    const height = 103
    data := make([]byte, 0)
    for i := range 10000 {
        data = append(data, strconv.Itoa(i+1)...)
        data = append(data, '\n')
        for i := range robots {
            robots[i].pos.add(robots[i].vel, width, height)
        }
        for y := range height {
            for x := range width {
                if has_robot(robots, x, y) {
                    data = append(data, '#')
                } else {
                    data = append(data, ' ')
                }
            }
            data = append(data, '\n')
        }
        data = append(data, "\n\n"...)
    }
    os.WriteFile("out.txt", data, 0666)
}


func main() {
	data, err := os.ReadFile(data_file)
	if err != nil {
		log.Fatal(err)
	}
    input := string(data)
    robots := parse_input(input)
    fmt.Println("Part 1: ", part_1(slices.Clone(robots)))
    part_2(robots)
    fmt.Println("Part 2: Generated out.txt")
}


