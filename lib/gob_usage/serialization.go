package gob_usage

import (
	"bytes"
	"encoding/gob"
	"log"
)

func Serialize(data interface{}) []byte {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	if err := e.Encode(data); err != nil {
		panic(err)
	}
	return b.Bytes()
}

func DeSerialize(serialized []byte, dest interface{}) error {
	r := bytes.NewReader(serialized)
	dec := gob.NewDecoder(r)
	err := dec.Decode(dest)
	if err != nil {
		log.Fatalln(err)
	}
	return nil
}
