package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	day_4_part_2()
}

func day_4_part_2() {
	lines := []string{}
	totalCards := 0
	scanFile("./inputs/task_4_part_2_demo.txt", func(s string) {
		lines = append(lines, s)
	})

	totalCards = len(lines)

	var get_matching_count = func(line string) int {
		matchedCount := 0
		game := strings.Split(line, ":")
		numbers := strings.Split(game[1], "|")

		winningNumbers := strings.Split(strings.Trim(numbers[0], " "), " ")
		actualNumbers := strings.Split(strings.Trim(numbers[1], " "), " ")

		for _, ch := range winningNumbers {
			if ch == "" {
				continue
			}
			if slices.Contains(actualNumbers, ch) {
				matchedCount++
			}
		}
		return matchedCount
	}

	// var recCount = func(posStart int, posEnd int) int {
	// 	total := 0
	// 	for i := posStart; i < posEnd; i++ {
	// 		card := lines[i]
	// 		winningCount := get_matching_count(card)
	// 		if winningCount != 0 {
	// 			total += recCount(i, winningCount)
	// 		}
	// 	}
	// }

	for i := 0; i < len(lines); i++ {
		card := lines[i]
		winningCount := get_matching_count(card)
		repetitions := lines[i+1:i+1+winningCount]
		for j := 0; j < len(repetitions) - 1; j++ {
			fmt.Println(repetitions)
			fmt.Println()
			rep := repetitions[j]
			winningCount := get_matching_count(rep)
			repetitions = repetitions[j+1:]
			repetitions = append(repetitions, lines[i+j:i+j+winningCount]...)
			totalCards++
		}
	}
	fmt.Println(totalCards)
}

func day_4_part_1() {
	total := 0
	scanFile("./inputs/task_4_part_1.txt", func(s string) {
		points := 0
		game := strings.Split(s, ":")
		// cardNumber := strings.Split(game[0], " ")[1]
		numbers := strings.Split(game[1], "|")

		winningNumbers := strings.Split(strings.Trim(numbers[0], " "), " ")
		actualNumbers := strings.Split(strings.Trim(numbers[1], " "), " ")

		for _, ch := range winningNumbers {
			if ch == "" {
				continue
			}
			if slices.Contains(actualNumbers, ch) {
				points = int(math.Max(float64(points*2), 1))
			}
		}
		total += points
	})
	fmt.Println(total)
}

type LineNumber struct {
	StartPos int
	EndPos   int
	Number   string
}

func isDigit(s string) bool {
	if _, err := strconv.Atoi(s); err == nil {
		return true
	}

	return false
}

func day_3_part_1() {
	total := 0
	lines := []string{}
	scanFile("./inputs/task_3_part_2.txt", func(s string) { lines = append(lines, s) })
	goods := [][][]int{}

	for i := 0; i < len(lines); i++ {
		goods = append(goods, [][]int{})
		for j := 0; j < len(lines[0]); j++ {
			goods[i] = append(goods[i], []int{})
		}
	}

	var isSymbol = func(i int, j int, num int) bool {
		if i < 0 || i >= len(lines) || j < 0 || j >= len(lines[0]) {
			return false
		}

		if lines[i][j] == '*' {
			goods[i][j] = append(goods[i][j], num)
		}

		return lines[i][j] != '.' && !(isDigit(string(lines[i][j])))
	}

	for i, line := range lines {
		j := 0
		for j < len(line) {
			start := j
			digit := ""
			for {
				if j >= len(line) {
					break
				}
				if !isDigit(string(line[j])) {
					break
				}
				digit += string(line[j])
				j++
			}

			if digit == "" {
				j++
				continue
			}

			num, _ := strconv.Atoi(digit)

			_ = isSymbol(i, start-1, num) || isSymbol(i, j, num)

			for k := start - 1; k < j+1; k++ {
				_ = isSymbol(i-1, k, num) || isSymbol(i+1, k, num)
			}
		}
	}

	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[0]); j++ {
			nums := goods[i][j]
			if lines[i][j] == '*' && len(nums) == 2 {
				total += nums[0] * nums[1]
			}
		}
	}
	// 	for j, char := range line {
	// 		if _, err := strconv.Atoi(string(char)); err == nil {
	// 			digit += string(char)
	// 		} else {
	// 			if digit != "" {
	// 				startPos := j - len(digit)
	// 				endPos := j
	// 				if (startPos > 0 && line[startPos-1] != '.') || (endPos < len(line) && line[endPos] != '.') {
	// 					num, _ := strconv.Atoi(digit)
	// 					total += num
	// 				} else {
	// 					for k := int(math.Max(0, float64(startPos-1))); k <= endPos; k++ {
	// 						isPrevLineMatch := i > 0 && isSymbol(string(lines[i-1][k]))
	// 						isNextLineMatch := i+1 < len(lines) && isSymbol(string(lines[i+1][k]))
	// 						if isPrevLineMatch || isNextLineMatch {
	// 							num, _ := strconv.Atoi(digit)
	// 							total += num
	// 							fmt.Println(num)
	// 							break
	// 						}
	// 					}
	// 				}

	// 				digit = ""
	// 			}
	// 		}
	// 	}

	// 	if digit != "" {
	// 		startPos := (len(line) - len(digit))
	// 		endPos := len(line) - 1
	// 		fmt.Println(startPos, endPos)
	// 		for k := int(math.Max(0, float64(startPos-1))); k <= endPos; k++ {
	// 			isPrevLineMatch := i > 0 && isSymbol(string(lines[i-1][k]))
	// 			isNextLineMatch := i+1 < len(lines) && isSymbol(string(lines[i+1][k]))
	// 			if isPrevLineMatch || isNextLineMatch {
	// 				num, _ := strconv.Atoi(digit)
	// 				total += num
	// 				fmt.Println(num)
	// 				break
	// 			}
	// 		}
	// 	}
	// 	digit = ""
	// }

	// symbolsPosByLine := [][]int{}
	// numbersPosByLine := [][]LineNumber{}
	// scanFile("./inputs/task_3_part_1_demo.txt", func(s string) {
	// 	symbolsPos := []int{}
	// 	numbersPos := []LineNumber{}
	// 	numberIndex := 0
	// 	for i, char := range s {
	// 		if _, err := strconv.Atoi(string(char)); err == nil {
	// 			if len(numbersPos) <= numberIndex {
	// 				numbersPos = append(numbersPos, LineNumber{
	// 					StartPos: i,
	// 				})
	// 			}
	// 			numbersPos[numberIndex].Number += string(char)
	// 		} else {
	// 			if len(numbersPos) > numberIndex {
	// 				numbersPos[numberIndex].EndPos = i - 1
	// 				numberIndex++
	// 			}
	// 			if char != '.' {
	// 				symbolsPos = append(symbolsPos, i)
	// 			}
	// 		}
	// 	}
	// 	symbolsPosByLine = append(symbolsPosByLine, symbolsPos)
	// 	numbersPosByLine = append(numbersPosByLine, numbersPos)
	// })

	// for i, numberLine := range numbersPosByLine {
	// 	// number_loop:
	// 	for _, number := range numberLine {
	// 		//Searching on the same line
	// 		symbolIndex := slices.IndexFunc(symbolsPosByLine[i], func(i int) bool {
	// 			return i == number.StartPos-1 || i == number.EndPos+1
	// 		})

	// 		if symbolIndex != -1 {
	// 			num, _ := strconv.Atoi(number.Number)
	// 			total += num
	// 			continue
	// 		}

	//Searching on previous line
	// if i > 0 {
	// 	startPos := number.StartPos
	// 	endPos := number.EndPos
	// 	index := 1

	// 	for j := i; j >= 0; j-- {
	// 		symbolIndex := slices.IndexFunc(symbolsPosByLine[j], func(i int) bool {
	// 			return (i >= startPos-1 && i <= number.EndPos-(1*index)) || (i >= number.StartPos+(1*index) && i <= endPos+1)
	// 		})

	// 		if symbolIndex != -1 {
	// 			fmt.Printf("number: %s, index: %d\n", number.Number, symbolIndex)
	// 			num, _ := strconv.Atoi(number.Number)
	// 			total += num
	// 			continue number_loop
	// 		} else {
	// 			startPos--
	// 			endPos++
	// 			index++
	// 		}
	// 	}
	// }

	// 		if i > 0 {
	// 			symbolIndex := slices.IndexFunc(symbolsPosByLine[i-1], func(i int) bool {
	// 				return i >= number.StartPos-1 && i <= number.EndPos+1
	// 			})

	// 			if symbolIndex != -1 {
	// 				num, _ := strconv.Atoi(number.Number)
	// 				total += num
	// 				continue
	// 			}
	// 		}

	// 		//Searching on next line
	// 		if i != len(symbolsPosByLine)-1 {
	// 			symbolIndex := slices.IndexFunc(symbolsPosByLine[i+1], func(i int) bool {
	// 				if number.Number == "9" {
	// 					fmt.Println(number.StartPos, number.EndPos)
	// 				}
	// 				return i >= number.StartPos-1 && i <= number.EndPos+1
	// 			})

	// 			if symbolIndex != -1 {
	// 				num, _ := strconv.Atoi(number.Number)
	// 				total += num
	// 				continue
	// 			}
	// 		}
	// 		fmt.Printf("No symbol: %s\n", number.Number)
	// 	}
	// }

	// fmt.Println(symbolsPosByLine)
	// fmt.Println(numbersPosByLine)
	fmt.Println(total)

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

	var getStrNum = func(passedStr string) int {
		num := -1
		for i, strNum := range strNums {
			if strings.Contains(passedStr, strNum) {
				num = i + 1
				break
			}
		}
		return num
	}

	scanFile("./inputs/task_1_part_2.txt", func(s string) {
		str := ""
		passedStr := ""

		for i := 0; i < len(s); i++ {
			_, err := strconv.Atoi(string(s[i]))
			passedStr += string(s[i])

			if len(passedStr) >= 3 {
				strNum := getStrNum(passedStr)
				if strNum != -1 {
					str += fmt.Sprint(strNum)
					break
				}
			}

			if err != nil {
				continue
			}

			str += string(s[i])
			break
		}

		passedStr = ""

		for i := len(s) - 1; i >= 0; i-- {
			_, err := strconv.Atoi(string(s[i]))
			passedStr = string(s[i]) + passedStr

			if len(passedStr) >= 3 {
				strNum := getStrNum(passedStr)

				if strNum != -1 {
					str += fmt.Sprint(strNum)
					break
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
