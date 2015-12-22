package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QSplitterHandle struct {
	QWidget
}

type QSplitterHandle_ITF interface {
	QWidget_ITF
	QSplitterHandle_PTR() *QSplitterHandle
}

func PointerFromQSplitterHandle(ptr QSplitterHandle_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSplitterHandle_PTR().Pointer()
	}
	return nil
}

func NewQSplitterHandleFromPointer(ptr unsafe.Pointer) *QSplitterHandle {
	var n = new(QSplitterHandle)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QSplitterHandle_") {
		n.SetObjectName("QSplitterHandle_" + qt.Identifier())
	}
	return n
}

func (ptr *QSplitterHandle) QSplitterHandle_PTR() *QSplitterHandle {
	return ptr
}

func NewQSplitterHandle(orientation core.Qt__Orientation, parent QSplitter_ITF) *QSplitterHandle {
	defer qt.Recovering("QSplitterHandle::QSplitterHandle")

	return NewQSplitterHandleFromPointer(C.QSplitterHandle_NewQSplitterHandle(C.int(orientation), PointerFromQSplitter(parent)))
}

func (ptr *QSplitterHandle) ConnectMouseMoveEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QSplitterHandle::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QSplitterHandle) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QSplitterHandle::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQSplitterHandleMouseMoveEvent
func callbackQSplitterHandleMouseMoveEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QSplitterHandle::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QSplitterHandle) ConnectMousePressEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QSplitterHandle::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QSplitterHandle) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QSplitterHandle::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQSplitterHandleMousePressEvent
func callbackQSplitterHandleMousePressEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QSplitterHandle::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QSplitterHandle) ConnectMouseReleaseEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QSplitterHandle::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QSplitterHandle) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QSplitterHandle::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQSplitterHandleMouseReleaseEvent
func callbackQSplitterHandleMouseReleaseEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QSplitterHandle::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QSplitterHandle) OpaqueResize() bool {
	defer qt.Recovering("QSplitterHandle::opaqueResize")

	if ptr.Pointer() != nil {
		return C.QSplitterHandle_OpaqueResize(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSplitterHandle) Orientation() core.Qt__Orientation {
	defer qt.Recovering("QSplitterHandle::orientation")

	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QSplitterHandle_Orientation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSplitterHandle) ConnectPaintEvent(f func(v *gui.QPaintEvent)) {
	defer qt.Recovering("connect QSplitterHandle::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QSplitterHandle) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QSplitterHandle::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQSplitterHandlePaintEvent
func callbackQSplitterHandlePaintEvent(ptrName *C.char, v unsafe.Pointer) bool {
	defer qt.Recovering("callback QSplitterHandle::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(v))
		return true
	}
	return false

}

func (ptr *QSplitterHandle) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QSplitterHandle::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QSplitterHandle) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QSplitterHandle::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQSplitterHandleResizeEvent
func callbackQSplitterHandleResizeEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSplitterHandle::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QSplitterHandle) SetOrientation(orientation core.Qt__Orientation) {
	defer qt.Recovering("QSplitterHandle::setOrientation")

	if ptr.Pointer() != nil {
		C.QSplitterHandle_SetOrientation(ptr.Pointer(), C.int(orientation))
	}
}

func (ptr *QSplitterHandle) SizeHint() *core.QSize {
	defer qt.Recovering("QSplitterHandle::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QSplitterHandle_SizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSplitterHandle) Splitter() *QSplitter {
	defer qt.Recovering("QSplitterHandle::splitter")

	if ptr.Pointer() != nil {
		return NewQSplitterFromPointer(C.QSplitterHandle_Splitter(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSplitterHandle) DestroyQSplitterHandle() {
	defer qt.Recovering("QSplitterHandle::~QSplitterHandle")

	if ptr.Pointer() != nil {
		C.QSplitterHandle_DestroyQSplitterHandle(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
