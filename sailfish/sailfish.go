// +build !sailfish,!sailfish_emulator

package sailfish

import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/quick"
	"strings"
	"unsafe"
)

func unpackStringList(s string) []string {
	if len(s) == 0 {
		return make([]string, 0)
	}
	return strings.Split(s, "¡¦!")
}

type SailfishApp struct {
	ptr unsafe.Pointer
}

type SailfishApp_ITF interface {
	SailfishApp_PTR() *SailfishApp
}

func (ptr *SailfishApp) SailfishApp_PTR() *SailfishApp {
	return ptr
}

func (ptr *SailfishApp) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *SailfishApp) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromSailfishApp(ptr SailfishApp_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.SailfishApp_PTR().Pointer()
	}
	return nil
}

func NewSailfishAppFromPointer(ptr unsafe.Pointer) (n *SailfishApp) {
	n = new(SailfishApp)
	n.SetPointer(ptr)
	return
}

func (ptr *SailfishApp) DestroySailfishApp() {}

func SailfishApp_Application(argc int, argv []string) *gui.QGuiApplication {

	return nil
}

func (ptr *SailfishApp) Application(argc int, argv []string) *gui.QGuiApplication {

	return nil
}

func SailfishApp_Main(argc int, argv []string) int {

	return 0
}

func (ptr *SailfishApp) Main(argc int, argv []string) int {

	return 0
}

func SailfishApp_CreateView() *quick.QQuickView {

	return nil
}

func (ptr *SailfishApp) CreateView() *quick.QQuickView {

	return nil
}

func SailfishApp_PathTo(filename string) *core.QUrl {

	return nil
}

func (ptr *SailfishApp) PathTo(filename string) *core.QUrl {

	return nil
}

func SailfishApp_PathToMainQml() *core.QUrl {

	return nil
}

func (ptr *SailfishApp) PathToMainQml() *core.QUrl {

	return nil
}

func init() {
	qt.ItfMap["sailfish.SailfishApp_ITF"] = SailfishApp{}
	qt.FuncMap["sailfish.SailfishApp_Application"] = SailfishApp_Application
	qt.FuncMap["sailfish.SailfishApp_Main"] = SailfishApp_Main
	qt.FuncMap["sailfish.SailfishApp_CreateView"] = SailfishApp_CreateView
	qt.FuncMap["sailfish.SailfishApp_PathTo"] = SailfishApp_PathTo
	qt.FuncMap["sailfish.SailfishApp_PathToMainQml"] = SailfishApp_PathToMainQml
}
