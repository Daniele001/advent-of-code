package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func twoNumbers(slice []int, r *int) {
	for i, j := 0, len(slice) - 1; i < len(slice) || j > 0;{
		if slice[i] + slice[j] > 2020 && slice[j] > slice[i] {
			j--
		} else if slice[i] + slice[j] < 2020 && slice[j] > slice[i] {
			i++
		}
		if slice[i] + slice[j] == 2020 {
			*r = slice[i] * slice[j]
			break
		}
	}
}

func threeNumbers(slice []int, r *int) {
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice); j++ {
			for k := 0; k < len(slice); k++ {
				if slice[i] + slice[j] + slice[k] == 2020 {
					*r = slice[i] * slice[j] * slice[k]
					return
				}
			}
		}
	}
}

func main() {
	var data []int
	data = make([]int, 0)
	buffer := bufio.NewScanner(os.Stdin)
	for buffer.Scan() {
		n, _ := strconv.Atoi(buffer.Text())
		data = append(data, n)
	}
	sort.Ints(data)
	var res int
	twoNumbers(data, &res)
	fmt.Println(res)
	threeNumbers(data, &res)
	fmt.Println(res)

}