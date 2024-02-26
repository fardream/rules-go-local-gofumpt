package anothersubofsub

import "github.com/davecgh/go-spew/spew"

func Dump(s any) string {
	return spew.Sdump(s)
}
