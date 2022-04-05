package fitfunction

import "strings"

// func to apply test data on modeled rules
func TestFit(test *[][]string, attribute_values *map[string]int, best_attribute_index int, title *[]string) int {
	classIndex := len(*title) - 1
	true_counts := 0
	for i := 0; i < len(*test); i++ {
		for k := range *attribute_values {
			s := strings.Split(k, ":")
			if (*test)[i][best_attribute_index] == s[0] {
				if (*test)[i][classIndex] == s[1] {
					true_counts++
					break
				}
			}
		}
	}
	return true_counts
}
