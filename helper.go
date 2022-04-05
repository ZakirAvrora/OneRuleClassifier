package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// timer struct that will be used for run time calculation
type Timer struct {
	Minutes     int64
	Seconds     int64
	MiliSeconds int64
	MuSecond    int64
	NanoSeconds int64
}

func TimeCount(durationInMin float64) *Timer {
	var timer Timer

	timer.Minutes = int64(durationInMin)
	durationInMin = (durationInMin - float64(timer.Minutes)) * 60

	timer.Seconds = int64(durationInMin)
	durationInMin = (durationInMin - float64(timer.Seconds)) * 1000

	timer.MiliSeconds = int64(durationInMin)
	durationInMin = (durationInMin - float64(timer.MiliSeconds)) * 1000

	timer.MuSecond = int64(durationInMin)
	durationInMin = (durationInMin - float64(timer.MuSecond)) * 1000

	timer.NanoSeconds = int64(durationInMin)

	return &timer
}

//func to parse file
func ParseFile(filename string) (*[]string, *[][]string, *[][]string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to open file")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var title []string
	var train [][]string
	var test [][]string

	for scanner.Scan() {
		s := strings.TrimSpace(scanner.Text())
		s = strings.ReplaceAll(s, " ", "")
		if len(s) > 0 {
			attributes := strings.Split(s[2:len(s)-1], ",")
			if rune(s[0]) == 'T' {
				title = append(title, attributes...)
				continue
			} else if rune(s[0]) == 'B' {
				test = append(test, attributes)
				continue
			} else if rune(s[0]) == 'A' {
				train = append(train, attributes)
			}
		}
	}
	file.Close()
	return &title, &train, &test
}

//func to ZeroRule
func ZeroRuleValue(train *[][]string, title *[]string) (string, int) {
	newMap := make(map[string]int)
	classIndex := len(*title) - 1
	for i := 0; i < len(*train); i++ {
		newMap[(*train)[i][classIndex]]++
	}

	maxKey := (*train)[0][classIndex]

	for key := range newMap {
		if newMap[key] > newMap[maxKey] {
			maxKey = key
		}
	}

	return maxKey, newMap[maxKey]
}
