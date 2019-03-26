package main

import (
	"fmt"
	memcacheWithTag "github.com/dakimura/memctag/memcache"
	"github.com/bradfitz/gomemcache/memcache"
)
// 	memcacheWithTag "github.com/dakimura/memctag/memcache"

//func main() {
//	mc := memcacheWithTag.NewClient("127.0.0.1:11211")
//
//	item := memcache.Item{Key: "1:message", Value: []byte("cache1")}
//	item2 := memcache.Item{Key: "2:message", Value: []byte("cache2")}
//	item3 := memcache.Item{Key: "3:message", Value: []byte("cache3")}
//	item4 := memcache.Item{Key: "4:message", Value: []byte("cache4")}
//	item5 := memcache.Item{Key: "5:message", Value: []byte("cache5")}
//	item6 := memcache.Item{Key: "6:message", Value: []byte("cache6")}
//	item7 := memcache.Item{Key: "7:message", Value: []byte("cache7")}
//	item8 := memcache.Item{Key: "8:message", Value: []byte("cache8")}
//	item9 := memcache.Item{Key: "9:message", Value: []byte("cache9")}
//	item10 := memcache.Item{Key: "10:message", Value: []byte("cache10")}
//
//	tags := []string{"1-tag"}
//
//	mc.SetWithTags(&item, tags)
//	mc.SetWithTags(&item2, tags)
//	mc.SetWithTags(&item3, tags)
//	mc.SetWithTags(&item4, tags)
//	mc.SetWithTags(&item5, tags)
//	mc.SetWithTags(&item6, tags)
//	mc.SetWithTags(&item7, tags)
//	mc.SetWithTags(&item8, tags)
//	mc.SetWithTags(&item9, tags)
//	mc.SetWithTags(&item10, tags)
//
//	printCache(mc.GetWithTags("1:message", tags))
//	printCache(mc.GetWithTags("2:message", tags))
//	printCache(mc.GetWithTags("3:message", tags))
//	printCache(mc.GetWithTags("4:message", tags))
//	printCache(mc.GetWithTags("5:message", tags))
//	printCache(mc.GetWithTags("6:message", tags))
//	printCache(mc.GetWithTags("7:message", tags))
//	printCache(mc.GetWithTags("8:message", tags))
//	printCache(mc.GetWithTags("9:message", tags))
//	printCache(mc.GetWithTags("10:message", tags))
//}



func main() {
	mc := memcacheWithTag.NewClient("127.0.0.1:11211")

	tags := []string{"1-tag", "2-tag"}
	tag := []string{"1-tag"}

	item := memcache.Item{Key: "1:message", Value: []byte("cache1")}
	item2 := memcache.Item{Key: "2:message", Value: []byte("cache2")}
	item3 := memcache.Item{Key: "3:message", Value: []byte("cache3")}

	mc.SetWithTags(&item, []string{"1-tag", "2-tag"})
	mc.SetWithTags(&item2, []string{"1-tag", "2-tag"})
	mc.SetWithTags(&item3, []string{"1-tag"})

	//fmt.Println("-------")
	mc.Delete("2-tag")

	printCache(mc.GetWithTags("1:message", tags)) //"1-tag", "2-tag"
	printCache(mc.GetWithTags("2:message", tags)) //"1-tag", "2-tag"
	printCache(mc.GetWithTags("3:message", tag))  //"1-tag"

	//fmt.Println("-------")
	//
	//mc.Delete("3:message")
	//printCache(mc.GetWithTags("3:message", tag))
}



func printCache(item *memcache.Item, err error) {
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(string(item.Value))
	}
}
