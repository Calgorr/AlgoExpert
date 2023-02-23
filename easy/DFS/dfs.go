package main

var graph [][]int
var visited []int

func DFS(graph [][]int) {
	for k, _ := range graph {
		dfsVisit(graph, k)
	}
}
func dfsVisit(graph [][]int, k int) {
	for _, v := range graph[k] {
		if visited[v] == 0 {
			visited[v] = 1
			dfsVisit(graph, v)
		}
	}
}
