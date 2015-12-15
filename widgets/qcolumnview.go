package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QColumnView struct {
	QAbstractItemView
}

type QColumnView_ITF interface {
	QAbstractItemView_ITF
	QColumnView_PTR() *QColumnView
}

func PointerFromQColumnView(ptr QColumnView_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QColumnView_PTR().Pointer()
	}
	return nil
}

func NewQColumnViewFromPointer(ptr unsafe.Pointer) *QColumnView {
	var n = new(QColumnView)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QColumnView_") {
		n.SetObjectName("QColumnView_" + qt.Identifier())
	}
	return n
}

func (ptr *QColumnView) QColumnView_PTR() *QColumnView {
	return ptr
}

func (ptr *QColumnView) ResizeGripsVisible() bool {
	defer qt.Recovering("QColumnView::resizeGripsVisible")

	if ptr.Pointer() != nil {
		return C.QColumnView_ResizeGripsVisible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QColumnView) SetResizeGripsVisible(visible bool) {
	defer qt.Recovering("QColumnView::setResizeGripsVisible")

	if ptr.Pointer() != nil {
		C.QColumnView_SetResizeGripsVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func NewQColumnView(parent QWidget_ITF) *QColumnView {
	defer qt.Recovering("QColumnView::QColumnView")

	return NewQColumnViewFromPointer(C.QColumnView_NewQColumnView(PointerFromQWidget(parent)))
}

func (ptr *QColumnView) ConnectCurrentChanged(f func(current *core.QModelIndex, previous *core.QModelIndex)) {
	defer qt.Recovering("connect QColumnView::currentChanged")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "currentChanged", f)
	}
}

func (ptr *QColumnView) DisconnectCurrentChanged() {
	defer qt.Recovering("disconnect QColumnView::currentChanged")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "currentChanged")
	}
}

//export callbackQColumnViewCurrentChanged
func callbackQColumnViewCurrentChanged(ptrName *C.char, current unsafe.Pointer, previous unsafe.Pointer) bool {
	defer qt.Recovering("callback QColumnView::currentChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "currentChanged")
	if signal != nil {
		defer signal.(func(*core.QModelIndex, *core.QModelIndex))(core.NewQModelIndexFromPointer(current), core.NewQModelIndexFromPointer(previous))
		return true
	}
	return false

}

func (ptr *QColumnView) IndexAt(point core.QPoint_ITF) *core.QModelIndex {
	defer qt.Recovering("QColumnView::indexAt")

	if ptr.Pointer() != nil {
		return core.NewQModelIndexFromPointer(C.QColumnView_IndexAt(ptr.Pointer(), core.PointerFromQPoint(point)))
	}
	return nil
}

func (ptr *QColumnView) PreviewWidget() *QWidget {
	defer qt.Recovering("QColumnView::previewWidget")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QColumnView_PreviewWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QColumnView) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QColumnView::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QColumnView) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QColumnView::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQColumnViewResizeEvent
func callbackQColumnViewResizeEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QColumnView::resizeEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "resizeEvent")
	if signal != nil {
		defer signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QColumnView) ConnectRowsInserted(f func(parent *core.QModelIndex, start int, end int)) {
	defer qt.Recovering("connect QColumnView::rowsInserted")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "rowsInserted", f)
	}
}

func (ptr *QColumnView) DisconnectRowsInserted() {
	defer qt.Recovering("disconnect QColumnView::rowsInserted")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "rowsInserted")
	}
}

//export callbackQColumnViewRowsInserted
func callbackQColumnViewRowsInserted(ptrName *C.char, parent unsafe.Pointer, start C.int, end C.int) bool {
	defer qt.Recovering("callback QColumnView::rowsInserted")

	var signal = qt.GetSignal(C.GoString(ptrName), "rowsInserted")
	if signal != nil {
		defer signal.(func(*core.QModelIndex, int, int))(core.NewQModelIndexFromPointer(parent), int(start), int(end))
		return true
	}
	return false

}

func (ptr *QColumnView) ConnectScrollContentsBy(f func(dx int, dy int)) {
	defer qt.Recovering("connect QColumnView::scrollContentsBy")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "scrollContentsBy", f)
	}
}

func (ptr *QColumnView) DisconnectScrollContentsBy() {
	defer qt.Recovering("disconnect QColumnView::scrollContentsBy")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "scrollContentsBy")
	}
}

//export callbackQColumnViewScrollContentsBy
func callbackQColumnViewScrollContentsBy(ptrName *C.char, dx C.int, dy C.int) bool {
	defer qt.Recovering("callback QColumnView::scrollContentsBy")

	var signal = qt.GetSignal(C.GoString(ptrName), "scrollContentsBy")
	if signal != nil {
		defer signal.(func(int, int))(int(dx), int(dy))
		return true
	}
	return false

}

func (ptr *QColumnView) ConnectScrollTo(f func(index *core.QModelIndex, hint QAbstractItemView__ScrollHint)) {
	defer qt.Recovering("connect QColumnView::scrollTo")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "scrollTo", f)
	}
}

func (ptr *QColumnView) DisconnectScrollTo() {
	defer qt.Recovering("disconnect QColumnView::scrollTo")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "scrollTo")
	}
}

//export callbackQColumnViewScrollTo
func callbackQColumnViewScrollTo(ptrName *C.char, index unsafe.Pointer, hint C.int) bool {
	defer qt.Recovering("callback QColumnView::scrollTo")

	var signal = qt.GetSignal(C.GoString(ptrName), "scrollTo")
	if signal != nil {
		defer signal.(func(*core.QModelIndex, QAbstractItemView__ScrollHint))(core.NewQModelIndexFromPointer(index), QAbstractItemView__ScrollHint(hint))
		return true
	}
	return false

}

func (ptr *QColumnView) ConnectSelectAll(f func()) {
	defer qt.Recovering("connect QColumnView::selectAll")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "selectAll", f)
	}
}

func (ptr *QColumnView) DisconnectSelectAll() {
	defer qt.Recovering("disconnect QColumnView::selectAll")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "selectAll")
	}
}

//export callbackQColumnViewSelectAll
func callbackQColumnViewSelectAll(ptrName *C.char) bool {
	defer qt.Recovering("callback QColumnView::selectAll")

	var signal = qt.GetSignal(C.GoString(ptrName), "selectAll")
	if signal != nil {
		defer signal.(func())()
		return true
	}
	return false

}

func (ptr *QColumnView) ConnectSetModel(f func(model *core.QAbstractItemModel)) {
	defer qt.Recovering("connect QColumnView::setModel")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setModel", f)
	}
}

func (ptr *QColumnView) DisconnectSetModel() {
	defer qt.Recovering("disconnect QColumnView::setModel")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setModel")
	}
}

//export callbackQColumnViewSetModel
func callbackQColumnViewSetModel(ptrName *C.char, model unsafe.Pointer) bool {
	defer qt.Recovering("callback QColumnView::setModel")

	var signal = qt.GetSignal(C.GoString(ptrName), "setModel")
	if signal != nil {
		defer signal.(func(*core.QAbstractItemModel))(core.NewQAbstractItemModelFromPointer(model))
		return true
	}
	return false

}

func (ptr *QColumnView) SetPreviewWidget(widget QWidget_ITF) {
	defer qt.Recovering("QColumnView::setPreviewWidget")

	if ptr.Pointer() != nil {
		C.QColumnView_SetPreviewWidget(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QColumnView) ConnectSetRootIndex(f func(index *core.QModelIndex)) {
	defer qt.Recovering("connect QColumnView::setRootIndex")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setRootIndex", f)
	}
}

func (ptr *QColumnView) DisconnectSetRootIndex() {
	defer qt.Recovering("disconnect QColumnView::setRootIndex")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setRootIndex")
	}
}

//export callbackQColumnViewSetRootIndex
func callbackQColumnViewSetRootIndex(ptrName *C.char, index unsafe.Pointer) bool {
	defer qt.Recovering("callback QColumnView::setRootIndex")

	var signal = qt.GetSignal(C.GoString(ptrName), "setRootIndex")
	if signal != nil {
		defer signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
		return true
	}
	return false

}

func (ptr *QColumnView) ConnectSetSelectionModel(f func(newSelectionModel *core.QItemSelectionModel)) {
	defer qt.Recovering("connect QColumnView::setSelectionModel")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setSelectionModel", f)
	}
}

func (ptr *QColumnView) DisconnectSetSelectionModel() {
	defer qt.Recovering("disconnect QColumnView::setSelectionModel")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setSelectionModel")
	}
}

//export callbackQColumnViewSetSelectionModel
func callbackQColumnViewSetSelectionModel(ptrName *C.char, newSelectionModel unsafe.Pointer) bool {
	defer qt.Recovering("callback QColumnView::setSelectionModel")

	var signal = qt.GetSignal(C.GoString(ptrName), "setSelectionModel")
	if signal != nil {
		defer signal.(func(*core.QItemSelectionModel))(core.NewQItemSelectionModelFromPointer(newSelectionModel))
		return true
	}
	return false

}

func (ptr *QColumnView) ConnectUpdatePreviewWidget(f func(index *core.QModelIndex)) {
	defer qt.Recovering("connect QColumnView::updatePreviewWidget")

	if ptr.Pointer() != nil {
		C.QColumnView_ConnectUpdatePreviewWidget(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "updatePreviewWidget", f)
	}
}

func (ptr *QColumnView) DisconnectUpdatePreviewWidget() {
	defer qt.Recovering("disconnect QColumnView::updatePreviewWidget")

	if ptr.Pointer() != nil {
		C.QColumnView_DisconnectUpdatePreviewWidget(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "updatePreviewWidget")
	}
}

//export callbackQColumnViewUpdatePreviewWidget
func callbackQColumnViewUpdatePreviewWidget(ptrName *C.char, index unsafe.Pointer) {
	defer qt.Recovering("callback QColumnView::updatePreviewWidget")

	var signal = qt.GetSignal(C.GoString(ptrName), "updatePreviewWidget")
	if signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
	}

}

func (ptr *QColumnView) DestroyQColumnView() {
	defer qt.Recovering("QColumnView::~QColumnView")

	if ptr.Pointer() != nil {
		C.QColumnView_DestroyQColumnView(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
