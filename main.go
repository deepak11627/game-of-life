package main

import (
	"bytes"
	"fmt"
	"reflect"
	"time"
)

type State struct {
	x int
	y int
}

func (s State) Display() {
	fmt.Printf(" %d,%d ", s.x, s.y)
}
func (s State) GetNeighbours(maxSize int) []State {
	var nb []State
	if s.x-1 >= 0 {
		nb = append(nb, State{s.x - 1, s.y})
	}
	if s.x-1 >= 0 && s.y-1 >= 0 {
		nb = append(nb, State{s.x - 1, s.y - 1})
	}
	if s.x-1 >= 0 && s.y+1 < maxSize {
		nb = append(nb, State{s.x - 1, s.y + 1})
	}
	if s.y-1 >= 0 {
		nb = append(nb, State{s.x, s.y - 1})
	}
	if s.y-1 >= 0 && s.x+1 < maxSize {
		nb = append(nb, State{s.x + 1, s.y - 1})
	}
	if s.y+1 < maxSize {
		nb = append(nb, State{s.x, s.y + 1})
	}
	if s.x+1 < maxSize {
		nb = append(nb, State{s.x + 1, s.y})
	}
	if s.y+1 < maxSize && s.x+1 < maxSize {
		nb = append(nb, State{s.x + 1, s.y + 1})
	}
	return nb
}

type GameOfLife struct {
	maxSize       int
	maxIterations int
	grid          map[State]struct{}
}

func NewGameOfLife(size, iterations int, initialState []State) *GameOfLife {
	grid := make(map[State]struct{}, len(initialState))
	for _, elem := range initialState {
		grid[elem] = struct{}{}
	}

	return &GameOfLife{
		maxSize:       size,
		maxIterations: iterations,
		grid:          grid,
	}
}

func (g *GameOfLife) ApplyRules() {
	counter := make(map[State]int, len(g.grid))
	for elem := range g.grid {
		//elem.Display()
		if _, ok := g.grid[elem]; !ok {
			counter[elem] = 0
		}
		nb := elem.GetNeighbours(g.maxSize)
		//fmt.Printf("\nneighbour of %d, %d \n %+v", elem.x, elem.y, nb)
		for _, n := range nb {
			//	n.Display()
			if _, ok := counter[n]; !ok {
				counter[n] = 1
			} else {
				counter[n]++
			}
		}
	}
	//	fmt.Printf("\n %+v", counter)

	for c := range counter {
		if counter[c] < 2 || counter[c] > 3 {
			//	fmt.Printf("\n remove %d, %d", c.x, c.y)
			delete(g.grid, c)
		}
		if counter[c] == 3 {
			//	fmt.Printf("\n adding %d, %d", c.x, c.y)
			g.grid[c] = struct{}{}
		}
	}
}

func (g *GameOfLife) Start() {
	previousState := make(map[State]struct{}, len(g.grid))
	i := 0
	for !reflect.DeepEqual(previousState, g.grid) && i < g.maxIterations {
		i++
		for k, v := range g.grid {
			previousState[k] = v
		}
		g.ApplyRules()
		fmt.Print("\x0c", g) // Clear screen and print field.
		time.Sleep(time.Second / 30)
	}
}

func (g *GameOfLife) String() string {
	var buf bytes.Buffer
	for i := 0; i <= g.maxSize; i++ {
		for j := 0; j <= g.maxSize; j++ {
			state := State{x: j, y: i}
			b := byte(' ')
			if _, ok := g.grid[state]; ok {
				b = '*'
			}
			buf.WriteByte(b)
		}
		buf.WriteByte('\n')
	}
	return buf.String()
}

func main() {
	initialState := []State{
		// {19, 20},
		// {20, 20},
		// {21, 20},
		{39, 40},
		{39, 41},
		{40, 39},
		{40, 40},
		{41, 40},
	}
	g := NewGameOfLife(80, 1500, initialState)
	g.Start()
}
