package swaggerFiles

import (
	"embed"
	"io/fs"
)

//go:embed swagger-ui/dist/*
var dist embed.FS

// FS holds embedded swagger ui files
var FS, _ = fs.Sub(dist, "swagger-ui/dist")
