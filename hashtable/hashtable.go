/*
Package hashtable provides a Hash Table
*/
package hashtable

import (
	"github.com/dorin131/go-data-structures/linkedlist"
)

// HashTable : a hash table data structure
type HashTable struct {
	data []*linkedlist.LinkedList
}

type listData struct {
	key   string
	value string
}

// New : a hash table constructor
func New() *HashTable {
	return &HashTable{
		make([]*linkedlist.LinkedList, 5),
	}
}

func hash(s string) int {
	return 0
}

func index(hash int) int {
	return 0
}

func (h *HashTable) Set(k, v string) *HashTable {
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

func (h *HashTable) Get(k string) (result string, ok bool) {
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
