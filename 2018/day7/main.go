package main

import (
	"fmt"
	"os"
)

func main() {
	part1(input, 26)
	part2(test, 6, 2, 0)
	part2(input, 26, 5, 60)
}

func part2(input []string, stepCount, workerCount int, delay int) {
	steps := setupSteps(input, stepCount)

	type worker struct {
		step int
		work int
	}

	workers := make([]worker, workerCount)

	var sec int
	for {
		fmt.Printf("second %d\n", sec)
		for i := range workers {
			w := &workers[i]
			if w.work == 0 {
				i, found := steps.next()
				if found {
					w.work = delay + i + 1
					w.step = i
				}
			}
		}

		var progress bool
		for i := range workers {
			w := &workers[i]
			if w.work > 0 {
				w.work--
				fmt.Printf("worker %d %c\n", i, w.step+'A')
				if w.work == 0 {
					steps.complete(w.step)
				}
				progress = true
			}
		}
		if !progress {
			break
		}
		sec++
	}
	fmt.Printf("Done in %d seconds\n", sec)
}

func part1(input []string, stepCount int) {
	steps := setupSteps(input, stepCount)

	for {
		// find first step with no parents

		i, found := steps.next()
		if !found {
			break
		}
		os.Stdout.Write([]byte{byte(i + 'A')})
		steps.complete(i)
	}
	fmt.Println()
}

type steps []step

func (s steps) next() (int, bool) {
	for i := range s {
		step := &s[i]
		if step.parents == 0 && !step.complete && !step.inProgress {

			step.inProgress = true
			return i, true
		}
	}

	return 0, false
}

func (s steps) complete(i int) {
	step := &(s)[i]
	step.complete = true

	// Clear parent bit for any children of this step
	for child := range s {
		if step.children&(1<<uint(child)) != 0 {
			s[child].parents &= ^(1 << uint(i))
		}
	}
}

type step struct {
	children   int32
	parents    int32
	complete   bool
	inProgress bool
}

func setupSteps(input []string, count int) steps {
	steps := make(steps, count)
	for _, line := range input {
		parent, child := line[5], line[36]

		steps[parent-'A'].children |= 1 << (child - 'A')
		steps[child-'A'].parents |= 1 << (parent - 'A')
	}

	return steps
}

var test = []string{
	"Step C must be finished before step A can begin.",
	"Step C must be finished before step F can begin.",
	"Step A must be finished before step B can begin.",
	"Step A must be finished before step D can begin.",
	"Step B must be finished before step E can begin.",
	"Step D must be finished before step E can begin.",
	"Step F must be finished before step E can begin.",
}

var input = []string{
	"Step T must be finished before step X can begin.",
	"Step G must be finished before step O can begin.",
	"Step X must be finished before step B can begin.",
	"Step I must be finished before step W can begin.",
	"Step N must be finished before step V can begin.",
	"Step K must be finished before step H can begin.",
	"Step S must be finished before step R can begin.",
	"Step P must be finished before step J can begin.",
	"Step L must be finished before step V can begin.",
	"Step D must be finished before step E can begin.",
	"Step J must be finished before step R can begin.",
	"Step U must be finished before step W can begin.",
	"Step M must be finished before step Q can begin.",
	"Step B must be finished before step F can begin.",
	"Step F must be finished before step E can begin.",
	"Step V must be finished before step Q can begin.",
	"Step C must be finished before step A can begin.",
	"Step H must be finished before step Z can begin.",
	"Step A must be finished before step Y can begin.",
	"Step O must be finished before step Y can begin.",
	"Step W must be finished before step Q can begin.",
	"Step E must be finished before step Y can begin.",
	"Step Y must be finished before step Z can begin.",
	"Step Q must be finished before step R can begin.",
	"Step R must be finished before step Z can begin.",
	"Step S must be finished before step E can begin.",
	"Step O must be finished before step W can begin.",
	"Step G must be finished before step B can begin.",
	"Step I must be finished before step N can begin.",
	"Step G must be finished before step I can begin.",
	"Step H must be finished before step R can begin.",
	"Step N must be finished before step C can begin.",
	"Step M must be finished before step W can begin.",
	"Step Y must be finished before step R can begin.",
	"Step T must be finished before step B can begin.",
	"Step G must be finished before step D can begin.",
	"Step J must be finished before step O can begin.",
	"Step I must be finished before step A can begin.",
	"Step J must be finished before step H can begin.",
	"Step T must be finished before step Y can begin.",
	"Step N must be finished before step H can begin.",
	"Step B must be finished before step V can begin.",
	"Step M must be finished before step R can begin.",
	"Step Y must be finished before step Q can begin.",
	"Step X must be finished before step J can begin.",
	"Step A must be finished before step E can begin.",
	"Step P must be finished before step Z can begin.",
	"Step P must be finished before step C can begin.",
	"Step N must be finished before step Q can begin.",
	"Step A must be finished before step O can begin.",
	"Step G must be finished before step X can begin.",
	"Step P must be finished before step U can begin.",
	"Step T must be finished before step S can begin.",
	"Step I must be finished before step V can begin.",
	"Step V must be finished before step H can begin.",
	"Step U must be finished before step F can begin.",
	"Step D must be finished before step Q can begin.",
	"Step D must be finished before step O can begin.",
	"Step G must be finished before step H can begin.",
	"Step I must be finished before step Z can begin.",
	"Step N must be finished before step D can begin.",
	"Step B must be finished before step Y can begin.",
	"Step J must be finished before step M can begin.",
	"Step V must be finished before step Y can begin.",
	"Step W must be finished before step Y can begin.",
	"Step E must be finished before step Z can begin.",
	"Step T must be finished before step N can begin.",
	"Step L must be finished before step U can begin.",
	"Step S must be finished before step A can begin.",
	"Step Q must be finished before step Z can begin.",
	"Step T must be finished before step F can begin.",
	"Step F must be finished before step Z can begin.",
	"Step J must be finished before step C can begin.",
	"Step X must be finished before step Y can begin.",
	"Step K must be finished before step V can begin.",
	"Step T must be finished before step I can begin.",
	"Step I must be finished before step O can begin.",
	"Step C must be finished before step W can begin.",
	"Step B must be finished before step Q can begin.",
	"Step W must be finished before step Z can begin.",
	"Step D must be finished before step H can begin.",
	"Step K must be finished before step A can begin.",
	"Step M must be finished before step E can begin.",
	"Step T must be finished before step U can begin.",
	"Step I must be finished before step J can begin.",
	"Step O must be finished before step Q can begin.",
	"Step M must be finished before step Z can begin.",
	"Step U must be finished before step C can begin.",
	"Step N must be finished before step F can begin.",
	"Step C must be finished before step H can begin.",
	"Step X must be finished before step E can begin.",
	"Step F must be finished before step O can begin.",
	"Step P must be finished before step O can begin.",
	"Step J must be finished before step A can begin.",
	"Step H must be finished before step Y can begin.",
	"Step A must be finished before step Q can begin.",
	"Step V must be finished before step Z can begin.",
	"Step S must be finished before step L can begin.",
	"Step H must be finished before step E can begin.",
	"Step X must be finished before step I can begin.",
	"Step O must be finished before step R can begin.",
}
