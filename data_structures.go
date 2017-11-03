package scheduler

type Event struct {
	name string
	capacity uint
}

type Candidate struct {
	name string
	capacity uint
	potentialEvents []*Event
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
	edges []*DirectedEdge
}

func (q VertexQueue) isEmpty() {
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
	v := VertexNode{element = v, next = nil}
	if (q.end == nil) {
		q.front = &v
		q.end = &v
	} else {
		q.end.next = &v
		q.end = &v
	}
}