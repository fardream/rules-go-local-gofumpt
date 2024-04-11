package subofsub

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"

	"mymodule.local/subdir"
)

type SubOfSub subdir.Subdir

func (s *SubOfSub) Spew() string {
	fmt.Println(s)

	return spew.Sdump(s)
}
