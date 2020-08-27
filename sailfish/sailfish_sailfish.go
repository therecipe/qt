// +build sailfish sailfish_emulator

package sailfish

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/internal"
	"github.com/therecipe/qt/quick"
	"unsafe"
)

type SailfishApp struct {
	internal.Internal
}

type SailfishApp_ITF interface {
	SailfishApp_PTR() *SailfishApp
}

func (ptr *SailfishApp) SailfishApp_PTR() *SailfishApp {
	return ptr
}

func (ptr *SailfishApp) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *SailfishApp) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromSailfishApp(ptr SailfishApp_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.SailfishApp_PTR().Pointer()
	}
	return nil
}

func (n *SailfishApp) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewSailfishAppFromPointer(ptr unsafe.Pointer) (n *SailfishApp) {
	n = new(SailfishApp)
	n.InitFromInternal(uintptr(ptr), "sailfish.SailfishApp")
	return
}

func (ptr *SailfishApp) DestroySailfishApp() {
}

func SailfishApp_Application(argc int, argv []string) *gui.QGuiApplication {

	return internal.CallLocalFunction([]interface{}{"", "", "sailfish.SailfishApp_Application", "", argc, argv}).(*gui.QGuiApplication)
}

func (ptr *SailfishApp) Application(argc int, argv []string) *gui.QGuiApplication {

	return internal.CallLocalFunction([]interface{}{"", "", "sailfish.SailfishApp_Application", "", argc, argv}).(*gui.QGuiApplication)
}

func SailfishApp_Main(argc int, argv []string) int {

	return int(internal.CallLocalFunction([]interface{}{"", "", "sailfish.SailfishApp_Main", "", argc, argv}).(float64))
}

func (ptr *SailfishApp) Main(argc int, argv []string) int {

	return int(internal.CallLocalFunction([]interface{}{"", "", "sailfish.SailfishApp_Main", "", argc, argv}).(float64))
}

func SailfishApp_CreateView() *quick.QQuickView {

	return internal.CallLocalFunction([]interface{}{"", "", "sailfish.SailfishApp_CreateView", ""}).(*quick.QQuickView)
}

func (ptr *SailfishApp) CreateView() *quick.QQuickView {

	return internal.CallLocalFunction([]interface{}{"", "", "sailfish.SailfishApp_CreateView", ""}).(*quick.QQuickView)
}

func SailfishApp_PathTo(filename string) *core.QUrl {

	return internal.CallLocalFunction([]interface{}{"", "", "sailfish.SailfishApp_PathTo", "", filename}).(*core.QUrl)
}

func (ptr *SailfishApp) PathTo(filename string) *core.QUrl {

	return internal.CallLocalFunction([]interface{}{"", "", "sailfish.SailfishApp_PathTo", "", filename}).(*core.QUrl)
}

func SailfishApp_PathToMainQml() *core.QUrl {

	return internal.CallLocalFunction([]interface{}{"", "", "sailfish.SailfishApp_PathToMainQml", ""}).(*core.QUrl)
}

func (ptr *SailfishApp) PathToMainQml() *core.QUrl {

	return internal.CallLocalFunction([]interface{}{"", "", "sailfish.SailfishApp_PathToMainQml", ""}).(*core.QUrl)
}

func init() {
	internal.ConstructorTable["sailfish.SailfishApp"] = NewSailfishAppFromPointer
}
