package subofsub

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"

	"example.com/localgofumpt/subdir"
)

type SubOfSub subdir.Subdir

func (s *SubOfSub) Spew() string {
	fmt.Println(s)

	return spew.Sdump(s)
}
