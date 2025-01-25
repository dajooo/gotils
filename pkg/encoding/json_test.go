package encoding

import (
	"testing"

	"github.com/matryer/is"
)

type testStruct struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestUnmarshalJSONWithValidInput(t *testing.T) {
	i := is.New(t)
	data := []byte(`{"name":"test","age":25}`)

	result, err := UnmarshalJSON[testStruct](data)
	i.NoErr(err)
	i.Equal(result.Name, "test")
	i.Equal(result.Age, 25)
}

func TestMarshalJSONWithValidStruct(t *testing.T) {
	i := is.New(t)
	data := testStruct{Name: "test", Age: 25}

	result, err := MarshalJSON(data)
	i.NoErr(err)
	i.Equal(string(result), `{"name":"test","age":25}`)
}

func TestMustUnmarshalJSONWithInvalidInput(t *testing.T) {
	i := is.New(t)

	defer func() {
		r := recover()
		i.True(r != nil)
	}()

	MustUnmarshalJSON[testStruct]([]byte(`invalid json`))
}

func TestMustMarshalJSONWithInvalidValue(t *testing.T) {
	i := is.New(t)

	defer func() {
		r := recover()
		i.True(r != nil)
	}()

	ch := make(chan int)
	MustMarshalJSON(ch)
}
