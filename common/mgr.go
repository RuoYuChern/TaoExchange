package common

import (
	"container/list"
	"sync"
)

type infraMgr struct {
	list *list.List
}

var mgr *infraMgr
var once sync.Once

func Get() *infraMgr {
	once.Do(func() {
		mgr = &infraMgr{list: list.New()}
	})
	return mgr
}

func (m *infraMgr) Add(c AutoCloseable) {
	m.list.PushBack(c)
}

func (m *infraMgr) Close() {
	for c := m.list.Front(); c != nil; c = c.Next() {
		if c != nil {
			ac := c.Value.(AutoCloseable)
			ac.AutoClose()
		}
	}
}
