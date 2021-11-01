package middleware

import (
	"context"
	"errors"
	"fmt"
	"runtime/debug"
)

var Recoverf = func(ctx context.Context, p interface{}) error {
	s := fmt.Sprint(p)
	debug.PrintStack()
	fmt.Println(s)
	return errors.New(s)
}
