package hw04lrucache

type Key string

type cacheItem struct {
	key   Key
	value interface{}
}

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func (l lruCache) Set(key Key, value interface{}) bool {
	mapItems := l.items
	listQueue := l.queue
	var answer bool
	if prtVal, ok := mapItems[key]; ok {
		prtVal.Value = cacheItem{value: value, key: key}
		listQueue.PushFront(prtVal)
		answer = true
	} else {
		answer = false
		mapItems[key] = listQueue.PushFront(cacheItem{value: value, key: key})
		if len(mapItems) > l.capacity {
			prtLast := listQueue.Back()
			listQueue.Remove(prtLast)
			delete(mapItems, prtLast.Value.(cacheItem).key)
		}
	}
	return answer
}

func (l lruCache) Get(key Key) (interface{}, bool) {
	mapItems := l.items
	listQueue := l.queue
	if prtVal, ok := mapItems[key]; ok {
		listQueue.PushFront(prtVal)
		value := prtVal.Value.(cacheItem).value
		return value, true
	}
	return nil, false
}

func (l *lruCache) Clear() {
	l.items = make(map[Key]*ListItem, l.capacity)
	l.queue = NewList()
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
