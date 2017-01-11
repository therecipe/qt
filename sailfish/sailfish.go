// +build !sailfish,!sailfish_emulator

package sailfish

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "sailfish_sailfish.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/quick"
	"unsafe"
)

type SailfishApp struct {
	ptr unsafe.Pointer
}

type SailfishApp_ITF interface {
	SailfishApp_PTR() *SailfishApp
}

func (p *SailfishApp) SailfishApp_PTR() *SailfishApp {
	return p
}

func (p *SailfishApp) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *SailfishApp) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromSailfishApp(ptr SailfishApp_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.SailfishApp_PTR().Pointer()
	}
	return nil
}

func NewSailfishAppFromPointer(ptr unsafe.Pointer) *SailfishApp {
	var n = new(SailfishApp)
	n.SetPointer(ptr)
	return n
}

func (ptr *SailfishApp) DestroySailfishApp() {
	C.free(ptr.Pointer())
	ptr.SetPointer(nil)
}

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
