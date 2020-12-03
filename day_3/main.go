package main

import (
	"bufio"
	"fmt"
	"os"
)

// This function return the number of trees for the slope right 1, down 1
func r1D1(t []string) (tree int) {
	index := 0
	for i := 1; i < len(t); i++{
		index += 1
		if index > 30 {
			index = index - 31
		}
		if string(t[i][index]) == "#" {
			tree++
		}
	}
	return
}

// This function return the number of trees for the slope right 3, down 1
func r3D1(t []string) (tree int) {
	index := 0
	for i := 1; i < len(t); i++{
		index += 3
		if index > 30 {
			index = index - 31
		}
		if string(t[i][index]) == "#" {
			tree++
		}
	}
	return
}

// This function return the number of trees for the slope right 5, down 1
func r5D1(t []string) (tree int) {
	index := 0
	for i := 1; i < len(t); i++{
		index += 5
		if index > 30 {
			index = index - 31
		}
		if string(t[i][index]) == "#" {
			tree++
		}
	}
	return
}

// This function return the number of trees for the slope right 7, down 1
func r7D1(t []string) (tree int) {
	index := 0
	for i := 1; i < len(t); i++{
		index += 7
		if index > 30 {
			index = index - 31
		}
		if string(t[i][index]) == "#" {
			tree++
		}
	}
	return
}

// This function return the number of trees for the slope right 1, down 2
func r1D2(t []string) (tree int) {
	index := 0
	for i := 2; i < len(t); i = i + 2{
		index += 1
		if index > 30 {
			index = index - 31
		}
		if string(t[i][index]) == "#" {
			tree++
		}
	}
	return
}

func main() {
	buffer := bufio.NewScanner(os.Stdin)
	var text []string
	for buffer.Scan() {
		text = append(text, buffer.Text())
	}
	mul := r1D1(text) * r3D1(text) * r5D1(text) * r7D1(text) * r1D2(text)
	fmt.Printf("%d\n%d\n", r3D1(text), mul)
}
