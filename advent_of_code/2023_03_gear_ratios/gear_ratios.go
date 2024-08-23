package gearratios

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode/utf8"
)

func Run() {
	file, err := os.Open("./2023_03_gear_ratios/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	processLines2(lines)
}

func processLines1(lines []string) {
	sum := 0
	for i, line := range lines {
		curr := ""
		for j, char := range line {
			if char >= 48 && char <= 57 {
				curr += string(char)
				continue
			}

			if curr != "" {
				if checkForSymbols1(i, j-1, utf8.RuneCountInString(curr), lines) {
					val, _ := strconv.ParseInt(curr, 10, 64)
					fmt.Println(val)
					sum += int(val)
				}
				curr = ""
			}
		}
		if curr != "" {
			if checkForSymbols1(i, utf8.RuneCountInString(line)-1, utf8.RuneCountInString(curr), lines) {
				val, _ := strconv.ParseInt(curr, 10, 64)
				fmt.Println(val)
				sum += int(val)
			}
			curr = ""
		}
	}
	fmt.Println(sum)
}

func checkForSymbols1(iVal int, jEnd int, numLen int, lines []string) bool {
	for i := iVal - 1; i <= iVal+1; i++ {
		for j := jEnd - numLen; j <= jEnd+1; j++ {
			if i < 0 || i >= len(lines) || j < 0 || j >= utf8.RuneCountInString(lines[i]) {
				continue
			}
			char := rune(lines[i][j])
			if char >= 48 && char <= 57 {
				continue
			}
			if char == 46 {
				continue
			}
			fmt.Println(string(char))
			return true
		}
	}

	return false
}

func processLines2(lines []string) {
    sum := 0
	for i, line := range lines {
		for j, char := range line {
			if char == 42 {
				sum += gearRatio(i, j, lines)
			}
		}
	}
    fmt.Println(sum)
	// FIND CONNECTED NUMBERS
}

func gearRatio(iGear int, jGear int, lines []string) int {
	// check left [j-1] -> looks left
	// check right [j+1] -> looks right
	// middle -> if left and/or right append

	var ratios []int

	for i := iGear - 1; i <= iGear+1; i++ {
		if i < 0 || i >= len(lines) {
			continue
		}
		left := ""
		right := ""
		j := jGear - 1
		for j >= 0 && lines[i][j] >= 48 && lines[i][j] <= 57 {
			left = string(lines[i][j]) + left
			j--
		}

		j = jGear + 1
		for j < utf8.RuneCountInString(lines[i]) && lines[i][j] >= 48 && lines[i][j] <= 57 {
			right += string(lines[i][j])
			j++
		}
		if lines[i][jGear] >= 48 && lines[i][jGear] <= 57 {
			val, _ := strconv.ParseInt(left+string(lines[i][jGear])+right, 10, 64)
			ratios = append(ratios, int(val))
		} else {
			if utf8.RuneCountInString(left) > 0 {
				val, _ := strconv.ParseInt(left, 10, 64)
				ratios = append(ratios, int(val))
			}
			if utf8.RuneCountInString(right) > 0 {
				val, _ := strconv.ParseInt(right, 10, 64)
				ratios = append(ratios, int(val))

			}

		}

	}
    if len(ratios) == 2{
        return ratios[0] * ratios[1]
    }
    return 0
}
