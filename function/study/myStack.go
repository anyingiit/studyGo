package study

import (
	"fmt"
)

type element interface {}

type sliceStack struct {
	element []element
}

func NewSeliceStack() *sliceStack {
	return &sliceStack{}
}

func (entry *sliceStack) Push(element element) {
	entry.element = append(entry.element,element)
}

func (entry *sliceStack) Size() int {
	return len(entry.element)
}

func (entry *sliceStack) Top() element {
	if entry.IsEmpty() {
		return nil
	}
	return entry.element[entry.Size()-1]
}

func (entry *sliceStack) Pop() element {
	if entry.IsEmpty() {
		return nil
	}
	lastElement:=entry.element[entry.Size()-1]
	entry.element[entry.Size()-1] = nil
	entry.element = entry.element[:entry.Size()-1]
	return lastElement
}

func (entry *sliceStack) Clear() bool{
	if entry.IsEmpty() {
		fmt.Print("Stack is empty!")
		return false
	}
	for i := 0; i < entry.Size(); i++ {
		entry.element[i] = nil
	}
	entry.element = make([]element,0)
	return true
}

func (entry *sliceStack) IsEmpty() bool {
	return entry.Size() == 0
}