// Contains efficient array that stores small-range values (up to uint64) in a bit array to optimize storage.
package bitarray

import (
	"errors"
	"fmt"
	"math"
	"unsafe"
)

// The biggest Item size that can fit in CompactArray.
type Item = uint64

// Provides a wrapper on top of BitSet to store items of arbitrary size (up to 64 bits each) efficiently.
type CompactArray struct {
	*BitSet    `json:",inline"`
	ItemSize   uint `json:"item,omitempty"`
	ItemsCount uint `json:"count,omitempty"`
}

func bitsNeeded(maxValue Item) uint {
	return uint(math.Ceil(math.Log2(float64(maxValue + 1))))
}

func (a *CompactArray) validateIndex(index int) {
	if index < 0 || uint(index) >= a.ItemsCount {
		panic(fmt.Sprintf("Attempt to access index [%v] in CompactArray of size [%v]", index, a.ItemsCount))
	}
}

func (a *CompactArray) validateValue(value Item) {
	maxValue := (uint64(1) << a.ItemSize) - 1
	if value > maxValue {
		panic(fmt.Sprintf("Value [%v] is too big. Max value is [%v].", value, maxValue))
	}
}

// Sets item at index to provided value.
func (a *CompactArray) SetItem(index int, value Item) {
	a.validateIndex(index)
	a.validateValue(value)
	bitIndex := uint(index) * a.ItemSize // #nosec G115
	x := Item(1)
	// #nosec G115
	for i := int(a.ItemSize - 1); i >= 0; i-- {
		if x&value != 0 {
			a.BitSet.Set(bitIndex + uint(i)) // #nosec G115

		} else {
			a.BitSet.Clear(bitIndex + uint(i)) // #nosec G115

		}

		x <<= 1
	}
}

// Gets Item at provided index.
func (a *CompactArray) GetItem(index int) Item {
	a.validateIndex(index)
	bitIndex := uint(index) * a.ItemSize // #nosec G115
	res := Item(0)
	x := Item(1)
	// #nosec G115
	for i := int(a.ItemSize - 1); i >= 0; i-- {
		// #nosec G115
		if a.BitSet.IsSet(bitIndex + uint(i)) {
			res |= x
		}

		x <<= 1
	}

	return res
}

// Gets all items stored in the array. The size of the returned array matches the ItemsCount it was initialized with.
func (a CompactArray) GetItems() []Item {
	res := make([]Item, 0, a.ItemsCount)
	// #nosec G115
	for i := 0; i < int(a.ItemsCount); i++ {
		res = append(res, a.GetItem(i)) // #nosec G115
	}

	return res
}

// Gets a string representation of the array contents. This is a relatively expensive operation.
func (a CompactArray) String() string {
	res := "["
	for _, item := range a.GetItems() {
		res += fmt.Sprintf("%v, ", item)
	}

	res += "]"
	return res
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (a *CompactArray) DeepCopyInto(out *CompactArray) {
	*out = *a
	if a.BitSet != nil {
		in, out := &a.BitSet, &out.BitSet
		*out = new(BitSet)
		if **in != nil {
			in, out := *in, *out
			*out = make([]Block, len(*in))
			copy(*out, *in)
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompactArray.
func (a *CompactArray) DeepCopy() *CompactArray {
	if a == nil {
		return nil
	}
	out := new(CompactArray)
	a.DeepCopyInto(out)
	return out
}

// Creates a new CompactArray of specified size.
func NewCompactArray(size uint, maxValue Item) (CompactArray, error) {
	itemSize := bitsNeeded(maxValue)
	/* #nosec */
	if uint64(itemSize) >= uint64(unsafe.Sizeof(Item(0))*8) {
		return CompactArray{}, errors.New("invalid itemSize, must be less than the size of Item")
	}

	return CompactArray{
		BitSet:     NewBitSet(size),
		ItemsCount: size,
		ItemSize:   itemSize,
	}, nil
}
