package gob_usage

import (
	"bytes"
	"encoding/gob"
	"log"
)

func RegisterType(t interface{}) {
	gob.Register(t)
}

func InterfaceDeserialise(serialized []byte, dest interface{}) error {
	r := bytes.NewReader(serialized)
	dec := gob.NewDecoder(r)
	err := dec.Decode(dest)
	if err != nil {
		log.Fatalln(err)
	}
	return nil
}
