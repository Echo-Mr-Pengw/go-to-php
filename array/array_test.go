package array

import "testing"

func TestParray_change_key_case(t *testing.T) {
	m := map[string]interface{}{
		"java": "java",
		"go": "go",
	}

	t.Log(Parray_change_key_case(m))
}