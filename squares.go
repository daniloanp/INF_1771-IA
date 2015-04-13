package main

// Square ..
type Square struct {
	p         Point
	Cost      int64
	neighbors []*Square
}

//Neighbors .. .
func (v *Square) Neighbors() []*Square {
	return v.neighbors
}

//Equals Compare two Squares
func (v *Square) Equals(v1 *Square) bool {
	neighbor := v1
	return neighbor.p.X == v.p.X && neighbor.p.Y == v.p.Y
}

//DistanceToNeighbor ...please, make some (at least simple and direct comment about it!
//this function searchs for the interessed neighbor(finding) in the current Square, 
//and return the cost of the neighbor(finding) 
func (v *Square) DistanceToNeighbor(finding *Square) int64 {
	for _, neighbor := range v.Neighbors() {
		if neighbor == finding {
			return neighbor.Cost
		}
	}
	return -1
}

// BFS ...
func (v *Square) BFS(finding *Square) []*Square {
	// var cost = 0
	//what is SquareQueue???????????????????????????????????????????????????????function,type?
	var Q = SquareQueue(make([]*Square, 0, 42*42))
	// var path = SquareQueue(make([]*Square, 0, 42*42))
	var D = make(map[*Square]bool)
	Q.add(v)
	D[v] = true

	for len(Q) > 0 {
		v := Q.get()
		if v == nil {
			break
		}
		for _, neighbor := range v.Neighbors() {
			if neighbor == nil {
				continue
			}
			if !D[neighbor] {
				if neighbor.Equals(finding) {
					return [](*Square){v, neighbor}
				}
				Q.add(neighbor)
				D[neighbor] = true
			}
		}
	}

	Q = nil
	return nil
}

func getMin(o map[*Square]bool, f map[*Square]int64) *Square {
	var (
		best *Square
		min  int64
	)
	for j, flag := range o {
		if !flag {
			continue
		}
		if f[j] < min {
			best = j
		}
	}
	return best
}
