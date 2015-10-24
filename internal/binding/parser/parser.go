package parser

import (
	"encoding/xml"
	"fmt"
	"runtime"

	"github.com/therecipe/qt/internal/utils"
)

var (
	AbstractMap     = make(map[string]bool)
	ClassMap        = make(map[string]*Class)
	SubnamespaceMap = make(map[string]bool)
)

func GetModule(s string) (m *Module) {

	switch runtime.GOOS {
	case "darwin", "linux":
		xml.Unmarshal([]byte(utils.Load(fmt.Sprintf("/usr/local/Qt5.5.1/Docs/Qt-5.5/qt%v/qt%v.index", s, s))), &m)

	case "windows":
		xml.Unmarshal([]byte(utils.Load(fmt.Sprintf("C:\\Qt\\Qt5.5.1\\Docs\\Qt-5.5\\qt%v\\qt%v.index", s, s))), &m)
	}

	m.prepare()

	return m
}
