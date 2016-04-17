# Cache Store

Cache Store

## Memcached Cache Store Usage

```go
import "github.com/qor/cache/memcached"

func main() {
  client := memcached.New(&Config{Hosts: []string{"127.0.0.1:11211"}, NameSpace: "qor_demo_v1"})

  // Save value `Hello World` with key `hello_world` into cache store
	err := client.Set("hello_world", "Hello World")

  // Get saved value with key `hello_world`
	result, err := client.Get("hello_world")

  // Save marshal value of user into cache store
	err := client.Set("user", user)

  // Unmarshal saved value into user2
	err := client.Unmarshal("user", &user2)

  // Fetch saved value with key `hello_world`, if haven't find, will save returned result of `func` into cached store with passed key
	result, err := client.Fetch("hello_world", func() interface{} {
    return "..."
  })

  // Delete saved value
	err := client.Delete(key string)
}
```

## Memory Cache Store Usage

```go
import "github.com/qor/cache/memory"

func main() {
  client := memory.New()
  // Same API with memcached cache store
}
```

## License

Released under the [MIT License](http://opensource.org/licenses/MIT).
