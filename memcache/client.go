package memcache

import (
	"math/rand"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
)

// NewClient returns a memcached client that wraps the original memcached client
func NewClient(server ...string) Client {
	rand.Seed(time.Now().UnixNano())
	return Client{
		Client: memcache.New(server...),
	}
}

// Client is a wrapper of the original memcached Client
type Client struct {
	*memcache.Client
}
