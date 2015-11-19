package widgets

//#include "qstatusbar.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QStatusBar struct {
	QWidget
}

type QStatusBar_ITF interface {
	QWidget_ITF
	QStatusBar_PTR() *QStatusBar
}

func PointerFromQStatusBar(ptr QStatusBar_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStatusBar_PTR().Pointer()
	}
	return nil
}

func NewQStatusBarFromPointer(ptr unsafe.Pointer) *QStatusBar {
	var n = new(QStatusBar)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QStatusBar_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QStatusBar) QStatusBar_PTR() *QStatusBar {
	return ptr
}

func (ptr *QStatusBar) IsSizeGripEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QStatusBar_IsSizeGripEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QStatusBar) SetSizeGripEnabled(v bool) {
	if ptr.Pointer() != nil {
		C.QStatusBar_SetSizeGripEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func NewQStatusBar(parent QWidget_ITF) *QStatusBar {
	return NewQStatusBarFromPointer(C.QStatusBar_NewQStatusBar(PointerFromQWidget(parent)))
}

func (ptr *QStatusBar) AddPermanentWidget(widget QWidget_ITF, stretch int) {
	if ptr.Pointer() != nil {
		C.QStatusBar_AddPermanentWidget(ptr.Pointer(), PointerFromQWidget(widget), C.int(stretch))
	}
}

func (ptr *QStatusBar) AddWidget(widget QWidget_ITF, stretch int) {
	if ptr.Pointer() != nil {
		C.QStatusBar_AddWidget(ptr.Pointer(), PointerFromQWidget(widget), C.int(stretch))
	}
}

func (ptr *QStatusBar) ClearMessage() {
	if ptr.Pointer() != nil {
		C.QStatusBar_ClearMessage(ptr.Pointer())
	}
}

func (ptr *QStatusBar) CurrentMessage() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QStatusBar_CurrentMessage(ptr.Pointer()))
	}
	return ""
}

func (ptr *QStatusBar) InsertPermanentWidget(index int, widget QWidget_ITF, stretch int) int {
	if ptr.Pointer() != nil {
		return int(C.QStatusBar_InsertPermanentWidget(ptr.Pointer(), C.int(index), PointerFromQWidget(widget), C.int(stretch)))
	}
	return 0
}

func (ptr *QStatusBar) InsertWidget(index int, widget QWidget_ITF, stretch int) int {
	if ptr.Pointer() != nil {
		return int(C.QStatusBar_InsertWidget(ptr.Pointer(), C.int(index), PointerFromQWidget(widget), C.int(stretch)))
	}
	return 0
}

func (ptr *QStatusBar) ConnectMessageChanged(f func(message string)) {
	if ptr.Pointer() != nil {
		C.QStatusBar_ConnectMessageChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "messageChanged", f)
	}
}

func (ptr *QStatusBar) DisconnectMessageChanged() {
	if ptr.Pointer() != nil {
		C.QStatusBar_DisconnectMessageChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "messageChanged")
	}
}

//export callbackQStatusBarMessageChanged
func callbackQStatusBarMessageChanged(ptrName *C.char, message *C.char) {
	qt.GetSignal(C.GoString(ptrName), "messageChanged").(func(string))(C.GoString(message))
}

func (ptr *QStatusBar) RemoveWidget(widget QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QStatusBar_RemoveWidget(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QStatusBar) ShowMessage(message string, timeout int) {
	if ptr.Pointer() != nil {
		C.QStatusBar_ShowMessage(ptr.Pointer(), C.CString(message), C.int(timeout))
	}
}

func (ptr *QStatusBar) DestroyQStatusBar() {
	if ptr.Pointer() != nil {
		C.QStatusBar_DestroyQStatusBar(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
