package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type integers []int

func newIntegers() integers {
	ints := integers{}

	for i := 1; i < 10; i++ {
		for _, j := range []int{10, 100, 1000} {
			ints = append(ints, i*j)
		}
	}
	return ints
}

func (ints integers) subset(startIndex int, count int) integers {
	return ints[startIndex : startIndex+count]
}

func (ints integers) split(splitIndex int) (integers, integers) {
	return ints[:splitIndex], ints[splitIndex:]
}

func (ints integers) print() {
	for i, v := range ints {
		fmt.Println(i, ":", v)
	}
}

func (ints *integers) reset() {
	for i, _ := range *ints {
		(*ints)[i] = 0
	}
}

func (ints integers) toString() string {
	i := []int(ints)
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(i)), ","), "[]")
}

func (ints integers) toByteSlice() []byte {
	return []byte(ints.toString())
}

func (ints integers) saveTofile(filename string) {
	ioutil.WriteFile(filename, ints.toByteSlice(), 0666)
}

func newIntegersFromFile(filename string) integers {
	bytes, _ := ioutil.ReadFile(filename)

	str := string(bytes)
	strs := strings.Split(str, ",")

	var ints = make([]int, len(strs))

	for i, s := range strs {
		ints[i], _ = strconv.Atoi(s)
	}

	return ints

}
