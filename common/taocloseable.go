package common

import (
	"container/list"
	"sync"
)

type TaoCloseable interface {
	AutoClose()
}

type InfraMgr struct {
	list *list.List
}

var mgr *InfraMgr
var once sync.Once

func Get() *InfraMgr {
	once.Do(func() {
		mgr = &InfraMgr{list: list.New()}
	})
	return mgr
}

func (m *InfraMgr) Add(c TaoCloseable) {
	m.list.PushBack(c)
}

func (m *InfraMgr) Close() {
	for c := m.list.Front(); c != nil; c = c.Next() {
		if c != nil {
			ac := c.Value.(TaoCloseable)
			ac.AutoClose()
		}
	}
}
