package scheduler

func (g *Graph) maxFlow(src *Vertex, dst *Vertex) {
	path DirectedEdgeNode;
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
	q VertexQueue
	prev := make(map[Vertex]*Edge)
	marked := make(map[Vertex]bool)
	for _, element := range g.vertices {
		prev[element] = nil
		marked[element] = false
	}
	path DirectedEdgeNode
	q.add(src)
	for ;!q.isEmpty(); {
		v := q.pop()
		for _, element := range q.edges {
			if marked[*(element.dst)] == false && element.cost > 0 {
				q.add(element.dst)
				marked[*(element.dst)] = true
				prev[*(element.dst)] = element
			}
		}
	}
	if prev[*dst] == nil {
		return nil
	}
	end := dst
	for ;end != nil; {
		path = DirectedEdgeNode{element = prev[*(q.element)], next = &path}
	}
	return &path
}