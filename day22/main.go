package day22

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type interval struct {
	lower int
	upper int
}

type cuboid struct {
	x interval
	y interval
	z interval
}

type instruction struct {
	on bool
	c  cuboid
}

func RunPuzzles() {
	file, err := os.Open("day22/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	instructions := []instruction{}
	for scanner.Scan() {
		i := instruction{}
		split1 := strings.Split(scanner.Text(), " ")
		if split1[0] == "on" {
			i.on = true
		} else {
			i.on = false
		}

		split2 := strings.Split(split1[1], ",")
		i.c = cuboid{interval{}, interval{}, interval{}}

		xsplit := strings.Split(split2[0][2:], "..")
		xmin, _ := strconv.Atoi(xsplit[0])
		xmax, _ := strconv.Atoi(xsplit[1])
		i.c.x = interval{xmin, xmax}

		ysplit := strings.Split(split2[1][2:], "..")
		ymin, _ := strconv.Atoi(ysplit[0])
		ymax, _ := strconv.Atoi(ysplit[1])
		i.c.y = interval{ymin, ymax}

		zsplit := strings.Split(split2[2][2:], "..")
		zmin, _ := strconv.Atoi(zsplit[0])
		zmax, _ := strconv.Atoi(zsplit[1])
		i.c.z = interval{zmin, zmax}

		instructions = append(instructions, i)
	}

	fmt.Println("Day 22:")
	fmt.Println("Puzzle 1:", puzzle1(instructions))
	fmt.Println("Puzzle 2:", puzzle2(instructions))
	fmt.Println()
}
