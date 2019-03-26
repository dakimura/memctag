## About
This is a memcached client library which can set tag(s) and delete by tag

## Installing

### Using *go get*

    $ go get github.com/dakimura/memctag/memcache

## Usage

    package main

    import (
        "fmt"
        "github.com/bradfitz/gomemcache/memcache"
        memcacheWithTag "github.com/dakimura/memctag/memcache"
    )


    func main(){
        // initialize memcached client
        mc := memcacheWithTag.NewClient("127.0.0.1:11211")

        tags := []string{"1-tag", "2-tag"}
        tag := []string{"1-tag"}

        item := memcache.Item{Key: "1:message", Value: []byte("cache1")}
        item2 := memcache.Item{Key: "2:message", Value: []byte("cache2")}
        item3 := memcache.Item{Key: "3:message", Value: []byte("cache3")}

        // set cache with tag(s)
        mc.SetWithTags(&item, tags)
        mc.SetWithTags(&item2, tags)
        mc.SetWithTags(&item3, tag)

        printCache(mc.GetWithTags("1:message", tags)) //"1-tag", "2-tag"
        printCache(mc.GetWithTags("2:message", tags)) //"1-tag", "2-tag"
        printCache(mc.GetWithTags("3:message", tag))  //"1-tag"

        // you can delete all caches which have been set with the tag
        mc.Delete("2-tag")

        printCache(mc.GetWithTags("1:message", tags)) // -> cache miss
        printCache(mc.GetWithTags("2:message", tags)) // -> cache miss
        printCache(mc.GetWithTags("3:message", tag))  // -> "cache3"

        // also you can delete it by the cache key
        mc.Delete("3:message")
        printCache(mc.GetWithTags("3:message", tag))

    }

    func printCache(item *memcache.Item, err error) {
        if err != nil {
            fmt.Println(err.Error())
        } else {
            fmt.Println(string(item.Value))
        }
    }

## gomemcache functionalities
This library is a wrapper of gomemcache. You can use all functionalities in gomemcache.

See https://godoc.org/github.com/bradfitz/gomemcache/memcache
