package linkedlist

import (
	"reflect"
	"testing"
)
var t1f = LinkedList{count: 1, head: &node{1, &node{2, nil}}}

var t0 = LinkedList{}
var t1 = LinkedList{count: 1, head: &node{1, nil}}
var t2 = LinkedList{count: 2, head: &node{1, &node{2, nil}}}
var t3 = LinkedList{count: 3, head: &node{1, &node{2, &node{3, nil}}}}

func TestLinkedList_Add(t *testing.T) {
	type args struct {
		item interface{}
	}
	tests := []struct {
		name   string
		fields LinkedList
		args   args
	}{
		{"add to empty list", t0, args{1}},
		{"add to list with 1 item", t1, args{2}},
		//{"add to list with 1 item but inconsistent count", t1f, args{2} },
		{"add to list with 2 items", t2, args{3}},
		{"add to list with 3 items", t3, args{4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{
				count: tt.fields.count,
				head:  tt.fields.head,
			}
			l.Add(tt.args.item)

			if l.count < 1 {
				t.Error("item was added to linked list, it cannot have size less than 1")
			}

			if l.count == 1 && l.head.v != tt.args.item {
				t.Error("item was added to linked list, the list is of size one, " +
					"but the head of the linked list has not received the value.")
			}

			if l.count == 1 && l.head.next != nil {
				t.Error("inconsistent count. Linked List size of 1 but head has a link to a non-nil node")
			}

			if l.count > 1 && l.head.next == nil {
				t.Error("failed to attach new item to head")
			}

		})
	}
}

func BenchmarkLinkedList_Add(b *testing.B) {
	for n := 0; n < b.N; n++ {
		t1.Add(n)
	}
}

func TestLinkedList_Size(t *testing.T) {
	tests := []struct {
		name   string
		fields LinkedList
		want   int64
	}{
		{"get size count 0", t0, 0},
		{"get size count 1", t1, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{
				count: tt.fields.count,
				head:  tt.fields.head,
			}
			if got := l.Size(); got != tt.want {
				t.Errorf("LinkedList.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_Get(t *testing.T) {
	type args struct {
		index int64
	}
	tests := []struct {
		name    string
		fields  LinkedList
		args    args
		want    interface{}
		wantErr bool
	}{
		{"get 0 from empty list", t0, args{0}, nil, true},
		{"get -1 from empty list", t0, args{-1}, nil, true},
		{"get 1 from empty list", t0, args{1}, nil, true},
		{"get 0 from list.count = 1", t1, args{0}, 1, false},
		{"get 1 from list.count = 1", t1, args{1}, nil, true},
		{"get 1 from list.count = 2", t2, args{1}, 2, false},
		{"get 2 from list.count = 3", t3, args{2}, 3, false},
		{"get 1 from list.count = 3", t3, args{1}, 2, false},
		{"get 0 from list.count = 3", t3, args{0}, 1, false},
		//{"get 0 from empty list", t0, args{1}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{
				count: tt.fields.count,
				head:  tt.fields.head,
			}
			got, err := l.Get(tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("LinkedList.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LinkedList.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
