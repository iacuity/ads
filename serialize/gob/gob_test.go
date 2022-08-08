package gob

import (
	"testing"

	"github.com/iacuity/ads/trie"
)

func TestSerializeTrie(t *testing.T) {
	stringTestSuit := []struct {
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

	nonExistsKeys := []string{
		"fishing",
		"rat",
		"rabbit",
		"Fox",
		"Cheetahh",
	}

	eTrie := trie.NewTrie[byte]()
	//initialise the trie
	for _, c := range stringTestSuit {
		if isNew := eTrie.Put([]byte(c.key), c.val); isNew != nil {
			t.Logf("%s is already exist.", c.key)
		}
	}

	byts, err := Serialize(eTrie)
	if nil != err {
		t.Fatalf("Error in serialize trie: %s\n", err.Error())
	}

	dTrie := trie.NewTrie[byte]()
	Deserialize(byts, dTrie)

	//verify the trie
	for _, c := range stringTestSuit {
		if value := dTrie.Get([]byte(c.key)); value != c.val {
			t.Fatalf("Invalid value for key '%s'. Expected %v and found %v", c.key, c.val, value)
		}
	}

	for _, key := range nonExistsKeys {
		if value := dTrie.Get([]byte(key)); nil != value {
			t.Fatalf("Found key: '%s', but it not expected to be present", key)
		}
	}
}
