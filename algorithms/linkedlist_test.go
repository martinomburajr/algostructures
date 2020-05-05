package main

import (
	"testing"
	"reflect"
)
var (
	node0 = &Node{id: "1", value: 1}
	node1 = &Node{id: "2", value: 2}
	node2 = &Node{id: "3", value: 3}

)
func Test_LinkedListGet(t *testing.T) {
	testList := LinkedList{}
	testList.root = node0
	testList.root.next = node1
	testList.root.next.next = node2

	type args struct {
		linkedList LinkedList
		index int
	}
	tests := []struct {
		name string
		args args
		want *Node
		wantErr bool
	} {
		{"empty list", args{LinkedList{}, 1}, &Node{}, true},
		{"list size [1] | index [1]", args{testList, 1}, &Node{}, true},
		{"list size [0] | index [0]", args{testList, 0}, node0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func (t *testing.T) {
			if tt.wantErr {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("the code did not panic")
					}
				}()
				return
			}
			got := tt.args.linkedList.Get(tt.args.index)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("\n%s\n\tinput: \t%v\n\twant: \t%#v\n\tgot: \t%#v\n", tt.name, tt.args, tt.want, got)
			}
		})
	}
}
