package main

import "fmt"

type DynamicArray struct {
	len,
	cap uint64
	arr []uint
}

func (da *DynamicArray) Append(newItem uint) {
	if da.len+1 > da.cap {
		da.cap *= 2
		newArr := make([]uint, da.len, da.cap)

		for i, value := range da.arr {
			newArr[i] = value
		}

		da.arr = newArr
	}

	da.arr = append(da.arr, newItem)
	da.len++
}

func (da *DynamicArray) Len() uint64 {
	return da.len
}

func (da *DynamicArray) Cap() uint64 {
	return da.cap
}

func (da *DynamicArray) Data() []uint {
	return da.arr
}

func New(cap uint64) *DynamicArray {
	return &DynamicArray{
		len: 0,
		cap: cap,
		arr: nil,
	}
}

func main() {
	myDynamicArray := New(2)

	for i := 1; i < 20; i++ {
		myDynamicArray.Append(uint(i))
	}

	fmt.Println(myDynamicArray.Len(), myDynamicArray.Cap(), myDynamicArray.Data())

	goDynamicArray := make([]uint, 0, 2)
	for i := 1; i < 20; i++ {
		goDynamicArray = append(goDynamicArray, uint(i))
	}

	fmt.Println(len(goDynamicArray), cap(goDynamicArray), goDynamicArray)
}
