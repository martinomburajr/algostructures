package linked

import "testing"

//Tests the LinkedList get method
var linkedList1 LinkedList

func init() {
	var c = Node {12, nil}
	var b = Node {4, &c}
	var a = Node {0, &b}

	linkedList1 = LinkedList{&a, &c, 3}
}

func TestAddTable(t *testing.T) {
	test1Func := func (val int) *LinkedList {
		list, e := linkedList1.Add(val)
		if e != nil {
			t.Fatalf("error: %v", e)
		}
		return list
	}

	tests := []struct {
		expectedSize int
		inputVal int
		input func(val int) *LinkedList
	} {
		{ 4, 9,test1Func},
	}

	for _,test :=  range tests  {
		hasNoSize := test.expectedSize != test.input(test.inputVal).Size
		_,val,_ := test.input(test.inputVal).Get(3)
		hasNoForm := test.inputVal != val
		if hasNoForm && hasNoSize {
			t.Error("lists are not of the same size and the correct element may have not eben added")
		}
	}
}
//func testGetTable(t *testing.T) {
//	//LinkedList.Get()
//	tests := []struct {
//		expected int
//		input LinkedList
//	}{
//		{2, LinkedList{Node{0,}}.Get(3)}
//	}
//}