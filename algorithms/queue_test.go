package main


import (
	"testing"
	"reflect"
)

var (
	queue0 = Queue{}
	queue1 = Queue{Head: &QueueNode{data: 0}}
	queue2 = Queue{Head: &QueueNode{data: 0, next: &QueueNode{data: 1}}}
	queue3 = Queue{Head: &QueueNode{data: 0, next: &QueueNode{data: 1, next: &QueueNode{data: 2}}}}
)

func TestQueueAdd(t *testing.T) {
	type args struct {
		queue Queue
		data int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"empty queue", args{queue: queue0, data: 123}, 1},
		{"queue 1", args{queue: queue1, data: 9}, 2},
		{"queue 2", args{queue: queue2, data: 5}, 3},
		{"queue 3", args{queue: queue3, data: 2}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			count := tt.args.queue.Add(tt.args.data)
			if count != tt.want {
				t.Errorf("Count Mismatch:\nwant:\t%d\ngot:\t%d", tt.want, count)
			}

			traversedNodes := tt.args.queue.traverse()
			addedNode := traversedNodes[len(traversedNodes)-1]

			if tt.args.data != addedNode.data {
				t.Errorf("Added Item Error:\nwant:\t%d\ngot:\t%d", tt.args.data, addedNode.data)
			}
		})
	}
}

func TestQueueRemove(t *testing.T) {
	// type args struct {
	// 	queue Queue
	// }
	// tests := []struct {
	// 	name string
	// 	args args
	// 	want *QueueNode //size
	// }{
	// 	{"empty queue", args{queue: queue0}, nil},
	// 	{"queue 1", args{queue: queue1}, &QueueNode{data: 0}},
	// 	{"queue 2", args{queue: queue2}, &QueueNode{data: 0}},
	// 	{"queue 3", args{queue: queue3}, &QueueNode{data: 0}},
	// }
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		removedItem := tt.args.queue.Remove()
	// 		if !reflect.DeepEqual(removedItem, tt.want) {
	// 			t.Errorf("Remove Item Error:\nwant:\t%#v\ngot:\t%#v", tt.want, removedItem)
	// 		}
	// 	})
	// }
	removedItem := queue3.Remove()
	removedItem = queue3.Remove()
	removedItem = queue3.Remove()
	t.Logf("PRINT: %s", queue3.print())
	want := &QueueNode{data: 2}
	if !reflect.DeepEqual(removedItem, &QueueNode{data: 2}) {
		t.Errorf("Remove Item Error:\nwant:\t%#v\ngot:\t%#v", want, removedItem)
	}
}