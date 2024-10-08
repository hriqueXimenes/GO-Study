package main

import "fmt"

type Cache struct {
	size    int
	storage map[string]Data
	order   []string
}

type Data struct {
	value interface{}
}

func main() {
	cache := Cache{
		storage: map[string]Data{},
		size:    3,
	}

	cache.Append("key1", "val1")
	cache.Append("key2", "val2")
	cache.Append("key3", "val3")
	cache.Append("key4", "val4")
	cache.Append("key5", "val5")
	cache.Append("key6", "val6")
	cache.Delete("key5")
	cache.Delete("key6")
	cache.Append("key1", "val1")
	cache.Append("key2", "val2")

	fmt.Println(cache.GetAllKeys())
}

func (c *Cache) Append(key, value string) {
	if len(c.storage) >= c.size && len(c.storage) >= 1 {
		oldestData := c.order[0]
		c.order = c.order[1:]
		delete(c.storage, oldestData)
	}

	c.storage[key] = Data{
		value: value,
	}

	c.order = append(c.order, key)
}

func (c *Cache) GetAllKeys() []string {
	return c.order
}

func (c *Cache) Delete(key string) {

	if _, found := c.storage[key]; found {
		delete(c.storage, key)

		for i, k := range c.order {
			if k == key {
				c.order = append(c.order[:i], c.order[i+1:]...)
				break
			}
		}
	}

}
