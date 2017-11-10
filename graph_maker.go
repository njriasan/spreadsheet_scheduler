package main

func graphConstructor(events []*Event, candidates []*Candidate) (g *Graph, source *Vertex, destination *Vertex) {
		destination = &(Vertex{name: "", edges: make([]*DirectedEdge, 0)})
		source = &(Vertex{name: "", edges: make([]*DirectedEdge, 0)})
		vertices := make([]*Vertex, 0)
		eventMap := make(map[string]*Vertex)
		vertices = append(vertices, source, destination)
		for _, element := range events {
			eVertex := &(Vertex{name: element.name, edges: make([]*DirectedEdge, 0)})
			forwardEdge := &(DirectedEdge{src: source, dst: eVertex, cost: element.capacity})
			reverseEdge := &(DirectedEdge{src: eVertex, dst: source, cost: 0})
			eVertex.edges = append(eVertex.edges, reverseEdge)
			source.edges = append(source.edges, forwardEdge)
			eventMap[element.name] = eVertex
			vertices = append(vertices, eVertex)
		}
		for _, element := range candidates {
			cVertex := &(Vertex{name: element.name, edges: make([]*DirectedEdge, 0)})
			forwardEdge := &(DirectedEdge{src: cVertex, dst: destination, cost: element.capacity})
			reverseEdge := &(DirectedEdge{src: destination, dst: cVertex, cost: 0})
			cVertex.edges = append(cVertex.edges, forwardEdge)
			destination.edges = append(destination.edges, reverseEdge)
			vertices = append(vertices, cVertex)
			for _, elem := range element.potentialEvents {
				dstEvent := eventMap[elem]
				fEdge := &(DirectedEdge{src: dstEvent, dst: cVertex, cost: 1})
				rEdge := &(DirectedEdge{src: cVertex, dst: dstEvent, cost: 0})
				cVertex.edges = append(cVertex.edges, rEdge)
				dstEvent.edges = append(dstEvent.edges, fEdge)
			}
		}
		g = &(Graph{vertices: vertices})
		return
}