package main

import (
	"fmt"
	"sort"
)

type People struct {
	Name string
}

type ByName []People

func (bn ByName) Len() int {
	return len(bn)
}
func (bn ByName) Less(i, j int) bool {
	return bn[i].Name < bn[j].Name
}
func (bn ByName) Swap(i, j int) {
	bn[i], bn[j] = bn[j], bn[i]
}

func main() {
	studyGroup := []People{{"Zeno"}, {"John"}, {"Al"}, {"Jenny"}}
	sort.Sort(ByName(studyGroup))
	fmt.Println(studyGroup)

	sort.Sort(sort.Reverse(ByName(studyGroup)))
	fmt.Println(studyGroup)
}
