package sort

import "testing"

func TestInsertion(t *testing.T) {
	for _, tc := range testCases {
		Insertion(tc.data)
		err := false
		for i := 0; i < len(tc.data); i++ {
			if tc.data[i] != tc.want[i] {
				err = true
			}
		}
		if err {
			t.Errorf("Got %v, want %v", tc.data, tc.want)
		}
	}
}
