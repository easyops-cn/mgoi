package mgoi

import (
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type Session interface {
	BuildInfo() (info BuildInfo, err error)
	Clone() Session
	Close()
	Copy() Session
	DB(name string) Database
	DatabaseNames() (names []string, err error)
	EnsureSafe(safe *Safe)
	FindRef(ref *DBRef) Query
	Fsync(async bool) error
	FsyncLock() error
	FsyncUnlock() error
	LiveServers() (addrs []string)
	Login(cred *Credential) error
	LogoutAll()
	Mode() Mode
	New() Session
	Ping() error
	Refresh()
	ResetIndexCache()
	Run(cmd interface{}, result interface{}) error
	Safe() (safe *Safe)
	SelectServers(tags ...bson.D)
	SetBatch(n int)
	SetBypassValidation(bypass bool)
	SetCursorTimeout(d time.Duration)
	SetMode(consistency Mode, refresh bool)
	SetPoolLimit(limit int)
	SetPoolTimeout(timeout time.Duration)
	SetPrefetch(p float64)
	SetSafe(safe *Safe)
	SetSocketTimeout(d time.Duration)
	SetSyncTimeout(d time.Duration)
}

func WrapSession(s *mgo.Session) Session {
	return &sessionWrapper{
		s: s,
	}
}

var _ Session = (*sessionWrapper)(nil)

type sessionWrapper struct {
	s *mgo.Session
}

func (s *sessionWrapper) BuildInfo() (info BuildInfo, err error) {
	return s.s.BuildInfo()
}

func (s *sessionWrapper) Clone() Session {
	return WrapSession(s.s.Clone())
}

func (s *sessionWrapper) Close() {
	s.s.Close()
}

func (s *sessionWrapper) Copy() Session {
	return WrapSession(s.s.Copy())
}

func (s *sessionWrapper) DB(name string) Database {
	db := s.s.DB(name)
	return WrapDatabase(db)
}

func (s *sessionWrapper) DatabaseNames() (names []string, err error) {
	return s.s.DatabaseNames()
}

func (s *sessionWrapper) EnsureSafe(safe *Safe) {
	s.s.EnsureSafe(safe)
}

func (s *sessionWrapper) FindRef(ref *DBRef) Query {
	query := s.s.FindRef(ref)
	return WrapQuery(query)
}

func (s *sessionWrapper) Fsync(async bool) error {
	return s.s.Fsync(async)
}

func (s *sessionWrapper) FsyncLock() error {
	return s.s.FsyncLock()
}

func (s *sessionWrapper) FsyncUnlock() error {
	return s.s.FsyncUnlock()
}

func (s *sessionWrapper) LiveServers() (addrs []string) {
	return s.s.LiveServers()
}

func (s *sessionWrapper) Login(cred *Credential) error {
	return s.s.Login(cred)
}

func (s *sessionWrapper) LogoutAll() {
	s.s.LogoutAll()
}

func (s *sessionWrapper) Mode() Mode {
	return s.s.Mode()
}

func (s *sessionWrapper) New() Session {
	return WrapSession(s.s.New())
}

func (s *sessionWrapper) Ping() error {
	return s.s.Ping()
}

func (s *sessionWrapper) Refresh() {
	s.s.Refresh()
}

func (s *sessionWrapper) ResetIndexCache() {
	s.s.ResetIndexCache()
}

func (s *sessionWrapper) Run(cmd interface{}, result interface{}) error {
	return s.s.Run(cmd, result)
}

func (s *sessionWrapper) Safe() (safe *Safe) {
	return s.s.Safe()
}

func (s *sessionWrapper) SelectServers(tags ...bson.D) {
	s.s.SelectServers(tags...)
}

func (s *sessionWrapper) SetBatch(n int) {
	s.s.SetBatch(n)
}

func (s *sessionWrapper) SetBypassValidation(bypass bool) {
	s.s.SetBypassValidation(bypass)
}

func (s *sessionWrapper) SetCursorTimeout(d time.Duration) {
	s.s.SetCursorTimeout(d)
}

func (s *sessionWrapper) SetMode(consistency Mode, refresh bool) {
	s.s.SetMode(consistency, refresh)
}

func (s *sessionWrapper) SetPoolLimit(limit int) {
	s.s.SetPoolLimit(limit)
}

func (s *sessionWrapper) SetPoolTimeout(timeout time.Duration) {
	s.s.SetPoolTimeout(timeout)
}

func (s *sessionWrapper) SetPrefetch(p float64) {
	s.s.SetPrefetch(p)
}

func (s *sessionWrapper) SetSafe(safe *Safe) {
	s.s.SetSafe(safe)
}

func (s *sessionWrapper) SetSocketTimeout(d time.Duration) {
	s.s.SetSocketTimeout(d)
}

func (s *sessionWrapper) SetSyncTimeout(d time.Duration) {
	s.s.SetSyncTimeout(d)
}
