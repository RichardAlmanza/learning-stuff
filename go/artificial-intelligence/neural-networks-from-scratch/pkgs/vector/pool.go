package vector

import "slices"

const (
	Free PoolNodeState = iota
	InUse
	Node
)

var _ Pool[float64] = &VectorPool[float64]{}

type PoolNodeState int

type Pool[T Number] interface {
	Make(size int) []T
	Clean()
}

type TreeNode[T Number] struct {
	Parent   *VectorPoolNode[T]
	Children []*VectorPoolNode[T]
}

type VectorPoolNode[T Number] struct {
	Pool       *VectorPool[T]
	MemoryNode *TreeNode[T]
	Before     *VectorPoolNode[T]
	After      *VectorPoolNode[T]
	Data       []T
	Start      int
	End        int
	State      PoolNodeState
}

func (pn *VectorPoolNode[T]) Allocate(size int) *VectorPoolNode[T] {
	defer pn.Pool.SortFreeNodes()

	if size > len(pn.Data) {
		panic("Memory Overflow!")
	}

	if pn.State != Free {
		panic("What are you doing! you're stealing!")
	}

	pn.State = Node

	newUsed := &VectorPoolNode[T]{
		Pool: pn.Pool,
		MemoryNode: &TreeNode[T]{
			Parent: pn,
		},
		Before: pn.Before,
		After:  pn.After,
		Data:   pn.Data[:size],
		Start:  0,
		End:    size,
		State:  InUse,
	}

	for i := 0; i < len(newUsed.Data); i++ {
		newUsed.Data[i] = 0
	}

	if len(pn.Data) > size {
		newFree := &VectorPoolNode[T]{
			Pool: pn.Pool,
			MemoryNode: &TreeNode[T]{
				Parent: pn,
			},
			Before: newUsed,
			After:  pn.After,
			Data:   pn.Data[size:],
			Start:  size,
			End:    pn.End,
			State:  Free,
		}
		newUsed.After = newFree

		pn.MemoryNode.Children = append(pn.MemoryNode.Children, newFree)
		pn.Pool.FreeNodes = append(pn.Pool.FreeNodes, newFree)
	}

	pn.MemoryNode.Children = append(pn.MemoryNode.Children, newUsed)
	return newUsed
}

type VectorPool[T Number] struct {
	Pool      []*VectorPoolNode[T]
	FreeNodes []*VectorPoolNode[T]
}

func (p *VectorPool[T]) findFree(size int) (int, *VectorPoolNode[T]) {
	if len(p.FreeNodes) == 0 {
		return 0, nil
	}

	current := len(p.FreeNodes) / 2
	left := 0
	exist := false
	right := len(p.FreeNodes) - 1

	for {
		currentNode := p.FreeNodes[current]
		length := len(currentNode.Data)
		switch {
		case length == size:
			return current, currentNode
		case length < size:
			left = current + 1
		default:
			right = current - 1
			exist = true
		}
		current = (left + right) / 2

		if left >= right {
			if !exist {
				return 0, nil
			}
			break
		}
	}

	for i := left; i < len(p.FreeNodes); i++ {
		if len(p.FreeNodes[i].Data) >= size {
			return i, p.FreeNodes[i]
		}
	}

	return 0, nil
}

func (p *VectorPool[T]) newNode(size int) *VectorPoolNode[T] {
	node := &VectorPoolNode[T]{
		Pool:       p,
		MemoryNode: &TreeNode[T]{},
		Data:       make([]T, size),
		Start:      0,
		End:        size,
		State:      Free,
	}

	p.Pool = append(p.Pool, node)
	p.FreeNodes = append(p.FreeNodes, node)
	return node
}

func (p *VectorPool[T]) SortFreeNodes() {
	for i := 0; i < len(p.FreeNodes); i++ {
		if p.FreeNodes[i].State != Free {
			p.FreeNodes[i] = nil
		}
	}

	slices.SortFunc(p.FreeNodes, func(a, b *VectorPoolNode[T]) int {
		switch {
		case a == nil:
			return 1
		case b == nil:
			return -1
		case len(a.Data) == len(b.Data):
			return 0
		case len(a.Data) < len(b.Data):
			return -1
		default:
			return 1

		}
	})

	counter := 0
	for _, n := range p.FreeNodes {
		if n != nil {
			counter++
		}
	}

	p.FreeNodes = slices.Clip(p.FreeNodes[:counter])
}
func (p *VectorPool[T]) Make(size int) []T {
	defer p.SortFreeNodes()

	if size <= 0 {
		return nil
	}

	_, node := p.findFree(size)

	if node != nil {
		return node.Allocate(size).Data
	}

	counter := 1

	for {
		if counter >= size {
			break
		}

		counter = counter << 1
	}
	node = p.newNode(counter)
	return node.Allocate(size).Data

}

func (p *VectorPool[T]) Clean() {
	p.FreeNodes = make([]*VectorPoolNode[T], len(p.Pool))
	copy(p.FreeNodes, p.Pool)

	for i := 0; i < len(p.Pool); i++ {
		p.Pool[i].MemoryNode.Children = nil
		p.Pool[i].State = Free

		for index := 0; index < len(p.Pool[i].Data); index++ {
			p.Pool[i].Data[index] = T(0)
		}
	}
}
