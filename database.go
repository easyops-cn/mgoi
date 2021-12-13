package mgoi

import (
	"github.com/globalsign/mgo"
)

type Database interface {
	AddUser(username, password string, readOnly bool) error
	C(name string) Collection
	CollectionNames() (names []string, err error)
	CreateView(view string, source string, pipeline interface{}, collation *Collation) error
	DropDatabase() error
	FindRef(ref *DBRef) Query
	// GridFS(prefix string) *GridFS
	Login(user, pass string) error
	Logout()
	RemoveUser(user string) error
	Run(cmd interface{}, result interface{}) error
	UpsertUser(user *User) error
	// With(s *Session) *Database
}

func WrapDatabase(d *mgo.Database) Database {
	return &databaseWrapper{
		d: d,
	}
}

type databaseWrapper struct {
	d *mgo.Database
}

func (d *databaseWrapper) AddUser(username, password string, readOnly bool) error {
	return d.d.AddUser(username, password, readOnly)
}

func (d *databaseWrapper) C(name string) Collection {
	return WrapCollection(d.d.C(name))
}

func (d *databaseWrapper) CollectionNames() (names []string, err error) {
	return d.d.CollectionNames()
}

func (d *databaseWrapper) CreateView(view string, source string, pipeline interface{}, collation *Collation) error {
	return d.d.CreateView(view, source, pipeline, collation)
}

func (d *databaseWrapper) DropDatabase() error {
	return d.d.DropDatabase()
}

func (d *databaseWrapper) FindRef(ref *DBRef) Query {
	return WrapQuery(d.d.FindRef(ref))
}

func (d *databaseWrapper) Login(user, pass string) error {
	return d.d.Login(user, pass)
}

func (d *databaseWrapper) Logout() {
	d.d.Logout()
}

func (d *databaseWrapper) RemoveUser(user string) error {
	return d.d.RemoveUser(user)
}

func (d *databaseWrapper) Run(cmd interface{}, result interface{}) error {
	return d.d.Run(cmd, result)
}

func (d *databaseWrapper) UpsertUser(user *User) error {
	return d.d.UpsertUser(user)
}
