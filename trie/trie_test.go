package trie

import (
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
