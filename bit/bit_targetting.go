package bit

import (
	"fmt"
	"log"
	"sort"
)

type TargettingEntity struct {
	Exclude uint8
	Ids     []uint64
}

type criteria struct {
	Id       uint64
	Criteria string
}

type TargettingRow struct {
	ExcludeRowFound  bool
	ExcludeRowId     uint64
	CriteriaRowFound bool
	CriteriaRowId    uint64
}

type BitTargettingContext struct {
	rowCount    uint64
	columnCount uint64
	rowIdMap    map[string]uint64
	bitMatrix   *BitMatrix
}

func NewCriteria(id uint64, name string) *criteria {
	return &criteria{
		Id:       id,
		Criteria: name,
	}
}

func NewBitTargettingContext() *BitTargettingContext {
	return &BitTargettingContext{
		rowCount: 0,
		rowIdMap: make(map[string]uint64),
	}
}

func (ctx *BitTargettingContext) registerRow(key string) {
	if _, found := ctx.rowIdMap[key]; !found {
		ctx.rowIdMap[key] = ctx.rowCount
		ctx.rowCount++
	}
}

func (c *criteria) getExcludeKeyName() string {
	return fmt.Sprintf("%s_exclude", c.Criteria)
}

func (c *criteria) getCriteriaIdKeyName() string {
	return fmt.Sprintf("%s_%d", c.Criteria, c.Id)
}

// create new row id if not already created for given entityType and entityId
func (ctx *BitTargettingContext) RegisterTargettingEntity(c *criteria) {
	ctx.registerRow(c.getExcludeKeyName())
	ctx.registerRow(c.getCriteriaIdKeyName())
}

func (ctx *BitTargettingContext) InitBitMatrix(entityCount uint64) error {
	var err error
	ctx.columnCount = entityCount
	ctx.bitMatrix, err = NewBitMatrix(ctx.rowCount, ctx.columnCount)
	return err
}

func (ctx *BitTargettingContext) GetCriteriaExcludeKeyRowId(c *criteria) (uint64, bool) {
	rowId, found := ctx.rowIdMap[c.getExcludeKeyName()]
	return rowId, found
}

func (ctx *BitTargettingContext) GetGetCriteriaIdRowId(c *criteria) (uint64, bool) {
	rowId, found := ctx.rowIdMap[c.getCriteriaIdKeyName()]
	return rowId, found
}

func (ctx *BitTargettingContext) SetBit(rowId, columnId uint64) bool {
	return ctx.bitMatrix.SetBit(rowId, columnId, true)
}

func (ctx *BitTargettingContext) PrintTargettingRows() {
	keys := make([]string, 0, len(ctx.rowIdMap))
	for key := range ctx.rowIdMap {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return ctx.rowIdMap[keys[i]] < ctx.rowIdMap[keys[j]]
	})

	for _, k := range keys {
		log.Printf("%25s: %d", k, ctx.rowIdMap[k])
	}
}

func (ctx *BitTargettingContext) PrintBitMatrix() {
	if nil != ctx.bitMatrix {
		ctx.bitMatrix.Print()
	}
}

func (ctx *BitTargettingContext) GetTargettingRow(c *criteria) *TargettingRow {
	targettingRow := &TargettingRow{}
	targettingRow.ExcludeRowId, targettingRow.ExcludeRowFound =
		ctx.GetCriteriaExcludeKeyRowId(c)
	targettingRow.CriteriaRowId, targettingRow.CriteriaRowFound =
		ctx.GetGetCriteriaIdRowId(c)
	return targettingRow
}

func (ctx *BitTargettingContext) EvaluateTargetting(rowPairs []*TargettingRow, columnIdxes []uint64) (resultBitArray *BitArray, err error) {
	//TODO: use columnIdxes
	resultBitArray, err = ctx.bitMatrix.CreateNewRow()
	if nil == err {
		resultBitArray.SetAll()
	}

	for _, rowPair := range rowPairs {
		if !rowPair.ExcludeRowFound {
			continue
		}

		excludeRow := ctx.bitMatrix.GetRow(rowPair.ExcludeRowId)
		// first perform bitwise XOR operation between critera row and exclusion row
		// then perform bitwise AND operation with resultBitArray
		if rowPair.CriteriaRowFound {
			criteriaRow := ctx.bitMatrix.GetRow(rowPair.CriteriaRowId)
			resultBitArray = resultBitArray.ANDWITHXOR(excludeRow, criteriaRow)
			// resultBitArray = resultBitArray.ANDWITHXORColumnIndexes(excludeRow, criteriaRow, columnIdxes)
		} else { // perform bitwise AND operation with exclusion row and resultBitArray
			resultBitArray = resultBitArray.AND(excludeRow)
			// resultBitArray = resultBitArray.ANDColumnIndexes(excludeRow, columnIdxes)
		}

		log.Println(resultBitArray.Print())
	}

	return resultBitArray, err
}
