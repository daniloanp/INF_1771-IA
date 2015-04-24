package main

import (
	"math"
)

// heuristicCostEstimate is, currently, just the distance between two pointers
func heuristicCostEstimate(origin *Square, goal *Square) int64 {
	var (
		dX       = float64(goal.Position.Row - origin.Position.Row)
		dY       = float64(goal.Position.Column - origin.Position.Column)
		distance = math.Sqrt(dX*dX + dY*dY)
	)

	if distance-math.Trunc(distance) > 0.5 {
		distance = distance + 1.0
	}
	return int64(math.Floor(distance))
}

// getMin is a auxiliary function that return
func getMin(o map[*Square]bool, f map[*Square]int64) *Square {
	var (
		best *Square
		min  int64 = 1<<63 - 1
	)
	for j, flag := range o {
		if !flag {
			continue
		}
		if f[j] <= min {
			best = j
		}
	}
	return best
}

// Return `reverse` Path.
func reconstructPath(cameFrom map[*Square]*Square, current *Square) ([]*Square, int64) {
	var path = make([]*Square, 1, 42*42)
	var duration = int64(0)
	path[0] = current
	duration = current.Cost

	for next, ok := cameFrom[current]; ok && next != nil; next, ok = cameFrom[next] {
		path = append(path, next)
		duration = duration + next.Cost
		current = next
	}

	// Fixing
	duration = duration - current.Cost

	return path, duration
}

//AStar ...
func (v *Square) AStar(goal *Square) ([]*Square, int64) {
	var (
		closedSet = make(map[*Square]bool)
		openSet   = map[*Square]bool{v: true}
		cameFrom  = make(map[*Square]*Square)
		gScore    = map[*Square]int64{v: 0}
		fScore    = map[*Square]int64{v: gScore[v] + heuristicCostEstimate(v, goal)}
	)
	for len(openSet) > 0 {
		var current = getMin(openSet, fScore)

		if current == goal {
			return reconstructPath(cameFrom, current)
		}

		delete(openSet, current)
		closedSet[current] = true

		for _, neighbor := range current.Neighbors() {
			if neighbor == nil || closedSet[neighbor] {
				continue
			}

			GScoreTry := gScore[current] + current.DistanceToNeighbor(neighbor)

			if !openSet[neighbor] || GScoreTry < gScore[neighbor] {
				cameFrom[neighbor] = current
				gScore[neighbor] = GScoreTry
				fScore[neighbor] = gScore[neighbor] + heuristicCostEstimate(neighbor, goal)
				//adding it to openSet
				openSet[neighbor] = true
			}
		}

	}
	return nil, 0
}
