package maps

import (
	"testing"

	"github.com/matryer/is"
)

func TestKeys(t *testing.T) {
	i := is.New(t)

	emptyMap := map[string]int{}
	emptyKeys := Keys(emptyMap)
	i.Equal(len(emptyKeys), 0)

	testMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	keys := Keys(testMap)
	i.Equal(len(keys), 3)
	i.True(contains(keys, "one"))
	i.True(contains(keys, "two"))
	i.True(contains(keys, "three"))
}

func TestValues(t *testing.T) {
	i := is.New(t)

	emptyMap := map[string]int{}
	emptyValues := Values(emptyMap)
	i.Equal(len(emptyValues), 0)

	testMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	values := Values(testMap)
	i.Equal(len(values), 3)
	i.True(contains(values, 1))
	i.True(contains(values, 2))
	i.True(contains(values, 3))
}

func TestEntries(t *testing.T) {
	i := is.New(t)

	emptyMap := map[string]int{}
	emptyEntries := Entries(emptyMap)
	i.Equal(len(emptyEntries), 0)

	testMap := map[string]int{
		"one": 1,
	}
	entries := Entries(testMap)
	i.Equal(len(entries), 1)
	i.Equal(entries[0].Key, "one")
	i.Equal(entries[0].Value, 1)
}

func TestMapEntries(t *testing.T) {
	i := is.New(t)

	testMap := map[string]int{
		"one": 1,
		"two": 2,
	}

	result := MapEntries(testMap, func(e Entry[string, int]) string {
		return e.Key + ":" + string(rune(e.Value+'0'))
	})

	i.Equal(len(result), 2)
	resultSet := makeSet(result)
	i.True(resultSet["one:1"])
	i.True(resultSet["two:2"])
}

func makeSet[T comparable](slice []T) map[T]bool {
	set := make(map[T]bool)
	for _, item := range slice {
		set[item] = true
	}
	return set
}

func contains[T comparable](slice []T, item T) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
