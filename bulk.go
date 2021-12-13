package mgoi

import "github.com/globalsign/mgo"

type Bulk interface {
	Insert(docs ...interface{})
	Remove(selectors ...interface{})
	RemoveAll(selectors ...interface{})
	Run() (*BulkResult, error)
	Unordered()
	Update(pairs ...interface{})
	UpdateAll(pairs ...interface{})
	Upsert(pairs ...interface{})
}

var _ Bulk = (*mgo.Bulk)(nil)
