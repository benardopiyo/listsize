package main

import (
    "github.com/01-edu/z01"
)

// NodeL defines a node in the linked list
type NodeL struct {
    Data interface{}
    Next *NodeL
}

// List defines the linked list with head and tail pointers
type List struct {
    Head *NodeL
    Tail *NodeL
}

// ListSize returns the number of elements in the linked list l.
func ListSize(l *List) int {
    count := 0
    current := l.Head
    for current != nil {
        count++
        current = current.Next
    }
    return count
}

// ListPushFront adds a new element to the front of the list.
func ListPushFront(l *List, data interface{}) {
    newNode := &NodeL{Data: data, Next: l.Head}
    l.Head = newNode
    if l.Tail == nil { // If the list was empty, update the tail
        l.Tail = newNode
    }
}

func printInt(n int) {
    // Handle negative numbers (not needed for this exercise but good practice)
    if n < 0 {
        z01.PrintRune('-')
        n = -n
    }
    if n == 0 {
        z01.PrintRune('0')
        z01.PrintRune('\n')
        return
    }

    // Convert integer to string and print using z01
    var stack []rune
    for n > 0 {
        stack = append(stack, rune('0'+(n%10)))
        n /= 10
    }
    for i := len(stack) - 1; i >= 0; i-- {
        z01.PrintRune(stack[i])
    }
    z01.PrintRune('\n')
}

func main() {
    list := &List{}

    ListPushFront(list, "Hello")
    ListPushFront(list, "2")
    ListPushFront(list, "you")
    ListPushFront(list, "man")

    size := ListSize(list)
    printInt(size)
}
