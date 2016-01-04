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

func (ptr *QColumnView) CreateColumn(index core.QModelIndex_ITF) *QAbstractItemView {
	defer qt.Recovering("QColumnView::createColumn")

	if ptr.Pointer() != nil {
		return NewQAbstractItemViewFromPointer(C.QColumnView_CreateColumn(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return nil
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
func callbackQColumnViewCurrentChanged(ptr unsafe.Pointer, ptrName *C.char, current unsafe.Pointer, previous unsafe.Pointer) bool {
	defer qt.Recovering("callback QColumnView::currentChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "currentChanged"); signal != nil {
		signal.(func(*core.QModelIndex, *core.QModelIndex))(core.NewQModelIndexFromPointer(current), core.NewQModelIndexFromPointer(previous))
		return true
	}
	return false

}

func (ptr *QColumnView) CurrentChanged(current core.QModelIndex_ITF, previous core.QModelIndex_ITF) {
	defer qt.Recovering("QColumnView::currentChanged")

	if ptr.Pointer() != nil {
		C.QColumnView_CurrentChanged(ptr.Pointer(), core.PointerFromQModelIndex(current), core.PointerFromQModelIndex(previous))
	}
}

func (ptr *QColumnView) CurrentChangedDefault(current core.QModelIndex_ITF, previous core.QModelIndex_ITF) {
	defer qt.Recovering("QColumnView::currentChanged")

	if ptr.Pointer() != nil {
		C.QColumnView_CurrentChangedDefault(ptr.Pointer(), core.PointerFromQModelIndex(current), core.PointerFromQModelIndex(previous))
	}
}

func (ptr *QColumnView) HorizontalOffset() int {
	defer qt.Recovering("QColumnView::horizontalOffset")

	if ptr.Pointer() != nil {
		return int(C.QColumnView_HorizontalOffset(ptr.Pointer()))
	}
	return 0
}

func (ptr *QColumnView) IndexAt(point core.QPoint_ITF) *core.QModelIndex {
	defer qt.Recovering("QColumnView::indexAt")

	if ptr.Pointer() != nil {
		return core.NewQModelIndexFromPointer(C.QColumnView_IndexAt(ptr.Pointer(), core.PointerFromQPoint(point)))
	}
	return nil
}

func (ptr *QColumnView) IsIndexHidden(index core.QModelIndex_ITF) bool {
	defer qt.Recovering("QColumnView::isIndexHidden")

	if ptr.Pointer() != nil {
		return C.QColumnView_IsIndexHidden(ptr.Pointer(), core.PointerFromQModelIndex(index)) != 0
	}
	return false
}

func (ptr *QColumnView) MoveCursor(cursorAction QAbstractItemView__CursorAction, modifiers core.Qt__KeyboardModifier) *core.QModelIndex {
	defer qt.Recovering("QColumnView::moveCursor")

	if ptr.Pointer() != nil {
		return core.NewQModelIndexFromPointer(C.QColumnView_MoveCursor(ptr.Pointer(), C.int(cursorAction), C.int(modifiers)))
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
func callbackQColumnViewResizeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QColumnView::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQColumnViewFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QColumnView) ResizeEvent(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QColumnView::resizeEvent")

	if ptr.Pointer() != nil {
		C.QColumnView_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QColumnView) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QColumnView::resizeEvent")

	if ptr.Pointer() != nil {
		C.QColumnView_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
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
func callbackQColumnViewRowsInserted(ptr unsafe.Pointer, ptrName *C.char, parent unsafe.Pointer, start C.int, end C.int) bool {
	defer qt.Recovering("callback QColumnView::rowsInserted")

	if signal := qt.GetSignal(C.GoString(ptrName), "rowsInserted"); signal != nil {
		signal.(func(*core.QModelIndex, int, int))(core.NewQModelIndexFromPointer(parent), int(start), int(end))
		return true
	}
	return false

}

func (ptr *QColumnView) RowsInserted(parent core.QModelIndex_ITF, start int, end int) {
	defer qt.Recovering("QColumnView::rowsInserted")

	if ptr.Pointer() != nil {
		C.QColumnView_RowsInserted(ptr.Pointer(), core.PointerFromQModelIndex(parent), C.int(start), C.int(end))
	}
}

func (ptr *QColumnView) RowsInsertedDefault(parent core.QModelIndex_ITF, start int, end int) {
	defer qt.Recovering("QColumnView::rowsInserted")

	if ptr.Pointer() != nil {
		C.QColumnView_RowsInsertedDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent), C.int(start), C.int(end))
	}
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
func callbackQColumnViewScrollContentsBy(ptr unsafe.Pointer, ptrName *C.char, dx C.int, dy C.int) {
	defer qt.Recovering("callback QColumnView::scrollContentsBy")

	if signal := qt.GetSignal(C.GoString(ptrName), "scrollContentsBy"); signal != nil {
		signal.(func(int, int))(int(dx), int(dy))
	} else {
		NewQColumnViewFromPointer(ptr).ScrollContentsByDefault(int(dx), int(dy))
	}
}

func (ptr *QColumnView) ScrollContentsBy(dx int, dy int) {
	defer qt.Recovering("QColumnView::scrollContentsBy")

	if ptr.Pointer() != nil {
		C.QColumnView_ScrollContentsBy(ptr.Pointer(), C.int(dx), C.int(dy))
	}
}

func (ptr *QColumnView) ScrollContentsByDefault(dx int, dy int) {
	defer qt.Recovering("QColumnView::scrollContentsBy")

	if ptr.Pointer() != nil {
		C.QColumnView_ScrollContentsByDefault(ptr.Pointer(), C.int(dx), C.int(dy))
	}
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
func callbackQColumnViewScrollTo(ptr unsafe.Pointer, ptrName *C.char, index unsafe.Pointer, hint C.int) {
	defer qt.Recovering("callback QColumnView::scrollTo")

	if signal := qt.GetSignal(C.GoString(ptrName), "scrollTo"); signal != nil {
		signal.(func(*core.QModelIndex, QAbstractItemView__ScrollHint))(core.NewQModelIndexFromPointer(index), QAbstractItemView__ScrollHint(hint))
	} else {
		NewQColumnViewFromPointer(ptr).ScrollToDefault(core.NewQModelIndexFromPointer(index), QAbstractItemView__ScrollHint(hint))
	}
}

func (ptr *QColumnView) ScrollTo(index core.QModelIndex_ITF, hint QAbstractItemView__ScrollHint) {
	defer qt.Recovering("QColumnView::scrollTo")

	if ptr.Pointer() != nil {
		C.QColumnView_ScrollTo(ptr.Pointer(), core.PointerFromQModelIndex(index), C.int(hint))
	}
}

func (ptr *QColumnView) ScrollToDefault(index core.QModelIndex_ITF, hint QAbstractItemView__ScrollHint) {
	defer qt.Recovering("QColumnView::scrollTo")

	if ptr.Pointer() != nil {
		C.QColumnView_ScrollToDefault(ptr.Pointer(), core.PointerFromQModelIndex(index), C.int(hint))
	}
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
func callbackQColumnViewSelectAll(ptr unsafe.Pointer, ptrName *C.char) bool {
	defer qt.Recovering("callback QColumnView::selectAll")

	if signal := qt.GetSignal(C.GoString(ptrName), "selectAll"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QColumnView) SelectAll() {
	defer qt.Recovering("QColumnView::selectAll")

	if ptr.Pointer() != nil {
		C.QColumnView_SelectAll(ptr.Pointer())
	}
}

func (ptr *QColumnView) SelectAllDefault() {
	defer qt.Recovering("QColumnView::selectAll")

	if ptr.Pointer() != nil {
		C.QColumnView_SelectAllDefault(ptr.Pointer())
	}
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
func callbackQColumnViewSetModel(ptr unsafe.Pointer, ptrName *C.char, model unsafe.Pointer) {
	defer qt.Recovering("callback QColumnView::setModel")

	if signal := qt.GetSignal(C.GoString(ptrName), "setModel"); signal != nil {
		signal.(func(*core.QAbstractItemModel))(core.NewQAbstractItemModelFromPointer(model))
	} else {
		NewQColumnViewFromPointer(ptr).SetModelDefault(core.NewQAbstractItemModelFromPointer(model))
	}
}

func (ptr *QColumnView) SetModel(model core.QAbstractItemModel_ITF) {
	defer qt.Recovering("QColumnView::setModel")

	if ptr.Pointer() != nil {
		C.QColumnView_SetModel(ptr.Pointer(), core.PointerFromQAbstractItemModel(model))
	}
}

func (ptr *QColumnView) SetModelDefault(model core.QAbstractItemModel_ITF) {
	defer qt.Recovering("QColumnView::setModel")

	if ptr.Pointer() != nil {
		C.QColumnView_SetModelDefault(ptr.Pointer(), core.PointerFromQAbstractItemModel(model))
	}
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
func callbackQColumnViewSetRootIndex(ptr unsafe.Pointer, ptrName *C.char, index unsafe.Pointer) bool {
	defer qt.Recovering("callback QColumnView::setRootIndex")

	if signal := qt.GetSignal(C.GoString(ptrName), "setRootIndex"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
		return true
	}
	return false

}

func (ptr *QColumnView) SetRootIndex(index core.QModelIndex_ITF) {
	defer qt.Recovering("QColumnView::setRootIndex")

	if ptr.Pointer() != nil {
		C.QColumnView_SetRootIndex(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QColumnView) SetRootIndexDefault(index core.QModelIndex_ITF) {
	defer qt.Recovering("QColumnView::setRootIndex")

	if ptr.Pointer() != nil {
		C.QColumnView_SetRootIndexDefault(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QColumnView) ConnectSetSelection(f func(rect *core.QRect, command core.QItemSelectionModel__SelectionFlag)) {
	defer qt.Recovering("connect QColumnView::setSelection")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setSelection", f)
	}
}

func (ptr *QColumnView) DisconnectSetSelection() {
	defer qt.Recovering("disconnect QColumnView::setSelection")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setSelection")
	}
}

//export callbackQColumnViewSetSelection
func callbackQColumnViewSetSelection(ptr unsafe.Pointer, ptrName *C.char, rect unsafe.Pointer, command C.int) {
	defer qt.Recovering("callback QColumnView::setSelection")

	if signal := qt.GetSignal(C.GoString(ptrName), "setSelection"); signal != nil {
		signal.(func(*core.QRect, core.QItemSelectionModel__SelectionFlag))(core.NewQRectFromPointer(rect), core.QItemSelectionModel__SelectionFlag(command))
	} else {
		NewQColumnViewFromPointer(ptr).SetSelectionDefault(core.NewQRectFromPointer(rect), core.QItemSelectionModel__SelectionFlag(command))
	}
}

func (ptr *QColumnView) SetSelection(rect core.QRect_ITF, command core.QItemSelectionModel__SelectionFlag) {
	defer qt.Recovering("QColumnView::setSelection")

	if ptr.Pointer() != nil {
		C.QColumnView_SetSelection(ptr.Pointer(), core.PointerFromQRect(rect), C.int(command))
	}
}

func (ptr *QColumnView) SetSelectionDefault(rect core.QRect_ITF, command core.QItemSelectionModel__SelectionFlag) {
	defer qt.Recovering("QColumnView::setSelection")

	if ptr.Pointer() != nil {
		C.QColumnView_SetSelectionDefault(ptr.Pointer(), core.PointerFromQRect(rect), C.int(command))
	}
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
func callbackQColumnViewSetSelectionModel(ptr unsafe.Pointer, ptrName *C.char, newSelectionModel unsafe.Pointer) {
	defer qt.Recovering("callback QColumnView::setSelectionModel")

	if signal := qt.GetSignal(C.GoString(ptrName), "setSelectionModel"); signal != nil {
		signal.(func(*core.QItemSelectionModel))(core.NewQItemSelectionModelFromPointer(newSelectionModel))
	} else {
		NewQColumnViewFromPointer(ptr).SetSelectionModelDefault(core.NewQItemSelectionModelFromPointer(newSelectionModel))
	}
}

func (ptr *QColumnView) SetSelectionModel(newSelectionModel core.QItemSelectionModel_ITF) {
	defer qt.Recovering("QColumnView::setSelectionModel")

	if ptr.Pointer() != nil {
		C.QColumnView_SetSelectionModel(ptr.Pointer(), core.PointerFromQItemSelectionModel(newSelectionModel))
	}
}

func (ptr *QColumnView) SetSelectionModelDefault(newSelectionModel core.QItemSelectionModel_ITF) {
	defer qt.Recovering("QColumnView::setSelectionModel")

	if ptr.Pointer() != nil {
		C.QColumnView_SetSelectionModelDefault(ptr.Pointer(), core.PointerFromQItemSelectionModel(newSelectionModel))
	}
}

func (ptr *QColumnView) SizeHint() *core.QSize {
	defer qt.Recovering("QColumnView::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QColumnView_SizeHint(ptr.Pointer()))
	}
	return nil
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
func callbackQColumnViewUpdatePreviewWidget(ptr unsafe.Pointer, ptrName *C.char, index unsafe.Pointer) {
	defer qt.Recovering("callback QColumnView::updatePreviewWidget")

	if signal := qt.GetSignal(C.GoString(ptrName), "updatePreviewWidget"); signal != nil {
		signal.(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
	}

}

func (ptr *QColumnView) UpdatePreviewWidget(index core.QModelIndex_ITF) {
	defer qt.Recovering("QColumnView::updatePreviewWidget")

	if ptr.Pointer() != nil {
		C.QColumnView_UpdatePreviewWidget(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QColumnView) VerticalOffset() int {
	defer qt.Recovering("QColumnView::verticalOffset")

	if ptr.Pointer() != nil {
		return int(C.QColumnView_VerticalOffset(ptr.Pointer()))
	}
	return 0
}

func (ptr *QColumnView) VisualRect(index core.QModelIndex_ITF) *core.QRect {
	defer qt.Recovering("QColumnView::visualRect")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QColumnView_VisualRect(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return nil
}

func (ptr *QColumnView) VisualRegionForSelection(selection core.QItemSelection_ITF) *gui.QRegion {
	defer qt.Recovering("QColumnView::visualRegionForSelection")

	if ptr.Pointer() != nil {
		return gui.NewQRegionFromPointer(C.QColumnView_VisualRegionForSelection(ptr.Pointer(), core.PointerFromQItemSelection(selection)))
	}
	return nil
}

func (ptr *QColumnView) DestroyQColumnView() {
	defer qt.Recovering("QColumnView::~QColumnView")

	if ptr.Pointer() != nil {
		C.QColumnView_DestroyQColumnView(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QColumnView) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QColumnView::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QColumnView) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QColumnView::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQColumnViewDragLeaveEvent
func callbackQColumnViewDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QColumnView::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQColumnViewFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QColumnView) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QColumnView::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QColumnView_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QColumnView) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QColumnView::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QColumnView_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QColumnView) ConnectCloseEditor(f func(editor *QWidget, hint QAbstractItemDelegate__EndEditHint)) {
	defer qt.Recovering("connect QColumnView::closeEditor")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEditor", f)
	}
}

func (ptr *QColumnView) DisconnectCloseEditor() {
	defer qt.Recovering("disconnect QColumnView::closeEditor")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEditor")
	}
}

//export callbackQColumnViewCloseEditor
func callbackQColumnViewCloseEditor(ptr unsafe.Pointer, ptrName *C.char, editor unsafe.Pointer, hint C.int) bool {
	defer qt.Recovering("callback QColumnView::closeEditor")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEditor"); signal != nil {
		signal.(func(*QWidget, QAbstractItemDelegate__EndEditHint))(NewQWidgetFromPointer(editor), QAbstractItemDelegate__EndEditHint(hint))
		return true
	}
	return false

}

func (ptr *QColumnView) CloseEditor(editor QWidget_ITF, hint QAbstractItemDelegate__EndEditHint) {
	defer qt.Recovering("QColumnView::closeEditor")

	if ptr.Pointer() != nil {
		C.QColumnView_CloseEditor(ptr.Pointer(), PointerFromQWidget(editor), C.int(hint))
	}
}

func (ptr *QColumnView) CloseEditorDefault(editor QWidget_ITF, hint QAbstractItemDelegate__EndEditHint) {
	defer qt.Recovering("QColumnView::closeEditor")

	if ptr.Pointer() != nil {
		C.QColumnView_CloseEditorDefault(ptr.Pointer(), PointerFromQWidget(editor), C.int(hint))
	}
}

func (ptr *QColumnView) ConnectCommitData(f func(editor *QWidget)) {
	defer qt.Recovering("connect QColumnView::commitData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "commitData", f)
	}
}

func (ptr *QColumnView) DisconnectCommitData() {
	defer qt.Recovering("disconnect QColumnView::commitData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "commitData")
	}
}

//export callbackQColumnViewCommitData
func callbackQColumnViewCommitData(ptr unsafe.Pointer, ptrName *C.char, editor unsafe.Pointer) bool {
	defer qt.Recovering("callback QColumnView::commitData")

	if signal := qt.GetSignal(C.GoString(ptrName), "commitData"); signal != nil {
		signal.(func(*QWidget))(NewQWidgetFromPointer(editor))
		return true
	}
	return false

}

func (ptr *QColumnView) CommitData(editor QWidget_ITF) {
	defer qt.Recovering("QColumnView::commitData")

	if ptr.Pointer() != nil {
		C.QColumnView_CommitData(ptr.Pointer(), PointerFromQWidget(editor))
	}
}

func (ptr *QColumnView) CommitDataDefault(editor QWidget_ITF) {
	defer qt.Recovering("QColumnView::commitData")

	if ptr.Pointer() != nil {
		C.QColumnView_CommitDataDefault(ptr.Pointer(), PointerFromQWidget(editor))
	}
}

func (ptr *QColumnView) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QColumnView::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QColumnView) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QColumnView::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQColumnViewDragEnterEvent
func callbackQColumnViewDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QColumnView::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQColumnViewFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QColumnView) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QColumnView::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QColumnView_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QColumnView) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QColumnView::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QColumnView_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QColumnView) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QColumnView::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QColumnView) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QColumnView::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQColumnViewDragMoveEvent
func callbackQColumnViewDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QColumnView::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQColumnViewFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QColumnView) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QColumnView::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QColumnView_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QColumnView) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QColumnView::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QColumnView_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QColumnView) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QColumnView::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QColumnView) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QColumnView::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQColumnViewDropEvent
func callbackQColumnViewDropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QColumnView::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQColumnViewFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QColumnView) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QColumnView::dropEvent")

	if ptr.Pointer() != nil {
		C.QColumnView_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QColumnView) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QColumnView::dropEvent")

	if ptr.Pointer() != nil {
		C.QColumnView_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QColumnView) ConnectEditorDestroyed(f func(editor *core.QObject)) {
	defer qt.Recovering("connect QColumnView::editorDestroyed")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "editorDestroyed", f)
	}
}

func (ptr *QColumnView) DisconnectEditorDestroyed() {
	defer qt.Recovering("disconnect QColumnView::editorDestroyed")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "editorDestroyed")
	}
}

//export callbackQColumnViewEditorDestroyed
func callbackQColumnViewEditorDestroyed(ptr unsafe.Pointer, ptrName *C.char, editor unsafe.Pointer) bool {
	defer qt.Recovering("callback QColumnView::editorDestroyed")

	if signal := qt.GetSignal(C.GoString(ptrName), "editorDestroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(editor))
		return true
	}
	return false

}

func (ptr *QColumnView) EditorDestroyed(editor core.QObject_ITF) {
	defer qt.Recovering("QColumnView::editorDestroyed")

	if ptr.Pointer() != nil {
		C.QColumnView_EditorDestroyed(ptr.Pointer(), core.PointerFromQObject(editor))
	}
}

func (ptr *QColumnView) EditorDestroyedDefault(editor core.QObject_ITF) {
	defer qt.Recovering("QColumnView::editorDestroyed")

	if ptr.Pointer() != nil {
		C.QColumnView_EditorDestroyedDefault(ptr.Pointer(), core.PointerFromQObject(editor))
	}
}

func (ptr *QColumnView) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QColumnView::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QColumnView) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QColumnView::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQColumnViewFocusInEvent
func callbackQColumnViewFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QColumnView::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQColumnViewFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QColumnView) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QColumnView::focusInEvent")

	if ptr.Pointer() != nil {
		C.QColumnView_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QColumnView) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QColumnView::focusInEvent")

	if ptr.Pointer() != nil {
		C.QColumnView_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QColumnView) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QColumnView::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QColumnView) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QColumnView::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQColumnViewFocusOutEvent
func callbackQColumnViewFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QColumnView::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQColumnViewFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QColumnView) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QColumnView::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QColumnView_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QColumnView) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QColumnView::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QColumnView_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QColumnView) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QColumnView::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QColumnView) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QColumnView::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQColumnViewInputMethodEvent
func callbackQColumnViewInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QColumnView::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQColumnViewFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QColumnView) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QColumnView::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QColumnView_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QColumnView) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QColumnView::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QColumnView_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QColumnView) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QColumnView::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QColumnView) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QColumnView::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQColumnViewKeyPressEvent
func callbackQColumnViewKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QColumnView::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQColumnViewFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QColumnView) KeyPressEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QColumnView::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QColumnView_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QColumnView) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QColumnView::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QColumnView_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QColumnView) ConnectKeyboardSearch(f func(search string)) {
	defer qt.Recovering("connect QColumnView::keyboardSearch")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyboardSearch", f)
	}
}

func (ptr *QColumnView) DisconnectKeyboardSearch() {
	defer qt.Recovering("disconnect QColumnView::keyboardSearch")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyboardSearch")
	}
}

//export callbackQColumnViewKeyboardSearch
func callbackQColumnViewKeyboardSearch(ptr unsafe.Pointer, ptrName *C.char, search *C.char) {
	defer qt.Recovering("callback QColumnView::keyboardSearch")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyboardSearch"); signal != nil {
		signal.(func(string))(C.GoString(search))
	} else {
		NewQColumnViewFromPointer(ptr).KeyboardSearchDefault(C.GoString(search))
	}
}

func (ptr *QColumnView) KeyboardSearch(search string) {
	defer qt.Recovering("QColumnView::keyboardSearch")

	if ptr.Pointer() != nil {
		C.QColumnView_KeyboardSearch(ptr.Pointer(), C.CString(search))
	}
}

func (ptr *QColumnView) KeyboardSearchDefault(search string) {
	defer qt.Recovering("QColumnView::keyboardSearch")

	if ptr.Pointer() != nil {
		C.QColumnView_KeyboardSearchDefault(ptr.Pointer(), C.CString(search))
	}
}

func (ptr *QColumnView) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QColumnView::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QColumnView) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QColumnView::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQColumnViewMouseDoubleClickEvent
func callbackQColumnViewMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QColumnView::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQColumnViewFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QColumnView) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QColumnView::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QColumnView_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QColumnView) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QColumnView::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QColumnView_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QColumnView) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QColumnView::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QColumnView) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QColumnView::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQColumnViewMouseMoveEvent
func callbackQColumnViewMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QColumnView::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQColumnViewFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QColumnView) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QColumnView::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QColumnView_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QColumnView) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QColumnView::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QColumnView_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QColumnView) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QColumnView::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QColumnView) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QColumnView::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQColumnViewMousePressEvent
func callbackQColumnViewMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QColumnView::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQColumnViewFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QColumnView) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QColumnView::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QColumnView_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QColumnView) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QColumnView::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QColumnView_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QColumnView) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QColumnView::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QColumnView) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QColumnView::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQColumnViewMouseReleaseEvent
func callbackQColumnViewMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QColumnView::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQColumnViewFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QColumnView) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QColumnView::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QColumnView_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QColumnView) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QColumnView::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QColumnView_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QColumnView) ConnectReset(f func()) {
	defer qt.Recovering("connect QColumnView::reset")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "reset", f)
	}
}

func (ptr *QColumnView) DisconnectReset() {
	defer qt.Recovering("disconnect QColumnView::reset")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "reset")
	}
}

//export callbackQColumnViewReset
func callbackQColumnViewReset(ptr unsafe.Pointer, ptrName *C.char) bool {
	defer qt.Recovering("callback QColumnView::reset")

	if signal := qt.GetSignal(C.GoString(ptrName), "reset"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QColumnView) Reset() {
	defer qt.Recovering("QColumnView::reset")

	if ptr.Pointer() != nil {
		C.QColumnView_Reset(ptr.Pointer())
	}
}

func (ptr *QColumnView) ResetDefault() {
	defer qt.Recovering("QColumnView::reset")

	if ptr.Pointer() != nil {
		C.QColumnView_ResetDefault(ptr.Pointer())
	}
}

func (ptr *QColumnView) ConnectRowsAboutToBeRemoved(f func(parent *core.QModelIndex, start int, end int)) {
	defer qt.Recovering("connect QColumnView::rowsAboutToBeRemoved")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "rowsAboutToBeRemoved", f)
	}
}

func (ptr *QColumnView) DisconnectRowsAboutToBeRemoved() {
	defer qt.Recovering("disconnect QColumnView::rowsAboutToBeRemoved")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "rowsAboutToBeRemoved")
	}
}

//export callbackQColumnViewRowsAboutToBeRemoved
func callbackQColumnViewRowsAboutToBeRemoved(ptr unsafe.Pointer, ptrName *C.char, parent unsafe.Pointer, start C.int, end C.int) bool {
	defer qt.Recovering("callback QColumnView::rowsAboutToBeRemoved")

	if signal := qt.GetSignal(C.GoString(ptrName), "rowsAboutToBeRemoved"); signal != nil {
		signal.(func(*core.QModelIndex, int, int))(core.NewQModelIndexFromPointer(parent), int(start), int(end))
		return true
	}
	return false

}

func (ptr *QColumnView) RowsAboutToBeRemoved(parent core.QModelIndex_ITF, start int, end int) {
	defer qt.Recovering("QColumnView::rowsAboutToBeRemoved")

	if ptr.Pointer() != nil {
		C.QColumnView_RowsAboutToBeRemoved(ptr.Pointer(), core.PointerFromQModelIndex(parent), C.int(start), C.int(end))
	}
}

func (ptr *QColumnView) RowsAboutToBeRemovedDefault(parent core.QModelIndex_ITF, start int, end int) {
	defer qt.Recovering("QColumnView::rowsAboutToBeRemoved")

	if ptr.Pointer() != nil {
		C.QColumnView_RowsAboutToBeRemovedDefault(ptr.Pointer(), core.PointerFromQModelIndex(parent), C.int(start), C.int(end))
	}
}

func (ptr *QColumnView) ConnectStartDrag(f func(supportedActions core.Qt__DropAction)) {
	defer qt.Recovering("connect QColumnView::startDrag")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "startDrag", f)
	}
}

func (ptr *QColumnView) DisconnectStartDrag() {
	defer qt.Recovering("disconnect QColumnView::startDrag")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "startDrag")
	}
}

//export callbackQColumnViewStartDrag
func callbackQColumnViewStartDrag(ptr unsafe.Pointer, ptrName *C.char, supportedActions C.int) {
	defer qt.Recovering("callback QColumnView::startDrag")

	if signal := qt.GetSignal(C.GoString(ptrName), "startDrag"); signal != nil {
		signal.(func(core.Qt__DropAction))(core.Qt__DropAction(supportedActions))
	} else {
		NewQColumnViewFromPointer(ptr).StartDragDefault(core.Qt__DropAction(supportedActions))
	}
}

func (ptr *QColumnView) StartDrag(supportedActions core.Qt__DropAction) {
	defer qt.Recovering("QColumnView::startDrag")

	if ptr.Pointer() != nil {
		C.QColumnView_StartDrag(ptr.Pointer(), C.int(supportedActions))
	}
}

func (ptr *QColumnView) StartDragDefault(supportedActions core.Qt__DropAction) {
	defer qt.Recovering("QColumnView::startDrag")

	if ptr.Pointer() != nil {
		C.QColumnView_StartDragDefault(ptr.Pointer(), C.int(supportedActions))
	}
}

func (ptr *QColumnView) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QColumnView::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QColumnView) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QColumnView::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQColumnViewTimerEvent
func callbackQColumnViewTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QColumnView::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQColumnViewFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QColumnView) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QColumnView::timerEvent")

	if ptr.Pointer() != nil {
		C.QColumnView_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QColumnView) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QColumnView::timerEvent")

	if ptr.Pointer() != nil {
		C.QColumnView_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QColumnView) ConnectUpdateGeometries(f func()) {
	defer qt.Recovering("connect QColumnView::updateGeometries")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "updateGeometries", f)
	}
}

func (ptr *QColumnView) DisconnectUpdateGeometries() {
	defer qt.Recovering("disconnect QColumnView::updateGeometries")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "updateGeometries")
	}
}

//export callbackQColumnViewUpdateGeometries
func callbackQColumnViewUpdateGeometries(ptr unsafe.Pointer, ptrName *C.char) bool {
	defer qt.Recovering("callback QColumnView::updateGeometries")

	if signal := qt.GetSignal(C.GoString(ptrName), "updateGeometries"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QColumnView) UpdateGeometries() {
	defer qt.Recovering("QColumnView::updateGeometries")

	if ptr.Pointer() != nil {
		C.QColumnView_UpdateGeometries(ptr.Pointer())
	}
}

func (ptr *QColumnView) UpdateGeometriesDefault() {
	defer qt.Recovering("QColumnView::updateGeometries")

	if ptr.Pointer() != nil {
		C.QColumnView_UpdateGeometriesDefault(ptr.Pointer())
	}
}

func (ptr *QColumnView) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QColumnView::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QColumnView) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QColumnView::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQColumnViewPaintEvent
func callbackQColumnViewPaintEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QColumnView::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQColumnViewFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QColumnView) PaintEvent(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QColumnView::paintEvent")

	if ptr.Pointer() != nil {
		C.QColumnView_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QColumnView) PaintEventDefault(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QColumnView::paintEvent")

	if ptr.Pointer() != nil {
		C.QColumnView_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QColumnView) ConnectContextMenuEvent(f func(e *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QColumnView::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QColumnView) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QColumnView::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQColumnViewContextMenuEvent
func callbackQColumnViewContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QColumnView::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(e))
	} else {
		NewQColumnViewFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(e))
	}
}

func (ptr *QColumnView) ContextMenuEvent(e gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QColumnView::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QColumnView_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(e))
	}
}

func (ptr *QColumnView) ContextMenuEventDefault(e gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QColumnView::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QColumnView_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(e))
	}
}

func (ptr *QColumnView) ConnectSetupViewport(f func(viewport *QWidget)) {
	defer qt.Recovering("connect QColumnView::setupViewport")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setupViewport", f)
	}
}

func (ptr *QColumnView) DisconnectSetupViewport() {
	defer qt.Recovering("disconnect QColumnView::setupViewport")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setupViewport")
	}
}

//export callbackQColumnViewSetupViewport
func callbackQColumnViewSetupViewport(ptr unsafe.Pointer, ptrName *C.char, viewport unsafe.Pointer) {
	defer qt.Recovering("callback QColumnView::setupViewport")

	if signal := qt.GetSignal(C.GoString(ptrName), "setupViewport"); signal != nil {
		signal.(func(*QWidget))(NewQWidgetFromPointer(viewport))
	} else {
		NewQColumnViewFromPointer(ptr).SetupViewportDefault(NewQWidgetFromPointer(viewport))
	}
}

func (ptr *QColumnView) SetupViewport(viewport QWidget_ITF) {
	defer qt.Recovering("QColumnView::setupViewport")

	if ptr.Pointer() != nil {
		C.QColumnView_SetupViewport(ptr.Pointer(), PointerFromQWidget(viewport))
	}
}

func (ptr *QColumnView) SetupViewportDefault(viewport QWidget_ITF) {
	defer qt.Recovering("QColumnView::setupViewport")

	if ptr.Pointer() != nil {
		C.QColumnView_SetupViewportDefault(ptr.Pointer(), PointerFromQWidget(viewport))
	}
}

func (ptr *QColumnView) ConnectWheelEvent(f func(e *gui.QWheelEvent)) {
	defer qt.Recovering("connect QColumnView::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QColumnView) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QColumnView::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQColumnViewWheelEvent
func callbackQColumnViewWheelEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QColumnView::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(e))
	} else {
		NewQColumnViewFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(e))
	}
}

func (ptr *QColumnView) WheelEvent(e gui.QWheelEvent_ITF) {
	defer qt.Recovering("QColumnView::wheelEvent")

	if ptr.Pointer() != nil {
		C.QColumnView_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(e))
	}
}

func (ptr *QColumnView) WheelEventDefault(e gui.QWheelEvent_ITF) {
	defer qt.Recovering("QColumnView::wheelEvent")

	if ptr.Pointer() != nil {
		C.QColumnView_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(e))
	}
}

func (ptr *QColumnView) ConnectChangeEvent(f func(ev *core.QEvent)) {
	defer qt.Recovering("connect QColumnView::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QColumnView) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QColumnView::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQColumnViewChangeEvent
func callbackQColumnViewChangeEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QColumnView::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(ev))
	} else {
		NewQColumnViewFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(ev))
	}
}

func (ptr *QColumnView) ChangeEvent(ev core.QEvent_ITF) {
	defer qt.Recovering("QColumnView::changeEvent")

	if ptr.Pointer() != nil {
		C.QColumnView_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(ev))
	}
}

func (ptr *QColumnView) ChangeEventDefault(ev core.QEvent_ITF) {
	defer qt.Recovering("QColumnView::changeEvent")

	if ptr.Pointer() != nil {
		C.QColumnView_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(ev))
	}
}

func (ptr *QColumnView) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QColumnView::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QColumnView) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QColumnView::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQColumnViewActionEvent
func callbackQColumnViewActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QColumnView::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQColumnViewFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QColumnView) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QColumnView::actionEvent")

	if ptr.Pointer() != nil {
		C.QColumnView_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QColumnView) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QColumnView::actionEvent")

	if ptr.Pointer() != nil {
		C.QColumnView_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QColumnView) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QColumnView::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QColumnView) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QColumnView::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQColumnViewEnterEvent
func callbackQColumnViewEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QColumnView::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQColumnViewFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QColumnView) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QColumnView::enterEvent")

	if ptr.Pointer() != nil {
		C.QColumnView_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QColumnView) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QColumnView::enterEvent")

	if ptr.Pointer() != nil {
		C.QColumnView_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QColumnView) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QColumnView::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QColumnView) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QColumnView::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQColumnViewHideEvent
func callbackQColumnViewHideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QColumnView::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQColumnViewFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QColumnView) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QColumnView::hideEvent")

	if ptr.Pointer() != nil {
		C.QColumnView_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QColumnView) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QColumnView::hideEvent")

	if ptr.Pointer() != nil {
		C.QColumnView_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QColumnView) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QColumnView::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QColumnView) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QColumnView::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQColumnViewLeaveEvent
func callbackQColumnViewLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QColumnView::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQColumnViewFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QColumnView) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QColumnView::leaveEvent")

	if ptr.Pointer() != nil {
		C.QColumnView_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QColumnView) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QColumnView::leaveEvent")

	if ptr.Pointer() != nil {
		C.QColumnView_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QColumnView) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QColumnView::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QColumnView) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QColumnView::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQColumnViewMoveEvent
func callbackQColumnViewMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QColumnView::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQColumnViewFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QColumnView) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QColumnView::moveEvent")

	if ptr.Pointer() != nil {
		C.QColumnView_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QColumnView) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QColumnView::moveEvent")

	if ptr.Pointer() != nil {
		C.QColumnView_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QColumnView) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QColumnView::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QColumnView) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QColumnView::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQColumnViewSetVisible
func callbackQColumnViewSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QColumnView::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QColumnView) SetVisible(visible bool) {
	defer qt.Recovering("QColumnView::setVisible")

	if ptr.Pointer() != nil {
		C.QColumnView_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QColumnView) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QColumnView::setVisible")

	if ptr.Pointer() != nil {
		C.QColumnView_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QColumnView) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QColumnView::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QColumnView) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QColumnView::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQColumnViewShowEvent
func callbackQColumnViewShowEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QColumnView::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQColumnViewFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QColumnView) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QColumnView::showEvent")

	if ptr.Pointer() != nil {
		C.QColumnView_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QColumnView) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QColumnView::showEvent")

	if ptr.Pointer() != nil {
		C.QColumnView_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QColumnView) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QColumnView::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QColumnView) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QColumnView::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQColumnViewCloseEvent
func callbackQColumnViewCloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QColumnView::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQColumnViewFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QColumnView) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QColumnView::closeEvent")

	if ptr.Pointer() != nil {
		C.QColumnView_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QColumnView) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QColumnView::closeEvent")

	if ptr.Pointer() != nil {
		C.QColumnView_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QColumnView) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QColumnView::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QColumnView) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QColumnView::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQColumnViewInitPainter
func callbackQColumnViewInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QColumnView::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQColumnViewFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QColumnView) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QColumnView::initPainter")

	if ptr.Pointer() != nil {
		C.QColumnView_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QColumnView) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QColumnView::initPainter")

	if ptr.Pointer() != nil {
		C.QColumnView_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QColumnView) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QColumnView::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QColumnView) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QColumnView::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQColumnViewKeyReleaseEvent
func callbackQColumnViewKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QColumnView::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQColumnViewFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QColumnView) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QColumnView::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QColumnView_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QColumnView) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QColumnView::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QColumnView_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QColumnView) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QColumnView::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QColumnView) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QColumnView::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQColumnViewTabletEvent
func callbackQColumnViewTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QColumnView::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQColumnViewFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QColumnView) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QColumnView::tabletEvent")

	if ptr.Pointer() != nil {
		C.QColumnView_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QColumnView) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QColumnView::tabletEvent")

	if ptr.Pointer() != nil {
		C.QColumnView_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QColumnView) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QColumnView::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QColumnView) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QColumnView::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQColumnViewChildEvent
func callbackQColumnViewChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QColumnView::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQColumnViewFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QColumnView) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QColumnView::childEvent")

	if ptr.Pointer() != nil {
		C.QColumnView_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QColumnView) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QColumnView::childEvent")

	if ptr.Pointer() != nil {
		C.QColumnView_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QColumnView) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QColumnView::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QColumnView) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QColumnView::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQColumnViewCustomEvent
func callbackQColumnViewCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QColumnView::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQColumnViewFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QColumnView) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QColumnView::customEvent")

	if ptr.Pointer() != nil {
		C.QColumnView_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QColumnView) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QColumnView::customEvent")

	if ptr.Pointer() != nil {
		C.QColumnView_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
