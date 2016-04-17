package memory

import (
	"encoding/json"
	"errors"
)

var ErrNotFound = errors.New("not found")

type Memory struct {
	values map[string][]byte
}

func New() *Memory {
	return &Memory{values: map[string][]byte{}}
}

func (memory *Memory) Get(key string) (string, error) {
	if value, ok := memory.values[key]; ok {
		return string(value), nil
	}
	return "", ErrNotFound
}

func (memory *Memory) Unmarshal(key string, object interface{}) error {
	if value, ok := memory.values[key]; ok {
		return json.Unmarshal(value, object)
	}
	return ErrNotFound
}

func convertToBytes(value interface{}) []byte {
	switch result := value.(type) {
	case string:
		return []byte(result)
	case []byte:
		return result
	default:
		bytes, _ := json.Marshal(value)
		return bytes
	}
}

func (memory *Memory) Set(key string, value interface{}) error {
	memory.values[key] = convertToBytes(value)
	return nil
}

func (memory *Memory) Fetch(key string, fc func() interface{}) (string, error) {
	if str, err := memory.Get(key); err == nil {
		return str, nil
	}
	results := convertToBytes(fc())
	return string(results), memory.Set(key, results)
}

func (memory *Memory) Delete(key string) error {
	delete(memory.values, key)
	return nil
}
