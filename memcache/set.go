package memcache

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/pkg/errors"
)

// SetWithTags sets a cache with key and tag(s). the cache can be invalidated by either of key or a tag.
func (c *Client) SetWithTags(item *memcache.Item, tags []string) error {
	keystr, err := c.generateKeyStr(item.Key, tags)
	if err != nil {
		return errors.Wrap(err, "failed to generate revision key string")
	}
	item.Key = keystr
	return c.Client.Set(item)
}

func (c *Client) generateKeyStr(k string, tags []string) (string, error) {
	keys := append(tags, k)
	revisions := make([]string, len(keys))

	keyValueMap, err := c.GetMulti(keys)
	if err != nil {
		return "", errors.Wrap(err, fmt.Sprintf("failed to get tag values. %v", tags))
	}

	for j := 0; j < len(keys); j++ {
		if revision, ok := keyValueMap[keys[j]]; ok {
			revisions[j] = convertByteToString(revision.Value)
		} else {
			rev := createRevision()

			revisions[j] = convertByteToString(rev)
			err = c.setRevision(keys[j], rev)
			if err != nil {
				return "", errors.Wrap(err, fmt.Sprintf("failed to set set revision cache. key=%v", k))
			}
		}
	}

	return strings.Join(revisions, ":"), nil
}

func createRevision() []byte {
	revision := make([]byte, 1)
	revision[0] = byte(rand.Intn(26) + 65)
	return revision
}

func (c *Client) setRevision(key string, revision []byte) error {
	err := c.Client.Set(&memcache.Item{Key: key, Value: revision})
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to set revision. key=%v", key))
	}
	return nil
}
