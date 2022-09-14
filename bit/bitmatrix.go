package bit

import (
	"bytes"
	"fmt"
	"log"
	"math"
)

const (
	maxBitmapSize = 0x01 << 32
	power         = 6                // 64 = 2^6. size of used data type for creating bit array.
	maskBits      = (1 << power) - 1 // SET ALL BITS
)

// represents the bitmap data structure
type BitArray struct {
	data []uint64
	size uint64
}

// return the BitArray of given size if size is less than maxBitmapSize
// else return the error
func NewBitArray(size uint64) (*BitArray, error) {
	var b *BitArray = nil
	var err error = nil

	for {
		if 0 == size {
			err = fmt.Errorf("Invalid size. Size should be greater than 0")
			break
		}

		if size > maxBitmapSize {
			err = fmt.Errorf("Invalid size. Size should be less than %d", maxBitmapSize)
			break
		}

		b = &BitArray{}
		b.data = make([]uint64, ((size-1)>>power)+1)
		b.size = size
		break
	}

	return b, err
}

// return the size of BitArray
func (b *BitArray) GetBitArraySize() uint64 {
	return b.size
}

// reset all bit to 0
func (b *BitArray) ResetAll() {
	for index := range b.data {
		b.data[index] = 0
	}
}

// set all bit to 1
func (b *BitArray) SetAll() {
	for index := range b.data {
		b.data[index] = uint64(math.MaxUint64)
	}
}

//return the bit value of given bit offset/position
func (b *BitArray) GetBit(offset uint64) bool {
	// safe check
	if b.size <= offset {
		return false
	}

	index, pos := offset>>power, offset&maskBits
	if (b.data[index])&(0x01<<pos) == 0 {
		return false
	}
	return true
}

// set the given bit value to provided bit offset/position
func (b *BitArray) SetBit(offset uint64, value bool) bool {
	// safe check
	if b.size <= offset {
		return false
	}

	index, pos := offset>>power, offset&maskBits
	if value == false {
		b.data[index] &^= 0x01 << pos
	} else {
		b.data[index] |= 0x01 << pos
	}

	return true
}

// perform the bitwise AND operation
func (b *BitArray) AND(t *BitArray) *BitArray {
	for index := range b.data {
		b.data[index] &= t.data[index]
	}

	return b
}

// perform the bitwise AND & XOR operation
func (b *BitArray) ANDWITHXOR(t1, t2 *BitArray) *BitArray {
	for index := range b.data {
		b.data[index] &= (t1.data[index] ^ t2.data[index])
	}

	return b
}

// perform the bitwise AND operation on given column indexes
func (b *BitArray) ANDColumnIndexes(t *BitArray, columnIdxes []uint64) *BitArray {
	for _, index := range columnIdxes {
		b.data[index] &= t.data[index]
	}
	return b
}

// perform the bitwise XOR operation on given column indexes
func (b *BitArray) ANDWITHXORColumnIndexes(t1, t2 *BitArray, columnIdxes []uint64) *BitArray {
	for _, index := range columnIdxes {
		b.data[index] &= (t1.data[index] ^ t2.data[index])
	}
	return b
}

// perform the bitwise OR operation
func (b *BitArray) OR(t *BitArray) *BitArray {
	for index := range b.data {
		b.data[index] |= t.data[index]
	}
	return b
}

// perform the bitwise XOR operation
func (b *BitArray) XOR() *BitArray {
	for index := range b.data {
		b.data[index] = ^b.data[index]
	}
	return b
}

// print the BitArray in binary format
func (b *BitArray) Print() string {
	var buff bytes.Buffer
	for index := len(b.data) - 1; index >= 0; index-- {
		buff.WriteString(StringReverse(fmt.Sprintf("%064b ", b.data[index])))
	}

	return buff.String()
}

type BitMatrix struct {
	bitArray []*BitArray
	rows     uint64
	columns  uint64
}

// return the BitMatrix by allocating all required memoory
// it's kind of representation of two dimensional
/*
   0 -> [array of uint64]
   1 -> [array of uint64]
   2 -> [array of uint64]
   3 -> [array of uint64]
   4 -> [array of uint64]
*/
func NewBitMatrix(rows, columns uint64) (*BitMatrix, error) {
	var err error = nil
	bitMatix := &BitMatrix{
		bitArray: make([]*BitArray, rows),
		rows:     rows,
		columns:  columns,
	}

	var index uint64 = 0
	var bits *BitArray = nil
	for ; index < rows; index++ {
		bits, err = NewBitArray(columns)
		if nil != err {
			break
		}

		bitMatix.bitArray[index] = bits
	}

	return bitMatix, err
}

// set the bit of given row and column position
func (matrix *BitMatrix) SetBit(rowId, colId uint64, value bool) bool {
	if rowId >= matrix.rows {
		return false
	}

	return matrix.bitArray[rowId].SetBit(colId, value)
}

// set all bit to 1 for given row position
func (matrix *BitMatrix) SetAllBitsForRow(rowId uint64, value bool) {
	matrix.bitArray[rowId].SetAll()
}

// return the BitArray of given row position
func (matrix *BitMatrix) GetRow(rowId uint64) *BitArray {
	return matrix.bitArray[rowId]
}

// create new row and return it
func (matrix *BitMatrix) CreateNewRow() (*BitArray, error) {
	return NewBitArray(matrix.columns)
}

func (matrix *BitMatrix) GetColumnIndexes(bitIndexes []uint64) (columnIdxes []uint64) {
	indexMap := make(map[uint64]bool)
	for _, bitIdx := range bitIndexes {
		idx := bitIdx >> power
		if _, found := indexMap[idx]; !found {
			columnIdxes = append(columnIdxes, idx)
			indexMap[idx] = true
		}
	}

	return columnIdxes
}

func (matrix *BitMatrix) Print() {
	var index uint64 = 0
	for ; index < matrix.rows; index++ {
		log.Printf("%6d:%s\n", index, matrix.bitArray[index].Print())
	}
}
