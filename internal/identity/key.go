package identity

import "github.com/dogmatiq/dogma"

// Key returns a handler's unique key.
func Key(h configurable) string {
	var c configurer
	h.Configure(&c)
	return c.key
}

var _ dogma.ProjectionConfigurer = (*configurer)(nil)

type configurable interface {
	Configure(dogma.ProjectionConfigurer)
}

type configurer struct {
	key string
}

func (c *configurer) Identity(name string, key string) {
	c.key = key
}

func (c *configurer) ConsumesEventType(m dogma.Message) {
}