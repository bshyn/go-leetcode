package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func arrayToLinkedList(arr []int) *ListNode {
	var head *ListNode
	var last *ListNode
	for index, value := range arr {
		if index == 0 {
			head = &ListNode{Val: value}
			last = head
		} else {
			new := ListNode{Val: value}
			last.Next = &new
			last = &new
		}
	}

	return head
}

func printAllValuesInLinkedList(n *ListNode) {
	print(n.Val)
	for node := *n; node.Next != nil; {
		node = *node.Next
		print(strconv.Itoa(node.Val))
	}
	fmt.Println()
}

func mergeSortedArrays(arr1 []int, arr2 []int) []int {
	var merged []int

	for ; len(arr1) != 0 && len(arr2) != 0; {
		if arr1[0] < arr2[0]{
			merged = append(merged, arr1[0])
			arr1 = remove(arr1, 0)
		} else {
			merged = append(merged, arr2[0])
			arr2 = remove(arr2, 0)
		}
	}

	for i := 0; i < len(arr1); i++ {
		merged = append(merged, arr1[i])
	}

	for i := 0; i < len(arr2); i++ {
		merged = append(merged, arr2[i])
	}

	return merged
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}
