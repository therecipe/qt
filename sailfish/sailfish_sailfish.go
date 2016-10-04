// +build sailfish sailfish_emulator

package sailfish

//#include <stdint.h>
//#include <stdlib.h>
//#include "sailfish_sailfish.h"
import "C"
import (
	"fmt"
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/quick"
	"runtime"
	"strings"
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
	defer qt.Recovering("SailfishApp::application")

	var argvC = C.CString(strings.Join(argv, "|"))
	defer C.free(unsafe.Pointer(argvC))
	var tmpValue = gui.NewQGuiApplicationFromPointer(C.SailfishApp_SailfishApp_Application(C.int(int32(argc)), argvC))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *SailfishApp) Application(argc int, argv []string) *gui.QGuiApplication {
	defer qt.Recovering("SailfishApp::application")

	var argvC = C.CString(strings.Join(argv, "|"))
	defer C.free(unsafe.Pointer(argvC))
	var tmpValue = gui.NewQGuiApplicationFromPointer(C.SailfishApp_SailfishApp_Application(C.int(int32(argc)), argvC))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func SailfishApp_Main(argc int, argv []string) int {
	defer qt.Recovering("SailfishApp::main")

	var argvC = C.CString(strings.Join(argv, "|"))
	defer C.free(unsafe.Pointer(argvC))
	return int(int32(C.SailfishApp_SailfishApp_Main(C.int(int32(argc)), argvC)))
}

func (ptr *SailfishApp) Main(argc int, argv []string) int {
	defer qt.Recovering("SailfishApp::main")

	var argvC = C.CString(strings.Join(argv, "|"))
	defer C.free(unsafe.Pointer(argvC))
	return int(int32(C.SailfishApp_SailfishApp_Main(C.int(int32(argc)), argvC)))
}

func SailfishApp_CreateView() *quick.QQuickView {
	defer qt.Recovering("SailfishApp::createView")

	var tmpValue = quick.NewQQuickViewFromPointer(C.SailfishApp_SailfishApp_CreateView())
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *SailfishApp) CreateView() *quick.QQuickView {
	defer qt.Recovering("SailfishApp::createView")

	var tmpValue = quick.NewQQuickViewFromPointer(C.SailfishApp_SailfishApp_CreateView())
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func SailfishApp_PathTo(filename string) *core.QUrl {
	defer qt.Recovering("SailfishApp::pathTo")

	var filenameC = C.CString(filename)
	defer C.free(unsafe.Pointer(filenameC))
	var tmpValue = core.NewQUrlFromPointer(C.SailfishApp_SailfishApp_PathTo(filenameC))
	runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
	return tmpValue
}

func (ptr *SailfishApp) PathTo(filename string) *core.QUrl {
	defer qt.Recovering("SailfishApp::pathTo")

	var filenameC = C.CString(filename)
	defer C.free(unsafe.Pointer(filenameC))
	var tmpValue = core.NewQUrlFromPointer(C.SailfishApp_SailfishApp_PathTo(filenameC))
	runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
	return tmpValue
}
