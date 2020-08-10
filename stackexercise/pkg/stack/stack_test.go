package stack

import (
    "errors"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

var numsSmall = []int{
	1, 4, 5, 3, 2,
}

var numsSmallNeg = []int{
	1, 4, 5, 3, -2,
}

type expected struct {
	sum    int
	avg    float64
	min    int
	max    int
	length int
}

func createSmallStack(n []int) *Stack {
	s := NewStack()
	for _, i := range n {
		s.Push(i)
	}
	return s
}

func testExpected(s *Stack, e expected) error {
	if m, _ := s.Min(); m != e.min {
		return fmt.Errorf("expected min %d, got %d", e.min, m)
	}
	if m, _ := s.Max(); m != e.max {
		return fmt.Errorf("expected max %d, got %d", e.max, m)
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

func TestInitStack(t *testing.T) {
	s := NewStack()
	_, err := s.Min()
	if err == nil {
		t.Errorf("shouldve received error for empty stack")
	}

	_, err = s.Pop()
	if err == nil {
		t.Errorf("shouldve received error for empty stack")
	}

}

func TestStackPush(t *testing.T) {
	e := expected{15, 3.0, 1, 5, 5}
	s := createSmallStack(numsSmall)

	err := testExpected(s, e)
	if err != nil {
		t.Errorf(err.Error())
	}

}

func TestStackPop(t *testing.T) {
	e := expected{13, 3.25, 1, 5, 4}
	s := createSmallStack(numsSmall)

	// remove the last added element (2)
	expectedLast := 2
	last, err := s.Pop()
	if err != nil {
		t.Errorf("the stack isnt empty, we shouldnt have error'd here")
	}

	if last != expectedLast {
		t.Errorf("expected %d, got %d", expectedLast, last)
	}

	err = testExpected(s, e)
	if err != nil {
		t.Errorf(err.Error())
	}

}
func TestStackPushNeg(t *testing.T) {
	e := expected{11, 2.2, -2, 5, 5}
	s := createSmallStack(numsSmallNeg)

	err := testExpected(s, e)
	if err != nil {
		t.Errorf(err.Error())
	}

}
func TestStackPopNeg(t *testing.T) {
	e := expected{13, 3.25, 1, 5, 4}
	s := createSmallStack(numsSmallNeg)

	// remove the last added element (2)
	expectedLast := -2
	last, err := s.Pop()
	if err != nil {
		t.Errorf("the stack isnt empty, we shouldnt have error'd here")
	}

	if last != expectedLast {
		t.Errorf("expected %d, got %d", expectedLast, last)
	}

	err = testExpected(s, e)
	if err != nil {
		t.Errorf(err.Error())
	}

	// removing a single element is the same as TestStackPop, so remove another
	e = expected{10, float64(10) / float64(3), 1, 5, 3}
	last, err = s.Pop()
	if err != nil {
		t.Errorf("the stack isnt empty, we shouldnt have error'd here")
	}
	expectedLast = 3
	err = testExpected(s, e)
	if err != nil {
		t.Errorf(err.Error())
	}

}

func TestDrain(t *testing.T) {
	s := createSmallStack(numsSmall)

	for i := s.Size(); i >= -1; i-- {
		_, err := s.Pop()
		if err != nil {
			if !errors.Is(err, ErrStackEmpty) {
				t.Errorf("shouldve only gotten a stack empty error only")
			}
		}
	}
}

/*
func TestDrain(t *testing.T) {
	s := createSmallStack(numsSmall)

	for i := 0; i < len(numsSmall)+1; i++ {
		_, err := s.Pop()
		if i == len(numsSmall) {
			if err == nil {
				t.Errorf("shouldve gotten an empty stack error")
			}
		} else {
			if err != nil {
				t.Errorf("i: %d, %v\n", i, err.Error())
			}
		}
	}
}
*/

func createRandomStack(n int) *Stack {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	s := NewStack()
	for i := 0; i < n; i++ {
		s.Push(r1.Intn(1024))
	}
	return s
}

func BenchmarkPush(b *testing.B) {
	_ = createRandomStack(b.N)
}

func BenchmarkPushB(b *testing.B) {
	s := NewStack()
	for i := 0; i < b.N; i++ {
		s.Push(b.N)
	}
}

func BenchmarkPush1000000(b *testing.B) {
	_ = createRandomStack(1000000)
}

func BenchmarkPop1000000(b *testing.B) {
	s := createRandomStack(1000000)
	for i := 0; i < b.N; i++ {
		_, _ = s.Pop()
	}
}
