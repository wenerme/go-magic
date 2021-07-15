package magic_test

import (
	"log"
	"os"
	"testing"

	"github.com/wenerme/go-magic"

	"github.com/stretchr/testify/assert"
)

func TestFile(t *testing.T) {
	assert.NotEmpty(t, magic.GetDefaultDir())
	log.Println("version", magic.Version(), "dir", magic.GetDefaultDir())

	mgc := magic.Open(magic.MAGIC_NONE)
	defer mgc.Close()
	log.Println("flags", mgc.GetFlags())
	assert.NoError(t, mgc.Load(""))

	log.Printf("file: %s", mgc.File(os.Args[0]))
	assert.Equal(t, 0, mgc.Errno())
	assert.NoError(t, mgc.Error())

	mgc.SetFlags(magic.MAGIC_MIME | magic.MAGIC_MIME_ENCODING)
	assert.Equal(t, "application/x-mach-binary; charset=binary", mgc.File(os.Args[0]))
	assert.Equal(t, 0, mgc.Errno())
	assert.NoError(t, mgc.Error())

	mgc.SetFlags(magic.MAGIC_ERROR)
	assert.Equal(t, "", mgc.File("./not-exists"))
	assert.Equal(t, 2, mgc.Errno())
	assert.EqualError(t, mgc.Error(), "cannot stat `./not-exists' (No such file or directory)")
	log.Println("flags", mgc.GetFlags())
}
