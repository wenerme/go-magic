package magic_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/wenerme/go-magic"

	"github.com/stretchr/testify/assert"
)

func TestFile(t *testing.T) {
	fmt.Println("version", magic.Version())

	mgc := magic.Open(magic.MAGIC_NONE)
	defer mgc.Close()
	fmt.Println(mgc.GetFlags())
	assert.NoError(t, mgc.Load(""))
	fmt.Printf("file: %s - error %#v errno %v\n", mgc.File(os.Args[0]), mgc.Error(), mgc.Errno())
	mgc.SetFlags(magic.MAGIC_MIME | magic.MAGIC_MIME_ENCODING)
	fmt.Printf("file: %s - error %#v errno %v\n", mgc.File(os.Args[0]), mgc.Error(), mgc.Errno())
}
