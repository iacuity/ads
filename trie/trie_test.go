package trie

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"runtime"
	"strings"
	"testing"
)

var (
	stringTestSuit = []struct {
		key string
		val interface{}
	}{
		{"fish", 0},
		{"cat", 1},
		{"dog", 2},
		{"cats", 3},
		{"caterpillar", 4},
		{"cat/t1", 5},
		{"cat/t2", 6},
		{"", 7},
		{"Camel", 8},
		{"Caterpillar", 9},
		{"Crab", 10},
		{"Cheetah", 11},
	}

	nonExistsKeys = []string{
		"fishing",
		"rat",
		"rabbit",
		"Fox",
		"Cheetahh",
	}
)

func TestNewStringTrie(t *testing.T) {
	trie := NewTrie[string]()

	//initialise the trie
	for _, c := range stringTestSuit {
		if isNew := trie.Put(strings.Split(c.key, "/"), c.val); isNew != nil {
			t.Logf("%s is already exist.", c.key)
		}
	}

	//verify the trie
	for _, c := range stringTestSuit {
		if value := trie.Get(strings.Split(c.key, "/")); value != c.val {
			t.Fatalf("Invalid value for key '%s'. Expected %v and found %v", c.key, c.val, value)
		}
	}

	for _, key := range nonExistsKeys {
		if value := trie.Get(strings.Split(key, "/")); nil != value {
			t.Fatalf("Found key: '%s', but it not expected to be present", key)
		}
	}
	/*
		actor := func(key []string, value interface{}) error {
			t.Logf("%v:%v\n", key, value)
			return nil
		}

		trie.Walk([]string{}, actor)
	*/
}

func TestNewByteTrie(t *testing.T) {
	trie := NewTrie[byte]()

	//initialise the trie
	for _, c := range stringTestSuit {
		if isNew := trie.Put([]byte(c.key), c.val); isNew != nil {
			t.Logf("%s is already exist.", c.key)
		}
	}

	//verify the trie
	for _, c := range stringTestSuit {
		if value := trie.Get([]byte(c.key)); value != c.val {
			t.Fatalf("Invalid value for key '%s'. Expected %v and found %v", c.key, c.val, value)
		}
	}

	for _, key := range nonExistsKeys {
		if value := trie.Get([]byte(key)); nil != value {
			t.Fatalf("Found key: '%s', but it not expected to be present", key)
		}
	}
	/*
		actor := func(key []byte, value interface{}) error {
			t.Logf("%s:%v\n", string(key), value)
			return nil
		}

		trie.Walk([]byte{}, actor)
	*/
}

func printMemoryUsage(t *testing.T, m1, m2 *runtime.MemStats) {
	t.Log(
		"Alloc:", (m2.Alloc-m1.Alloc)/1000000,
		"TotalAlloc:", (m2.TotalAlloc-m1.TotalAlloc)/1000000,
		"HeapAlloc:", (m2.HeapAlloc-m1.HeapAlloc)/1000000,
	)
}

func getMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func TestTrieMemorySize(t *testing.T) {
	var m1, m2 runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&m1)
	//	byteTrie := NewTrie[byte]()
	myMap := make(map[string]int)
	for i := 0; i < 10,000,000; i++ {
		//	byteTrie.Put([]byte(getMD5Hash(fmt.Sprintf("dc%de9-2fda-%d-bd2c-%da8284b9c4%d", i, i, i, i))), i)
		myMap[getMD5Hash(fmt.Sprintf("dc%de9-2fda-%d-bd2c-%da8284b9c4%d", i, i, i, i))] = i
	}

	runtime.ReadMemStats(&m2)
	printMemoryUsage(t, &m1, &m2)

	for i := 0; i < 1000000; i++ {
		if val, foud := myMap[getMD5Hash(fmt.Sprintf("dc%de9-2fda-%d-bd2c-%da8284b9c4%d", i, i, i, i))]; !foud || i != val {
			t.Fatalf("Invalid Value")
		}
	}
}
