package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var operations = map[string]func([]int) float64{
	"AVG": avg,
	"SUM": sum,
	"MED": med,
}

func main() {
	op := inputOperation()
	nums := inputNumbers()

	operation := operations[op]
	var res float64 = 0

	if operation != nil {
		res = operation(nums)
	}

	fmt.Printf("Calculation result:%f", res)
}

func inputOperation() string {
	var operation string
	for {
		fmt.Println("Input operation:")
		fmt.Scan(&operation)
		switch operation {
		case "AVG":
			return operation
		case "SUM":
			return operation
		case "MED":
			return operation
		default:
			fmt.Println("Incorrect input, supported operations is AVG,SUM,MED")
		}
	}
}

func inputNumbers() []int {
	for {
		fmt.Println("Input nums divided by `,`")
		var input string
		fmt.Scan(&input)
		inputNums := strings.Split(input, ",")
		fmt.Println(inputNums)
		nums, err := stringsToInts(inputNums)
		if err != nil {
			fmt.Printf("Incorrect number %s\n", err)
			continue
		}
		return nums
	}
}

func stringsToInts(strings []string) ([]int, error) {
	nums := []int{}
	for _, value := range strings {
		num, err := strconv.Atoi(value)
		if err != nil {
			return nums, err
		}
		nums = append(nums, num)
	}
	return nums, nil
}

func avg(nums []int) float64 {
	var num float64

	num = sum(nums)

	return float64(num / float64(len(nums)))
}

func sum(nums []int) float64 {
	var sum int

	for _, value := range nums {
		sum += value
	}

	return float64(sum)
}

func med(nums []int) float64 {
	sort.Ints(nums)
	indx := len(nums) / 2

	if len(nums)%2 == 0 {
		return float64(float64(nums[indx]+nums[indx-1]) / 2)
	}

	return float64(nums[indx])
}
