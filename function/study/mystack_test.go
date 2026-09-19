package study

import "testing"

// TestSliceStackPushPopOrder exercises sliceStack (myStack.go) end to end:
// push three elements, then confirm Pop() unwinds them in LIFO order and
// that Size/IsEmpty agree with the stack's contents at each step. This
// repository had no *_test.go file at all before this pass.
func TestSliceStackPushPopOrder(t *testing.T) {
	s := NewSeliceStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	if got := s.Size(); got != 3 {
		t.Fatalf("Size() = %d, want 3", got)
	}
	if got := s.Pop(); got != 3 {
		t.Fatalf("Pop() = %v, want 3 (LIFO order)", got)
	}
	if got := s.Pop(); got != 2 {
		t.Fatalf("Pop() = %v, want 2 (LIFO order)", got)
	}
	if s.IsEmpty() {
		t.Fatalf("IsEmpty() = true with one element left, want false")
	}
	if got := s.Pop(); got != 1 {
		t.Fatalf("Pop() = %v, want 1 (LIFO order)", got)
	}
	if !s.IsEmpty() {
		t.Fatalf("IsEmpty() = false after popping every pushed element, want true")
	}
}
