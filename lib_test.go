package magic_test

import (
	"log"
	"os"
	"runtime"
	"testing"

	"github.com/wenerme/go-magic"

	"github.com/stretchr/testify/assert"
)

func TestFile(t *testing.T) {
	log.Println("version", magic.Version(), "dir", magic.GetDefaultDir())
	// assert.NotEmpty(t, magic.GetDefaultDir())

	mgc := magic.Open(magic.MAGIC_NONE)
	defer mgc.Close()
	log.Println("flags", mgc.GetFlags())
	assert.NoError(t, mgc.Load(""))

	log.Printf("file: %s", mgc.File(os.Args[0]))
	assert.Equal(t, 0, mgc.Errno())
	assert.NoError(t, mgc.Error())

	mgc.SetFlags(magic.MAGIC_MIME | magic.MAGIC_MIME_ENCODING)
	if runtime.GOOS == "darwin" {
		assert.Equal(t, "application/x-mach-binary; charset=binary", mgc.File(os.Args[0]))
	} else {
		assert.Equal(t, "application/x-executable; charset=binary", mgc.File(os.Args[0]))
	}
	assert.Equal(t, 0, mgc.Errno())
	assert.NoError(t, mgc.Error())

	mgc.SetFlags(magic.MAGIC_ERROR)
	assert.Equal(t, "", mgc.File("./not-exists"))
	assert.Equal(t, 2, mgc.Errno())
	assert.EqualError(t, mgc.Error(), "cannot stat `./not-exists' (No such file or directory)")
	log.Println("flags", mgc.GetFlags())
}
