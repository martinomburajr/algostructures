package algostructures

type Getter interface {
	Get() interface{}
}

type Adder interface {
	Add(item interface{})
}

type Sizer interface {
	Size() int64
}