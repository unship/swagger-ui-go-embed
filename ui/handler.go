package ui

import (
	"net/http"

	"github.com/philips/go-bindata-assetfs"
)

var (
	Handler =http.FileServer(&assetfs.AssetFS{
		// TODO: custom start swagger.json URL
		Asset:     Asset,
		AssetDir:  AssetDir,
		// use CMD go-bindata -prefix swagger-ui/dist to trim prefix or here, not both
		//Prefix: "swagger-ui/dist",
	})
)