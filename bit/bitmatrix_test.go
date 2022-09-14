package bit

import (
	"testing"
)

func TestNewBitArray(t *testing.T) {
	var size uint64 = 0
	if _, err := NewBitArray(size); nil == err {
		t.Fatal("Failed to create BitArray of size: ", size)
	}

	// checking for one element
	size = 0x01 << 0
	if bitArray, err := NewBitArray(size); nil != err {
		t.Fatal("Failed to create BitArray of size: ", size)
	} else if 1 != len(bitArray.data) {
		t.Fatal("Invalid BitArray. Expected to be length 1, found with len: ", len(bitArray.data))
	}

	// checking for boundry condition for one size
	size = 63
	if bitArray, err := NewBitArray(size); nil != err {
		t.Fatal("Failed to creat bitmap with size: ", size)
	} else if 1 != len(bitArray.data) {
		t.Fatal("Invalid BitArray. Expected to be length 1, found with len: ", len(bitArray.data))
	}

	size = 0x01 << 6
	if bitArray, err := NewBitArray(size); nil != err {
		t.Fatal("Failed to create BitArray of size: ", size)
	} else if 1 != len(bitArray.data) {
		t.Fatal("Invalid BitArray. Expected to be length 1, found with len: ", len(bitArray.data))
	}

	size = 0x01 << 7
	if bitArray, err := NewBitArray(size); nil != err {
		t.Fatal("Failed to create BitArray of size: ", size)
	} else if 2 != len(bitArray.data) {
		t.Fatal("Invalid BitArray. Expected to be length 2, found with len: ", len(bitArray.data))
	}

	size = 129
	if bitArray, err := NewBitArray(size); nil != err {
		t.Fatal("Failed to create BitArray of size: ", size)
	} else {
		if 3 != len(bitArray.data) {
			t.Fatal("Invalid BitArray. Expected to be size 3, found with size: ", len(bitArray.data))
		}
		if size != bitArray.GetBitArraySize() {
			t.Fatalf("Invalid BitArray. Expected to be length %d, found with len: %d", size, len(bitArray.data))
		}
	}

	// -ve test case exceeding limit of BitArray size
	size = 0x01 << 33
	if _, err := NewBitArray(size); nil == err {
		t.Fatal("Failed to create BitArray of size: ", size)
	}
}

func TestNewBitArraySetBitAndGetBit(t *testing.T) {
	var size uint64 = 0x01 << 7
	bitArray, err := NewBitArray(size)
	if nil != err {
		t.Fatal("Failed to creat bitmap with size: ", size)
	}

	var setBits []uint64 = []uint64{0, 1, 3, 7, 15, 25, 28, 31, 32, 63, 64, 100, 127}
	var resetBits []uint64 = []uint64{2, 4, 6, 8, 20, 34, 68, 99, 126}

	// initially all bits are unset
	for _, val := range setBits {
		if bitArray.GetBit(val) {
			t.Fatal("Expected bit value as false, but found true")
		}

		// setting the bit
		bitArray.SetBit(val, true)
	}
	for _, val := range resetBits {
		if bitArray.GetBit(val) {
			t.Fatal("Expected bit value as false, but found true")
		}

		// resetting the bit
		bitArray.SetBit(val, false)
	}

	for _, val := range setBits {
		if !bitArray.GetBit(val) {
			t.Fatal("Expected bit value as true, but found false")
		}

		// resetting the bit
		bitArray.SetBit(val, false)
	}
	for _, val := range resetBits {
		if bitArray.GetBit(val) {
			t.Fatal("Expected bit value as false, but found true")
		}

		// setting the bit
		bitArray.SetBit(val, true)
	}

	for _, val := range setBits {
		if bitArray.GetBit(val) {
			t.Fatal("Expected bit value as false, but found true")
		}
	}
	for _, val := range resetBits {
		if !bitArray.GetBit(val) {
			t.Fatal("Expected bit value as true, but found false")
		}
	}
}

func TestNewBitArraySetAndResetAll(t *testing.T) {
	var size uint64 = 0x01 << 7 // create bit array of size 127 bits
	bitArray, err := NewBitArray(size)
	if nil != err {
		t.Fatal("Failed to create BitArray of size: ", size)
	}

	// by default all bits are reset
	var index []uint64 = []uint64{0, 1, 3, 7, 15, 25, 28, 31, 32, 63, 64, 100, 127}
	for _, val := range index {
		if bitArray.GetBit(val) {
			t.Fatal("Expected bit value as false, but found true")
		}
	}

	// set all bits
	bitArray.SetAll()
	for _, val := range index {
		if !bitArray.GetBit(val) {
			t.Fatal("Expected bit value as true, but found false")
		}
	}

	//reset all bits
	bitArray.ResetAll()
	for _, val := range index {
		if bitArray.GetBit(val) {
			t.Fatal("Expected bit value as false, but found true")
		}
	}
}
