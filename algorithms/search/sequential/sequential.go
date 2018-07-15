package sequential

//Sequentially searches for the search item. Iterating sequentially over each element
//If the item is not found, it returns a value of -1
//It is imperative to check if the return value is -1
func SequentialSearch(arr []int, searchItem int) (index int) {
	for i := range arr {
		if arr[i] == searchItem {
			index = i;
			return index
		}
	}
	index = -1
	return index
}