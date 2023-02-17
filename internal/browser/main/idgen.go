package browsermain

type idgen struct {
	free []int
	next int
}

func (g *idgen) alloc() int {
	if len(g.free) == 0 {
		ret := g.next
		g.next++
		return ret
	}
	ret := g.free[len(g.free)-1]
	g.free = g.free[:len(g.free)-1]
	return ret
}

func (g *idgen) release(n int) {
	g.free = append(g.free, n)
}
