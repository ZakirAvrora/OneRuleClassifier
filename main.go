package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"

	fit "github.com/ZakirAvrora/OneRuleClassifier/fitFunction"
)

func main() {

	start := time.Now()

	filePath := os.Args[1]

	if len(os.Args) != 2 {
		log.Fatalf("Specify single filename path")
	}

	title, train, test := ParseFile(filePath) // parse file returns pointer

	best_attribute, best_accuracy, attribute_values, attibute_counts := fit.TrainFit(train, title)
	durationTraining := time.Since(start).Minutes()

	startTest := time.Now()
	test_trueCounts := fit.TestFit(test, attribute_values, best_attribute, title)
	durationTesting := time.Since(startTest).Minutes()

	//--------------- Printing the result on the console ------------------------------------
	PrintText := fmt.Sprintf("%v: (%v/%v)\n", (*title)[best_attribute], best_accuracy*float64(len(*train)), len(*train))

	keys := make([]string, 0, len(*attibute_counts))
	for k := range *attibute_counts {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, v1 := range keys {
		for k2, v2 := range *attribute_values {
			s := strings.Split(k2, ":")
			if v1 == s[0] {
				PrintText += fmt.Sprintf("  %v -> %v (%v/%v)\n", s[0], s[1], v2, (*attibute_counts)[v1])
			}
		}
	}

	ans, res := ZeroRuleValue(train, title)
	PrintText += fmt.Sprintf("  [ZeroRule] -> %v (%v/%v)\n\n\n", ans, res, len(*train))

	timerTrain := TimeCount(durationTraining)

	timerTest := TimeCount(durationTesting)

	PrintText += fmt.Sprintf("Training time: %v minutes %v seconds %v milliseconds %v microseconds %v nanoseconds\n\n",
		timerTrain.Minutes, timerTrain.Seconds, timerTrain.MiliSeconds, timerTrain.MuSecond, timerTrain.NanoSeconds)

	PrintText += fmt.Sprintf("Accuracy on test data: %v/%v\n\n", test_trueCounts, len(*test))

	PrintText += fmt.Sprintf("Testing time: %v minutes %v seconds %v milliseconds %v microseconds %v nanoseconds",
		timerTest.Minutes, timerTest.Seconds, timerTest.MiliSeconds, timerTest.MuSecond, timerTest.NanoSeconds)

	fmt.Println(PrintText)
}
