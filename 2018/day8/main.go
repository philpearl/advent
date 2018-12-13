package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

func run() error {
	input, err := read("input.txt")
	if err != nil {
		return err
	}
	d := datastream{data: input}
	d.nodes = make([]node, d.countNodes())
	d.data = input

	tree := d.readNode()
	fmt.Println(tree.metaValue())
	fmt.Println(tree.value())

	return nil
}

func sum(vv []int8) (total int) {
	for _, v := range vv {
		total += int(v)
	}
	return total
}

type node struct {
	meta     []int8
	children []node
}

func (n *node) metaValue() (total int) {
	for i := range n.children {
		total += n.children[i].metaValue()
	}
	return total + sum(n.meta)
}

func (n *node) value() (total int) {
	if len(n.children) == 0 {
		return sum(n.meta)
	}
	for _, i := range n.meta {
		if i == 0 || int(i) > len(n.children) {
			continue
		}
		total += n.children[i-1].value()
	}
	return total
}

type datastream struct {
	data  []int8
	nodes []node
}

func (d *datastream) countNodes() (count int) {
	childrenCount := d.next()
	metaCount := d.next()
	for i := 0; i < int(childrenCount); i++ {
		count += d.countNodes()
	}
	d.data = d.data[metaCount:]
	return count + 1
}

func (d *datastream) readNode() node {
	childrenCount := d.next()
	metaCount := d.next()
	children := d.nodes[:childrenCount]
	d.nodes = d.nodes[childrenCount:]
	for i := range children {
		children[i] = d.readNode()
	}
	meta := d.data[:metaCount]
	d.data = d.data[metaCount:]
	return node{
		meta:     meta,
		children: children,
	}
}

func (d *datastream) next() int8 {
	if len(d.data) == 0 {
		return -1
	}

	v := d.data[0]
	d.data = d.data[1:]
	return v
}

func read(filename string) ([]int8, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	values := make([]int8, 0, len(content))

	var current int8
	for _, c := range content {
		if c >= '0' && c <= '9' {
			current = current*10 + int8(c-'0')
		} else {
			values = append(values, current)
			current = 0
		}
	}
	return append(values, current), nil
}
