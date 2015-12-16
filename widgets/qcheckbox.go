package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
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
	for len(n.ObjectName()) < len("QCheckBox_") {
		n.SetObjectName("QCheckBox_" + qt.Identifier())
	}
	return n
}

func (ptr *QCheckBox) QCheckBox_PTR() *QCheckBox {
	return ptr
}

func (ptr *QCheckBox) IsTristate() bool {
	defer qt.Recovering("QCheckBox::isTristate")

	if ptr.Pointer() != nil {
		return C.QCheckBox_IsTristate(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QCheckBox) SetTristate(y bool) {
	defer qt.Recovering("QCheckBox::setTristate")

	if ptr.Pointer() != nil {
		C.QCheckBox_SetTristate(ptr.Pointer(), C.int(qt.GoBoolToInt(y)))
	}
}

func NewQCheckBox(parent QWidget_ITF) *QCheckBox {
	defer qt.Recovering("QCheckBox::QCheckBox")

	return NewQCheckBoxFromPointer(C.QCheckBox_NewQCheckBox(PointerFromQWidget(parent)))
}

func NewQCheckBox2(text string, parent QWidget_ITF) *QCheckBox {
	defer qt.Recovering("QCheckBox::QCheckBox")

	return NewQCheckBoxFromPointer(C.QCheckBox_NewQCheckBox2(C.CString(text), PointerFromQWidget(parent)))
}

func (ptr *QCheckBox) CheckState() core.Qt__CheckState {
	defer qt.Recovering("QCheckBox::checkState")

	if ptr.Pointer() != nil {
		return core.Qt__CheckState(C.QCheckBox_CheckState(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCheckBox) ConnectCheckStateSet(f func()) {
	defer qt.Recovering("connect QCheckBox::checkStateSet")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "checkStateSet", f)
	}
}

func (ptr *QCheckBox) DisconnectCheckStateSet() {
	defer qt.Recovering("disconnect QCheckBox::checkStateSet")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "checkStateSet")
	}
}

//export callbackQCheckBoxCheckStateSet
func callbackQCheckBoxCheckStateSet(ptrName *C.char) bool {
	defer qt.Recovering("callback QCheckBox::checkStateSet")

	var signal = qt.GetSignal(C.GoString(ptrName), "checkStateSet")
	if signal != nil {
		defer signal.(func())()
		return true
	}
	return false

}

func (ptr *QCheckBox) MinimumSizeHint() *core.QSize {
	defer qt.Recovering("QCheckBox::minimumSizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QCheckBox_MinimumSizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QCheckBox) ConnectMouseMoveEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QCheckBox::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QCheckBox) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QCheckBox::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQCheckBoxMouseMoveEvent
func callbackQCheckBoxMouseMoveEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QCheckBox::mouseMoveEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent")
	if signal != nil {
		defer signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QCheckBox) ConnectNextCheckState(f func()) {
	defer qt.Recovering("connect QCheckBox::nextCheckState")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "nextCheckState", f)
	}
}

func (ptr *QCheckBox) DisconnectNextCheckState() {
	defer qt.Recovering("disconnect QCheckBox::nextCheckState")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "nextCheckState")
	}
}

//export callbackQCheckBoxNextCheckState
func callbackQCheckBoxNextCheckState(ptrName *C.char) bool {
	defer qt.Recovering("callback QCheckBox::nextCheckState")

	var signal = qt.GetSignal(C.GoString(ptrName), "nextCheckState")
	if signal != nil {
		defer signal.(func())()
		return true
	}
	return false

}

func (ptr *QCheckBox) ConnectPaintEvent(f func(v *gui.QPaintEvent)) {
	defer qt.Recovering("connect QCheckBox::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QCheckBox) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QCheckBox::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQCheckBoxPaintEvent
func callbackQCheckBoxPaintEvent(ptrName *C.char, v unsafe.Pointer) bool {
	defer qt.Recovering("callback QCheckBox::paintEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "paintEvent")
	if signal != nil {
		defer signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(v))
		return true
	}
	return false

}

func (ptr *QCheckBox) SetCheckState(state core.Qt__CheckState) {
	defer qt.Recovering("QCheckBox::setCheckState")

	if ptr.Pointer() != nil {
		C.QCheckBox_SetCheckState(ptr.Pointer(), C.int(state))
	}
}

func (ptr *QCheckBox) SizeHint() *core.QSize {
	defer qt.Recovering("QCheckBox::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QCheckBox_SizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QCheckBox) ConnectStateChanged(f func(state int)) {
	defer qt.Recovering("connect QCheckBox::stateChanged")

	if ptr.Pointer() != nil {
		C.QCheckBox_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QCheckBox) DisconnectStateChanged() {
	defer qt.Recovering("disconnect QCheckBox::stateChanged")

	if ptr.Pointer() != nil {
		C.QCheckBox_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQCheckBoxStateChanged
func callbackQCheckBoxStateChanged(ptrName *C.char, state C.int) {
	defer qt.Recovering("callback QCheckBox::stateChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "stateChanged")
	if signal != nil {
		signal.(func(int))(int(state))
	}

}

func (ptr *QCheckBox) DestroyQCheckBox() {
	defer qt.Recovering("QCheckBox::~QCheckBox")

	if ptr.Pointer() != nil {
		C.QCheckBox_DestroyQCheckBox(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
