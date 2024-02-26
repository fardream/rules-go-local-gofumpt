package subdir

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"

	"example.com/localgofumpt/subdir/anothersubofsub"
)

type Subdir struct{}

func (s Subdir) String() string {
	return anothersubofsub.Dump(s) + fmt.Sprint("abc", 5) + spew.Sdump(s)
}
