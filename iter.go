package mgoi

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type Iter interface {
	All(result interface{}) error
	Close() error
	Done() bool
	Err() error
	For(result interface{}, f func() error) (err error)
	Next(result interface{}) bool
	State() (int64, []bson.Raw)
	Timeout() bool
}

var _ Iter = (*mgo.Iter)(nil)
