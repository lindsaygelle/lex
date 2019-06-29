// Package lexicon provides interface for managing collections of abstract data in an Map-like structure.
package lexicon

import (
	"fmt"

	"github.com/gellel/slice"
)

var (
	_ lexicon = (*Lexicon)(nil)
)

// New instantiates a new, empty Lexicon pointer.
func New() *Lexicon {
	return &Lexicon{}
}

// NewLexicon instantiates a empty or populated Lexicon pointer. Takes an argument of 0-N maps.
func NewLexicon(m ...map[string]interface{}) *Lexicon {
	lexicon := Lexicon{}
	for _, m := range m {
		for k, v := range m {
			lexicon[k] = v
		}
	}
	return &lexicon
}

type lexicon interface {
	Add(key string, value interface{}) *Lexicon
	Del(key string) bool
	Each(f func(key string, value interface{})) *Lexicon
	Empty() bool
	Fetch(key string) interface{}
	Get(key string) (interface{}, bool)
	Has(key string) bool
	Intersection(lexicon *Lexicon) *Lexicon
	Keys() *slice.String
	Len() int
	Map(f func(key string, value interface{}) interface{}) *Lexicon
	Merge(lexicon ...*Lexicon) *Lexicon
	Mesh(m ...map[string]interface{}) *Lexicon
	Peek(key string) string
	Values() *slice.Slice
}

// Lexicon is a map-like object whose methods are used to perform traversal and mutation operations by key-value pair.
type Lexicon map[string]interface{}

// Add method adds one element to the Lexicon using the key reference and returns the modified Lexicon.
func (pointer *Lexicon) Add(key string, value interface{}) *Lexicon {
	(*pointer)[key] = value
	return pointer
}

// Del method removes a entry from the Lexicon if it exists. Returns a boolean to confirm if it succeeded.
func (pointer *Lexicon) Del(key string) bool {
	ok := pointer.Has(key)
	if ok {
		delete(*pointer, key)
		ok = (pointer.Has(key) == false)
	}
	return ok
}

// Each method executes a provided function once for each Lexicon element.
func (pointer *Lexicon) Each(f func(key string, value interface{})) *Lexicon {
	for key, value := range *pointer {
		f(key, value)
	}
	return pointer
}

// Empty returns a boolean indicating whether the Lexicon contains zero values.
func (pointer *Lexicon) Empty() bool {
	return pointer.Len() == 0
}

// Fetch retrieves the interface held by the argument key. Returns nil if key does not exist.
func (pointer *Lexicon) Fetch(key string) interface{} {
	value, _ := (*pointer)[key]
	return value
}

// Get returns the value held at the argument key and a boolean indicating if it was successfully retrieved.
func (pointer *Lexicon) Get(key string) (interface{}, bool) {
	value, ok := (*pointer)[key]
	return value, ok
}

// Has method checks that a given key exists in the Lexicon.
func (pointer *Lexicon) Has(key string) bool {
	_, ok := pointer.Get(key)
	return ok
}

// Intersection method returns a new Lexicon containing the shared keys between two Lexicons.
func (pointer *Lexicon) Intersection(lexicon *Lexicon) *Lexicon {
	m := &Lexicon{}
	pointer.Each(func(key string, _ interface{}) {
		if ok := lexicon.Has(key); ok {
			m.Add(key, true)
		}
	})
	lexicon.Each(func(key string, _ interface{}) {
		if ok := pointer.Has(key); ok {
			m.Add(key, 1)
		}
	})
	return m
}

// Keys method returns a slice.String of the Lexicon's own property names, in the same order as we get with a normal loop.
func (pointer *Lexicon) Keys() *slice.String {
	s := slice.NewString()
	for key := range *pointer {
		s.Append(key)
	}
	return s
}

// Len method returns the number of keys in the Lexicon.
func (pointer *Lexicon) Len() int {
	return len(*pointer)
}

// Map method executes a provided function once for each Lexicon element and sets the returned value to the current key.
func (pointer *Lexicon) Map(f func(key string, value interface{}) interface{}) *Lexicon {
	for key, value := range *pointer {
		pointer.Add(key, f(key, value))
	}
	return pointer
}

// Merge merges N number of Lexicons.
func (pointer *Lexicon) Merge(lexicon ...*Lexicon) *Lexicon {
	for _, lexicon := range lexicon {
		lexicon.Each(func(key string, value interface{}) {
			pointer.Add(key, value)
		})
	}
	return pointer
}

// Mesh merges a collection maps to the Lexicon.
func (pointer *Lexicon) Mesh(m ...map[string]interface{}) *Lexicon {
	for _, m := range m {
		for k, v := range m {
			pointer.Add(k, v)
		}
	}
	return pointer
}

// Peek returns the string value of the element assigned to the argument key.
func (pointer *Lexicon) Peek(key string) string {
	return fmt.Sprintf("%v", pointer.Fetch(key))
}

// Values method returns a slice.Slice pointer of the Lexicon's own enumerable property values, in the same order as that provided by a for...in loop.
func (pointer *Lexicon) Values() *slice.Slice {
	s := &slice.Slice{}
	for _, value := range *pointer {
		s.Append(value)
	}
	return s
}
