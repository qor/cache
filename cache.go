package cache

type Store interface {
	Get(key string) (string, error)
	Load(key string, object interface{}) error
	Set(key string, value interface{}) error
	Fetch(key string, fc func() interface{}) (string, error)
	Delete(key string) error
}
