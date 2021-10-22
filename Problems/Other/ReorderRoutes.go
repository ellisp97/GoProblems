package Problems

/*
There are n cities numbered from 0 to n - 1 and n - 1 roads such that there is only one way to travel between two different cities (this network form a tree).

Roads are represented by connections where connections[i] = [ai, bi] represents a road from city ai to city bi.
This year, there will be a big event in the capital (city 0), and many people want to travel to this city.

TASK:  reorienting some roads such that each city can visit the city 0. Return the minimum number of edges changed.



Input: n = 6, connections = [[0,1],[1,3],[2,3],[4,0],[4,5]]
Output: 3
Explanation: Change the direction of edges [4,5] -> [5,4], [0,1] -> [1,0], [1, 3] -> [3,1]
*/

type Edge struct {
	cityA int
	cityB int
}

func MinReorder(n int, connections [][]int) int {
	// Start at city 0
	// Recursively check its neighbours
	// Count ongoing egdes
	edges := map[Edge]bool{}

	for _, v := range connections {
		edges[Edge{cityA: v[0], cityB: v[1]}] = true
	}

	neighbours := make(map[int][]int, n)
	visited := map[int]bool{}
	changes := 0

	for _, v := range connections {
		neighbours[v[0]] = append(neighbours[v[0]], v[1])
		neighbours[v[1]] = append(neighbours[v[1]], v[0])
	}

	var dfs func(this int)
	dfs = func(city int) {
		for _, neighbour := range neighbours[city] {
			if visited[neighbour] {
				continue
			}

			edge := Edge{cityA: neighbour, cityB: city}
			// check if the neighbour can reach city (0)
			if !edges[edge] { // if it cant edge is going backwards increment changes
				changes++
			}
			visited[neighbour] = true
			dfs(neighbour)
		}
	}
	visited[0] = true
	dfs(0)
	return changes

}
