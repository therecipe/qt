package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/gui"
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
	for len(n.ObjectName()) < len("QStatusBar_") {
		n.SetObjectName("QStatusBar_" + qt.Identifier())
	}
	return n
}

func (ptr *QStatusBar) QStatusBar_PTR() *QStatusBar {
	return ptr
}

func (ptr *QStatusBar) IsSizeGripEnabled() bool {
	defer qt.Recovering("QStatusBar::isSizeGripEnabled")

	if ptr.Pointer() != nil {
		return C.QStatusBar_IsSizeGripEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QStatusBar) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QStatusBar::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QStatusBar) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QStatusBar::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQStatusBarPaintEvent
func callbackQStatusBarPaintEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QStatusBar::paintEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "paintEvent")
	if signal != nil {
		defer signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QStatusBar) SetSizeGripEnabled(v bool) {
	defer qt.Recovering("QStatusBar::setSizeGripEnabled")

	if ptr.Pointer() != nil {
		C.QStatusBar_SetSizeGripEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func NewQStatusBar(parent QWidget_ITF) *QStatusBar {
	defer qt.Recovering("QStatusBar::QStatusBar")

	return NewQStatusBarFromPointer(C.QStatusBar_NewQStatusBar(PointerFromQWidget(parent)))
}

func (ptr *QStatusBar) AddPermanentWidget(widget QWidget_ITF, stretch int) {
	defer qt.Recovering("QStatusBar::addPermanentWidget")

	if ptr.Pointer() != nil {
		C.QStatusBar_AddPermanentWidget(ptr.Pointer(), PointerFromQWidget(widget), C.int(stretch))
	}
}

func (ptr *QStatusBar) AddWidget(widget QWidget_ITF, stretch int) {
	defer qt.Recovering("QStatusBar::addWidget")

	if ptr.Pointer() != nil {
		C.QStatusBar_AddWidget(ptr.Pointer(), PointerFromQWidget(widget), C.int(stretch))
	}
}

func (ptr *QStatusBar) ClearMessage() {
	defer qt.Recovering("QStatusBar::clearMessage")

	if ptr.Pointer() != nil {
		C.QStatusBar_ClearMessage(ptr.Pointer())
	}
}

func (ptr *QStatusBar) CurrentMessage() string {
	defer qt.Recovering("QStatusBar::currentMessage")

	if ptr.Pointer() != nil {
		return C.GoString(C.QStatusBar_CurrentMessage(ptr.Pointer()))
	}
	return ""
}

func (ptr *QStatusBar) InsertPermanentWidget(index int, widget QWidget_ITF, stretch int) int {
	defer qt.Recovering("QStatusBar::insertPermanentWidget")

	if ptr.Pointer() != nil {
		return int(C.QStatusBar_InsertPermanentWidget(ptr.Pointer(), C.int(index), PointerFromQWidget(widget), C.int(stretch)))
	}
	return 0
}

func (ptr *QStatusBar) InsertWidget(index int, widget QWidget_ITF, stretch int) int {
	defer qt.Recovering("QStatusBar::insertWidget")

	if ptr.Pointer() != nil {
		return int(C.QStatusBar_InsertWidget(ptr.Pointer(), C.int(index), PointerFromQWidget(widget), C.int(stretch)))
	}
	return 0
}

func (ptr *QStatusBar) ConnectMessageChanged(f func(message string)) {
	defer qt.Recovering("connect QStatusBar::messageChanged")

	if ptr.Pointer() != nil {
		C.QStatusBar_ConnectMessageChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "messageChanged", f)
	}
}

func (ptr *QStatusBar) DisconnectMessageChanged() {
	defer qt.Recovering("disconnect QStatusBar::messageChanged")

	if ptr.Pointer() != nil {
		C.QStatusBar_DisconnectMessageChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "messageChanged")
	}
}

//export callbackQStatusBarMessageChanged
func callbackQStatusBarMessageChanged(ptrName *C.char, message *C.char) {
	defer qt.Recovering("callback QStatusBar::messageChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "messageChanged")
	if signal != nil {
		signal.(func(string))(C.GoString(message))
	}

}

func (ptr *QStatusBar) RemoveWidget(widget QWidget_ITF) {
	defer qt.Recovering("QStatusBar::removeWidget")

	if ptr.Pointer() != nil {
		C.QStatusBar_RemoveWidget(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QStatusBar) ConnectResizeEvent(f func(e *gui.QResizeEvent)) {
	defer qt.Recovering("connect QStatusBar::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QStatusBar) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QStatusBar::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQStatusBarResizeEvent
func callbackQStatusBarResizeEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QStatusBar::resizeEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "resizeEvent")
	if signal != nil {
		defer signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QStatusBar) ConnectShowEvent(f func(v *gui.QShowEvent)) {
	defer qt.Recovering("connect QStatusBar::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QStatusBar) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QStatusBar::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQStatusBarShowEvent
func callbackQStatusBarShowEvent(ptrName *C.char, v unsafe.Pointer) bool {
	defer qt.Recovering("callback QStatusBar::showEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "showEvent")
	if signal != nil {
		defer signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(v))
		return true
	}
	return false

}

func (ptr *QStatusBar) ShowMessage(message string, timeout int) {
	defer qt.Recovering("QStatusBar::showMessage")

	if ptr.Pointer() != nil {
		C.QStatusBar_ShowMessage(ptr.Pointer(), C.CString(message), C.int(timeout))
	}
}

func (ptr *QStatusBar) DestroyQStatusBar() {
	defer qt.Recovering("QStatusBar::~QStatusBar")

	if ptr.Pointer() != nil {
		C.QStatusBar_DestroyQStatusBar(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
