package lru

type LRU struct {
	// whatever fields you want here
}

// Return a new LRU with capacity to store limit bytes.
func NewLru(limit int) *LRU {
	return new(LRU)
}

// Return the maximum number of bytes that your LRU can store.
func (lru *LRU) MaxStorage() int {
	return 0
}

// Return the number of bytes of storage remaining in your LRU.
func (lru *LRU) RemainingStorage() int {
	return 0
}

// Return the value associated with the specified key from the LRU.
//
// Use `ok=true` to indicate the value was returned successfully, or `ok=false`
// to indicate some issue (e.g. no binding exists for that key)
//
// Additionally, each call to `Get` should update the binding in some way
// to ensure that it is the most-recently-used.
func (lru *LRU) Get(key string) (value []byte, ok bool) {
	return nil, false
}

// Remove the binding with specified key from the LRU, and return the value
// that was bound to it. Use `ok=true` to indicate the value was removed
// and returned successfully, or `ok=false` to indicate some issue
// (e.g. no binding exists for that key)
func (lru *LRU) Remove(key string) (value []byte, ok bool) {
	return nil, false
}

// Add a binding to the LRU with the specified key and value.
//
// Use len(key) and len(value) to give the number of bytes that this new binding
// would consume, and ensure that there is enough space in the LRU to
// accommodate it.
//
// If the LRU is not large enough to accommodate the binding, return false.
//
// Otherwise, evict the least-recently-used binding as many t
func (lru *LRU) Set(key string, value []byte) bool {
	return false
}

// Return the number of bindings currently stored in the LRU.
func (lru *LRU) Len() int {
	return 0
}
