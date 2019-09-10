package UGlobalHotkey

//#include "ugh.h"
//#include <stdlib.h>
import "C"

import (
	"runtime"
	"unsafe"

	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type UGlobalHotkeys struct {
	QWidget widgets.QWidget
}

type UGlobalHotkeys_ITF interface {
	widgets.QWidget_ITF
	UGlobalHotkeys_PTR() *UGlobalHotkeys
}

func (ptr *UGlobalHotkeys) UGlobalHotkeys_PTR() *UGlobalHotkeys {
	return ptr
}

func (ptr *UGlobalHotkeys) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QWidget.QWidget_PTR().Pointer()
	}
	return nil
}

func (ptr *UGlobalHotkeys) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QWidget.QWidget_PTR().SetPointer(p)
	}
}

func PointerFromUGlobalHotkeys(ptr UGlobalHotkeys_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.UGlobalHotkeys_PTR().Pointer()
	}
	return nil
}

func NewUGlobalHotkeysFromPointer(ptr unsafe.Pointer) (n *UGlobalHotkeys) {
	n = new(UGlobalHotkeys)
	n.SetPointer(ptr)
	return
}

func NewUGlobalHotkeys(parent widgets.QWidget_ITF) *UGlobalHotkeys {
	tmpValue := NewUGlobalHotkeysFromPointer(C.UGlobalHotkeys_NewUGlobalHotkeys(widgets.PointerFromQWidget(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.QWidget.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *UGlobalHotkeys) RegisterHotkey(keySeq string, id uint) {
	if ptr.Pointer() != nil {
		var keySeqC *C.char
		if keySeq != "" {
			keySeqC = C.CString(keySeq)
			defer C.free(unsafe.Pointer(keySeqC))
		}
		C.UGlobalHotkeys_RegisterHotkey(ptr.Pointer(), C.struct_UGlobalHotkeys_PackedString{data: keySeqC, len: C.longlong(len(keySeq))}, C.ulong(uint32(id)))
	}
}

func (ptr *UGlobalHotkeys) UnregisterHotkey(id uint) {
	if ptr.Pointer() != nil {
		C.UGlobalHotkeys_UnregisterHotkey(ptr.Pointer(), C.ulong(uint32(id)))
	}
}

func (ptr *UGlobalHotkeys) UnregisterAllHotkeys() {
	if ptr.Pointer() != nil {
		C.UGlobalHotkeys_UnregisterAllHotkeys(ptr.Pointer())
	}
}

func (ptr *UGlobalHotkeys) DestroyUGlobalHotkeys() {
	if ptr.Pointer() != nil {
		C.UGlobalHotkeys_DestroyUGlobalHotkeys(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackUGlobalHotkeys_Activated
func callbackUGlobalHotkeys_Activated(ptr unsafe.Pointer, id C.ulong) {
	if signal := qt.GetSignal(ptr, "activated"); signal != nil {
		(*(*func(uint))(signal))(uint(uint32(id)))
	}

}

func (ptr *UGlobalHotkeys) ConnectActivated(f func(id uint)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "activated") {
			C.UGlobalHotkeys_ConnectActivated(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "activated"); signal != nil {
			f := func(id uint) {
				(*(*func(uint))(signal))(id)
				f(id)
			}
			qt.ConnectSignal(ptr.Pointer(), "activated", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "activated", unsafe.Pointer(&f))
		}
	}
}

func (ptr *UGlobalHotkeys) DisconnectActivated() {
	if ptr.Pointer() != nil {
		C.UGlobalHotkeys_DisconnectActivated(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "activated")
	}
}

func (ptr *UGlobalHotkeys) Activated(id uint) {
	if ptr.Pointer() != nil {
		C.UGlobalHotkeys_Activated(ptr.Pointer(), C.ulong(uint32(id)))
	}
}
