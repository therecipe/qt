package parser

import (
	"encoding/xml"
	"fmt"
	"path/filepath"

	"github.com/therecipe/qt/internal/utils"
)

const (
	SIGNAL = "signal"
	SLOT   = "slot"

	IMPURE = "impure"
	PURE   = "pure"

	MOC              = "main"
	PLAIN            = "plain"
	CONSTRUCTOR      = "constructor"
	COPY_CONSTRUCTOR = "copy-constructor"
	MOVE_CONSTRUCTOR = "move-constructor"
	DESTRUCTOR       = "destructor"

	CONNECT    = "Connect"
	DISCONNECT = "Disconnect"
	CALLBACK   = "callback"

	GETTER = "getter"
	SETTER = "setter"

	VOID = "void"

	TILDE = "~"
)

var (
	ClassMap        = make(map[string]*Class)
	SubnamespaceMap = make(map[string]bool)
)

func GetModule(s string) *Module {

	if s == "sailfish" {
		var m = sailfishModule()
		m.Prepare()
		return m
	}

	var m = new(Module)
	if utils.UsePkgConfig() {
		switch utils.LinuxDistro() {
		case "arch":
			{
				xml.Unmarshal([]byte(utils.Load(filepath.Join("/usr", "share", "doc", "qt", fmt.Sprintf("qt%v", s), fmt.Sprintf("qt%v.index", s)))), &m)
			}
		case "ubuntu":
			{
				xml.Unmarshal([]byte(utils.Load(filepath.Join("/usr", "share", "qt5", "doc", fmt.Sprintf("qt%v", s), fmt.Sprintf("qt%v.index", s)))), &m)
			}
		}
	} else {
		xml.Unmarshal([]byte(utils.Load(filepath.Join(utils.QT_DIR(), "Docs", "Qt-5.7", fmt.Sprintf("qt%v", s), fmt.Sprintf("qt%v.index", s)))), &m)
	}

	m.Prepare()
	return m
}
