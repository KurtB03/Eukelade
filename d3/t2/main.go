package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

func main() {

	dat, err := ioutil.ReadFile("d3/t2/input.txt")
	check(err)

	//fmt.Println(string(dat))

	s := strings.Split(string(dat), "\n")
	result := 1
	count := 0
	j := 0
	//m := (len(s) * 3) / 30
	for _, t := range s {

		//t := strings.Repeat(s, m)
		if t[j] == '#' {
			count++
			//fmt.Println(replaceAtIndex(t, 'X', j))
		} else {
			//fmt.Println(replaceAtIndex(t, '0', j))
		}
		j = (j + 1) % 31
	}
	result = result * count
	fmt.Println(count)
	j = 0
	count = 0
	for _, t := range s {

		//t := strings.Repeat(s, m)
		if t[j] == '#' {
			count++
			//fmt.Println(replaceAtIndex(t, 'X', j))
		} else {
			//fmt.Println(replaceAtIndex(t, '0', j))
		}
		j = (j + 3) % 31
	}
	result = result * count
	fmt.Println(count)
	j = 0
	count = 0
	for _, t := range s {

		//t := strings.Repeat(s, m)
		if t[j] == '#' {
			count++
			//fmt.Println(replaceAtIndex(t, 'X', j))
		} else {
			//fmt.Println(replaceAtIndex(t, '0', j))
		}
		j = (j + 5) % 31
	}
	result = result * count
	fmt.Println(count)
	j = 0
	count = 0
	for _, t := range s {

		//t := strings.Repeat(s, m)
		if t[j] == '#' {
			count++
			//fmt.Println(replaceAtIndex(t, 'X', j))
		} else {
			//fmt.Println(replaceAtIndex(t, '0', j))
		}
		j = (j + 7) % 31
	}
	result = result * count
	fmt.Println(count)
	j = 0
	count = 0
	for k, t := range s {
		if k%2 == 0 {
			if t[j] == '#' {
				count++
				//fmt.Println(replaceAtIndex(t, 'X', j))
			} else {
				//fmt.Println(replaceAtIndex(t, '0', j))
			}
			//fmt.Println(k)
		}
		//t := strings.Repeat(s, m)

		j = (j + 1) % 31
	}
	result = result * count
	fmt.Println(count)
	j = 0

	fmt.Println(result)
}
