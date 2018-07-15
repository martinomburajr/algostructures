package linked

import (
	"errors"
	"log"
)

type ILinkedList interface {
	//Adds to the tail
	Add(val int) (*LinkedList,error)
	Remove(index int) (*LinkedList,error)
	Get(index int) (*LinkedList,int,error)
	Replace(index int, val int) (*LinkedList,error)
}

type LinkedList struct {
	Root *Node
	Tail *Node
	Size int
}

type Node struct {
	Val int
	Next *Node
}

//Adds elements to the Linked List
func (l LinkedList)Add(val int) (*LinkedList, error) {
	if l.Size == 0 {
		l.Root.Val = val
		l.Size++
		return &l, nil
	}else if l.Size > 0 {
		temp := l.Root
		for i := 0; i < l.Size; i++{
			if i == (l.Size-1) {
				temp.Next = &Node{val, &Node{0, nil}}
				l.Size++
				return &l, nil
			}
			temp = temp.Next
		}
		return &l, errors.New("unexpected error, check arguments, list size etc")
	}else {
		return &l, errors.New("size is negative")
	}
}

func (l LinkedList) Get(index int) (*LinkedList, int,error) {
	if index >= l.Size {
		log.Panic("index cannot be greater than the linkedlists size, hint: check the linkedlists")
		return &l, -1, errors.New("index cannot be greater than the linkedlists size, hint: check the linkedlists")
	}
	temp := l.Root
	for i:=0; i < index; i++ {
		if i == index-1 {
			return &l, temp.Val, nil
		}
		temp = temp.Next
	}
	return &l, -1, errors.New("unable to get")
}

func (l LinkedList) Remove(index int) (*LinkedList,error) {
	if index >= l.Size {
		log.Panic("index cannot be greater than the linkedlists size, hint: check the linkedlists")
	}
	temp := l.Root
	if index == 0 {
		l.Root = l.Root.Next
		l.Size--
		return &l, nil
	} else if index == 1 {
		l.Root.Next = l.Root.Next.Next
		l.Size--
		return &l, nil
	} else {
		for i:=0; i < index; i++ {
			if i == (index-2) {
				temp.Next = temp.Next.Next
				temp.Next.Next = &Node{0, nil}
				l.Size--
				return &l, nil
			}
			temp = temp.Next
		}
	}
	return &l, errors.New("was unable to delete")
}

func (l LinkedList) Replace(index int, val int) (*LinkedList,error) {
	if index >= l.Size || index < 0 {
		log.Panic("index cannot be greater than the linkedlists size, hint: check the linkedlists")
	}

	if index == 0 {
		l.Root.Val = val
		return &l,nil
	} else if index == 1 {
		l.Root.Next.Val = val
		return &l,nil
	} else {
		//iterate to the index
		temp := l.Root
		for i:=0; i < index; i++ {
			if i == index {
				temp.Val = val
				return &l,nil
			}
			temp = temp.Next
 		}
		return &l, errors.New("unable to replace item in list")
	}
}

//[0,1,2,3]


func main() {
	
}
