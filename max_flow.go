package main

func (g *Graph) maxFlow(src *Vertex, dst *Vertex) {
	var path *DirectedEdgeNode;
	for path = g.bfsPath(src, dst); path != nil; path = g.bfsPath(src, dst) {
		for ; path != nil; path = path.next {
			path.element.cost -= 1
			match := path.element.src
			for _, element := range path.element.dst.edges {
				if element.dst == match {
					element.cost += 1
				}
			}
		}
	}
}

func (g *Graph) bfsPath(src *Vertex, dst *Vertex) *DirectedEdgeNode {
	var q *VertexQueue
	q = &VertexQueue{nil, nil}
	prev := make(map[*Vertex]*DirectedEdge)
	marked := make(map[*Vertex]bool)
	for _, element := range g.vertices {
		prev[element] = nil
		marked[element] = false
	}
	marked[src] = true
	var path *DirectedEdgeNode
	q.add(src)
	for ;!q.isEmpty(); {
		v := q.pop()
		for _, element := range v.edges {
			if marked[element.dst] == false && element.cost > 0 {
				q.add(element.dst)
				marked[element.dst] = true
				prev[element.dst] = element
			}
		}
	}
	if prev[dst] == nil {
		return nil
	}
	end := prev[dst]
	for ;end != nil; {
		path = &(DirectedEdgeNode{element: end, next: path})
		end = prev[end.src]
	}
	return path
}