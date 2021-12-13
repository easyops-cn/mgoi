package mgoi

import (
	"github.com/globalsign/mgo"
)

type Collection interface {
	Bulk() Bulk
	Count() (n int, err error)
	Create(info *CollectionInfo) error
	DropAllIndexes() error
	DropCollection() error
	DropIndex(key ...string) error
	DropIndexName(name string) error
	EnsureIndex(index Index) error
	EnsureIndexKey(key ...string) error
	Find(query interface{}) Query
	FindId(id interface{}) Query
	Indexes() (indexes []Index, err error)
	Insert(docs ...interface{}) error
	// NewIter(session *Session, firstBatch []bson.Raw, cursorId int64, err error) *Iter
	// Pipe(pipeline interface{}) Pipe
	Remove(selector interface{}) error
	RemoveAll(selector interface{}) (info *ChangeInfo, err error)
	RemoveId(id interface{}) error
	Repair() Iter
	Update(selector interface{}, update interface{}) error
	UpdateAll(selector interface{}, update interface{}) (info *ChangeInfo, err error)
	UpdateId(id interface{}, update interface{}) error
	Upsert(selector interface{}, update interface{}) (info *ChangeInfo, err error)
	UpsertId(id interface{}, update interface{}) (info *ChangeInfo, err error)
	// Watch(pipeline interface{}, options ChangeStreamOptions) (*ChangeStream, error)
	// With(s *Session) *Collection
}

func WrapCollection(c *mgo.Collection) Collection {
	return &collectionWrapper{
		c: c,
	}
}

type collectionWrapper struct {
	c *mgo.Collection
}

func (c *collectionWrapper) Bulk() Bulk {
	return c.c.Bulk()
}

func (c *collectionWrapper) Count() (n int, err error) {
	return c.c.Count()
}

func (c *collectionWrapper) Create(info *CollectionInfo) error {
	return c.c.Create(info)
}

func (c *collectionWrapper) DropAllIndexes() error {
	return c.c.DropAllIndexes()
}

func (c *collectionWrapper) DropCollection() error {
	return c.c.DropCollection()
}

func (c *collectionWrapper) DropIndex(key ...string) error {
	return c.c.DropIndex(key...)
}

func (c *collectionWrapper) DropIndexName(name string) error {
	return c.c.DropIndexName(name)
}

func (c *collectionWrapper) EnsureIndex(index Index) error {
	return c.c.EnsureIndex(index)
}

func (c *collectionWrapper) EnsureIndexKey(key ...string) error {
	return c.c.EnsureIndexKey(key...)
}

func (c *collectionWrapper) Find(query interface{}) Query {
	return WrapQuery(c.c.Find(query))
}

func (c *collectionWrapper) FindId(id interface{}) Query {
	return WrapQuery(c.c.FindId(id))
}

func (c *collectionWrapper) Indexes() (indexes []Index, err error) {
	return c.c.Indexes()
}

func (c *collectionWrapper) Insert(docs ...interface{}) error {
	return c.c.Insert(docs...)
}

func (c *collectionWrapper) Remove(selector interface{}) error {
	return c.c.Remove(selector)
}

func (c *collectionWrapper) RemoveAll(selector interface{}) (info *ChangeInfo, err error) {
	return c.c.RemoveAll(selector)
}

func (c *collectionWrapper) RemoveId(id interface{}) error {
	return c.c.RemoveId(id)
}

func (c *collectionWrapper) Repair() Iter {
	return c.c.Repair()
}

func (c *collectionWrapper) Update(selector interface{}, update interface{}) error {
	return c.c.Update(selector, update)
}

func (c *collectionWrapper) UpdateAll(selector interface{}, update interface{}) (info *ChangeInfo, err error) {
	return c.c.UpdateAll(selector, update)
}

func (c *collectionWrapper) UpdateId(id interface{}, update interface{}) error {
	return c.c.UpdateId(id, update)
}

func (c *collectionWrapper) Upsert(selector interface{}, update interface{}) (info *ChangeInfo, err error) {
	return c.c.Upsert(selector, update)
}

func (c *collectionWrapper) UpsertId(id interface{}, update interface{}) (info *ChangeInfo, err error) {
	return c.c.UpsertId(id, update)
}
