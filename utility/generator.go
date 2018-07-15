package utility

import ("math/rand"
	"golang.org/x/tools/container/intsets"
)

//Generates a list with random numbers of a given size containing a single unique number that you
//specify in the @uniqueNumber param
func generateListWUniqueNumber(size int, uniqueNumber int) []int {
	list := make([]int, size)
	i := 0
	for i < size {
		intn := rand.Intn(intsets.MaxInt)
		if intn != uniqueNumber {
			list[i] = intn
			i++
		}
	}
	int2 := rand.Intn(size)
	list[int2] = uniqueNumber
	return list
}

//Generates a list of random numbers of a given size
func generateList(size int) []int {
	list := make([]int, size)
	i := 0
	for i < size {
		randInt := rand.Int()
		list[i] = randInt
		i++
	}
	return list
}

