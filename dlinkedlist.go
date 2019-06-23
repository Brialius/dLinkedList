package dlinkedlist

import (
	"errors"
	"fmt"
)

type LinkedList struct {
	len         int
	first, last *ListItem
}

type ListItem struct {
	value      interface{}
	prev, next *ListItem
	list       *LinkedList
}

func (l *LinkedList) String() string {
	if l != nil {
		return fmt.Sprintf("Addr: (%p), Len: %d, First: (%v), Last: (%v)", &l, l.len, l.first, l.last)
	} else {
		return "Empty list"
	}
}

func (l *LinkedList) Len() int {
	return l.len
}

func (l *LinkedList) First() *ListItem {
	return l.first
}

func (l *LinkedList) Last() *ListItem {
	return l.last
}

func (l *LinkedList) PushFront(v interface{}) {
	newItem := ListItem{
		list:  l,
		value: v,
		prev:  nil,
		next:  l.first,
	}

	if l.len == 0 {
		l.last = &newItem
	}

	if newItem.next != nil {
		newItem.next.prev = &newItem
	}

	l.first = &newItem
	l.len++
}

func (l *LinkedList) PushBack(v interface{}) {
	newItem := ListItem{
		list:  l,
		value: v,
		prev:  l.last,
		next:  nil,
	}

	if l.len == 0 {
		l.first = &newItem
	}

	if newItem.prev != nil {
		newItem.prev.next = &newItem
	}

	l.last = &newItem
	l.len++
}

func (item *ListItem) Value() interface{} {
	return item.value
}

func (item *ListItem) Next() *ListItem {
	return item.next
}

func (item *ListItem) Prev() *ListItem {
	return item.prev
}

func (item *ListItem) Remove() error {
	if item != nil {
		if item.prev != nil {
			item.prev.next = item.next
		} else {
			item.list.first = item.next
		}

		if item.next != nil {
			item.next.prev = item.prev
		} else {
			item.list.last = item.prev
		}

		if item.list.len > 0 {
			item.list.len--
		}
		return nil
	}
	return errors.New("can't delete empty element")
}
