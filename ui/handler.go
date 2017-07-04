package ui

import (
	"net/http"

	"github.com/philips/go-bindata-assetfs"
)

var (
	Handler =http.FileServer(&assetfs.AssetFS{
		Asset:     Asset,
		AssetDir:  AssetDir,
		//Prefix: "swagger-ui/dist",
		// use go-bindata -prefix swagger-ui/dist to trim prefix or here, not both
	})
)
