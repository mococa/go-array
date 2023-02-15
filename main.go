package goarray

/* Types */

type GoArray[T comparable] struct {
	data []T
}

/* Methods */

/*
Instantiates a new GoArray

initial value is optional
*/
func NewGoArray[T comparable](initial_value ...[]T) *GoArray[T] {
	if len(initial_value) > 0 {
		return &GoArray[T]{data: initial_value[0]}
	}

	return &GoArray[T]{data: []T{}}
}

/*
Inserts an item on end of array
*/
func (arr *GoArray[T]) Push(value T) {
	arr.data = append(arr.data, value)
}

/*
Returns the occurence item on array based on given position
*/
func (arr *GoArray[T]) At(position int) (data T) {
	if position > len(arr.data) {
		return
	}

	for i, item := range arr.data {
		if i == position {
			return item
		}
	}

	return
}

/*
Removes occurence item on array where the array is the given position
*/
func (arr *GoArray[T]) RemoveAt(position int) {
	if position > len(arr.data) {
		return
	}

	for i := range arr.data {
		if i == position {
			arr.data = append(arr.data[:i], arr.data[i+1:]...)
		}
	}
}

/*
Removes all the occurences on array where given comparing function returns true
*/
func (arr *GoArray[T]) RemoveWhere(fn func(item T) bool) {
	for i := 0; i < len(arr.data); i++ {
		if fn(arr.data[i]) {
			arr.RemoveAt(i)
		}
	}
}

/*
Returns the first occurence item on array based on given comparing function

if it doesn't find anything, it returns -1
*/
func (arr *GoArray[T]) Find(fn func(item T) bool) (value T) {
	for i := 0; i < len(arr.data); i++ {
		if fn(arr.data[i]) {
			return arr.data[i]
		}
	}

	return
}

/*
Returns the first occurence index on array based on given comparing function

if it doesn't find anything, it returns -1
*/
func (arr *GoArray[T]) FindIndex(fn func(item T) bool) (idx int) {
	for i := 0; i < len(arr.data); i++ {
		if fn(arr.data[i]) {
			return i
		}
	}

	return -1
}

/*
Returns a new array filtered given predicate function
*/
func (arr *GoArray[T]) Filter(fn func(item T) bool) []T {
	arr_copy := make([]T, len(arr.data))

	copy(arr_copy, arr.data)

	go_arr_copy := NewGoArray[T]()

	for i := 0; i < len(arr.data); i++ {
		if fn(arr.data[i]) {
			go_arr_copy.Push(arr.data[i])
		}
	}

	return go_arr_copy.data
}

/*
Returns a modified version of the array based on given predicate function
*/
func (arr *GoArray[T]) Map(fn func(item T) any) []any {
	arr_copy := make([]any, len(arr.data))

	for i, e := range arr.data {
		arr_copy[i] = fn(e)
	}

	return arr_copy

}

/*
Iterates each item in array, not returning a value
*/
func (arr *GoArray[T]) Each(fn func(item T)) {
	for _, e := range arr.data {
		fn(e)
	}
}
