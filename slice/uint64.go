package slice

import (
	"fmt"
	"math/rand"
)

// IndexOfUInt64 gets the index of an uint64 element in an uint64 slice
func IndexOfUInt64(x []uint64, y uint64) int {
	for i, v := range x {
		if v == y {
			return i
		}
	}
	return -1
}

// ContainsUInt64 checks whether an uint64 element is in an uint64 slice
func ContainsUInt64(x []uint64, y uint64) bool {
	return IndexOfUInt64(x, y) != -1
}

// EqualsUInt64s checks whether two uint64 slice has the same elements
func EqualsUInt64s(x []uint64, y []uint64) bool {
	if len(x) != len(y) {
		return false
	}

	for i := 0; i < len(x); i++ {
		if x[i] != y[i] {
			return false
		}
	}

	return true
}

// CopyUInt64s makes a new uint64 slice that copies the content of the given uint64 slice
func CopyUInt64s(x []uint64) []uint64 {
	return append([]uint64{}, x...)
}

// CutUInt64s cuts an uint64 slice by removing the elements starts from i and ends at j-1
func CutUInt64s(x []uint64, i, j int) ([]uint64, error) {
	if i < 0 || j > len(x) {
		return x, fmt.Errorf("out of bound")
	}
	if i >= j {
		return x, fmt.Errorf("%d must be smaller than %d", i, j)
	}
	return append(x[:i], x[j:]...), nil
}

// RemoveUInt64 removes an uint64 from a given uint64 slice by value
func RemoveUInt64(x []uint64, y uint64) []uint64 {
	index := IndexOfUInt64(x, y)
	if index != -1 {
		return append(x[:index], x[(index+1):]...)
	}
	return x
}

// RemoveUInt64At removes an uint64 from a given uint64 slice by index
func RemoveUInt64At(x []uint64, index int) ([]uint64, error) {
	if index < 0 || index > len(x) {
		return x, fmt.Errorf("out of bound")
	}
	return append(x[:index], x[(index+1):]...), nil
}

// InsertUInt64At inserts an uint64 value into a given uint64 slice at given index
func InsertUInt64At(x []uint64, y uint64, index int) ([]uint64, error) {
	if index < 0 || index > len(x) {
		return x, fmt.Errorf("out of bound")
	}
	x = append(x, 0)
	copy(x[index+1:], x[index:])
	x[index] = y
	return x, nil
}

// InsertUInt64sAt inserts a uint64 slice into a given uint64 slice at given index
func InsertUInt64sAt(x []uint64, y []uint64, index int) ([]uint64, error) {
	if index < 0 || index > len(x) {
		return x, fmt.Errorf("out of bound")
	}
	return append(x[:index], append(y, x[index:]...)...), nil
}

// PopFirstUInt64 pops the first value of an uint64 slice
func PopFirstUInt64(x []uint64) (uint64, []uint64, error) {
	if len(x) == 0 {
		return 0, nil, fmt.Errorf("no value to pop")
	}
	return x[0], x[1:], nil
}

// PopLastUInt64 pops the last value of an uint64 slice
func PopLastUInt64(x []uint64) (uint64, []uint64, error) {
	if len(x) == 0 {
		return 0, nil, fmt.Errorf("no value to pop")
	}
	return x[len(x)-1], x[:len(x)-1], nil
}

// FilterUInt64s filters an uint64 slice by the given filter function
func FilterUInt64s(x []uint64, filter func(uint64) bool) []uint64 {
	y := x[:0]
	for _, v := range x {
		if filter(v) {
			y = append(y, v)
		}
	}
	return y
}

// ReverseUInt64s reverses an uint64 slice
func ReverseUInt64s(x []uint64) []uint64 {
	for i := len(x)/2 - 1; i >= 0; i-- {
		opp := len(x) - 1 - i
		x[i], x[opp] = x[opp], x[i]
	}
	return x
}

// ShuffleUInt64s shuffles an uint64 slice
func ShuffleUInt64s(x []uint64) []uint64 {
	for i := len(x) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		x[i], x[j] = x[j], x[i]
	}
	return x
}

// MergeUInt64s merges two uint64 slice with specific excluded values
func MergeUInt64s(x []uint64, y []uint64, excludes ...uint64) []uint64 {
	traceMap := make(map[uint64]bool)
	result := make([]uint64, 0)
	for _, ex := range excludes {
		traceMap[ex] = true
	}
	// We preserve the order by x and then y
	for _, v := range x {
		if !traceMap[v] {
			traceMap[v] = true
			result = append(result, v)
		}
	}

	for _, v := range y {
		if !traceMap[v] {
			traceMap[v] = true
			result = append(result, v)
		}
	}
	return result
}
