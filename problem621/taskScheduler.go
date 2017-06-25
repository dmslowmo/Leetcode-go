package main

import (
	"fmt"
	"sort"
)

// A data structure to hold key/value pairs
type Pair struct {
	Key   byte
	Value int
}

// A slice of pairs that implements sort.Interface to sort by values
type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

func leastInterval(tasks []byte, n int) int {
	uniqueTasks := make(map[byte]int)
	for i:=0;i<len(tasks);i++ {
		if count, ok := uniqueTasks[tasks[i]]; ok {
			uniqueTasks[tasks[i]] = count + 1
		} else {
			uniqueTasks[tasks[i]] = 1
		}
	}

	// sort the sequence based on values ie occurrences
	p := make(PairList, len(uniqueTasks))

	i := 0
	for k, v := range uniqueTasks {
		p[i] = Pair{k, v}
		i++
	}

	sort.Sort(sort.Reverse(p))

	//for _, member := range p {
	//	fmt.Printf("%s: %d\n", string(member.Key), member.Value)
	//}

	taskNoMap := make(map[byte][]int)
	compactSequence := []byte{}
	taskNo := 0
	zeros := 0
	subSeq := []byte{}
	for {
		for i:=0;i<len(p);i++ {
			if len(subSeq) > n {
				subSeq = []byte{}
				sort.Sort(sort.Reverse(p))
				break
			}
			if zeros == len(p) {
				break
			}
			if p[i].Value <= 0 {
				continue
			} else {
				if contains(subSeq, p[i].Key) {
					compactSequence = append(compactSequence, '-')
					subSeq = append(subSeq, '-')
				} else {
					compactSequence = append(compactSequence, p[i].Key)
					p[i].Value = p[i].Value - 1
					if p[i].Value == 0 {
						zeros++
					}
					subSeq = append(subSeq, p[i].Key)
					updateTaskNoMap(taskNoMap, p, i, taskNo)
				}
				taskNo++
			}
		}
		if zeros == len(p) {
			break
		}
	}

	for i := 0; i < len(compactSequence); i++ {
		fmt.Print(string(compactSequence[i]))
	}
	fmt.Println()
	fmt.Println(len(compactSequence))
	return len(compactSequence)
}

func updateTaskNoMap(taskNoMap map[byte][]int, p PairList, i int, taskNo int) {
	if _, ok := taskNoMap[p[i].Key]; ok {
		taskNoMap[p[i].Key] = append(taskNoMap[p[i].Key], taskNo)
	} else {
		taskNoMap[p[i].Key] = []int{taskNo}
	}
}

func contains(s []byte, e byte) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(leastInterval([]byte{'A','A','A','B','B','B'}, 2) == 8)
	fmt.Println()
	fmt.Println(leastInterval([]byte{'A','A','A','B','B','B'}, 1) == 6)
	fmt.Println()
	fmt.Println(leastInterval([]byte{'A','B','B','B'}, 2) == 7)
	fmt.Println()
	fmt.Println(leastInterval([]byte{'A','B','B','B'}, 0) == 4)
	//fmt.Println()
	////longInput := []byte{'G','C','A','H','A','G','G','F','G','J','H','C','A','G','E','A','H','E','F','D','B','D','H','H','E','G','F','B','C','G','F','H','J','F','A','C','G','D','I','J','A','G','D','F','B','F','H','I','G','J','G','H','F','E','H','J','C','E','H','F','C','E','F','H','H','I','G','A','G','D','C','B','I','D','B','C','J','I','B','G','C','H','D','I','A','B','A','J','C','E','B','F','B','J','J','D','D','H','I','I','B','A','E','H','J','J','A','J','E','H','G','B','F','C','H','C','B','J','B','A','H','B','D','I','F','A','E','J','H','C','E','G','F','G','B','G','C','G','A','H','E','F','H','F','C','G','B','I','E','B','J','D','B','B','G','C','A','J','B','J','J','F','J','C','A','G','J','E','G','J','C','D','D','A','I','A','J','F','H','J','D','D','D','C','E','D','D','F','B','A','J','D','I','H','B','A','F','E','B','J','A','H','D','E','I','B','H','C','C','C','G','C','B','E','A','G','H','H','A','I','A','B','A','D','A','I','E','C','C','D','A','B','H','D','E','C','A','H','B','I','A','B','E','H','C','B','A','D','H','E','J','B','J','A','B','G','J','J','F','F','H','I','A','H','F','C','H','D','H','C','C','E','I','G','J','H','D','E','I','J','C','C','H','J','C','G','I','E','D','E','H','J','A','H','D','A','B','F','I','F','J','J','H','D','I','C','G','J','C','C','D','B','E','B','E','B','G','B','A','C','F','E','H','B','D','C','H','F','A','I','A','E','J','F','A','E','B','I','G','H','D','B','F','D','B','I','B','E','D','I','D','F','A','E','H','B','I','G','F','D','E','B','E','C','C','C','J','J','C','H','I','B','H','F','H','F','D','J','D','D','H','H','C','D','A','J','D','F','D','G','B','I','F','J','J','C','C','I','F','G','F','C','E','G','E','F','D','A','I','I','H','G','H','H','A','J','D','J','G','F','G','E','E','A','H','B','G','A','J','J','E','I','H','A','G','E','C','D','I','B','E','A','G','A','C','E','B','J','C','B','A','D','J','E','J','I','F','F','C','B','I','H','C','F','B','C','G','D','A','A','B','F','C','D','B','I','I','H','H','J','A','F','J','F','J','F','H','G','F','D','J','G','I','E','B','C','G','I','F','F','J','H','H','G','A','A','J','C','G','F','B','A','A','E','E','A','E','I','G','F','D','B','I','F','A','B','J','F','F','J','B','F','J','F','J','F','I','E','J','H','D','G','G','D','F','G','B','J','F','J','A','J','E','G','H','I','E','G','D','I','B','D','J','A','A','G','A','I','I','A','A','I','I','H','E','C','A','G','I','F','F','C','D','J','J','I','A','A','F','C','J','G','C','C','H','E','A','H','F','B','J','G','I','A','A','H','G','B','E','G','D','I','C','G','J','C','C','I','H','B','D','J','H','B','J','H','B','F','J','E','J','A','G','H','B','E','H','B','F','F','H','E','B','E','G','H','J','G','J','B','H','C','H','A','A','B','E','I','H','B','I','D','J','J','C','D','G','I','J','G','J','D','F','J','E','F','D','E','B','D','B','C','B','B','C','C','I','F','D','E','I','G','G','I','B','H','G','J','A','A','H','I','I','H','A','I','F','C','D','A','C','G','E','G','E','E','H','D','C','G','D','I','A','G','G','D','A','H','H','I','F','E','I','A','D','H','B','B','G','I','C','G','B','I','I','D','F','F','C','C','A','I','E','A','E','J','A','H','C','D','A','C','B','G','H','G','J','G','I','H','B','A','C','H','I','D','D','C','F','G','B','H','E','B','B','H','C','B','G','G','C','F','B','E','J','B','B','I','D','H','D','I','I','A','A','H','G','F','B','J','F','D','E','G','F','A','G','G','D','A','B','B','B','J','A','F','H','H','D','C','J','I','A','H','G','C','J','I','F','J','C','A','E','C','H','J','H','H','F','G','E','A','C','F','J','H','D','G','G','D','D','C','B','H','B','C','E','F','B','D','J','H','J','J','J','A','F','F','D','E','F','C','I','B','H','H','D','E','A','I','A','B','F','G','F','F','I','E','E','G','A','I','D','F','C','H','E','C','G','H','F','F','H','J','H','G','A','E','H','B','G','G','D','D','D','F','I','A','F','F','D','E','H','J','E','D','D','A','J','F','E','E','E','F','I','D','A','F','F','J','E','I','J','D','D','G','A','C','G','G','I','E','G','E','H','E','D','E','J','B','G','I','J','C','H','C','C','A','A','B','C','G','B','D','I','D','E','H','J','J','B','F','E','J','H','H','I','G','B','D'}
	////fmt.Println(len(longInput))
	fmt.Println(leastInterval([]byte{'G','C','A','H','A','G','G','F','G','J','H','C','A','G','E','A','H','E','F','D','B','D','H','H','E','G','F','B','C','G','F','H','J','F','A','C','G','D','I','J','A','G','D','F','B','F','H','I','G','J','G','H','F','E','H','J','C','E','H','F','C','E','F','H','H','I','G','A','G','D','C','B','I','D','B','C','J','I','B','G','C','H','D','I','A','B','A','J','C','E','B','F','B','J','J','D','D','H','I','I','B','A','E','H','J','J','A','J','E','H','G','B','F','C','H','C','B','J','B','A','H','B','D','I','F','A','E','J','H','C','E','G','F','G','B','G','C','G','A','H','E','F','H','F','C','G','B','I','E','B','J','D','B','B','G','C','A','J','B','J','J','F','J','C','A','G','J','E','G','J','C','D','D','A','I','A','J','F','H','J','D','D','D','C','E','D','D','F','B','A','J','D','I','H','B','A','F','E','B','J','A','H','D','E','I','B','H','C','C','C','G','C','B','E','A','G','H','H','A','I','A','B','A','D','A','I','E','C','C','D','A','B','H','D','E','C','A','H','B','I','A','B','E','H','C','B','A','D','H','E','J','B','J','A','B','G','J','J','F','F','H','I','A','H','F','C','H','D','H','C','C','E','I','G','J','H','D','E','I','J','C','C','H','J','C','G','I','E','D','E','H','J','A','H','D','A','B','F','I','F','J','J','H','D','I','C','G','J','C','C','D','B','E','B','E','B','G','B','A','C','F','E','H','B','D','C','H','F','A','I','A','E','J','F','A','E','B','I','G','H','D','B','F','D','B','I','B','E','D','I','D','F','A','E','H','B','I','G','F','D','E','B','E','C','C','C','J','J','C','H','I','B','H','F','H','F','D','J','D','D','H','H','C','D','A','J','D','F','D','G','B','I','F','J','J','C','C','I','F','G','F','C','E','G','E','F','D','A','I','I','H','G','H','H','A','J','D','J','G','F','G','E','E','A','H','B','G','A','J','J','E','I','H','A','G','E','C','D','I','B','E','A','G','A','C','E','B','J','C','B','A','D','J','E','J','I','F','F','C','B','I','H','C','F','B','C','G','D','A','A','B','F','C','D','B','I','I','H','H','J','A','F','J','F','J','F','H','G','F','D','J','G','I','E','B','C','G','I','F','F','J','H','H','G','A','A','J','C','G','F','B','A','A','E','E','A','E','I','G','F','D','B','I','F','A','B','J','F','F','J','B','F','J','F','J','F','I','E','J','H','D','G','G','D','F','G','B','J','F','J','A','J','E','G','H','I','E','G','D','I','B','D','J','A','A','G','A','I','I','A','A','I','I','H','E','C','A','G','I','F','F','C','D','J','J','I','A','A','F','C','J','G','C','C','H','E','A','H','F','B','J','G','I','A','A','H','G','B','E','G','D','I','C','G','J','C','C','I','H','B','D','J','H','B','J','H','B','F','J','E','J','A','G','H','B','E','H','B','F','F','H','E','B','E','G','H','J','G','J','B','H','C','H','A','A','B','E','I','H','B','I','D','J','J','C','D','G','I','J','G','J','D','F','J','E','F','D','E','B','D','B','C','B','B','C','C','I','F','D','E','I','G','G','I','B','H','G','J','A','A','H','I','I','H','A','I','F','C','D','A','C','G','E','G','E','E','H','D','C','G','D','I','A','G','G','D','A','H','H','I','F','E','I','A','D','H','B','B','G','I','C','G','B','I','I','D','F','F','C','C','A','I','E','A','E','J','A','H','C','D','A','C','B','G','H','G','J','G','I','H','B','A','C','H','I','D','D','C','F','G','B','H','E','B','B','H','C','B','G','G','C','F','B','E','J','B','B','I','D','H','D','I','I','A','A','H','G','F','B','J','F','D','E','G','F','A','G','G','D','A','B','B','B','J','A','F','H','H','D','C','J','I','A','H','G','C','J','I','F','J','C','A','E','C','H','J','H','H','F','G','E','A','C','F','J','H','D','G','G','D','D','C','B','H','B','C','E','F','B','D','J','H','J','J','J','A','F','F','D','E','F','C','I','B','H','H','D','E','A','I','A','B','F','G','F','F','I','E','E','G','A','I','D','F','C','H','E','C','G','H','F','F','H','J','H','G','A','E','H','B','G','G','D','D','D','F','I','A','F','F','D','E','H','J','E','D','D','A','J','F','E','E','E','F','I','D','A','F','F','J','E','I','J','D','D','G','A','C','G','G','I','E','G','E','H','E','D','E','J','B','G','I','J','C','H','C','C','A','A','B','C','G','B','D','I','D','E','H','J','J','B','F','E','J','H','H','I','G','B','D'}, 1) == 1000)
	fmt.Println()
	//in := []byte{'F','J','J','A','J','F','C','H','J','B','E','G','G','F','A','C','I','F','J','C','J','C','H','C','A','D','G','H','B','F','G','C','C','A','E','B','H','J','E','I','F','D','E','A','C','D','B','D','J','J','C','F','D','D','J','H','A','E','C','D','J','D','G','G','B','C','E','G','H','I','D','H','F','E','I','B','D','E','I','E','C','J','G','I','D','E','D','J','C','A','C','C','D','I','J','B','D','H','H','J','G','B','G','A','H','E','H','E','D','E','J','E','J','C','F','C','J','G','B','C','I','I','H','F','A','D','G','F','C','C','F','G','C','J','B','B','I','C','J','J','E','G','H','C','I','G','J','I','G','G','J','G','G','E','G','B','I','J','B','H','D','H','G','F','C','H','C','D','A','G','B','H','H','B','J','C','A','F','J','G','F','E','B','F','E','B','B','A','E','F','E','H','I','I','C','G','J','D','H','E','F','G','G','D','E','B','F','J','J','J','D','H','E','B','D','J','I','F','C','I','E','H','F','E','G','D','E','C','F','E','D','E','A','I','E','A','D','H','G','C','I','E','G','A','H','I','G','G','A','G','F','H','J','D','F','A','G','H','B','J','A','H','B','H','C','G','F','A','C','C','B','I','G','G','B','C','J','J','I','E','G','D','I','J','I','C','G','A','J','G','F','J','F','C','F','G','J','I','E','B','G','F','A','D','A','I','A','E','H','F','D','D','C','B','J','I','J','H','I','C','D','A','G','F','I','B','E','D','C','J','G','I','H','E','C','E','I','I','B','B','H','J','C','F','I','D','B','F','H','F','A','C','A','A','B','D','C','A','G','B','G','F','E','G','A','A','A','C','J','H','H','G','C','C','B','C','E','B','E','F','I','E','E','D','I','H','G','F','A','H','B','J','B','G','H','C','C','B','G','C','B','A','E','G','A','J','G','D','C','I','G','F','G','G','A','J','E','I','D','E','A','F','A','H','C','E','D','D','D','H','I','F','F','A','F','A','A','C','J','D','J','H','I','F','A','C','B','C','A','C','C','H','A','J','I','B','A','I','F','J','C','I','B','C','E','E','E','J','G','F','E','I','A','A','E','B','J','H','H','H','A','H','J','E','F','E','F','G','J','D','I','D','I','F','B','J','D','A','A','D','F','G','B','J','H','F','A','D','H','C','B','A','J','H','I','F','H','E','G','B','A','F','F','A','C','D','G','I','I','J','H','H','C','J','G','B','A','D','B','F','J','D','I','A','F','F','F','F','A','E','B','C','G','H','E','B','B','A','G','D','C','C','E','A','C','F','G','A','I','F','B','H','J','G','C','B','H','D','A','H','B','H','H','C','A','F','I','C','F','A','C','J','I','H','H','F','B','B','D','E','C','J','F','C','E','A','J','E','C','A','E','B','A','J','F','J','J','J','H','H','C','I','E','G','G','H','J','J','H','H','H','J','H','A','G','I','C','E','C','D','G','G','F','H','D','G','H','A','E','I','D','A','H','G','E','A','B','F','I','C','A','F','B','A','I','F','G','I','F','D','A','B','J','B','D','F','G','J','J','A','A','C','H','G','F','B','I','I','J','A','H','D','F','E','F','J','B','F','C','G','E','A','G','H','E','H','H','F','I','G','C','C','G','J','B','H','F','H','D','I','B','D','I','F','H','I','D','F','G','G','E','A','C','A','G','H','G','H','J','F','D','F','G','D','D','C','J','C','J','G','G','G','G','H','H','G','D','E','H','G','C','B','F','I','F','C','H','J','I','A','F','D','C','F','C','E','E','D','D','C','G','B','F','E','J','C','I','E','D','B','B','I','I','I','H','C','E','C','J','F','G','A','I','J','D','I','C','G','F','I','E','I','E','F','A','G','E','J','A','I','A','D','A','G','J','F','E','D','I','A','E','J','I','C','J','B','F','B','E','C','E','F','G','E','J','J','I','E','D','F','C','H','H','B','G','D','I','I','F','B','G','C','F','J','B','G','J','H','D','G','C','C','I','I','E','I','B','H','B','I','G','F','H','G','C','J','D','C','E','G','F','C','H','D','A','C','D','H','B','C','H','I','B','A','J','C','B','D','J','D','H','F','B','A','G','G','J','I','E','F','A','D','H','D','B','C','A','H','F','G','B','F','H','B','H','I','J','D','H','I','B','C','D','G','A','E','A','A','I','F','I','F','B','B','I','F','A','E','I','A','B','G','C','F','I','A','F','I','D','H','B','I','I','B','J','F','E','B','B','B','D','C','J','E','J','J','G','D','F','F','F','G','I','H','J','J','G','D','G','F'}
	//fmt.Println(len(in))
	fmt.Println(leastInterval([]byte{'F','J','J','A','J','F','C','H','J','B','E','G','G','F','A','C','I','F','J','C','J','C','H','C','A','D','G','H','B','F','G','C','C','A','E','B','H','J','E','I','F','D','E','A','C','D','B','D','J','J','C','F','D','D','J','H','A','E','C','D','J','D','G','G','B','C','E','G','H','I','D','H','F','E','I','B','D','E','I','E','C','J','G','I','D','E','D','J','C','A','C','C','D','I','J','B','D','H','H','J','G','B','G','A','H','E','H','E','D','E','J','E','J','C','F','C','J','G','B','C','I','I','H','F','A','D','G','F','C','C','F','G','C','J','B','B','I','C','J','J','E','G','H','C','I','G','J','I','G','G','J','G','G','E','G','B','I','J','B','H','D','H','G','F','C','H','C','D','A','G','B','H','H','B','J','C','A','F','J','G','F','E','B','F','E','B','B','A','E','F','E','H','I','I','C','G','J','D','H','E','F','G','G','D','E','B','F','J','J','J','D','H','E','B','D','J','I','F','C','I','E','H','F','E','G','D','E','C','F','E','D','E','A','I','E','A','D','H','G','C','I','E','G','A','H','I','G','G','A','G','F','H','J','D','F','A','G','H','B','J','A','H','B','H','C','G','F','A','C','C','B','I','G','G','B','C','J','J','I','E','G','D','I','J','I','C','G','A','J','G','F','J','F','C','F','G','J','I','E','B','G','F','A','D','A','I','A','E','H','F','D','D','C','B','J','I','J','H','I','C','D','A','G','F','I','B','E','D','C','J','G','I','H','E','C','E','I','I','B','B','H','J','C','F','I','D','B','F','H','F','A','C','A','A','B','D','C','A','G','B','G','F','E','G','A','A','A','C','J','H','H','G','C','C','B','C','E','B','E','F','I','E','E','D','I','H','G','F','A','H','B','J','B','G','H','C','C','B','G','C','B','A','E','G','A','J','G','D','C','I','G','F','G','G','A','J','E','I','D','E','A','F','A','H','C','E','D','D','D','H','I','F','F','A','F','A','A','C','J','D','J','H','I','F','A','C','B','C','A','C','C','H','A','J','I','B','A','I','F','J','C','I','B','C','E','E','E','J','G','F','E','I','A','A','E','B','J','H','H','H','A','H','J','E','F','E','F','G','J','D','I','D','I','F','B','J','D','A','A','D','F','G','B','J','H','F','A','D','H','C','B','A','J','H','I','F','H','E','G','B','A','F','F','A','C','D','G','I','I','J','H','H','C','J','G','B','A','D','B','F','J','D','I','A','F','F','F','F','A','E','B','C','G','H','E','B','B','A','G','D','C','C','E','A','C','F','G','A','I','F','B','H','J','G','C','B','H','D','A','H','B','H','H','C','A','F','I','C','F','A','C','J','I','H','H','F','B','B','D','E','C','J','F','C','E','A','J','E','C','A','E','B','A','J','F','J','J','J','H','H','C','I','E','G','G','H','J','J','H','H','H','J','H','A','G','I','C','E','C','D','G','G','F','H','D','G','H','A','E','I','D','A','H','G','E','A','B','F','I','C','A','F','B','A','I','F','G','I','F','D','A','B','J','B','D','F','G','J','J','A','A','C','H','G','F','B','I','I','J','A','H','D','F','E','F','J','B','F','C','G','E','A','G','H','E','H','H','F','I','G','C','C','G','J','B','H','F','H','D','I','B','D','I','F','H','I','D','F','G','G','E','A','C','A','G','H','G','H','J','F','D','F','G','D','D','C','J','C','J','G','G','G','G','H','H','G','D','E','H','G','C','B','F','I','F','C','H','J','I','A','F','D','C','F','C','E','E','D','D','C','G','B','F','E','J','C','I','E','D','B','B','I','I','I','H','C','E','C','J','F','G','A','I','J','D','I','C','G','F','I','E','I','E','F','A','G','E','J','A','I','A','D','A','G','J','F','E','D','I','A','E','J','I','C','J','B','F','B','E','C','E','F','G','E','J','J','I','E','D','F','C','H','H','B','G','D','I','I','F','B','G','C','F','J','B','G','J','H','D','G','C','C','I','I','E','I','B','H','B','I','G','F','H','G','C','J','D','C','E','G','F','C','H','D','A','C','D','H','B','C','H','I','B','A','J','C','B','D','J','D','H','F','B','A','G','G','J','I','E','F','A','D','H','D','B','C','A','H','F','G','B','F','H','B','H','I','J','D','H','I','B','C','D','G','A','E','A','A','I','F','I','F','B','B','I','F','A','E','I','A','B','G','C','F','I','A','F','I','D','H','B','I','I','B','J','F','E','B','B','B','D','C','J','E','J','J','G','D','F','F','F','G','I','H','J','J','G','D','G','F'}, 8) == 1000)
}
