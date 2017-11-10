package scheduler

type Event struct {
	name string
	capacity uint
}

type Candidate struct {
	name string
	capacity uint
	potentialEvents []string
}

type VertexNode struct {
	element *Vertex
	next *VertexNode
}

type DirectedEdgeNode struct {
	element *DirectedEdge
	next *DirectedEdgeNode
}

type VertexQueue struct {
	front *VertexNode
	end *VertexNode
}

type Vertex struct {
	name string
	edges []*DirectedEdge
}

type DirectedEdge struct {
	src *Vertex
	dst *Vertex
	cost uint
}	

type Graph struct {
	vertices []*Vertex
}

func (q VertexQueue) isEmpty() bool {
	return q.front == nil
}

func (q VertexQueue) pop() *Vertex {
	if (q.front == q.end) {
		q.end = nil
	}
	temp := q.front
	q.front = q.front.next
	return temp.element
}

func (q VertexQueue) add(v *Vertex) {
	val := &(VertexNode{element: v, next: nil})
	if (q.end == nil) {
		q.front = val
		q.end = val
	} else {
		q.end.next = val
		q.end = val
	}
}