package february

import "sort"

// 2349. Design a Number Container System
// Medium

// Design a number container system that can do the following:

// Insert or Replace a number at the given index in the system.
// Return the smallest index for the given number in the system.
// Implement the NumberContainers class:

// NumberContainers() Initializes the number container system.
// void change(int index, int number) Fills the container at index with the number.
// If there is already a number at that index, replace it.
// int find(int number) Returns the smallest index for the given number,
// or -1 if there is no index that is filled by number in the system.

// Example 1:

// Input
// ["NumberContainers", "find", "change", "change", "change", "change", "find", "change", "find"]
// [[], [10], [2, 10], [1, 10], [3, 10], [5, 10], [10], [1, 20], [10]]
// Output
// [null, -1, null, null, null, null, 1, null, 2]

type group struct {
	s        []int
	isSorted bool
}

type NumberContainers struct {
	index map[int]int
	group map[int]group
}

func Constructor() NumberContainers {
	return NumberContainers{
		index: make(map[int]int),
		group: make(map[int]group),
	}
}

func (nc *NumberContainers) Change(index int, number int) {
	if prev, exists := nc.index[index]; exists {
		if prev != number {
			nc.remove(index, prev)
			nc.add(index, number)
		}
	} else {
		nc.add(index, number)
	}
}

func (nc *NumberContainers) add(index, number int) {
	nc.index[index] = number
	nc.group[number] = group{
		s:        append(nc.group[number].s, index),
		isSorted: false,
	}
}

func (nc *NumberContainers) remove(index, prev int) {
	g := nc.group[prev]
	for i, val := range g.s {
		if val == index {
			g.s = append(g.s[:i], g.s[i+1:]...)
			nc.group[prev] = g
			break
		}
	}
	delete(nc.index, index)
}

func (nc *NumberContainers) Find(number int) int {
	g, exists := nc.group[number]
	if !exists || len(g.s) == 0 {
		return -1
	}

	if !g.isSorted {
		sort.Ints(g.s)
		g.isSorted = true
	}

	nc.group[number] = g

	return g.s[0]
}
