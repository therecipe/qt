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

type QStatusBarITF interface {
	QWidgetITF
	QStatusBarPTR() *QStatusBar
}

func PointerFromQStatusBar(ptr QStatusBarITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStatusBarPTR().Pointer()
	}
	return nil
}

func QStatusBarFromPointer(ptr unsafe.Pointer) *QStatusBar {
	var n = new(QStatusBar)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QStatusBar_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QStatusBar) QStatusBarPTR() *QStatusBar {
	return ptr
}

func (ptr *QStatusBar) IsSizeGripEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QStatusBar_IsSizeGripEnabled(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QStatusBar) SetSizeGripEnabled(v bool) {
	if ptr.Pointer() != nil {
		C.QStatusBar_SetSizeGripEnabled(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(v)))
	}
}

func NewQStatusBar(parent QWidgetITF) *QStatusBar {
	return QStatusBarFromPointer(unsafe.Pointer(C.QStatusBar_NewQStatusBar(C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func (ptr *QStatusBar) AddPermanentWidget(widget QWidgetITF, stretch int) {
	if ptr.Pointer() != nil {
		C.QStatusBar_AddPermanentWidget(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(widget)), C.int(stretch))
	}
}

func (ptr *QStatusBar) AddWidget(widget QWidgetITF, stretch int) {
	if ptr.Pointer() != nil {
		C.QStatusBar_AddWidget(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(widget)), C.int(stretch))
	}
}

func (ptr *QStatusBar) ClearMessage() {
	if ptr.Pointer() != nil {
		C.QStatusBar_ClearMessage(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QStatusBar) CurrentMessage() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QStatusBar_CurrentMessage(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QStatusBar) InsertPermanentWidget(index int, widget QWidgetITF, stretch int) int {
	if ptr.Pointer() != nil {
		return int(C.QStatusBar_InsertPermanentWidget(C.QtObjectPtr(ptr.Pointer()), C.int(index), C.QtObjectPtr(PointerFromQWidget(widget)), C.int(stretch)))
	}
	return 0
}

func (ptr *QStatusBar) InsertWidget(index int, widget QWidgetITF, stretch int) int {
	if ptr.Pointer() != nil {
		return int(C.QStatusBar_InsertWidget(C.QtObjectPtr(ptr.Pointer()), C.int(index), C.QtObjectPtr(PointerFromQWidget(widget)), C.int(stretch)))
	}
	return 0
}

func (ptr *QStatusBar) ConnectMessageChanged(f func(message string)) {
	if ptr.Pointer() != nil {
		C.QStatusBar_ConnectMessageChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "messageChanged", f)
	}
}

func (ptr *QStatusBar) DisconnectMessageChanged() {
	if ptr.Pointer() != nil {
		C.QStatusBar_DisconnectMessageChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "messageChanged")
	}
}

//export callbackQStatusBarMessageChanged
func callbackQStatusBarMessageChanged(ptrName *C.char, message *C.char) {
	qt.GetSignal(C.GoString(ptrName), "messageChanged").(func(string))(C.GoString(message))
}

func (ptr *QStatusBar) RemoveWidget(widget QWidgetITF) {
	if ptr.Pointer() != nil {
		C.QStatusBar_RemoveWidget(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(widget)))
	}
}

func (ptr *QStatusBar) ShowMessage(message string, timeout int) {
	if ptr.Pointer() != nil {
		C.QStatusBar_ShowMessage(C.QtObjectPtr(ptr.Pointer()), C.CString(message), C.int(timeout))
	}
}

func (ptr *QStatusBar) DestroyQStatusBar() {
	if ptr.Pointer() != nil {
		C.QStatusBar_DestroyQStatusBar(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
