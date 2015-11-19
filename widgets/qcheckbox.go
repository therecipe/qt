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

type QCheckBox_ITF interface {
	QAbstractButton_ITF
	QCheckBox_PTR() *QCheckBox
}

func PointerFromQCheckBox(ptr QCheckBox_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCheckBox_PTR().Pointer()
	}
	return nil
}

func NewQCheckBoxFromPointer(ptr unsafe.Pointer) *QCheckBox {
	var n = new(QCheckBox)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QCheckBox_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QCheckBox) QCheckBox_PTR() *QCheckBox {
	return ptr
}

func (ptr *QCheckBox) IsTristate() bool {
	if ptr.Pointer() != nil {
		return C.QCheckBox_IsTristate(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QCheckBox) SetTristate(y bool) {
	if ptr.Pointer() != nil {
		C.QCheckBox_SetTristate(ptr.Pointer(), C.int(qt.GoBoolToInt(y)))
	}
}

func NewQCheckBox(parent QWidget_ITF) *QCheckBox {
	return NewQCheckBoxFromPointer(C.QCheckBox_NewQCheckBox(PointerFromQWidget(parent)))
}

func NewQCheckBox2(text string, parent QWidget_ITF) *QCheckBox {
	return NewQCheckBoxFromPointer(C.QCheckBox_NewQCheckBox2(C.CString(text), PointerFromQWidget(parent)))
}

func (ptr *QCheckBox) CheckState() core.Qt__CheckState {
	if ptr.Pointer() != nil {
		return core.Qt__CheckState(C.QCheckBox_CheckState(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCheckBox) SetCheckState(state core.Qt__CheckState) {
	if ptr.Pointer() != nil {
		C.QCheckBox_SetCheckState(ptr.Pointer(), C.int(state))
	}
}

func (ptr *QCheckBox) ConnectStateChanged(f func(state int)) {
	if ptr.Pointer() != nil {
		C.QCheckBox_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QCheckBox) DisconnectStateChanged() {
	if ptr.Pointer() != nil {
		C.QCheckBox_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQCheckBoxStateChanged
func callbackQCheckBoxStateChanged(ptrName *C.char, state C.int) {
	qt.GetSignal(C.GoString(ptrName), "stateChanged").(func(int))(int(state))
}

func (ptr *QCheckBox) DestroyQCheckBox() {
	if ptr.Pointer() != nil {
		C.QCheckBox_DestroyQCheckBox(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
