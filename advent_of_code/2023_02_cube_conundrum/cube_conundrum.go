package cubeconundrum

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Run() {
	file, err := os.Open("./2023_02_cube_conundrum/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
        fmt.Println("power - ", processLinePart2(scanner.Text()) )
        sum += processLinePart2(scanner.Text()) 
	}
	fmt.Println(sum)
}

func processLinePart1(line string) bool {
	games := strings.Split(strings.Split(line, ":")[1], ";")
	for _, game := range games {
		if !processGamePart1(game) {
			return false
		}
	}
	return true
}

func processGamePart1(game string) bool {
	if processColourPart1(game, "red", 12) && processColourPart1(game, "green", 13) && processColourPart1(game, "blue", 13) {
		return true
	}
	return false
}

func processColourPart1(game string, c string, total int64) bool {
	re := regexp.MustCompile(fmt.Sprintf(`\d+ %s`, c))
	colourString := string(re.Find([]byte(game))[:])
	if len(colourString) > 0 {
		i, err := strconv.ParseInt(strings.Split(colourString, " ")[0], 10, 8)
		if err == nil && i > total {
			return false
		}
	}
	return true
}

func processLinePart2(line string) int {
    fmt.Println("red - ", processColourPart2(line, "red") )
    fmt.Println(processColourPart2(line, "green") )
    fmt.Println(processColourPart2(line, "blue") )
	return processColourPart2(line, "red") * processColourPart2(line, "green") * processColourPart2(line, "blue")
}

func processColourPart2(line string, c string) int {
	re := regexp.MustCompile(fmt.Sprintf(`\d+ %s`, c))
	hits := re.FindAll([]byte(line), -1)
	m := -1
	for _, hit := range hits {
		i, err := strconv.ParseInt(strings.Split(string(hit), " ")[0], 10, 8)
        ii := int(i) 
		if err != nil || (ii < m){
			continue
		}
		m = ii

	}
	if m == -1 {
		m = 0
	}
	return m
}
