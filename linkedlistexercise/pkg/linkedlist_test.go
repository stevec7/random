package linkedlistexercise

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateList(t *testing.T) {
	ll := NewLinkedList()
	for i := 0; i < 10; i++ {
		node := NewNode(i*2)
		ll.Add(node)
	}
	t.Run("size", func(t *testing.T){
		require.Equal(t, 10, ll.Size())
	})
}

func TestCreateEmptyList(t *testing.T) {
	ll := NewLinkedList()
	require.Equal(t, true, ll.Empty())

	ll.Apply(func(i int) int{
		return i*i
	})
	require.Equal(t, true, ll.Empty())

	node := ll.Pop()
	require.Nil(t, node)
}

func testDouble(i int) int {
	return i*2
}


func TestApply(t *testing.T) {
	seedData := []int{1, 2, 4, 8}
	expected := []int{2, 4, 8, 16}
	expected2 := []int{4, 16, 64, 256}
	ll := NewLinkedList()
	for _, s := range seedData {
		node := NewNode(s)
		ll.Add(node)
	}

	t.Run("double", func(t *testing.T){
		ll.Apply(testDouble)
		require.Equal(t, expected, ll.Values())
	})

	t.Run("anon", func(t *testing.T){
		ll.Apply(func(i int) int{
			return i*i
		})
		require.Equal(t, expected2, ll.Values())
	})
}

func TestInsertion(t *testing.T) {
	ll := NewLinkedList()
	for i := 0; i < 10; i++ {
		node := NewNode(i*2)
		ll.Add(node)
	}

	t.Run("push", func(t *testing.T){
		ll.Push(NewNode(66))
		require.Equal(t, 0, ll.Index(66))
	})

	t.Run("push_index", func(t *testing.T){
		ll.PushIndex(NewNode(67), 3)
		require.Equal(t, 3, ll.Index(67))
	})

	t.Run("push_index_head", func(t *testing.T){
		ll.PushIndex(NewNode(71), 0)
		idx := ll.Index(71)
		require.Equal(t, 0, idx)
	})

	t.Run("push_index_tail", func(t *testing.T){
		ll.PushIndex(NewNode(73), ll.Size())
		idx := ll.Index(73)
		require.Equal(t, ll.Size()-1, idx)
	})

	t.Run("push_index_bounds", func(t *testing.T){
		ll.PushIndex(NewNode(68), 20)
		require.Equal(t, 4, ll.Index(67))
		idx := ll.Index(68)
		require.Equal(t, -1, idx)
	})
}

func TestRemoval(t *testing.T) {
	ll := NewLinkedList()
	for i := 0; i < 10; i++ {
		node := NewNode(i*2)
		ll.Add(node)
	}

	t.Run("pop", func(t *testing.T){
		node := ll.Pop()
		require.Equal(t, 0, node.data)
	})

	t.Run("pop_index", func(t *testing.T){
		node := ll.PopIndex(3)
		require.Equal(t, 8, node.data)
	})

	t.Run("pop_index_head", func(t *testing.T){
		node := ll.PopIndex(0)
		require.Equal(t, 2, node.data)
	})

	t.Run("pop_index_tail", func(t *testing.T){
		node := ll.PopIndex(ll.Size()-1)
		require.Equal(t, 18, node.data)
	})

	t.Run("pop_index_bounds", func(t *testing.T){
		node := ll.PopIndex(ll.Size())
		require.Nil(t, node)
	})
}

func TestReverse(t *testing.T) {
	ll := NewLinkedList()
	for i := 0; i < 10; i++ {
		node := NewNode(i*2)
		ll.Add(node)
	}

	llr := NewLinkedList()
	for i := 9; i > -1; i-- {
		node := NewNode(i*2)
		llr.Add(node)
	}

	ll.Reverse()
	require.Equal(t, llr.Values(), ll.Values())

}
