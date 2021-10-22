package main

import "fmt"

type Stack struct {
	val []int
}

func (s *Stack) pop() int {
	ele := s.val[len(s.val)-1]
	s.val = s.val[:len(s.val)-1]
	return ele
}

func (s *Stack) push(n int) {
	s.val = append(s.val, n)
}

func DFS(friends [][]bool, personsCount int, visited map[int]bool, v int) {
	for i := 0; i < personsCount; i++ {
		if i != v && friends[i][v] && !visited[i] {
			visited[i] = true
			DFS(friends, personsCount, visited, i)
		}
	}
}

func FriendCircles(friends [][]bool, personsCount int) {
	visited := map[int]bool{}
	count := 0
	for i := 0; i < personsCount; i++ {
		if !visited[i] {
			visited[i] = true
			DFS(friends, personsCount, visited, i)
			count++
		}
	}
	fmt.Println(count)
}

// func main() {
// 	friends := [][]bool{
// 		{true, true, false, false},
// 		{true, true, true, false},
// 		{false, true, true, false},
// 		{false, false, false, true}}
// 	FriendCircles(friends, 4)
// }
