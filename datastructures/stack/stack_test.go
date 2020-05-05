package stack

import (
	"testing"
	"reflect"
)

func TestSize(t *testing.T) {
	tests := []struct {
		name string
		stack Stack
		want int
	} {
		{ "empty", Stack{}, 0},
		{ "size-1", Stack{head: &Node{value: 0}}, 1},
		{ "size-2", Stack{head: &Node{value: 0, next: &Node{value: 1}}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.stack.Size()
			if got != tt.want {
				t.Errorf("want: %d | got: %d", tt.want, got)
			}
		})
	}
}

func TestPush(t *testing.T) {
	type args struct {
		data int
	}
	tests := []struct {
		name string
		stack Stack
		args args
		wantSize int
	} {
		{ "empty", Stack{}, args{data: 0}, 1},
		{ "size-1", Stack{head: &Node{value: 0}}, args{data: 1}, 2},
		{ "size-2", Stack{head: &Node{value: 0, next: &Node{value: 1}}}, args{data: 2}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.stack.Push(tt.args.data)
			gotSize := tt.stack.Size()

			if gotSize != tt.wantSize {
				t.Errorf("wantSize: %d | gotSize: %d", tt.wantSize, gotSize)
			}
			if tt.stack.head.value != tt.args.data {
				t.Errorf("wantHeadValue: %d | gotHeadValue: %d", tt.args.data, tt.stack.head.value)
			}
		})
	}
}

func TestPop(t *testing.T) {
	node0 := &Node{value: 0}
	node1 := &Node{value: 0, next: &Node{value: 1}}

	tests := []struct {
		name string
		stack Stack
		want *Node
		wantSize int
	} {
		{ "empty", Stack{}, nil, 0},
		{ "size-1", Stack{head: node0}, node0, 0},
		{ "size-2", Stack{head: node1}, node1, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPop := tt.stack.Pop()
			gotSize := tt.stack.Size()

			if gotSize != tt.wantSize {
				t.Errorf("wantSize: %d | gotSize: %d", tt.wantSize, gotSize)
			}
			if tt.want != nil {
				if tt.want.value != gotPop.value {
					t.Errorf("wantHeadValue: %d | gotHeadValue: %d", tt.want.value, gotPop.value)
				}
			}
		})
	}
}

func TestPeek(t *testing.T) {
	node0 := &Node{value: 0}
	node1 := &Node{value: 0, next: &Node{value: 1}}

	tests := []struct {
		name string
		stack Stack
		want *Node
	} {
		{ "empty", Stack{}, nil},
		{ "size-1", Stack{head: node0}, node0},
		{ "size-2", Stack{head: node1}, node1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPop := tt.stack.Peek()

			if !reflect.DeepEqual(tt.want, gotPop) {
				t.Errorf("wantHeadValue: %#v | gotHeadValue: %#v", tt.want, gotPop)
			}
		})
	}
}

func TestClearStack(t *testing.T) {
	node0 := &Node{value: 0}
	nodeSep2 := &Node{value: 2}
	nodeSep1 := &Node{value: 0, next: nodeSep2}

	tests := []struct {
		name string
		stack Stack
		want []*Node
	} {
		{ "empty", Stack{}, []*Node{}},
		{ "size-1", Stack{head: node0}, []*Node{node0}},
		{ "size-2", Stack{head: nodeSep1}, []*Node{nodeSep2, nodeSep1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotList := tt.stack.ClearStack()

			if len(gotList) != len(tt.want) {
				t.Errorf("List len not equal\ngot: %d | want: %d", len(gotList), len(tt.want))
			}
			for i := range gotList {
				if gotList[i].value != tt.want[i].value {
					t.Errorf("List items not equal\ngot: %#v | want: %#v", gotList, tt.want)
				}
			}
		})
	}
}