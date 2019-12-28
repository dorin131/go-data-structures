package hashtable

import (
	"testing"
)

func TestHashTableBasic(t *testing.T) {
	h := New()
	if h.data[0] != nil {
		t.Error("Expected hash data to be nil")
	}
	h.Set("Dorin", "great")

	if h.data[0] == nil {
		t.Error("Expected hash data to be NOT nil")
	}

	if result, ok := h.Get("Dorin"); !ok {
		t.Error("Expected Get to return a value")
	} else {
		t.Logf("Got %s\n", result)
	}
}

func TestHashTableSetGet(t *testing.T) {
	var h *HashTable
	testCases := []struct {
		key   string
		value string
	}{
		{"Dorin", "Great"},
		{"SomeKey", "SomeValue"},
		{"__", "%"},
		{" ", ""},
		{"", " "},
		{"", ""},
	}

	for i, test := range testCases {
		h = New()
		h.Set(test.key, test.value)
		result, ok := h.Get(test.key)
		if !ok {
			t.Errorf("[%d]: Could not get value from key \"%s\"", i, test.key)
		}
		if result != test.value {
			t.Errorf(
				"[%d]: Incorrect value for key \"%s\", got = \"%s\", expected = \"%s\"",
				i, test.key, result, test.value,
			)
		}
	}
}
