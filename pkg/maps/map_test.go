package maps

import (
	"strings"
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

func TestEntryOf(t *testing.T) {
	i := is.New(t)

	entry1 := EntryOf("age", 25)
	i.Equal(entry1.Key, "age")
	i.Equal(entry1.Value, 25)

	entry2 := EntryOf(1, "one")
	i.Equal(entry2.Key, 1)
	i.Equal(entry2.Value, "one")

	entry3 := EntryOf(true, 3.14)
	i.Equal(entry3.Key, true)
	i.Equal(entry3.Value, 3.14)

	entry4 := EntryOf[string, any]("key", 123)
	i.Equal(entry4.Key, "key")
	i.Equal(entry4.Value, 123)

	entries := []Entry[string, int]{
		EntryOf("one", 1),
		EntryOf("two", 2),
	}
	m := FromEntries(entries)
	i.Equal(m["one"], 1)
	i.Equal(m["two"], 2)
}

func TestFromEntries(t *testing.T) {
	i := is.New(t)

	entries1 := []Entry[string, int]{
		{Key: "one", Value: 1},
		{Key: "two", Value: 2},
		{Key: "three", Value: 3},
	}
	result1 := FromEntries(entries1)
	i.Equal(len(result1), 3)
	i.Equal(result1["one"], 1)
	i.Equal(result1["two"], 2)
	i.Equal(result1["three"], 3)

	emptyEntries := []Entry[string, int]{}
	emptyResult := FromEntries(emptyEntries)
	i.Equal(len(emptyResult), 0)

	entries2 := []Entry[string, string]{
		{Key: "key", Value: "first"},
		{Key: "key", Value: "second"},
	}
	result2 := FromEntries(entries2)
	i.Equal(len(result2), 1)
	i.Equal(result2["key"], "second")

	entries3 := []Entry[int, bool]{
		{Key: 1, Value: true},
		{Key: 2, Value: false},
	}
	result3 := FromEntries(entries3)
	i.Equal(len(result3), 2)
	i.Equal(result3[1], true)
	i.Equal(result3[2], false)
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

func TestFilter(t *testing.T) {
	i := is.New(t)

	testMap1 := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
	}

	evenValues := Filter(testMap1, func(e Entry[string, int]) bool {
		return e.Value%2 == 0
	})
	i.Equal(len(evenValues), 2)
	i.Equal(evenValues["two"], 2)
	i.Equal(evenValues["four"], 4)

	longKeys := Filter(testMap1, func(e Entry[string, int]) bool {
		return len(e.Key) > 3
	})
	i.Equal(len(longKeys), 2)
	i.Equal(longKeys["three"], 3)
	i.Equal(longKeys["four"], 4)

	emptyMap := map[string]int{}
	filteredEmpty := Filter(emptyMap, func(e Entry[string, int]) bool {
		return true
	})
	i.Equal(len(filteredEmpty), 0)

	testMap2 := map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}

	withO := Filter(testMap2, func(e Entry[int, string]) bool {
		return strings.Contains(e.Value, "o")
	})
	i.Equal(len(withO), 2)
	i.Equal(withO[1], "one")
	i.Equal(withO[2], "two")

	nothing := Filter(testMap1, func(e Entry[string, int]) bool {
		return false
	})
	i.Equal(len(nothing), 0)
}

func TestJoinEntries(t *testing.T) {
	i := is.New(t)

	testMap := map[string]string{
		"name": "john",
		"age":  "30",
	}

	result1 := JoinEntries(testMap, ":")
	i.Equal(len(result1), 2)
	resultSet1 := makeSet(result1)
	i.True(resultSet1["name:john"])
	i.True(resultSet1["age:30"])

	result2 := JoinEntries(testMap, "=")
	i.Equal(len(result2), 2)
	resultSet2 := makeSet(result2)
	i.True(resultSet2["name=john"])
	i.True(resultSet2["age=30"])

	result3 := JoinEntries(testMap, "")
	i.Equal(len(result3), 2)
	resultSet3 := makeSet(result3)
	i.True(resultSet3["namejohn"])
	i.True(resultSet3["age30"])
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
