package trie

// data types supported by the trie
type trieType interface{ rune | byte | string }

// trie data structure
type Trie[T trieType] struct {
	Value      interface{}
	ChildNodes map[T]*Trie[T]
}

// create new object of trie
func NewTrie[T trieType]() *Trie[T] {
	return &Trie[T]{}
}

// insert the object into trie
func (trie *Trie[T]) Put(keys []T, value interface{}) interface{} {
	if nil == trie {
		return nil
	}

	node := trie
	for _, key := range keys {
		var child *Trie[T]
		child = node.ChildNodes[key]
		if nil == child {
			if nil == node.ChildNodes {
				node.ChildNodes = map[T]*Trie[T]{}
			}
			child = new(Trie[T])
			node.ChildNodes[key] = child
		}
		node = child
	}

	oldVal := node.Value
	node.Value = value
	return oldVal
}

// get the object from trie
func (trie *Trie[T]) Get(keys []T) interface{} {
	if nil == trie {
		return nil
	}

	node := trie
	for _, key := range keys {
		node = node.ChildNodes[key]
		if nil == node { //expected node for the key
			return nil
		}
	}

	return node.Value
}

// walk the trie nodes. useful to print the trie
func (trie *Trie[T]) Walk(keys []T, actOn func(key []T, value interface{}) error) error {
	if nil != trie.Value {
		if err := actOn(keys, trie.Value); err != nil {
			return err
		}
	}

	for index, child := range trie.ChildNodes {
		keyToSend := append(keys, T(index))
		if err := child.Walk(keyToSend, actOn); nil != err {
			return err
		}
	}

	return nil
}
