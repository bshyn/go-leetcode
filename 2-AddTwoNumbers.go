package main

import (
	"strconv"
)

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	firstListNumbers := strconv.Itoa(l1.Val)
	secondListNumbers := strconv.Itoa(l2.Val)

	for node := *l1; node.Next != nil; {
		node = *node.Next
		firstListNumbers = firstListNumbers + strconv.Itoa(node.Val)
	}
	for node := *l2; node.Next != nil; {
		node = *node.Next
		secondListNumbers = secondListNumbers + strconv.Itoa(node.Val)
	}

	var longestString *string
	if len(firstListNumbers) > len(secondListNumbers) {
		longestString = &firstListNumbers
	} else {
		longestString = &secondListNumbers
	}

	sumToNext := false
	var head *ListNode
	var last *ListNode

	for index, _ := range *longestString {
		var firstNumber int
		var secondNumber int
		if len(firstListNumbers) <= index {
			firstNumber = 0
		} else {
			firstNumber, _ = strconv.Atoi(string(firstListNumbers[index]))
		}

		if len(secondListNumbers) <= index {
			secondNumber = 0
		} else {
			secondNumber, _ = strconv.Atoi(string(secondListNumbers[index]))
		}
		result := firstNumber + secondNumber

		if sumToNext {
			result++
			sumToNext = false
		}

		if result > 9 {
			result = result - 10
			sumToNext = true
		}

		if index == 0 {
			head = &ListNode{Val: result}
			last = head
		} else {
			new := ListNode{Val: result}
			last.Next = &new
			last = &new
		}
	}

	if sumToNext {
		new := ListNode{Val: 1}
		last.Next = &new
		last = &new
	}

	return head
}
