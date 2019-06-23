package dlinkedlist

import (
	"fmt"
	"testing"
)

func TestLinkedList(t *testing.T) {
	l := &LinkedList{}

	t.Run("First nil", func(t *testing.T) {
		want := ListItem{}
		if got := l.First(); got != nil {
			t.Errorf("LinkedList.First() = %v, want %v", got, want)
		}
	})

	t.Run("Last nil", func(t *testing.T) {
		want := ListItem{}
		if got := l.First(); got != nil {
			t.Errorf("LinkedList.First() = %v, want %v", got, want)
		}
	})

	checkLinkedListPointers(l, t)
	//[1]
	l.PushBack(1)
	checkLinkedListPointers(l, t)
	//[1 2]
	l.PushBack(2)
	checkLinkedListPointers(l, t)
	//[3 1 2]
	l.PushFront(3)
	checkLinkedListPointers(l, t)
	//[4 3 1 2]
	l.PushFront(4)
	checkLinkedListPointers(l, t)
	//[5 4 3 1 2]
	l.PushFront(5)
	checkLinkedListPointers(l, t)
	//[5 4 3 1 2 6]
	l.PushBack(6)
	checkLinkedListPointers(l, t)
	//[5 4 3 1 2 6 7]
	l.PushBack(7)
	checkLinkedListPointers(l, t)
	//[5 4 3 1 2 6 7 8]
	l.PushBack(8)
	checkLinkedListPointers(l, t)
	//[5 4 3 2 6 7 8]
	_ = l.First().Next().Next().Next().Remove()
	checkLinkedListPointers(l, t)
	//["Rabbit", 5 4 3 2 6 7 8]
	l.PushFront("Rabbit")
	checkLinkedListPointers(l, t)
	//["Rabbit", 5 4 3 2 6 7 8, "Horse"]
	l.PushBack("Horse")
	checkLinkedListPointers(l, t)

	t.Run("Check length", func(t *testing.T) {
		want := 9
		if got := l.Len(); got != want {
			t.Errorf("LinkedList.Len() = %v, want %v", got, want)
		}
	})

	// Remove all elements
	t.Run("Remove all elements", func(t *testing.T) {
		ll := l.len
		for i := 0; i < ll; i++ {
			err := l.First().Remove()
			checkLinkedListPointers(l, t)
			if err != nil {
				t.Errorf("Cant remove element: %v", err)
			}
		}
	})

	// []
	t.Run("Expect error during remove operation", func(t *testing.T) {
		if err := l.First().Remove(); err == nil {
			t.Errorf("error expected in Remove() method")
		}
	})

	t.Run("First nil", func(t *testing.T) {
		want := ListItem{}
		if got := l.First(); got != nil {
			t.Errorf("LinkedList.First() = %v, want %v", got, want)
		}
	})

	t.Run("Last nil", func(t *testing.T) {
		want := ListItem{}
		if got := l.First(); got != nil {
			t.Errorf("LinkedList.First() = %v, want %v", got, want)
		}
	})
}

func checkLinkedListPointers(list *LinkedList, t *testing.T) {
	f := list.first
	l := list.last
	directOrder := make([]interface{}, 0, list.len)
	reverseOrder := make([]interface{}, 0, list.len)

	for {
		if f != nil && l != nil {
			directOrder = append(directOrder, f.value)
			reverseOrder = append(reverseOrder, l.value)
			f = f.next
			l = l.prev
		} else {
			break
		}
	}

	t.Run("Check pointers", func(t *testing.T) {
		fmt.Println(directOrder)
		if len(directOrder) != len(reverseOrder) {
			t.Errorf("lenght of directOrder: (%d) and reverseOrder: (%d) are different", len(directOrder), len(reverseOrder))
		}

		for i, el := range directOrder {
			if el != reverseOrder[len(reverseOrder)-i-1] {
				t.Errorf("directOrder(%v) and reverseOrder(%v) slices are different", directOrder, reverseOrder)
			}
		}
	})
}
