package msgpack_usage

import "github.com/vmihailenco/msgpack/v5"

func Marshall(data interface{}) []byte {
	b, err := msgpack.Marshal(data)
	if err != nil {
		panic(err)
	}
	return b
}

func Unmarshall(marshalled []byte) interface{} {
	var i interface{}
	err := msgpack.Unmarshal(marshalled, &i)
	if err != nil {
		panic(err)
	}
	return i
}
