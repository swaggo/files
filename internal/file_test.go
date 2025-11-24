package internal_test

import (
	"context"
	"io/fs"
	"io/ioutil"
	"os"
	"testing"

	"gotest.tools/assert"

	"github.com/swaggo/files/v2/internal"
)

func TestNewWebDAVFile(t *testing.T) {
	t.Parallel()

	webdavFS := internal.NewWebDAVFileSystemFromFS(dist)
	assert.Equal(t, webdavFS.FileSystem, dist)

	file, err := webdavFS.OpenFile(context.Background(), "testfiles/test.txt", 777, os.ModeAppend)
	assert.NilError(t, err)

	_, err = internal.NewWebDAVFile(webdavFS.FileSystem, file)
	assert.NilError(t, err)
}

func TestClose(t *testing.T) {
	t.Parallel()

	webdavFS := internal.NewWebDAVFileSystemFromFS(dist)
	assert.Equal(t, webdavFS.FileSystem, dist)

	file, err := webdavFS.OpenFile(context.Background(), "testfiles/test.txt", 777, os.ModeAppend)
	assert.NilError(t, err)
	assert.NilError(t, file.Close())
}

func TestRead(t *testing.T) {
	t.Parallel()

	webdavFS := internal.NewWebDAVFileSystemFromFS(dist)
	assert.Equal(t, webdavFS.FileSystem, dist)

	file, err := webdavFS.OpenFile(context.Background(), "testfiles/test.txt", 777, os.ModeAppend)
	assert.NilError(t, err)

	b, err := ioutil.ReadAll(file)
	assert.NilError(t, err)
	assert.Equal(t, string(b), "1\n")
}

func TestReaddir(t *testing.T) {
	t.Parallel()

	inside, _ := fs.Sub(dist, "testfiles")
	webdavFS := internal.NewWebDAVFileSystemFromFS(inside)
	assert.Equal(t, webdavFS.FileSystem, inside)

	file, err := webdavFS.OpenFile(context.Background(), "test.txt", 777, os.ModeAppend)
	assert.NilError(t, err)

	fileInfos, err := file.Readdir(2)
	assert.NilError(t, err)
	assert.Equal(t, len(fileInfos), 2)
}

func TestStatFile(t *testing.T) {
	t.Parallel()

	webdavFS := internal.NewWebDAVFileSystemFromFS(dist)
	assert.Equal(t, webdavFS.FileSystem, dist)

	file, err := webdavFS.OpenFile(context.Background(), "testfiles/test.txt", 777, os.ModeAppend)
	assert.NilError(t, err)

	fileInfo, err := os.Stat("testfiles/test.txt")
	assert.NilError(t, err)

	fileInfoActual, err := file.Stat()
	assert.NilError(t, err)

	assert.Equal(t, fileInfo.Size(), fileInfoActual.Size())
}

func TestWrite(t *testing.T) {
	t.Parallel()

	webdavFS := internal.NewWebDAVFileSystemFromFS(dist)
	assert.Equal(t, webdavFS.FileSystem, dist)

	file, err := webdavFS.OpenFile(context.Background(), "testfiles/test.txt", 777, os.ModeAppend)
	assert.NilError(t, err)

	_, err = file.Write([]byte{})
	assert.Equal(t, err, internal.ErrReadOnly)
}

func TestSeek(t *testing.T) {
	t.Parallel()

	webdavFS := internal.NewWebDAVFileSystemFromFS(dist)
	assert.Equal(t, webdavFS.FileSystem, dist)

	file, err := webdavFS.OpenFile(context.Background(), "testfiles/test.txt", 777, os.ModeAppend)
	assert.NilError(t, err)

	res, err := file.Seek(0, 0)
	assert.NilError(t, err)
	assert.Equal(t, res, int64(0))

}
