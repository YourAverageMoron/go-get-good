package trebuchet

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
		i, err := processLine(scanner.Text())
		if err != nil {
			panic(err)
		}
		fmt.Println(i)
        sum += i
	}
    fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func processLine(line string) (int64, error) {
	l := 0
	r := utf8.RuneCountInString(line) - 1
	n1 := -1
	n2 := -1
	for l <= utf8.RuneCountInString(line) - 1 || r >=0 {
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
