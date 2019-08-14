package yamloader

import (
	"io/ioutil"
	"log"
)

type Yamloader struct {
	FileName    string
	FileContent []byte
	hasLoaded   bool
}

func (y *Yamloader) SetFileName(filename string) {
	y.hasLoaded = false
	if filename == "" {
		log.Fatalf("[ERR]: %s", "No name provided")
	} else {
		y.FileName = filename
	}
}

func (y *Yamloader) LoadFile(filename string) {
	if y.FileName == "" && filename != "" {
		y.SetFileName(filename)
	}
	if y.FileName != "" {
		file, e := ioutil.ReadFile(filename)
		if e != nil {
			log.Fatalf("[ERR]: %v", e)
		}
		y.FileContent = file
		y.hasLoaded = true
	} else {
		log.Fatalf("[ERR] %s", "No name provided")
	}
}

func (y Yamloader) GetFileContent() []byte {
	return y.FileContent
}
