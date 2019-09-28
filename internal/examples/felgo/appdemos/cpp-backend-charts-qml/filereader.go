package main

import (
	"github.com/therecipe/qt/core"
)

type FileReader struct {
	core.QObject

	_ func(fileName string) string `slot:"readFile,auto"`
}

func (fr *FileReader) readFile(fileName string) string {
	var content string
	file := core.NewQFile2(fileName)
	if file.Open(core.QIODevice__ReadOnly) {
		stream := core.NewQTextStream2(file)
		content = stream.ReadAll()
	}
	return content
}
