package main

import (
	"fmt"
	"studyGo/function/study"
)

func main(){
	//fmt.Print("Hello World!")
	//study.SortSelice()
	//stack:= study.NewSliceEntry()
	//stack.Push(1)`
	//fmt.Println(*stack)
	mySlice:=study.NewSeliceStack()
	for i := 0; i < 100; i++ {
		mySlice.Push(i)
	}
	//for i := mySlice.Size(); i > 0; i-- {
	//	fmt.Print(mySlice.Pop(),"\t\n")
	//}
	mySlice.Clear()
	fmt.Println("is empty?\t",mySlice.IsEmpty())
}