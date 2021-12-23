package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// 	input := `dc-end
	// HN-start
	// start-kj
	// dc-start
	// dc-HN
	// LN-dc
	// HN-end
	// kj-sa
	// kj-HN
	// kj-dc`
	if err := part1(input); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	if err := part2(input); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func part1(input string) error {
	g := parseInput(input)

	g.paths("start")
	fmt.Println(g.pathCount)

	return nil
}

func part2(input string) error {
	g := parseInput(input)

	g.allowTwice = true
	g.paths("start")
	fmt.Println(g.pathCount)

	return nil
}

type graph struct {
	edges      map[string][]string
	pathCount  int
	allowTwice bool
}

func (g *graph) paths(path string) {
	current := path[strings.LastIndexByte(path, '-')+1:]

	edges := g.edges[current]
	for _, next := range edges {
		if next == "start" {
			continue
		}
		if next == "end" {
			// path is complete
			// fmt.Println(path + "-end")
			g.pathCount++
			continue
		}
		if next[0] >= 'a' && next[0] <= 'z' {
			if g.allowTwice && !strings.Contains(path, "-2-") {
				// haven't used up two cave visits
				if idx := strings.Index(path, next); idx != -1 {
					// We've seen this once before
					if strings.Contains(path[idx+len(next):], next) {
						// already in twice
						continue
					}
					// Flag that we've used up our double visit
					next = "2-" + next
				}
			} else {
				// double visit used up, so no more repeat visits to small
				// caves (or not allowing double visits)
				if strings.Contains(path, next) {
					continue
				}
			}
		}
		// Need to copy complete path to avoid overwritting backing array
		g.paths(path + "-" + next)
	}
}

func parseInput(in string) graph {
	g := graph{
		edges: make(map[string][]string),
	}
	for _, line := range strings.Split(in, "\n") {
		parts := strings.Split(line, "-")
		a, b := parts[0], parts[1]
		g.edges[a] = append(g.edges[a], b)
		g.edges[b] = append(g.edges[b], a)
	}
	return g
}

var input = `start-kc
pd-NV
start-zw
UI-pd
HK-end
UI-kc
pd-ih
ih-end
start-UI
kc-zw
end-ks
MF-mq
HK-zw
LF-ks
HK-kc
ih-HK
kc-pd
ks-pd
MF-pd
UI-zw
ih-NV
ks-HK
MF-kc
zw-NV
NV-ks`
