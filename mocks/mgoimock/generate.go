package mgoimock

//go:generate mockgen -destination mocks.go -package mgoimock github.com/easyops-cn/mgoi Session,Database,Collection,Query,Iter,Bulk
