package cuckoo

import (
	"strconv"
	"testing"
)

func TestCuckooFilter(t *testing.T) {
	maxKeys := 10000000
	cuckooFilter := NewCuckooFilter(uint(maxKeys))
	for i := 0; i < maxKeys; i++ {
		if sucess := cuckooFilter.Insert(strconv.Itoa(i)); !sucess {
			t.Fatalf("Failed to insert key at index: %d\n", i)
		}
	}

	t.Log("cuckooFilter.Size(): ", cuckooFilter.Size())
	t.Log("cuckooFilter.SizeInBytes(): ", cuckooFilter.SizeInBytes()/1000)

	for i := 0; i < maxKeys; i++ {
		if found := cuckooFilter.Contain(strconv.Itoa(i)); !found {
			t.Fatalf("Expected key to present at index: %d\n", i)
		}
	}

	for i := maxKeys; i < maxKeys+10000; i++ {
		if found := cuckooFilter.Contain(strconv.Itoa(i)); found {
			t.Logf("Expected key to not present at index: %d\n", i)
		}
	}
}
