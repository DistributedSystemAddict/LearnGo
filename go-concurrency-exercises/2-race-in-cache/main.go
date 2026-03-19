//////////////////////////////////////////////////////////////////////
//
// Given is some code to cache key-value pairs from a database into
// the main memory (to reduce access time). Note that golang's map are
// not entirely thread safe. Multiple readers are fine, but multiple
// writers are not. Change the code to make this thread safe.
//

package main

import (
	"container/list"
	"sync"
	"testing"
)

// CacheSize determines how big the cache can grow
const CacheSize = 100

var ch = make(chan struct{}, CacheSize)

var m sync.RWMutex

// KeyStoreCacheLoader is an interface for the KeyStoreCache
type KeyStoreCacheLoader interface {
	// Load implements a function where the cache should gets it's content from
	Load(string) string
}

type page struct {
	Key   string
	Value string
}

// KeyStoreCache is a LRU cache for string key-value pairs
type KeyStoreCache struct {
	cache map[string]*list.Element
	pages list.List
	load  func(string) string
}

// New creates a new KeyStoreCache
func New(load KeyStoreCacheLoader) *KeyStoreCache {
	return &KeyStoreCache{
		load:  load.Load,
		cache: make(map[string]*list.Element),
	}
}

var inFlight = make(map[string]chan struct{})

// Get gets the key from cache, loads it from the source if needed
func (k *KeyStoreCache) Get(key string) string {
	m.Lock()
	e, ok := k.cache[key]
	if ok {
		k.pages.MoveToFront(e)
		val := e.Value.(page).Value
		m.Unlock()
		return val
	}
	if ch, ok := inFlight[key]; ok {
		m.Unlock()        // Nhả lock ra cho goroutine khác chạy
		<-ch              // Đứng chờ thằng kia lấy DB xong (nó sẽ đóng channel)
		return k.Get(key) // Nó lấy xong rồi thì mình gọi lại Get để bốc từ Cache ra
	}

	// 3. Nếu chưa có ai đi lấy, MÌNH SẼ LÀ ĐẠI DIỆN đi lấy!
	// Tạo 1 channel đánh dấu là "Tôi đang đi lấy key này nhé, ai hỏi thì bảo chờ ở channel này"
	ch := make(chan struct{})
	inFlight[key] = ch
	m.Unlock() // Nhả lock để gọi DB không bị nghẽn hệ thống

	// Miss - load from database and save it in cache
	val := k.load(key)
	m.Lock()
	defer m.Unlock()
	if e, ok := k.cache[key]; ok {
		k.pages.MoveToFront(e)
		return e.Value.(page).Value
	}
	p := page{key, val}
	// if cache is full remove the least used item
	if len(k.cache) >= CacheSize {
		end := k.pages.Back()
		// remove from map
		delete(k.cache, end.Value.(page).Key)
		// remove from list
		k.pages.Remove(end)
	}
	k.pages.PushFront(p)
	k.cache[key] = k.pages.Front()

	delete(inFlight, key)
	close(ch)
	return p.Value
}

// Loader implements KeyStoreLoader
type Loader struct {
	DB *MockDB
}

// Load gets the data from the database
func (l *Loader) Load(key string) string {
	val, err := l.DB.Get(key)
	if err != nil {
		panic(err)
	}

	return val
}

func run(t *testing.T) (*KeyStoreCache, *MockDB) {
	loader := Loader{
		DB: GetMockDB(),
	}
	cache := New(&loader)

	RunMockServer(cache, t)

	return cache, loader.DB
}

func main() {
	run(nil)
}
