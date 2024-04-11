package subdir

import (
	"fmt"

	"mymodule.local/subdir/anothersubofsub"

	"github.com/davecgh/go-spew/spew"
)

type Subdir struct{}

func (s Subdir) String() string {
	return anothersubofsub.Dump(s) + fmt.Sprint("abc", 5) + spew.Sdump(s)
}
