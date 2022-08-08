package gob

import (
	"bytes"
	"encoding/gob"
)

func Serialize(o any) ([]byte, error) {
	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)
	err := enc.Encode(o)
	return buff.Bytes(), err
}

func Deserialize(byts []byte, o any) error {
	reader := bytes.NewReader(byts)
	dec := gob.NewDecoder(reader)
	return dec.Decode(o)
}
