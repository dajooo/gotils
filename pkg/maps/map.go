package maps

type Entry[K comparable, V any] struct {
	Key   K
	Value V
}

func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func Values[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

func Entries[K comparable, V any](m map[K]V) []Entry[K, V] {
	entries := make([]Entry[K, V], 0, len(m))
	for k, v := range m {
		entries = append(entries, Entry[K, V]{Key: k, Value: v})
	}
	return entries
}

func MapEntries[K comparable, V any, O any](m map[K]V, f func(Entry[K, V]) O) []O {
	result := make([]O, 0, len(m))
	for k, v := range m {
		result = append(result, f(Entry[K, V]{Key: k, Value: v}))
	}
	return result
}

func Filter[K comparable, V any](m map[K]V, f func(Entry[K, V]) bool) map[K]V {
	result := make(map[K]V, len(m))
	for k, v := range m {
		if f(Entry[K, V]{Key: k, Value: v}) {
			result[k] = v
		}
	}
	return result
}

func JoinEntries(m map[string]string, delimiter string) []string {
	return MapEntries(m, func(e Entry[string, string]) string {
		return e.Key + delimiter + e.Value
	})
}
