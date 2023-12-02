package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	day_1_part_2()
}

func day_2_part_1() {
	total := 0
	const MAX_RED = 12
	const MAX_BLUE = 14
	const MAX_GREEN = 13

	scanFile("./inputs/task_2_part_1.txt", func(s string) {
		semiSep := strings.Index(s, ":")
		if semiSep == -1 {
			log.Fatal("Something wrong with identifying game number")
		}

		splitedGame := strings.Split(s[:semiSep], " ")

		gameNum, err := strconv.Atoi(splitedGame[1])

		if err != nil {
			log.Fatal(err)
		}

		subsets := strings.Split(s[semiSep+1:], ";")

		for _, str := range subsets {
			cubesSections := strings.Split(str, ",")
			colorsCount := map[string]int{
				"red":   0,
				"green": 0,
				"blue":  0,
			}

			for _, cube := range cubesSections {
				cubeSplited := strings.Split(strings.Trim(cube, " "), " ")
				cubeNum, err := strconv.Atoi(cubeSplited[0])

				if err != nil {
					log.Fatal(err)
				}

				if colorsCount[cubeSplited[1]] < cubeNum {
					colorsCount[cubeSplited[1]] = cubeNum
				}
			}

			if colorsCount["green"] > MAX_GREEN || colorsCount["red"] > MAX_RED || colorsCount["blue"] > MAX_BLUE {
				return
			}
		}

		total += gameNum
	})

	fmt.Println(total)
}

func day_2_part_2() {
	total := 0
	scanFile("./inputs/task_2_part_2.txt", func(s string) {
		semiSep := strings.Index(s, ":")
		if semiSep == -1 {
			log.Fatal("Something wrong with identifying game number")
		}

		subsets := strings.Split(s[semiSep+1:], ";")

		colorsCount := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		for _, str := range subsets {
			cubesSections := strings.Split(str, ",")

			for _, cube := range cubesSections {
				cubeSplited := strings.Split(strings.Trim(cube, " "), " ")
				cubeNum, err := strconv.Atoi(cubeSplited[0])

				if err != nil {
					log.Fatal(err)
				}

				if colorsCount[cubeSplited[1]] < cubeNum {
					colorsCount[cubeSplited[1]] = cubeNum
				}
			}
		}

		total += colorsCount["red"] * colorsCount["green"] * colorsCount["blue"]
	})

	fmt.Println(total)
}

func day_1_part_2() {
	total := 0
	strNums := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	scanFile("./inputs/task_1_part_2.txt", func(s string) {
		str := ""
		passedStr := ""

	out:
		for i := 0; i < len(s); i++ {
			_, err := strconv.Atoi(string(s[i]))
			passedStr += string(s[i])

			if len(passedStr) >= 3 {
				for i, strNum := range strNums {
					if strings.Contains(passedStr, strNum) {
						str += fmt.Sprint(i + 1)
						break out
					}
				}
			}

			if err != nil {
				continue
			}

			str += string(s[i])
			break
		}
		passedStr = ""

	out_reverse:
		for i := len(s) - 1; i >= 0; i-- {
			_, err := strconv.Atoi(string(s[i]))
			passedStr = string(s[i]) + passedStr

			if len(passedStr) >= 3 {
				for i, strNum := range strNums {
					if strings.Contains(passedStr, strNum) {
						str += fmt.Sprint(i + 1)
						break out_reverse
					}
				}
			}

			if err != nil {
				continue
			}
			str += string(s[i])
			break
		}

		num, _ := strconv.Atoi(str)
		total += num
	})
	fmt.Println(total)
}

func day_1_part_1() {
	total := 0
	scanFile("./inputs/task_1_part_1.txt", func(s string) {
		str := ""

		for i := 0; i < len(s); i++ {
			_, err := strconv.Atoi(string(s[i]))
			if err != nil {
				continue
			}

			str += string(s[i])
			break
		}

		for i := len(s) - 1; i >= 0; i-- {
			_, err := strconv.Atoi(string(s[i]))

			if err != nil {
				continue
			}
			str += string(s[i])
			break
		}

		num, _ := strconv.Atoi(str)
		total += num
	})
	fmt.Println(total)
}

type cbFunc func(string)

func scanFile(filePath string, cb cbFunc) {
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		cb(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
