package main

import "log"

//Matrix ..
type Matrix struct {
	data []*Square
	m    uint64
	n    uint64
}

//NewMatrix
//i'm sorry being boring but it for your best
//sometimes is ok to use lazy name, but you must to say what is your intention
//c'mon man, don't throw out everything you learned in Modular
func NewMatrix(n, m uint64) *Matrix {
	return &Matrix{
		data: make([]*Square, n*m),
		m:    m,//Coluns
		n:    n,//Row
	}
}

const (
	TOP    = 0
	RIGHT  = 1
	BOTTOM = 2
	LEFT   = 3
)

//note: the matrix is a Slice , then:
//get(i,0)-> i*mat.m + 0 -> walks i times the m coluns-in the slice-reaching the (i,j) element
//get(i,j) -> i*mat.m + j -> walks i times the m coluns-in the slice- and walks j elements, to reach(i,j)element 
func (mat *Matrix) get(i, j uint64) *Square {
	if i*mat.m+j < uint64(len(mat.data)) {
		return mat.data[i*mat.m+j]
	} else {
		log.Fatalln("Oops:", i*mat.m+j, i, j)
	}
	return nil
}

//i-> row
//j-> column
//v -> data
func (mat *Matrix) set(i uint64, j uint64, v *Square) {
	mat.data[i*mat.m+j] = v
}




func getCost(env *Environment, c string, x, y uint64) int64 {
	if c == "S" {//what is exactly S? it's the end? or start ? it's not defined anywhere in Json 
		return 0
	}
	for _, g := range env.Grounds {//given a ground c, searches in env the ground returning the cost of c 
		//if t.Position.X == uint64(x) && t.Position.Y == uint64(y) {
		if g.ID == c {
			return g.Cost
		}
	}
	// Temple?
	if c != "_" {
		log.Fatalln("oops: Invalid Square at", x, y, "!!!")
	}
	for _, t := range env.Temples {//return cost of the current temple
		if t.Position.X == uint64(x) && t.Position.Y == uint64(y) {
			return t.Cost
		}
	}
	return 0
}
//why get or build? why don't you use the get already defined!?
func getOrBuild(env *Environment, ref *Matrix, g string, x, y uint64) *Square {
	if v := ref.get(x, y); v != nil {
		return v
	}
	s := new(Square)
	s.Cost = getCost(env, g, x, y)
	s.p = Point{X: x, Y: y}
	s.neighbors = make([]*Square, 4)
	ref.set(x, y, s)
	return s
}

// buildGraphFromEnv build a Graph and return the initial and the final Square
func buildGraphFromEnv(env *Environment) (*Square, *Square) {
	ref := NewMatrix(42, 42)
	for x, l := range env.Map {
		for y, c := range l {
			var x, y = uint64(x), uint64(y)
//how can you use the variable if you didn't initialize in any instance?
//i'm sorry asking, but are you testing it?
			v := getOrBuild(env, ref, c, x, y)
			s := v //why this? why not just using s or v? even because you don't use v in the code anymore
			if x > 0 { // has top neighbor
				s.neighbors[TOP] = getOrBuild(env, ref, c, x-1, y)// this shouldn't be... x, y-1) instead ...x-1,y)? after all, you will the neighbor on top, this mean y -1 by convention
			}
			if y > 0 { // has  left neighbor
				s.neighbors[LEFT] = getOrBuild(env, ref, c, x, y-1)
			}
			if x < uint64(len(env.Map)-1) { // has bottom neighbor
				s.neighbors[BOTTOM] = getOrBuild(env, ref, c, x+1, y)
			}
			if y < uint64(len(l)-1) { // has right neighbor
				s.neighbors[RIGHT] = getOrBuild(env, ref, c, x, y+1)
			}
		}
	}
	return ref.get(env.Start.X, env.Start.Y), ref.get(env.End.X, env.End.Y)
}
