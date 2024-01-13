package funcs

import (
	"log"
	"net/url"
	"path"
)

func FormatfileName(fname string) string {
	return path.Base(fname)
}

func FormatMirroredDirName(rawurl string) string {
	parsedURl, err := url.Parse(rawurl)
	if err != nil {
		log.Fatal("[MIRRORING ERROR]", err)
	}
	return parsedURl.Hostname()
}
