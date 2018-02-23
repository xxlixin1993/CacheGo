package lru

import (
	"testing"
)

func TestAddAndGet(t *testing.T) {
	var (
		key   = "foo"
		value = "bar"
	)
	cache := NewCache(0)

	ok := cache.Add(key, value)
	if ok != true {
		t.Fatalf("cache add fail, key is %s, value is %s", key, value)
	}

	getV, getOk := cache.Get(key)
	if getOk != true {
		t.Fatalf("cache get fail, key is %s, value is %s", key, value)
	}
	if getV != value {
		t.Fatalf("cache get fail, get value is %s, value is %s", getV, value)
	}
}

func TestDelete(t *testing.T) {
	cache := NewCache(0)

	cache.Add("1", "1")
	cache.Add("2", "2")
	cache.Add("3", "3")

	ok := cache.Delete("1")
	if ok != true {
		t.Fatalf("cache delete fail, delete key is 1, value is 1")
	}

	two, _ := cache.Get("2")
	three, _ := cache.Get("3")

	if two != "2" || three != "3" {
		t.Fatalf("cache delete fail, delete key is 1, value is 1. but anthor key has deleted")
	}

}

func TestLen(t *testing.T) {
	cache := NewCache(0)

	cache.Add("1", "1")
	cache.Add("2", "2")
	cache.Add("3", "3")
	cache.Add("4", "4")

	len := cache.Len()
	if len != 4 {
		t.Fatalf("cache get length fail, expected 4 but output %d", len)
	}
}

// TODO test list cases