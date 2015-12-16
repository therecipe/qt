package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QSplitter struct {
	QFrame
}

type QSplitter_ITF interface {
	QFrame_ITF
	QSplitter_PTR() *QSplitter
}

func PointerFromQSplitter(ptr QSplitter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSplitter_PTR().Pointer()
	}
	return nil
}

func NewQSplitterFromPointer(ptr unsafe.Pointer) *QSplitter {
	var n = new(QSplitter)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QSplitter_") {
		n.SetObjectName("QSplitter_" + qt.Identifier())
	}
	return n
}

func (ptr *QSplitter) QSplitter_PTR() *QSplitter {
	return ptr
}

func (ptr *QSplitter) ChildrenCollapsible() bool {
	defer qt.Recovering("QSplitter::childrenCollapsible")

	if ptr.Pointer() != nil {
		return C.QSplitter_ChildrenCollapsible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSplitter) HandleWidth() int {
	defer qt.Recovering("QSplitter::handleWidth")

	if ptr.Pointer() != nil {
		return int(C.QSplitter_HandleWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSplitter) IndexOf(widget QWidget_ITF) int {
	defer qt.Recovering("QSplitter::indexOf")

	if ptr.Pointer() != nil {
		return int(C.QSplitter_IndexOf(ptr.Pointer(), PointerFromQWidget(widget)))
	}
	return 0
}

func (ptr *QSplitter) OpaqueResize() bool {
	defer qt.Recovering("QSplitter::opaqueResize")

	if ptr.Pointer() != nil {
		return C.QSplitter_OpaqueResize(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSplitter) Orientation() core.Qt__Orientation {
	defer qt.Recovering("QSplitter::orientation")

	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QSplitter_Orientation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSplitter) SetChildrenCollapsible(v bool) {
	defer qt.Recovering("QSplitter::setChildrenCollapsible")

	if ptr.Pointer() != nil {
		C.QSplitter_SetChildrenCollapsible(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QSplitter) SetHandleWidth(v int) {
	defer qt.Recovering("QSplitter::setHandleWidth")

	if ptr.Pointer() != nil {
		C.QSplitter_SetHandleWidth(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QSplitter) SetOpaqueResize(opaque bool) {
	defer qt.Recovering("QSplitter::setOpaqueResize")

	if ptr.Pointer() != nil {
		C.QSplitter_SetOpaqueResize(ptr.Pointer(), C.int(qt.GoBoolToInt(opaque)))
	}
}

func (ptr *QSplitter) SetOrientation(v core.Qt__Orientation) {
	defer qt.Recovering("QSplitter::setOrientation")

	if ptr.Pointer() != nil {
		C.QSplitter_SetOrientation(ptr.Pointer(), C.int(v))
	}
}

func NewQSplitter(parent QWidget_ITF) *QSplitter {
	defer qt.Recovering("QSplitter::QSplitter")

	return NewQSplitterFromPointer(C.QSplitter_NewQSplitter(PointerFromQWidget(parent)))
}

func NewQSplitter2(orientation core.Qt__Orientation, parent QWidget_ITF) *QSplitter {
	defer qt.Recovering("QSplitter::QSplitter")

	return NewQSplitterFromPointer(C.QSplitter_NewQSplitter2(C.int(orientation), PointerFromQWidget(parent)))
}

func (ptr *QSplitter) AddWidget(widget QWidget_ITF) {
	defer qt.Recovering("QSplitter::addWidget")

	if ptr.Pointer() != nil {
		C.QSplitter_AddWidget(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QSplitter) ConnectChangeEvent(f func(ev *core.QEvent)) {
	defer qt.Recovering("connect QSplitter::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QSplitter) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QSplitter::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQSplitterChangeEvent
func callbackQSplitterChangeEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QSplitter::changeEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "changeEvent")
	if signal != nil {
		defer signal.(func(*core.QEvent))(core.NewQEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QSplitter) ConnectChildEvent(f func(c *core.QChildEvent)) {
	defer qt.Recovering("connect QSplitter::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QSplitter) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QSplitter::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQSplitterChildEvent
func callbackQSplitterChildEvent(ptrName *C.char, c unsafe.Pointer) bool {
	defer qt.Recovering("callback QSplitter::childEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "childEvent")
	if signal != nil {
		defer signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(c))
		return true
	}
	return false

}

func (ptr *QSplitter) Count() int {
	defer qt.Recovering("QSplitter::count")

	if ptr.Pointer() != nil {
		return int(C.QSplitter_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSplitter) GetRange(index int, min int, max int) {
	defer qt.Recovering("QSplitter::getRange")

	if ptr.Pointer() != nil {
		C.QSplitter_GetRange(ptr.Pointer(), C.int(index), C.int(min), C.int(max))
	}
}

func (ptr *QSplitter) Handle(index int) *QSplitterHandle {
	defer qt.Recovering("QSplitter::handle")

	if ptr.Pointer() != nil {
		return NewQSplitterHandleFromPointer(C.QSplitter_Handle(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QSplitter) InsertWidget(index int, widget QWidget_ITF) {
	defer qt.Recovering("QSplitter::insertWidget")

	if ptr.Pointer() != nil {
		C.QSplitter_InsertWidget(ptr.Pointer(), C.int(index), PointerFromQWidget(widget))
	}
}

func (ptr *QSplitter) IsCollapsible(index int) bool {
	defer qt.Recovering("QSplitter::isCollapsible")

	if ptr.Pointer() != nil {
		return C.QSplitter_IsCollapsible(ptr.Pointer(), C.int(index)) != 0
	}
	return false
}

func (ptr *QSplitter) MinimumSizeHint() *core.QSize {
	defer qt.Recovering("QSplitter::minimumSizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QSplitter_MinimumSizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSplitter) Refresh() {
	defer qt.Recovering("QSplitter::refresh")

	if ptr.Pointer() != nil {
		C.QSplitter_Refresh(ptr.Pointer())
	}
}

func (ptr *QSplitter) ConnectResizeEvent(f func(v *gui.QResizeEvent)) {
	defer qt.Recovering("connect QSplitter::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QSplitter) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QSplitter::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQSplitterResizeEvent
func callbackQSplitterResizeEvent(ptrName *C.char, v unsafe.Pointer) bool {
	defer qt.Recovering("callback QSplitter::resizeEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "resizeEvent")
	if signal != nil {
		defer signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(v))
		return true
	}
	return false

}

func (ptr *QSplitter) RestoreState(state core.QByteArray_ITF) bool {
	defer qt.Recovering("QSplitter::restoreState")

	if ptr.Pointer() != nil {
		return C.QSplitter_RestoreState(ptr.Pointer(), core.PointerFromQByteArray(state)) != 0
	}
	return false
}

func (ptr *QSplitter) SaveState() *core.QByteArray {
	defer qt.Recovering("QSplitter::saveState")

	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QSplitter_SaveState(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSplitter) SetCollapsible(index int, collapse bool) {
	defer qt.Recovering("QSplitter::setCollapsible")

	if ptr.Pointer() != nil {
		C.QSplitter_SetCollapsible(ptr.Pointer(), C.int(index), C.int(qt.GoBoolToInt(collapse)))
	}
}

func (ptr *QSplitter) SetStretchFactor(index int, stretch int) {
	defer qt.Recovering("QSplitter::setStretchFactor")

	if ptr.Pointer() != nil {
		C.QSplitter_SetStretchFactor(ptr.Pointer(), C.int(index), C.int(stretch))
	}
}

func (ptr *QSplitter) SizeHint() *core.QSize {
	defer qt.Recovering("QSplitter::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QSplitter_SizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSplitter) ConnectSplitterMoved(f func(pos int, index int)) {
	defer qt.Recovering("connect QSplitter::splitterMoved")

	if ptr.Pointer() != nil {
		C.QSplitter_ConnectSplitterMoved(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "splitterMoved", f)
	}
}

func (ptr *QSplitter) DisconnectSplitterMoved() {
	defer qt.Recovering("disconnect QSplitter::splitterMoved")

	if ptr.Pointer() != nil {
		C.QSplitter_DisconnectSplitterMoved(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "splitterMoved")
	}
}

//export callbackQSplitterSplitterMoved
func callbackQSplitterSplitterMoved(ptrName *C.char, pos C.int, index C.int) {
	defer qt.Recovering("callback QSplitter::splitterMoved")

	var signal = qt.GetSignal(C.GoString(ptrName), "splitterMoved")
	if signal != nil {
		signal.(func(int, int))(int(pos), int(index))
	}

}

func (ptr *QSplitter) Widget(index int) *QWidget {
	defer qt.Recovering("QSplitter::widget")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QSplitter_Widget(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QSplitter) DestroyQSplitter() {
	defer qt.Recovering("QSplitter::~QSplitter")

	if ptr.Pointer() != nil {
		C.QSplitter_DestroyQSplitter(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
