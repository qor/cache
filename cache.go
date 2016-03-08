package cache

type Option map[string]string

type Store interface {
	Read(key string) string
	Write(key string, value string)
	Fetch(key string, fc func() string)
	Delete(key string)
}
