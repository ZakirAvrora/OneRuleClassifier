package fitfunction

import "strings"

// function to count particluar attribute
func AttrbuteCount(train *[][]string, colNum int, attribute string) int {
	count := 0
	for i := 0; i < len(*train); i++ {
		if (*train)[i][colNum] == attribute {
			count++
		}
	}
	return count
}

// function to select best attribute-class pairs
func SelectBestAtrValues(dict *map[string]int) {
	for key1, val1 := range *dict {
		for key2, val2 := range *dict {
			s1 := strings.Split(key1, ":")
			s2 := strings.Split(key2, ":")
			if s1[0] == s2[0] {
				if val1 > val2 {
					delete(*dict, key2)
				} else if val2 > val1 {
					delete(*dict, key1)
				} else {
					if s1[1] != s2[1] {
						delete(*dict, key2)
					}
				}
			}
		}
	}
}

// function to train data
func TrainFit(train *[][]string, title *[]string) (int, float64, *map[string]int, *map[string]int) {

	classIndex := len(*title) - 1
	var best_accuracy float64 = 0
	var best_attrubute_index int
	var attribute_values *map[string]int
	var best_attibute_counts *map[string]int

	for i := 0; i < len(*title)-1; i++ {
		attibute_counts := make(map[string]int)
		var true_counts float64 = 0
		result := make(map[string]int)
		for j := 0; j < len(*train); j++ {
			result[(*train)[j][i]+":"+(*train)[j][classIndex]]++
		}

		SelectBestAtrValues(&result)

		for k, val := range result {
			s := strings.Split(k, ":")
			attibute_counts[s[0]] = AttrbuteCount(train, i, s[0])
			true_counts += float64(val)
		}

		var accuracy float64 = true_counts / float64(len(*train))
		if best_accuracy < accuracy {
			best_attrubute_index = i
			best_accuracy = accuracy
			attribute_values = &result
			best_attibute_counts = &attibute_counts
		}
	}
	return best_attrubute_index, best_accuracy, attribute_values, best_attibute_counts
}
