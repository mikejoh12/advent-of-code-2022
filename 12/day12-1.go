package main

import (
	"container/list"
	"fmt"
	"image"
	"os"
	"strings"
	"unicode"
)

type Graph struct {
	startPos *Node
	nodes    map[image.Point]*Node
}

type Node struct {
	id      image.Point
	visited bool
	char    rune
	edges   map[image.Point]*Node
}

func NewGraph(s string) *Graph {
	rows := strings.Split(s, "\n")
	g := Graph{nodes: make(map[image.Point]*Node)}

	for y, row := range rows {
		for x, r := range row {
			n := Node{
				id:    image.Pt(x, y),
				char:  r,
				edges: make(map[image.Point]*Node),
			}
			g.nodes[image.Pt(x, y)] = &n
			if r == 'S' {
				g.startPos = &n
			}
		}
	}

	var directions = [4]image.Point{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}

	for y, row := range rows {
		for x, r := range row {
			for _, dir := range directions {
				nPos := image.Pt(x, y).Add(dir)
				if nNode, ok := g.nodes[nPos]; ok {
					switch {
					case r == 'S':
						g.nodes[image.Pt(x, y)].edges[image.Pt(nPos.X, nPos.Y)] = nNode
					case (r == 'y' || r == 'z') && nNode.char == 'E':
						g.nodes[image.Pt(x, y)].edges[image.Pt(nPos.X, nPos.Y)] = nNode
					case nNode.char != 'S' && unicode.IsLower(r) && unicode.IsLower(nNode.char) && (nNode.char <= r+1):
						g.nodes[image.Pt(x, y)].edges[image.Pt(nPos.X, nPos.Y)] = nNode
					}
				}
			}
		}
	}

	return &g
}

func (g *Graph) AddNode(id image.Point, c rune) {
	g.nodes[id] = &Node{
		id:    id,
		char:  c,
		edges: make(map[image.Point]*Node),
	}
}

func (g *Graph) bfs(start *Node, endChar rune) int {
	start.visited = true
	var levels = map[*Node]int{start: 0}
	levels[start] = 0
	queue := list.New()
	queue.PushBack(g.startPos)

	for queue.Len() > 0 {
		current := queue.Front()

		if current.Value.(*Node).char == endChar {
			return levels[current.Value.(*Node)]
		}
		for _, child := range current.Value.(*Node).edges {
			if !child.visited {
				queue.PushBack(child)
				levels[child] = levels[current.Value.(*Node)] + 1
				child.visited = true
			}
		}
		queue.Remove(current)
	}
	return -1
}

func main() {
	s, _ := os.ReadFile("input.txt")
	g := NewGraph(string(s))
	fmt.Println("Part 1:", g.bfs(g.startPos, 'E'))
}
