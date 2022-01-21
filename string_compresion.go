package main

import "fmt"

func out(str string, repeatition int) string {
	tmp_str := ""
	for i := 0; i <= repeatition; i++ {
		tmp_str = tmp_str + str
	}
	return tmp_str
}

func decompress(compressed_arr []byte) {
	result_str := ""
	i := 0
	for i <= (len(compressed_arr) - 2) {
		result_str = result_str + out(string(compressed_arr[i]), int(compressed_arr[i+1]))
		i = i + 2
	}
	fmt.Println(result_str)
}

func main() {
	my_sting := "abbbbcddde"
	char_arr := []byte(my_sting)
	compressed_arr := []byte{}
	count := 1
	for i := 0; i <= (len(char_arr) - 1); i++ {
		if i == (len(char_arr) - 1) {
			compressed_arr = append(compressed_arr, char_arr[i])
			compressed_arr = append(compressed_arr, byte(count))
		} else {
			if char_arr[i] == char_arr[i+1] {
				count++
			} else {
				compressed_arr = append(compressed_arr, char_arr[i])
				compressed_arr = append(compressed_arr, byte(count))
				count = 1
			}
		}
	}
	fmt.Println(compressed_arr)

	decompress(compressed_arr)
}
