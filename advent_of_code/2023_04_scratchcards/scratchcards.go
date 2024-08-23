package scratchcards

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func Solution1() {
	file, err := os.Open("./2023_04_scratchcards/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		count := numberOfMatches(line)
		if count == 0 {
			continue
		}
		sum += int(math.Pow(2, float64(count-1)))
	}
	fmt.Println(sum)
}

func Solution2() {
	file, err := os.Open("./2023_04_scratchcards/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	cards := make([]int, len(lines))
	for i, line := range lines {
		cards[i]++
		matches := numberOfMatches(line)
		for ix := i + 1; ix <= i+matches; ix++ {
			cards[ix] += 1 * cards[i]
		}
	}
	sum := 0
	for _, val := range cards {
		sum += val
	}
	fmt.Println(sum)
}

func numberOfMatches(line string) int {
	winningAndActual := strings.Split(strings.Split(line, ":")[1], " | ")
	s := map[string]struct{}{}
	for _, val := range strings.Split(winningAndActual[0], " ") {
		trimmed := strings.Trim(val, " ")
		if trimmed == "" {
			continue
		}
		s[trimmed] = struct{}{}
	}
	count := 0
	for _, val := range strings.Split(winningAndActual[1], " ") {
		trimmed := strings.Trim(val, " ")
		if trimmed == "" {
			continue
		}
		_, exists := s[trimmed]
		if exists {
			count++
		}
	}
	return count
}
