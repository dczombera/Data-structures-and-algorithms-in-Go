package ternary_search_trie

import "testing"

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
