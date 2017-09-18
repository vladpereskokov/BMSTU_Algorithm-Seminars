package linkedList

import (
	"testing"
)

func TestSinglyListInsertAfter(t *testing.T) {
	list := New()

	Insert(list, 3)
	Insert(list, 5)
	Insert(list, 6)

	array := []int{3, 5, 6}
	for _, value := range array {
		if value != list.Head.Data {
			t.Errorf("Don't match: %d and %d\n", value, list.Head.Data)
		}

		list.Head = list.Head.Next
	}
}

func TestSinglyListHas(t *testing.T) {
	list := New()

	Insert(list, 3)
	Insert(list, 5)
	Insert(list, 6)

	if !Has(list, 3) {
		t.Error("Don't match: not found")
	}
}

func TestSinglyListNotHas(t *testing.T) {
	list := New()

	Insert(list, 3)
	Insert(list, 5)
	Insert(list, 6)

	if Has(list, 0) {
		t.Error("Don't match: not found")
	}
}

func TestSinglyListLoop(t *testing.T) {
	list := New()

	Insert(list, 3)
	Insert(list, 5)
	Insert(list, 6)

	list.Head.Next.Next.Next = list.Head

	if !IsLoopList(list) {
		t.Error("Don't match: loop not found")
	}
}

func TestSinglyListNoLoop(t *testing.T) {
	list := New()

	Insert(list, 3)
	Insert(list, 5)
	Insert(list, 6)

	if IsLoopList(list) {
		t.Error("Don't match: loop found")
	}
}
