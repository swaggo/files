package internal_test

import (
	"context"
	"embed"
	"io/ioutil"
	"os"
	"testing"

	"gotest.tools/assert"

	"github.com/swaggo/files/v2/internal"
)

//go:embed testfiles
var dist embed.FS

func TestNewWebDAVFileSystemFromFS(t *testing.T) {
	t.Parallel()

	webdavFS := internal.NewWebDAVFileSystemFromFS(dist)
	assert.Equal(t, webdavFS.FileSystem, dist)
}

func TestOpenFile(t *testing.T) {
	t.Parallel()

	webdavFS := internal.NewWebDAVFileSystemFromFS(dist)
	file, err := webdavFS.OpenFile(context.Background(), "testfiles/test.txt", 777, os.ModeAppend)
	assert.NilError(t, err)

	b, err := ioutil.ReadAll(file)
	assert.NilError(t, err)
	assert.Equal(t, string(b), "1\n")

	_, err = webdavFS.OpenFile(context.Background(), "testfiles/nofile.txt", 777, os.ModeAppend)
	assert.ErrorContains(t, err, "file does not exist")
}

func TestStat(t *testing.T) {
	t.Parallel()

	webdavFS := internal.NewWebDAVFileSystemFromFS(dist)
	fileInfo, err := webdavFS.Stat(context.Background(), "testfiles/test.txt")
	assert.NilError(t, err)
	assert.Equal(t, fileInfo.Size(), int64(2))

	_, err = webdavFS.Stat(context.Background(), "testfiles/nofile.txt")
	assert.ErrorContains(t, err, "file does not exist")
}

func TestRemoveAll(t *testing.T) {
	t.Parallel()

	webdavFS := internal.NewWebDAVFileSystemFromFS(dist)
	err := webdavFS.RemoveAll(context.Background(), "")
	assert.Equal(t, err, internal.ErrReadOnly)
}

func TestMkdir(t *testing.T) {
	t.Parallel()

	webdavFS := internal.NewWebDAVFileSystemFromFS(dist)
	err := webdavFS.Mkdir(context.Background(), "", os.ModeAppend)
	assert.Equal(t, err, internal.ErrReadOnly)
}

func TestRename(t *testing.T) {
	t.Parallel()

	webdavFS := internal.NewWebDAVFileSystemFromFS(dist)
	err := webdavFS.Rename(context.Background(), "", "")
	assert.Equal(t, err, internal.ErrReadOnly)
}
