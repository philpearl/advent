package main

import (
	"container/list"
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

var passcode = "veumntbg"

func hashPath(path string) [4]bool {
	sum := md5.Sum([]byte(passcode + path))
	hash := hex.EncodeToString(sum[:])
	var open [4]bool
	for i := 0; i < 4; i++ {
		if hash[i] > 'a' {
			open[i] = true
		}
	}
	return open
}

type location struct{ x, y int }

type state struct {
	location
	path string
}

func (p *problem) addPossibleSteps(s state) {
	// Up, down, left, right
	open := hashPath(s.path)
	if open[0] {
		p.addStep(state{location: location{x: s.x, y: s.y - 1}, path: s.path + "U"})
	}
	if open[1] {
		p.addStep(state{location: location{x: s.x, y: s.y + 1}, path: s.path + "D"})
	}
	if open[2] {
		p.addStep(state{location: location{x: s.x - 1, y: s.y}, path: s.path + "L"})
	}
	if open[3] {
		p.addStep(state{location: location{x: s.x + 1, y: s.y}, path: s.path + "R"})
	}
}

func (p *problem) addStep(s state) {
	if s.x < 0 || s.y < 0 || s.x > 3 || s.y > 3 {
		return
	}

	// if _, seen := p.pathsSeen[s.path]; seen {
	// 	return
	// }

	// p.pathsSeen[s.path] = struct{}{}

	p.possiblePaths.PushBack(s)
}

func (p *problem) nextStep() (state, bool) {
	e := p.possiblePaths.Front()
	if e == nil {
		return state{}, false
	}
	p.possiblePaths.Remove(e)
	return e.Value.(state), true
}

type problem struct {
	possiblePaths *list.List
	pathsSeen     map[string]struct{}
}

func atDestination(s state) bool {
	return s.x == 3 && s.y == 3
}

func main() {

	longestPath := ""

	p := problem{
		possiblePaths: list.New(),
		pathsSeen:     map[string]struct{}{},
	}

	p.addStep(state{location: location{0, 0}, path: ""})

	for {
		s, ok := p.nextStep()
		if !ok {
			fmt.Printf("nothing left\n")
			break
		}

		if atDestination(s) {
			if len(s.path) > len(longestPath) {
				longestPath = s.path
			}
			continue
		}

		p.addPossibleSteps(s)
	}

	fmt.Printf("Longest path is %d\n", len(longestPath))
}
