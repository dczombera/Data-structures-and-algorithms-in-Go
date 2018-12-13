package trie

import (
	"testing"
)

type testCase struct {
	data []KV
	want []KV
}

type KV struct {
	key string
	val int
}

var testCases = []testCase{
	{[]KV{{"han", 1}, {"luke", 2}, {"r2d2", 3}}, []KV{{"han", 1}, {"luke", 2}, {"r2d2", 3}}},
	{[]KV{{"foo", 42}, {"bar", 0}, {"bar", 42}, {"foo", 0}}, []KV{{"bar", 42}, {"foo", 0}}},
	{[]KV{{"han", -42}, {"shot", 42}, {"first", 0}, {"second", 0}, {"jumped", 42}}, []KV{{"second", 0}, {"han", -42}, {"jumped", 42}, {"shot", 42}, {"first", 0}}},
}

var errorTestCases = []testCase{
	{[]KV{{"hola", 1}, {"hello", 2}}, []KV{{"bye", -1}, {"adios", -1}}},
}

func TestTrie(t *testing.T) {
	for _, tc := range testCases {
		trie := Constructor()
		for _, d := range tc.data {
			trie.Put(d.key, d.val)
		}

		for _, w := range tc.want {
			v, err := trie.Get(w.key)
			if err != nil {
				t.Errorf("key %v with value %v not found in trie", w.key, w.val)
			}
			if w.val != v {
				t.Errorf("Got %v value for key %v, want value %v", v, w.key, w.val)
			}
		}
	}
}

func TestTrieError(t *testing.T) {
	for _, tc := range errorTestCases {
		trie := Constructor()
		for _, d := range tc.data {
			trie.Put(d.key, d.val)
		}

		for _, w := range tc.want {
			v, err := trie.Get(w.key)
			if err == nil {
				t.Errorf("key %v with value %v found in trie, don't want it to exist", w.key, v)
			}
		}
	}
}

func TestCollecteys(t *testing.T) {
	type Key struct {
		pre  string
		keys []string
	}
	testCases := []struct {
		data []KV
		want []Key
	}{
		{
			data: []KV{{"she", 1}, {"sea", 1}, {"see", 1}, {"shell", 1}, {"shell", 1}, {"shellsort", 1}, {"hello", 1}, {"he", 1}},
			want: []Key{{"she", []string{"she", "shell", "shellsort"}}, {"sea", []string{"sea"}}, {"shell", []string{"shell", "shellsort"}}, {"he", []string{"he", "hello"}}, {"hello", []string{"hello"}}, {"s", []string{"sea", "see", "she", "shell", "shellsort"}}},
		},
		{
			data: []KV{{"Han", 1}, {"shot", 1}, {"first", 1}, {",", 1}, {"he", 1}, {"said", 1}, {".", 1}},
			want: []Key{{"", []string{",", ".", "Han", "first", "he", "said", "shot"}}},
		},
	}

	for _, tc := range testCases {
		trie := Constructor()
		for _, d := range tc.data {
			trie.Put(d.key, d.val)
		}

		for _, w := range tc.want {
			keys := trie.KeysWithPrefix(w.pre)
			for i, k := range w.keys {
				if keys[i] != k {
					t.Errorf("Got key %v in collection of keys with prefix %v, want %v", keys[i], w.pre, k)
				}
			}
		}
	}
}

func TestKeysThatMatch(t *testing.T) {
	type Key struct {
		pattern string
		keys    []string
	}
	testCases := []struct {
		data []KV
		want []Key
	}{
		{
			data: []KV{{"she", 1}, {"sea", 1}, {"see", 1}, {"shell", 1}, {"shell", 1}, {"shellsort", 1}, {"hello", 1}, {"he", 1}},
			want: []Key{{"s.e", []string{"see", "she"}}, {".ea", []string{"sea"}}, {".hel.", []string{"shell"}}, {"he", []string{"he"}}, {"he..o", []string{"hello"}}, {"...", []string{"sea", "see", "she"}}},
		},
		{
			data: []KV{{"Han", 1}, {"shot", 1}, {"first", 1}, {",", 1}, {"he", 1}, {"said", 1}, {".", 1}},
			want: []Key{{".", []string{",", "."}}},
		},
	}

	for _, tc := range testCases {
		trie := Constructor()
		for _, d := range tc.data {
			trie.Put(d.key, d.val)
		}

		for _, w := range tc.want {
			keys := trie.KeysThatMatch(w.pattern)
			for i, k := range w.keys {
				if keys[i] != k {
					t.Errorf("Got key %v in collection of keys that match %v, want %v", keys[i], w.pattern, k)
				}
			}
		}
	}
}
