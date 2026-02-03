package vector_test

import (
	"testing"

	"github.com/RichardAlmanza/learning-stuff/go/artificial-intelligence/neural-networks-from-scratch/pkgs/vector"
)

func TestVectorPool_Make(t *testing.T) {
	pool := &vector.VectorPool[int]{}

	data1 := pool.Make(5)
	if len(data1) != 5 {
		t.Errorf("Expected data length to be 5, got %d", len(data1))
	}
	for _, v := range data1 {
		if v != 0 {
			t.Errorf("Expected all elements to be 0, got %v", data1)
		}
	}

	data2 := pool.Make(10)
	if len(data2) != 10 {
		t.Errorf("Expected data length to be 10, got %d", len(data2))
	}
	for _, v := range data2 {
		if v != 0 {
			t.Errorf("Expected all elements to be 0, got %v", data2)
		}
	}

	for i := 0; i < len(data1); i++ {
		data1[i] = i
		data2[i] = i
	}

	pool.Clean()
	data1 = pool.Make(5)
	if len(data1) != 5 {
		t.Errorf("Expected data length to be 5, got %d", len(data1))
	}
	for _, v := range data1 {
		if v != 0 {
			t.Errorf("Expected all elements to be 0, got %v", data1)
		}
	}
}

func TestVectorPool_Clean(t *testing.T) {
	pool := &vector.VectorPool[int]{}

	data := pool.Make(5)
	if len(data) != 5 {
		t.Errorf("Expected data length to be 5, got %d", len(data))
	}
	for _, v := range data {
		if v != 0 {
			t.Errorf("Expected all elements to be 0, got %v", data)
		}
	}

	pool.Clean()
	data = pool.Make(5)
	if len(data) != 5 {
		t.Errorf("Expected data length to be 5, got %d", len(data))
	}
	for _, v := range data {
		if v != 0 {
			t.Errorf("Expected all elements to be 0, got %v", data)
		}
	}
}

func TestVectorPool_Allocate(t *testing.T) {
	pool := &vector.VectorPool[int]{}

	pool.Make(17)
	nodes := pool.Pool[0].MemoryNode.Children
	var node *vector.VectorPoolNode[int]

	for _, n := range nodes {
		if n.State == vector.Free {
			node = n
		}
	}
	data := node.Allocate(3).Data
	if len(data) != 3 {
		t.Errorf("Expected data length to be 3, got %d", len(data))
	}
	for _, v := range data {
		if v != 0 {
			t.Errorf("Expected all elements to be 0, got %v", data)
		}
	}

	nodes = node.MemoryNode.Children

	for _, n := range nodes {
		if n.State == vector.Free {
			node = n
		}
	}
	data = node.Allocate(2).Data
	if len(data) != 2 {
		t.Errorf("Expected data length to be 2, got %d", len(data))
	}
	for _, v := range data {
		if v != 0 {
			t.Errorf("Expected all elements to be 0, got %v", data)
		}
	}
}
