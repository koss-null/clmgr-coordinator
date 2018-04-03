package common

// Counter implements some integer counters with various capacity
type (
	// default
	counter struct {
		n int
	}

	counter8 struct {
		n int8
	}

	counter32 struct {
		n int32
	}

	counter64 struct {
		n int64
	}

	Counter interface {
		Count() interface{}
	}
)

// Provides ability to create global counters
var counters map[string]Counter

// Set counter if it's new
func SetGlobalCounter(id string, c Counter) bool {
	_, ok := counters[id]
	if !ok {
		counters[id] = c
	}
	return ok
}

// Set counter anyway
func ResetGlobalCounter(id string, c Counter) {
	counters[id] = c
}

func Count(id string) interface{} {
	if _, ok := counters[id]; ok {
		return counters[id].Count()
	}
	counters[id] = CommonCounter()
	return counters[id].Count()
}

func CommonCounter() Counter {
	return &counter{0}
}

func Counter8() Counter {
	return &counter8{0}
}

func Counter32() Counter {
	return &counter32{0}
}

func Counter64() Counter {
	return &counter64{0}
}

func (c *counter) Count() interface{} {
	c.n++
	return c.n
}

func (c *counter8) Count() interface{} {
	c.n++
	return c.n
}

func (c *counter32) Count() interface{} {
	c.n++
	return c.n
}

func (c *counter64) Count() interface{} {
	c.n++
	return c.n
}
