package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Skill is a type for the priority queue
type Skill struct {
	name     string
	progress float64
	index    int
}

// PriorityQueue is a slice of Skill pointers
type PriorityQueue []*Skill

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].progress > pq[j].progress
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index, pq[j].index = i, j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	skill := x.(*Skill)
	skill.index = n
	*pq = append(*pq, skill)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	skill := old[n-1]
	old[n-1] = nil
	skill.index = -1
	*pq = old[0 : n-1]
	return skill
}

func main() {
	// Reading from input.txt
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	graph := make(map[string][]string)
	progress := make(map[string]float64)
	inDegree := make(map[string]int)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "->") {
			parts := strings.Split(line, "->")
			from, to := strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
			graph[from] = append(graph[from], to)
			inDegree[to]++
			if _, exists := inDegree[from]; !exists {
				inDegree[from] = 0
			}
		} else if strings.Contains(line, "=") {
			parts := strings.Split(line, "=")
			skill := strings.TrimSpace(parts[0])
			prog, _ := strconv.ParseFloat(strings.TrimSpace(parts[1]), 64)
			progress[skill] = prog
			if _, exists := inDegree[skill]; !exists {
				inDegree[skill] = 0
			}
		}
	}

	// Initialize priority queue
	pq := &PriorityQueue{}
	heap.Init(pq)

	// Add nodes with in-degree 0 to the priority queue
	for node, deg := range inDegree {
		if deg == 0 {
			heap.Push(pq, &Skill{name: node, progress: progress[node]})
		}
	}

	// Process nodes and output topological order
	var topologicalOrder []string

	for pq.Len() > 0 {
		skill := heap.Pop(pq).(*Skill)
		topologicalOrder = append(topologicalOrder, skill.name)

		for _, neighbor := range graph[skill.name] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				heap.Push(pq, &Skill{name: neighbor, progress: progress[neighbor]})
			}
		}
	}

	fmt.Println("Topological Order:")
	for _, skill := range topologicalOrder {
		fmt.Println(skill)
	}
}
