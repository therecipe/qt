package parser

import (
	"encoding/xml"
	"fmt"
	"os"
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
	switch {
	case utils.UseHomeBrew():
		{
			xml.Unmarshal([]byte(utils.Load(filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "therecipe", "qt", "internal", "binding", "files", "docs", "5.7.0", fmt.Sprintf("qt%v.index", s)))), &m)
		}

	case utils.UsePkgConfig():
		{
			xml.Unmarshal([]byte(utils.Load(filepath.Join(utils.QT_DOC_DIR(), fmt.Sprintf("qt%v", s), fmt.Sprintf("qt%v.index", s)))), &m)
		}

	default:
		{
			xml.Unmarshal([]byte(utils.Load(filepath.Join(utils.QT_DIR(), "Docs", "Qt-5.7", fmt.Sprintf("qt%v", s), fmt.Sprintf("qt%v.index", s)))), &m)
		}
	}

	m.Prepare()
	return m
}
