// +build sailfish sailfish_emulator

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
	"runtime"
	"strings"
	"unsafe"
)

func cGoUnpackString(s C.struct_QtSailfish_PackedString) string {
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_QtSailfish_PackedString) []byte {
	if int(s.len) == -1 {
		return []byte(C.GoString(s.data))
	}
	return C.GoBytes(unsafe.Pointer(s.data), C.int(s.len))
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

func (ptr *SailfishApp) DestroySailfishApp() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func SailfishApp_Application(argc int, argv []string) *gui.QGuiApplication {
	argvC := C.CString(strings.Join(argv, "|"))
	defer C.free(unsafe.Pointer(argvC))
	tmpValue := gui.NewQGuiApplicationFromPointer(C.SailfishApp_SailfishApp_Application(C.int(int32(argc)), argvC))
	return tmpValue
}

func (ptr *SailfishApp) Application(argc int, argv []string) *gui.QGuiApplication {
	argvC := C.CString(strings.Join(argv, "|"))
	defer C.free(unsafe.Pointer(argvC))
	tmpValue := gui.NewQGuiApplicationFromPointer(C.SailfishApp_SailfishApp_Application(C.int(int32(argc)), argvC))
	return tmpValue
}

func SailfishApp_Main(argc int, argv []string) int {
	argvC := C.CString(strings.Join(argv, "|"))
	defer C.free(unsafe.Pointer(argvC))
	return int(int32(C.SailfishApp_SailfishApp_Main(C.int(int32(argc)), argvC)))
}

func (ptr *SailfishApp) Main(argc int, argv []string) int {
	argvC := C.CString(strings.Join(argv, "|"))
	defer C.free(unsafe.Pointer(argvC))
	return int(int32(C.SailfishApp_SailfishApp_Main(C.int(int32(argc)), argvC)))
}

func SailfishApp_CreateView() *quick.QQuickView {
	tmpValue := quick.NewQQuickViewFromPointer(C.SailfishApp_SailfishApp_CreateView())
	return tmpValue
}

func (ptr *SailfishApp) CreateView() *quick.QQuickView {
	tmpValue := quick.NewQQuickViewFromPointer(C.SailfishApp_SailfishApp_CreateView())
	return tmpValue
}

func SailfishApp_PathTo(filename string) *core.QUrl {
	var filenameC *C.char
	if filename != "" {
		filenameC = C.CString(filename)
		defer C.free(unsafe.Pointer(filenameC))
	}
	tmpValue := core.NewQUrlFromPointer(C.SailfishApp_SailfishApp_PathTo(C.struct_QtSailfish_PackedString{data: filenameC, len: C.longlong(len(filename))}))
	runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
	return tmpValue
}

func (ptr *SailfishApp) PathTo(filename string) *core.QUrl {
	var filenameC *C.char
	if filename != "" {
		filenameC = C.CString(filename)
		defer C.free(unsafe.Pointer(filenameC))
	}
	tmpValue := core.NewQUrlFromPointer(C.SailfishApp_SailfishApp_PathTo(C.struct_QtSailfish_PackedString{data: filenameC, len: C.longlong(len(filename))}))
	runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
	return tmpValue
}

func SailfishApp_PathToMainQml() *core.QUrl {
	tmpValue := core.NewQUrlFromPointer(C.SailfishApp_SailfishApp_PathToMainQml())
	runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
	return tmpValue
}

func (ptr *SailfishApp) PathToMainQml() *core.QUrl {
	tmpValue := core.NewQUrlFromPointer(C.SailfishApp_SailfishApp_PathToMainQml())
	runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
	return tmpValue
}
