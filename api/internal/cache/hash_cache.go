package cache

import (
	"github.com/bjvanbemmel/benkyou/internal/utils"
)

type HashCache[T any] struct {
	internal Cache[string, T]
}

func NewHash[T any]() HashCache[T] {
	return HashCache[T]{
		internal: Cache[string, T]{},
	}
}

func (h *HashCache[T]) Insert(t T) error {
	return h.internal.Insert(utils.Hash(t), t)
}

func (h *HashCache[T]) Get(t T) (T, error) {
	return h.internal.Get(utils.Hash(t))
}

func (h *HashCache[T]) Delete(t T) error {
	return h.internal.Delete(utils.Hash(t))
}
