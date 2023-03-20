package swaggerFiles

import (
	"testing"
)

func TestOpenFiles(t *testing.T) {
	for _, filename := range []string{
		"favicon-16x16.png",
		"favicon-32x32.png",
		"index.css",
		"index.html",
		"oauth2-redirect.html",
		"swagger-initializer.js",
		"swagger-ui.css",
		"swagger-ui.js",
		"swagger-ui-bundle.js",
		"swagger-ui-standalone-preset.js",
	} {

		_, err := FS.Open(filename)
		if err != nil {
			t.Fatalf("unable to open %s: %s", filename, err)
		}

	}
}
