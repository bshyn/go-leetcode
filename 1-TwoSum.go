package main

func twoSum(nums []int, target int) []int {
	numbersMap := make(map[int]int)
	for index, currNumber := range nums {
		numToFind := target - currNumber
		if _, isPresent := numbersMap[numToFind]; isPresent {
			return []int{numbersMap[numToFind], index}
		}
		numbersMap[currNumber] = index
	}
	return nil
}
