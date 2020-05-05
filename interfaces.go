package algostructures

//Getter contains behavior relating to an object that given a location (index) can retrieve contents from that location
type Getter interface {
	Get(index int64) interface{}
}

//Adder contains behavior relating to an object that has the ability to add to itself
type Adder interface {
	Add(item interface{})
}

//Sizer contains behavior that gives a type the ability to report on its size
type Sizer interface {
	Size() int64
}