package cuckoo

import (
	"github.com/linvon/cuckoo-filter"
)

type CuckooFilter struct {
	filter *cuckoo.Filter
}

func NewCuckooFilter(maxNumKeys uint) *CuckooFilter {
	return &CuckooFilter{
		filter: cuckoo.NewFilter(4, 24, maxNumKeys, cuckoo.TableTypePacked),
	}
}

func (c *CuckooFilter) Insert(key string) bool {
	return c.filter.Add([]byte(key))
}

func (c *CuckooFilter) Remove(key string) bool {
	return c.filter.Delete([]byte(key))
}

func (c *CuckooFilter) Contain(key string) bool {
	return c.filter.Contain([]byte(key))
}

func (c *CuckooFilter) Size() uint {
	return c.filter.Size()
}

func (c *CuckooFilter) SizeInBytes() uint {
	return c.filter.SizeInBytes()
}

func (c *CuckooFilter) Reset() {
	c.filter.Reset()
}

func (c *CuckooFilter) Encode() ([]byte, error) {
	return c.filter.Encode()
}

func Decode(b []byte) (c *CuckooFilter, err error) {
	var filter *cuckoo.Filter
	if filter, err = cuckoo.Decode(b); nil != err {
		return
	}

	c = &CuckooFilter{
		filter: filter,
	}

	return
}
