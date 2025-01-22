package encoding

import (
	"encoding/json"

	"dario.lol/gotils/pkg/gotils"
)

func UnmarshalJSON[T any](data []byte, unmarshaler ...json.Unmarshaler) (T, error) {
	var result T
	if len(unmarshaler) > 0 {
		return result, unmarshaler[0].UnmarshalJSON(data)
	}
	return result, json.Unmarshal(data, &result)
}

func MustUnmarshalJSON[T any](data []byte, unmarshaler ...json.Unmarshaler) T {
	return gotils.Must(UnmarshalJSON[T](data, unmarshaler...))
}

func MarshalJSON[T any](v T, marshaler ...json.Marshaler) ([]byte, error) {
	if len(marshaler) > 0 {
		return marshaler[0].MarshalJSON()
	}
	return json.Marshal(v)
}

func MustMarshalJSON[T any](v T, marshaler ...json.Marshaler) []byte {
	result, err := MarshalJSON(v, marshaler...)
	if err != nil {
		panic(err)
	}
	return result
}
