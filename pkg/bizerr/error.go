package bizerr

import "fmt"

type Error struct {
	Msg  string
	Code int
}

func (e Error) Error() string {
	return fmt.Sprintf("code: %v, msg: %s", e.Code, e.Msg)
}
