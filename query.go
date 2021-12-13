package mgoi

import (
	"time"

	"github.com/globalsign/mgo"
)

type Query interface {
	All(result interface{}) error
	Apply(change Change, result interface{}) (info *ChangeInfo, err error)
	Batch(n int) Query
	Collation(collation *Collation) Query
	Comment(comment string) Query
	Count() (n int, err error)
	Distinct(key string, result interface{}) error
	Explain(result interface{}) error
	For(result interface{}, f func() error) error
	Hint(indexKey ...string) Query
	Iter() Iter
	Limit(n int) Query
	LogReplay() Query
	MapReduce(job *MapReduce, result interface{}) (info *MapReduceInfo, err error)
	One(result interface{}) (err error)
	Prefetch(p float64) Query
	Select(selector interface{}) Query
	SetMaxScan(n int) Query
	SetMaxTime(d time.Duration) Query
	Skip(n int) Query
	Snapshot() Query
	Sort(fields ...string) Query
	Tail(timeout time.Duration) Iter
}

func WrapQuery(q *mgo.Query) Query {
	return &queryWrapper{
		qq: q,
	}
}

var _ Query = (*queryWrapper)(nil)

type queryWrapper struct {
	qq *mgo.Query
}

func (q *queryWrapper) All(result interface{}) error {
	return q.qq.All(result)
}

func (q *queryWrapper) Apply(change Change, result interface{}) (info *ChangeInfo, err error) {
	return q.qq.Apply(change, result)
}

func (q *queryWrapper) Batch(n int) Query {
	q.qq.Batch(n)
	return q
}

func (q *queryWrapper) Collation(collation *Collation) Query {
	q.qq.Collation(collation)
	return q
}

func (q *queryWrapper) Comment(comment string) Query {
	q.qq.Comment(comment)
	return q
}

func (q *queryWrapper) Count() (n int, err error) {
	return q.qq.Count()
}

func (q *queryWrapper) Distinct(key string, result interface{}) error {
	return q.qq.Distinct(key, result)
}

func (q *queryWrapper) Explain(result interface{}) error {
	return q.qq.Explain(result)
}

func (q *queryWrapper) For(result interface{}, f func() error) error {
	return q.qq.For(result, f)
}

func (q *queryWrapper) Hint(indexKey ...string) Query {
	q.qq.Hint(indexKey...)
	return q
}

func (q *queryWrapper) Iter() Iter {
	return q.qq.Iter()
}

func (q *queryWrapper) Limit(n int) Query {
	q.qq.Limit(n)
	return q
}

func (q *queryWrapper) LogReplay() Query {
	q.qq.LogReplay()
	return q
}

func (q *queryWrapper) MapReduce(job *MapReduce, result interface{}) (info *MapReduceInfo, err error) {
	return q.qq.MapReduce(job, result)
}

func (q *queryWrapper) One(result interface{}) (err error) {
	return q.qq.One(result)
}

func (q *queryWrapper) Prefetch(p float64) Query {
	q.qq.Prefetch(p)
	return q
}

func (q *queryWrapper) Select(selector interface{}) Query {
	q.qq.Select(selector)
	return q
}

func (q *queryWrapper) SetMaxScan(n int) Query {
	q.qq.SetMaxScan(n)
	return q
}

func (q *queryWrapper) SetMaxTime(d time.Duration) Query {
	q.qq.SetMaxTime(d)
	return q
}

func (q *queryWrapper) Skip(n int) Query {
	q.qq.Skip(n)
	return q
}

func (q *queryWrapper) Snapshot() Query {
	q.qq.Snapshot()
	return q
}

func (q *queryWrapper) Sort(fields ...string) Query {
	q.qq.Sort(fields...)
	return q
}

func (q *queryWrapper) Tail(timeout time.Duration) Iter {
	return q.qq.Tail(timeout)
}
