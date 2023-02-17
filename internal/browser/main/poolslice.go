package browsermain

// TODO(cleanup): it may make sense to move these to a separate package; this
// is more of a general data structure, it's only in this package because
// we happen to use it for keeping track of grain iframes.

// A PoolSlice wraps a slice, adding support for tracking free vs. in-use
// slots and inserting new elements into free slots.
type PoolSlice[T any] struct {
	Items  []T            // The slice
	idxGen indexGenerator // Keeps track of free slots
}

// Add adds T to the slice, returning its index. Will use a free slot if
// available, otherwise will expand the slice.
func (s *PoolSlice[T]) Add(value T) int {
	index := s.idxGen.alloc()
	if index == len(s.Items) {
		s.Items = append(s.Items, value)
	} else {
		s.Items[index] = value
	}
	return index
}

// Remove removes the item at the index, replacing it with the zero value and
// marking its slot for re-use.
func (s *PoolSlice[T]) Remove(index int) {
	var zero T
	s.Items[index] = zero
	s.idxGen.release(index)
}

// An indexGenerator allocates slice indexes, allowing re-use. This lets you insert
// remove items within a slice dynamically, re-using empty slots when available.
type indexGenerator struct {
	free []int // list of free slots
	next int   // next index to use when no slots are free
}

// alloc allocates a free index. Will use an empty slot if available.
func (g *indexGenerator) alloc() int {
	if len(g.free) == 0 {
		ret := g.next
		g.next++
		return ret
	}
	ret := g.free[len(g.free)-1]
	g.free = g.free[:len(g.free)-1]
	return ret
}

// release marks an index as free.
func (g *indexGenerator) release(index int) {
	g.free = append(g.free, index)
}
