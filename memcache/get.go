package memcache

import (
	"fmt"
	"strings"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/pkg/errors"
)

// GetWithTags gets a cache with a key and tag(s). Tags cannot be omitted.
func (c *Client) GetWithTags(k string, tags []string) (item *memcache.Item, err error) {
	key, err := c.getKeyStr(k, tags)
	if err == memcache.ErrCacheMiss {
		return nil, memcache.ErrCacheMiss
	}

	return c.Client.Get(key)
}

func (c *Client) getKeyStr(key string, tags []string) (string, error) {
	keys := append(tags, key)
	revisions := make([]string, len(keys))

	keyValueMap, err := c.GetMulti(keys)
	if err != nil {
		return "", errors.Wrap(err, fmt.Sprintf("failed to get tag values. %v", tags))
	}

	if len(keyValueMap) != len(keys) {
		return "", memcache.ErrCacheMiss
	}

	for i := 0; i < len(keys); i++ {
		revisions[i] = convertByteToString(keyValueMap[keys[i]].Value)
	}

	return strings.Join(revisions, ":"), nil
}
