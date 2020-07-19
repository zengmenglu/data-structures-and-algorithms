package dijkstra

var inf = 0x7FF00000

// Dijkstra algorithm calculate the minimum distance between source and destination.
// n is the number of nodes.
// edge stores the distance between two nodes. If nodes has no direct connection, the distance is infinite initially.
// src and dst is the nodes' index of source and destination.
func Dijkstra(n int, edge [][]int, src, dst int) int {
	// init
	minDistance := make([]int, n) // record the min distance between src and node i
	noVisit := make(map[int]bool) // record nodes not been visited yet
	for i := 0; i < n; i++ {
		minDistance[i] = edge[src][i]
		noVisit[i] = true
	}
	delete(noVisit, src)

	for i := 0; i < n; i++ {
		// find
		distance := inf
		curNode := 0
		for j := range noVisit {
			if minDistance[j] < distance {
				curNode = j
				distance = minDistance[j]
			}
		}
		delete(noVisit, curNode)

		// update
		for k := range noVisit {
			if distance+edge[curNode][k] < minDistance[k] {
				minDistance[k] = distance + edge[curNode][k]
			}
		}
	}
	return minDistance[dst]
}
