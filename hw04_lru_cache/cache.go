package hw04lrucache

type Key string

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
		prtVal.Value = value
		listQueue.PushFront(prtVal)
		answer = true
	} else {
		answer = false
		mapItems[key] = listQueue.PushFront(value)
		if len(mapItems) > l.capacity {
			prtLast := listQueue.Back()
			listQueue.Remove(prtLast)
			var keyDel Key
			for k, v := range mapItems {
				if v == prtLast {
					keyDel = k
				}
			}
			delete(mapItems, keyDel)
		}
	}
	return answer
}

func (l lruCache) Get(key Key) (interface{}, bool) {
	mapItems := l.items
	listQueue := l.queue
	if prtVal, ok := mapItems[key]; ok {
		listQueue.PushFront(prtVal)
		value := prtVal.Value
		return value, true
	} else {
		return nil, false
	}
}

func (l lruCache) Clear() {
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
