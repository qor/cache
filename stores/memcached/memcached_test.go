package memcached

import (
	"reflect"
	"testing"
)

func TestPlainText(t *testing.T) {
	memcached := New(&Config{Hosts: []string{"127.0.0.1:11211"}})

	if err := memcached.Set("hello_world", "Hello World"); err != nil {
		t.Errorf("No error should happen when saving plain text into memcached")
	}

	if value, err := memcached.Get("hello_world"); err != nil || value != "Hello World" {
		t.Errorf("found value: %v", value)
	}

	if err := memcached.Delete("hello_world"); err != nil {
		t.Errorf("failed to delete value: %v", err)
	}

	if value, err := memcached.Get("hello_world"); err == nil || value == "Hello World" {
		t.Errorf("the key should been deleted")
	}
}

func TestUnmarshal(t *testing.T) {
	memcached := New(&Config{Hosts: []string{"127.0.0.1:11211"}})

	type result struct {
		Name  string
		Value string
	}

	r1 := result{Name: "result_name_1", Value: "result_value_1"}
	if err := memcached.Set("unmarshal", r1); err != nil {
		t.Errorf("No error should happen when saving struct into memcached")
	}

	var r2 result
	if err := memcached.Unmarshal("unmarshal", &r2); err != nil || !reflect.DeepEqual(r1, r2) {
		t.Errorf("found value: %#v", r2)
	}

	if err := memcached.Delete("unmarshal"); err != nil {
		t.Errorf("failed to delete value: %v", err)
	}

	if err := memcached.Unmarshal("unmarshal", &r2); err == nil {
		t.Errorf("the key should been deleted")
	}
}
