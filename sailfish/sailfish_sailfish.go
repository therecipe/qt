// +build sailfish sailfish_emulator

package sailfish

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "sailfish_sailfish.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/quick"
	"runtime"
	"strings"
	"unsafe"
)

func cGoUnpackString(s C.struct_QtSailfish_PackedString) string {
	if len := int(s.len); len == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
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

func NewSailfishAppFromPointer(ptr unsafe.Pointer) *SailfishApp {
	var n = new(SailfishApp)
	n.SetPointer(ptr)
	return n
}

func (ptr *SailfishApp) DestroySailfishApp() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func SailfishApp_Application(argc int, argv []string) *gui.QGuiApplication {
	var argvC = C.CString(strings.Join(argv, "|"))
	defer C.free(unsafe.Pointer(argvC))
	var tmpValue = gui.NewQGuiApplicationFromPointer(C.SailfishApp_SailfishApp_Application(C.int(int32(argc)), argvC))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *SailfishApp) Application(argc int, argv []string) *gui.QGuiApplication {
	var argvC = C.CString(strings.Join(argv, "|"))
	defer C.free(unsafe.Pointer(argvC))
	var tmpValue = gui.NewQGuiApplicationFromPointer(C.SailfishApp_SailfishApp_Application(C.int(int32(argc)), argvC))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func SailfishApp_Main(argc int, argv []string) int {
	var argvC = C.CString(strings.Join(argv, "|"))
	defer C.free(unsafe.Pointer(argvC))
	return int(int32(C.SailfishApp_SailfishApp_Main(C.int(int32(argc)), argvC)))
}

func (ptr *SailfishApp) Main(argc int, argv []string) int {
	var argvC = C.CString(strings.Join(argv, "|"))
	defer C.free(unsafe.Pointer(argvC))
	return int(int32(C.SailfishApp_SailfishApp_Main(C.int(int32(argc)), argvC)))
}

func SailfishApp_CreateView() *quick.QQuickView {
	var tmpValue = quick.NewQQuickViewFromPointer(C.SailfishApp_SailfishApp_CreateView())
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *SailfishApp) CreateView() *quick.QQuickView {
	var tmpValue = quick.NewQQuickViewFromPointer(C.SailfishApp_SailfishApp_CreateView())
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func SailfishApp_PathTo(filename string) *core.QUrl {
	var filenameC *C.char
	if filename != "" {
		filenameC = C.CString(filename)
		defer C.free(unsafe.Pointer(filenameC))
	}
	var tmpValue = core.NewQUrlFromPointer(C.SailfishApp_SailfishApp_PathTo(C.struct_QtSailfish_PackedString{data: filenameC, len: C.longlong(len(filename))}))
	runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
	return tmpValue
}

func (ptr *SailfishApp) PathTo(filename string) *core.QUrl {
	var filenameC *C.char
	if filename != "" {
		filenameC = C.CString(filename)
		defer C.free(unsafe.Pointer(filenameC))
	}
	var tmpValue = core.NewQUrlFromPointer(C.SailfishApp_SailfishApp_PathTo(C.struct_QtSailfish_PackedString{data: filenameC, len: C.longlong(len(filename))}))
	runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
	return tmpValue
}
