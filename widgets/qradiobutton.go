package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QRadioButton struct {
	QAbstractButton
}

type QRadioButton_ITF interface {
	QAbstractButton_ITF
	QRadioButton_PTR() *QRadioButton
}

func PointerFromQRadioButton(ptr QRadioButton_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRadioButton_PTR().Pointer()
	}
	return nil
}

func NewQRadioButtonFromPointer(ptr unsafe.Pointer) *QRadioButton {
	var n = new(QRadioButton)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QRadioButton_") {
		n.SetObjectName("QRadioButton_" + qt.Identifier())
	}
	return n
}

func (ptr *QRadioButton) QRadioButton_PTR() *QRadioButton {
	return ptr
}

func NewQRadioButton(parent QWidget_ITF) *QRadioButton {
	defer qt.Recovering("QRadioButton::QRadioButton")

	return NewQRadioButtonFromPointer(C.QRadioButton_NewQRadioButton(PointerFromQWidget(parent)))
}

func NewQRadioButton2(text string, parent QWidget_ITF) *QRadioButton {
	defer qt.Recovering("QRadioButton::QRadioButton")

	return NewQRadioButtonFromPointer(C.QRadioButton_NewQRadioButton2(C.CString(text), PointerFromQWidget(parent)))
}

func (ptr *QRadioButton) MinimumSizeHint() *core.QSize {
	defer qt.Recovering("QRadioButton::minimumSizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QRadioButton_MinimumSizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QRadioButton) ConnectMouseMoveEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QRadioButton::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QRadioButton) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QRadioButton::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQRadioButtonMouseMoveEvent
func callbackQRadioButtonMouseMoveEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QRadioButton::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QRadioButton) ConnectPaintEvent(f func(v *gui.QPaintEvent)) {
	defer qt.Recovering("connect QRadioButton::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QRadioButton) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QRadioButton::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQRadioButtonPaintEvent
func callbackQRadioButtonPaintEvent(ptrName *C.char, v unsafe.Pointer) bool {
	defer qt.Recovering("callback QRadioButton::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(v))
		return true
	}
	return false

}

func (ptr *QRadioButton) SizeHint() *core.QSize {
	defer qt.Recovering("QRadioButton::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QRadioButton_SizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QRadioButton) DestroyQRadioButton() {
	defer qt.Recovering("QRadioButton::~QRadioButton")

	if ptr.Pointer() != nil {
		C.QRadioButton_DestroyQRadioButton(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
