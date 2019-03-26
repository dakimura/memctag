package memcache

import "unsafe"

// this is the fastest way to convert a byte array to string
// per http://chidakiyo.hatenablog.com/entry/2017/12/12/%E5%BF%98%E3%82%8C%E3%81%8C%E3%81%A1%E3%81%AAGo%E3%81%A7byte%E3%82%92string%E3%81%AB%E5%A4%89%E6%8F%9B%E3%81%99%E3%82%8B%E6%96%B9%E6%B3%95%E3%82%92%E3%83%99%E3%83%B3%E3%83%81%E3%83%9E%E3%83%BC
func convertByteToString(data []byte) string {
	return *(*string)(unsafe.Pointer(&data))
}
