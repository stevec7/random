package stackv2

import (
	"fmt"
	"testing"
)

var nodes = []*Node{
	NewNode(5),
	NewNode(10),
	NewNode(3),
	NewNode(12),
}

var nodesNeg = []*Node{
	NewNode(5),
	NewNode(3),
	NewNode(10),
	NewNode(-7),
}

type expected struct {
	sum    int
	avg    float64
	min    int
	max    int
	length int
}

func testExpected(s *Stack, e expected) error {
	min, _ := s.Min()
	max, _ := s.Max()

	if min.data != e.min {
		return fmt.Errorf("expected min %d, got %d", e.min, min.data)
	}
	if max.data != e.max {
		return fmt.Errorf("expected max %d, got %d", e.max, max.data)
	}
	if m := s.Avg(); m != e.avg {
		return fmt.Errorf("expected avg %f, got %f", e.avg, m)
	}
	if m := s.Sum(); m != e.sum {
		return fmt.Errorf("expected sum %d, got %d", e.sum, m)
	}
	if m := s.Size(); m != e.length {
		return fmt.Errorf("expected length %d, got %d", e.length, m)
	}
	return nil
}

func TestNewStack(t *testing.T) {
	s := NewStack()
	e := expected{
		sum:    30,
		avg:    7.5,
		min:    3,
		max:    12,
		length: 4,
	}

	for _, n := range nodes {
		s.Push(n)
	}

	err := testExpected(s, e)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}

func TestNegStack(t *testing.T) {
	s := NewStack()
	e := expected{
		sum:    11,
		avg:    2.75,
		min:    -7,
		max:    10,
		length: 4,
	}

	for _, n := range nodesNeg {
		s.Push(n)
	}

	err := testExpected(s, e)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}

func TestPop(t *testing.T) {
	s := NewStack()
	e := expected{
		sum:    18,
		avg:    6,
		min:    3,
		max:    10,
		length: 3,
	}

	for _, n := range nodesNeg {
		s.Push(n)
	}

	n, err := s.Pop()
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if n.data != -7 {
		t.Errorf("expected '%d', got '%d'", -7, n.data)
	}

	err = testExpected(s, e)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}
