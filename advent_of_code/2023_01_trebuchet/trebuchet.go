package trebuchet

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"unicode/utf8"
)

func Run() {
	file, err := os.Open("./2023_01_trebuchet/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := int64(0)
	for scanner.Scan() {
		i, err := processLine2(scanner.Text())
		if err != nil {
			panic(err)
		}
		sum += i
	}
	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func processLine1(line string) (int64, error) {
	l := 0
	r := utf8.RuneCountInString(line) - 1
	n1 := -1
	n2 := -1
	for l <= utf8.RuneCountInString(line)-1 || r >= 0 {
		if line[l] >= 48 && line[l] <= 57 && n1 == -1 {
			n1 = int(line[l]) - 48
		}
		if line[r] >= 48 && line[r] <= 57 && n2 == -1 {
			n2 = int(line[r]) - 48
		}
		l++
		r--
	}
	return strconv.ParseInt(fmt.Sprintf("%d%d", n1, n2), 10, 8)
}

func processLine2(line string) (int64, error) {
	regex := "zero|one|two|three|four|five|six|seven|eight|nine|0|1|2|3|4|5|6|7|8|9"
	r, err := regexp.Compile(regex)
	lineR := reverse(line)
	rr, err := regexp.Compile(reverse(regex))
	if err != nil {
		return 0, err
	}

	return strconv.ParseInt(convertStringToInt(r.FindString(line))+convertStringToInt(reverse(rr.FindString(lineR))), 10, 64)

}

func convertStringToInt(s string) string {
	if s == "one" {
		return "1"
	}
	if s == "two" {
		return "2"
	}
	if s == "three" {
		return "3"
	}
	if s == "four" {
		return "4"
	}
	if s == "five" {
		return "5"
	}
	if s == "six" {
		return "6"
	}
	if s == "seven" {
		return "7"
	}
	if s == "eight" {
		return "8"
	}
	if s == "nine" {
		return "9"
	}
	return s
}

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}
