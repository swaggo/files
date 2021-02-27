package files // import "github.com/swaggo/files"

import (
	"embed"
	"io/fs"
)

//go:embed dist
var dist embed.FS

// UI holds embedded swagger ui files
var UI, _ = fs.Sub(dist, "dist")
