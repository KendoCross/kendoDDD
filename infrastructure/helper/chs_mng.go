package helper

import "sync"

var cacheCoreChs map[string]chan<- interface{} = make(map[string]chan<- interface{})
var cacheMuCh *sync.RWMutex = &sync.RWMutex{}

func AddChCache(uid string, ch chan<- interface{}) {
	cacheMuCh.Lock()
	defer cacheMuCh.Unlock()
	cacheCoreChs[uid] = ch
}

func GetChCache(uid string) (ch chan<- interface{}, has bool) {
	cacheMuCh.RLock()
	defer cacheMuCh.RUnlock()
	ch, has = cacheCoreChs[uid]
	return
}

func DelChCache(uid string) {
	cacheMuCh.Lock()
	defer cacheMuCh.Unlock()
	delete(cacheCoreChs, uid)
}
