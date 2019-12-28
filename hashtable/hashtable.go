/*
Package hashtable provides a Hash Table
*/
package hashtable

import (
	"math"

	"github.com/dorin131/go-data-structures/linkedlist"
)

const arrayLength = 100

// HashTable : a hash table data structure
type HashTable struct {
	data [arrayLength]*linkedlist.LinkedList
}

type listData struct {
	key   string
	value interface{}
}

// New : a hash table constructor
func New() *HashTable {
	return &HashTable{
		[arrayLength]*linkedlist.LinkedList{},
	}
}

func hash(s string) int {
	h := 0
	// char[0] * 31^n-1 + char[1] * 31^n-2 + ... + char[n-1]
	for pos, char := range s {
		h += int(char) * int(math.Pow(31, float64(len(s)-pos+1)))
	}
	return h
}

func index(hash int) int {
	return hash % arrayLength
}

// Set : set a key and value
func (h *HashTable) Set(k string, v interface{}) *HashTable {
	index := index(hash(k))

	if h.data[index] == nil {
		h.data[index] = linkedlist.New()
		h.data[index].Append(listData{k, v})
	} else {
		node := h.data[index].Head
		for {
			if node != nil {
				d := node.Data.(listData)
				if d.key == k {
					d.value = v
					break
				}
			} else {
				h.data[index].Append(listData{k, v})
				break
			}
			node = node.Next
		}
	}
	return h
}

// Get : get a value by key
func (h *HashTable) Get(k string) (result interface{}, ok bool) {
	index := index(hash(k))
	linkedList := h.data[index]

	if linkedList == nil {
		return "", false
	}
	node := linkedList.Head
	for {
		if node != nil {
			d := node.Data.(listData)
			if d.key == k {
				return d.value, true
			}
		} else {
			return "", false
		}
		node = node.Next
	}
}
