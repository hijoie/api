package cachekey

import (
	"fmt"
	"strings"
)

func GetCacheKey(a ...interface{}) string {
	var ss []string
	for _, v := range a {
		ss = append(ss,fmt.Sprint(v))
	}
	return strings.Join(ss,":")
}

const (
	UserVeriCode = "user:verifycode:"
)
