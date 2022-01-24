package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortString(str string) string {
	str_arr := strings.Split(str, "")
	sort.Strings(str_arr)
	return (strings.Join(str_arr, ""))
}

func find_anagram(words_string []string) {
	my_map := make(map[string][]string)
	sorted_array := []string{}
	for _, val := range words_string {
		sorted_val := sortString(val)
		sorted_array = append(sorted_array, sorted_val)
	}
	fmt.Println(sorted_array)
	for i := 0; i < len(sorted_array); i++ {
		for j := i + 1; j < len(sorted_array); j++ {
			if sorted_array[i] == sorted_array[j] {
				if my_map[sorted_array[i]] == nil {
					my_map[sorted_array[i]] = append(my_map[sorted_array[i]], words_string[i], words_string[j])
				} else {
					my_map[sorted_array[i]] = append(my_map[sorted_array[i]], words_string[j])
				}
			} else {
				if my_map[sorted_array[i]] == nil {
					my_map[sorted_array[i]] = append(my_map[sorted_array[i]], words_string[i])
				}
			}
		}
	}
	for _, val := range my_map {
		fmt.Println(val)
	}
}

func main() {
	words_string := []string{"cat", "act", "boat", "coat", "arm", "ram", "bare", "bear"}
	find_anagram(words_string)
}
