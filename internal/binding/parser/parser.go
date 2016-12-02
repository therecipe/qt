package parser

import (
	"encoding/xml"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/therecipe/qt/internal/utils"
	"github.com/Sirupsen/logrus"
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

func GetModule(s string) (m *Module, err error) {
	s = strings.ToLower(s)
	fields := logrus.Fields{
		"module": fmt.Sprintf("qt/%v", s),
		"func": "GetModule",
	}
	utils.Log.WithFields(fields).Debug("loading module")

	var goPath = utils.MustGoPath()

	if s == "sailfish" {
		m = sailfishModule()
		m.Prepare()
		return m, nil
	}

	m = new(Module)
	switch {
	case utils.UseHomeBrew(), utils.UseMsys2():
		{
			err = xml.Unmarshal([]byte(utils.Load(filepath.Join(goPath, "src", "github.com", "therecipe", "qt", "internal", "binding", "files", "docs", "5.7.0", fmt.Sprintf("qt%v.index", s)))), &m)
		}

	case utils.UsePkgConfig():
		{
			err = xml.Unmarshal([]byte(utils.Load(filepath.Join(utils.QT_DOC_DIR(), fmt.Sprintf("qt%v", s), fmt.Sprintf("qt%v.index", s)))), &m)
		}

	default:
		{
			err = xml.Unmarshal([]byte(utils.Load(filepath.Join(utils.QT_DIR(), "Docs", "Qt-5.7", fmt.Sprintf("qt%v", s), fmt.Sprintf("qt%v.index", s)))), &m)
		}
	}
	if err != nil {
		return nil, err
	}

	m.Prepare()
	return m, nil
}
