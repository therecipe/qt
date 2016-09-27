package parser

import (
	"encoding/xml"
	"fmt"
	"runtime"

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

	switch runtime.GOOS {
	case "darwin", "linux":
		{
			xml.Unmarshal([]byte(utils.Load(fmt.Sprintf("%v/Docs/Qt-5.7/qt%v/qt%v.index", utils.QtInstallDir(), s, s))), &m)
		}

	case "windows":
		{
			xml.Unmarshal([]byte(utils.Load(fmt.Sprintf("%v\\Docs\\Qt-5.7\\qt%v\\qt%v.index", utils.QtInstallDir(), s, s))), &m)
		}
	}

	m.Prepare()

	return m
}
