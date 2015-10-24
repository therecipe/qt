package widgets

//#include "qcheckbox.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QCheckBox struct {
	QAbstractButton
}

type QCheckBoxITF interface {
	QAbstractButtonITF
	QCheckBoxPTR() *QCheckBox
}

func PointerFromQCheckBox(ptr QCheckBoxITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCheckBoxPTR().Pointer()
	}
	return nil
}

func QCheckBoxFromPointer(ptr unsafe.Pointer) *QCheckBox {
	var n = new(QCheckBox)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QCheckBox_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QCheckBox) QCheckBoxPTR() *QCheckBox {
	return ptr
}

func (ptr *QCheckBox) IsTristate() bool {
	if ptr.Pointer() != nil {
		return C.QCheckBox_IsTristate(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QCheckBox) SetTristate(y bool) {
	if ptr.Pointer() != nil {
		C.QCheckBox_SetTristate(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(y)))
	}
}

func NewQCheckBox(parent QWidgetITF) *QCheckBox {
	return QCheckBoxFromPointer(unsafe.Pointer(C.QCheckBox_NewQCheckBox(C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func NewQCheckBox2(text string, parent QWidgetITF) *QCheckBox {
	return QCheckBoxFromPointer(unsafe.Pointer(C.QCheckBox_NewQCheckBox2(C.CString(text), C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func (ptr *QCheckBox) CheckState() core.Qt__CheckState {
	if ptr.Pointer() != nil {
		return core.Qt__CheckState(C.QCheckBox_CheckState(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QCheckBox) SetCheckState(state core.Qt__CheckState) {
	if ptr.Pointer() != nil {
		C.QCheckBox_SetCheckState(C.QtObjectPtr(ptr.Pointer()), C.int(state))
	}
}

func (ptr *QCheckBox) ConnectStateChanged(f func(state int)) {
	if ptr.Pointer() != nil {
		C.QCheckBox_ConnectStateChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QCheckBox) DisconnectStateChanged() {
	if ptr.Pointer() != nil {
		C.QCheckBox_DisconnectStateChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQCheckBoxStateChanged
func callbackQCheckBoxStateChanged(ptrName *C.char, state C.int) {
	qt.GetSignal(C.GoString(ptrName), "stateChanged").(func(int))(int(state))
}

func (ptr *QCheckBox) DestroyQCheckBox() {
	if ptr.Pointer() != nil {
		C.QCheckBox_DestroyQCheckBox(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
