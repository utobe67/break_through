package lrucache

import (
	"code.byted.org/gopkg/pkg/testing/assert"
	"fmt"
	"testing"
)

func TestCache(t *testing.T) {
	var val interface{}
	var found bool
	cache := NewCache(3)
	cache.Set("a", 1)
	cache.Set("b", 2)
	cache.Set("c", 3)
	fmt.Printf("%+v\n", cache.Keys())
	val, found = cache.Get("a")
	fmt.Printf("%+v\n", cache.Keys())
	cache.Set("d", 4)
	val, found = cache.Get("b")
	fmt.Printf("%+v\n", cache.Keys())
	assert.Equal(t, false, found)
	cache.Set("e", 5)
	val, found = cache.Get("c")
	assert.Equal(t, false, found)
	val, found = cache.Get("a")
	assert.Equal(t, 1, val)
}
