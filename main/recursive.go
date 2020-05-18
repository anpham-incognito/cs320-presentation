package main

import "fmt"

var arr []byte

func dequy(currentIndex int, n int) {
	if currentIndex == n {
		fmt.Println(arr)
		return
	}
	arr[currentIndex] = 0
	dequy(currentIndex+1, n)
	arr[currentIndex] = 1
	dequy(currentIndex+1, n)
}

func recursive2n(n int) {
	arr = make([]byte, n)
	dequy(0, n)
}
