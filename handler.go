package swaggerFiles

import (
	"golang.org/x/net/webdav"

	"github.com/swaggo/files/v2/internal"
)

var (
	// Handler is used to server files through an http.Handler
	Handler *webdav.Handler
)

func init() {
	Handler = &webdav.Handler{
		FileSystem: internal.NewWebDAVFileSystemFromFS(FS),
		LockSystem: webdav.NewMemLS(),
	}
}
