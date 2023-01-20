#!/bin/sh

set -e

GEN_DIR="./gen"
EMBED_FILE="$GEN_DIR/embed.go"

mkdir -p "$GEN_DIR"

cat <<EOF > "$EMBED_FILE"
package gen

import (
	"embed"
	"mime"
	"net/http"
)

//go:embed docs/*
var openAPI embed.FS

func init() {
	if err := mime.AddExtensionType(".svg", "image/svg+xml"); err != nil {
		panic(err)
	}
}

func MustGetOpenAPIHandler() http.Handler {
	return http.FileServer(http.FS(openAPI))
}
EOF