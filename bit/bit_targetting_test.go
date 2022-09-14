package bit

import (
	"testing"
)

func TestBitTargettingExample(t *testing.T) {

	const (
		fuelType      = "fuel_type"      // 1. All 2. Petrol, 3. Disel, 4. CNG, 5. EV
		manufacturers = "manufacturers"  // 1. All 2. Tata, 3. Mahindra, 4. Ford, 5. Maruti
		parkingSensor = "parking_sensor" // Yes / No
	)

	testsuit := []struct {
		Id            uint64           `json:"id"`
		FuelType      TargettingEntity `json:"fuel_type"`
		Manufacturers TargettingEntity `json:"manufacturers"`
		ParkingSensor bool             `json:"parking_sensor"`
	}{
		{1,
			TargettingEntity{Ids: []uint64{2, 4}},
			TargettingEntity{Exclude: 1, Ids: []uint64{3, 5}},
			true,
		},
		{2,
			TargettingEntity{Ids: []uint64{5}},
			TargettingEntity{Ids: []uint64{2, 3, 5}},
			true,
		},
		{3,
			TargettingEntity{Ids: []uint64{5}},
			TargettingEntity{Ids: []uint64{2, 3, 5}},
			false,
		},
		{4,
			TargettingEntity{Ids: []uint64{1}},
			TargettingEntity{Ids: []uint64{1}},
			true,
		},
	}

	// create the bit targetting context
	targettingContext := NewBitTargettingContext()

	for _, data := range testsuit {
		for _, id := range data.FuelType.Ids {
			targettingContext.RegisterTargettingEntity(NewCriteria(id, fuelType))
		}
		for _, id := range data.Manufacturers.Ids {
			targettingContext.RegisterTargettingEntity(NewCriteria(id, manufacturers))
		}

		if data.ParkingSensor {
			targettingContext.RegisterTargettingEntity(NewCriteria(1, parkingSensor))
		} else {
			targettingContext.RegisterTargettingEntity(NewCriteria(0, parkingSensor))
		}
	}

	targettingContext.InitBitMatrix((uint64)(len(testsuit)))

	for idx, data := range testsuit {
		if len(data.FuelType.Ids) > 0 {
			ct := NewCriteria(0, fuelType)
			if 1 == data.FuelType.Exclude {
				if excludeRowId, found := targettingContext.GetCriteriaExcludeKeyRowId(ct); found {
					targettingContext.SetBit(excludeRowId, (uint64)(idx))
				}
			}
			for _, id := range data.FuelType.Ids {
				ct.Id = id
				if criteriaRowId, found := targettingContext.GetGetCriteriaIdRowId(ct); found {
					targettingContext.SetBit(criteriaRowId, (uint64)(idx))
				}
			}
		}

		if len(data.Manufacturers.Ids) > 0 {
			ct := NewCriteria(0, manufacturers)
			if 1 == data.Manufacturers.Exclude {
				if excludeRowId, found := targettingContext.GetCriteriaExcludeKeyRowId(ct); found {
					targettingContext.SetBit(excludeRowId, (uint64)(idx))
				}
			}
			for _, id := range data.Manufacturers.Ids {
				ct.Id = id
				if criteriaRowId, found := targettingContext.GetGetCriteriaIdRowId(ct); found {
					targettingContext.SetBit(criteriaRowId, (uint64)(idx))
				}
			}
		}

		if data.ParkingSensor {
			ct := NewCriteria(1, parkingSensor)
			if criteriaRowId, found := targettingContext.GetGetCriteriaIdRowId(ct); found {
				targettingContext.SetBit(criteriaRowId, (uint64)(idx))
			}
		} else {
			ct := NewCriteria(0, parkingSensor)
			if criteriaRowId, found := targettingContext.GetGetCriteriaIdRowId(ct); found {
				targettingContext.SetBit(criteriaRowId, (uint64)(idx))
			}
		}
	}

	// targettingContext.PrintTargettingRows()
	// targettingContext.PrintBitMatrix()
	targettingRows := make([]*TargettingRow, 0)
	targettingRows = append(targettingRows, targettingContext.GetTargettingRow(NewCriteria(2, fuelType)))
	targettingRows = append(targettingRows, targettingContext.GetTargettingRow(NewCriteria(2, manufacturers)))
	targettingRows = append(targettingRows, targettingContext.GetTargettingRow(NewCriteria(1, parkingSensor)))

	resultBitArray, err := targettingContext.EvaluateTargetting(targettingRows, []uint64{})
	if nil == err {
		t.Log(resultBitArray.Print())
	}
}
