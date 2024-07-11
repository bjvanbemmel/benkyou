package cache

import (
	"errors"
	"sync"
)

type Cache[K comparable, V any] struct {
	store map[K]V
	mutex *sync.Mutex
}

func New[K comparable, V any]() Cache[K, V] {
	return Cache[K, V]{
		store: make(map[K]V),
		mutex: &sync.Mutex{},
	}
}

func (c *Cache[K, V]) Insert(k K, v V) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if _, ok := c.store[k]; ok {
		return errors.New("store already contains item with given key")
	}

	c.store[k] = v

	return nil
}

func (c *Cache[K, V]) Get(k K) (V, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	v, ok := c.store[k]
	if !ok {
		return v, errors.New("store contains no such item")
	}

	return v, nil
}

func (c *Cache[K, V]) Delete(k K) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if _, err := c.Get(k); err != nil {
		return err
	}

	delete(c.store, k)

	return nil
}
