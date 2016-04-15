# Cache Store

Cache Store

```go
I18n := i18n.New(database.New(DB))
I18n.CacheStore = memcached.New(memcached.Config{})
```

## License

Released under the [MIT License](http://opensource.org/licenses/MIT).
