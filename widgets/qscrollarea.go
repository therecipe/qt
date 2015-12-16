package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QScrollArea struct {
	QAbstractScrollArea
}

type QScrollArea_ITF interface {
	QAbstractScrollArea_ITF
	QScrollArea_PTR() *QScrollArea
}

func PointerFromQScrollArea(ptr QScrollArea_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScrollArea_PTR().Pointer()
	}
	return nil
}

func NewQScrollAreaFromPointer(ptr unsafe.Pointer) *QScrollArea {
	var n = new(QScrollArea)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QScrollArea_") {
		n.SetObjectName("QScrollArea_" + qt.Identifier())
	}
	return n
}

func (ptr *QScrollArea) QScrollArea_PTR() *QScrollArea {
	return ptr
}

func (ptr *QScrollArea) Alignment() core.Qt__AlignmentFlag {
	defer qt.Recovering("QScrollArea::alignment")

	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QScrollArea_Alignment(ptr.Pointer()))
	}
	return 0
}

func (ptr *QScrollArea) SetAlignment(v core.Qt__AlignmentFlag) {
	defer qt.Recovering("QScrollArea::setAlignment")

	if ptr.Pointer() != nil {
		C.QScrollArea_SetAlignment(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QScrollArea) SetWidget(widget QWidget_ITF) {
	defer qt.Recovering("QScrollArea::setWidget")

	if ptr.Pointer() != nil {
		C.QScrollArea_SetWidget(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QScrollArea) SetWidgetResizable(resizable bool) {
	defer qt.Recovering("QScrollArea::setWidgetResizable")

	if ptr.Pointer() != nil {
		C.QScrollArea_SetWidgetResizable(ptr.Pointer(), C.int(qt.GoBoolToInt(resizable)))
	}
}

func (ptr *QScrollArea) WidgetResizable() bool {
	defer qt.Recovering("QScrollArea::widgetResizable")

	if ptr.Pointer() != nil {
		return C.QScrollArea_WidgetResizable(ptr.Pointer()) != 0
	}
	return false
}

func NewQScrollArea(parent QWidget_ITF) *QScrollArea {
	defer qt.Recovering("QScrollArea::QScrollArea")

	return NewQScrollAreaFromPointer(C.QScrollArea_NewQScrollArea(PointerFromQWidget(parent)))
}

func (ptr *QScrollArea) EnsureVisible(x int, y int, xmargin int, ymargin int) {
	defer qt.Recovering("QScrollArea::ensureVisible")

	if ptr.Pointer() != nil {
		C.QScrollArea_EnsureVisible(ptr.Pointer(), C.int(x), C.int(y), C.int(xmargin), C.int(ymargin))
	}
}

func (ptr *QScrollArea) EnsureWidgetVisible(childWidget QWidget_ITF, xmargin int, ymargin int) {
	defer qt.Recovering("QScrollArea::ensureWidgetVisible")

	if ptr.Pointer() != nil {
		C.QScrollArea_EnsureWidgetVisible(ptr.Pointer(), PointerFromQWidget(childWidget), C.int(xmargin), C.int(ymargin))
	}
}

func (ptr *QScrollArea) FocusNextPrevChild(next bool) bool {
	defer qt.Recovering("QScrollArea::focusNextPrevChild")

	if ptr.Pointer() != nil {
		return C.QScrollArea_FocusNextPrevChild(ptr.Pointer(), C.int(qt.GoBoolToInt(next))) != 0
	}
	return false
}

func (ptr *QScrollArea) ConnectResizeEvent(f func(v *gui.QResizeEvent)) {
	defer qt.Recovering("connect QScrollArea::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QScrollArea) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QScrollArea::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQScrollAreaResizeEvent
func callbackQScrollAreaResizeEvent(ptrName *C.char, v unsafe.Pointer) bool {
	defer qt.Recovering("callback QScrollArea::resizeEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "resizeEvent")
	if signal != nil {
		defer signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(v))
		return true
	}
	return false

}

func (ptr *QScrollArea) ConnectScrollContentsBy(f func(dx int, dy int)) {
	defer qt.Recovering("connect QScrollArea::scrollContentsBy")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "scrollContentsBy", f)
	}
}

func (ptr *QScrollArea) DisconnectScrollContentsBy() {
	defer qt.Recovering("disconnect QScrollArea::scrollContentsBy")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "scrollContentsBy")
	}
}

//export callbackQScrollAreaScrollContentsBy
func callbackQScrollAreaScrollContentsBy(ptrName *C.char, dx C.int, dy C.int) bool {
	defer qt.Recovering("callback QScrollArea::scrollContentsBy")

	var signal = qt.GetSignal(C.GoString(ptrName), "scrollContentsBy")
	if signal != nil {
		defer signal.(func(int, int))(int(dx), int(dy))
		return true
	}
	return false

}

func (ptr *QScrollArea) SizeHint() *core.QSize {
	defer qt.Recovering("QScrollArea::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QScrollArea_SizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScrollArea) TakeWidget() *QWidget {
	defer qt.Recovering("QScrollArea::takeWidget")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QScrollArea_TakeWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScrollArea) Widget() *QWidget {
	defer qt.Recovering("QScrollArea::widget")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QScrollArea_Widget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScrollArea) DestroyQScrollArea() {
	defer qt.Recovering("QScrollArea::~QScrollArea")

	if ptr.Pointer() != nil {
		C.QScrollArea_DestroyQScrollArea(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
