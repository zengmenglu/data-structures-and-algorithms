package dijkstra

import (
	"testing"
)

func TestDijkstra(t *testing.T) {
	edge := [][]int{
		{0, 1, inf, inf, inf, 6},
		{1, 0, 2, inf, inf, inf, 7},
		{inf, 2, 0, 3, 8, inf},
		{inf, inf, 3, 0, 4, inf},
		{inf, inf, 8, 4, 0, 5},
		{6, 7, inf, inf, 5, 0},
	}
	t.Logf("%t", Dijkstra(6, edge, 0, 0) == 0)
	t.Logf("%t", Dijkstra(6, edge, 0, 1) == 1)
	t.Logf("%t", Dijkstra(6, edge, 0, 2) == 3)
	t.Logf("%t", Dijkstra(6, edge, 0, 3) == 6)
	t.Logf("%t", Dijkstra(6, edge, 0, 4) == 10)
	t.Logf("%t", Dijkstra(6, edge, 0, 5) == 6)
}
