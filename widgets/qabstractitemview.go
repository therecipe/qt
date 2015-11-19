package widgets

//#include "qabstractitemview.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QAbstractItemView struct {
	QAbstractScrollArea
}

type QAbstractItemView_ITF interface {
	QAbstractScrollArea_ITF
	QAbstractItemView_PTR() *QAbstractItemView
}

func PointerFromQAbstractItemView(ptr QAbstractItemView_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractItemView_PTR().Pointer()
	}
	return nil
}

func NewQAbstractItemViewFromPointer(ptr unsafe.Pointer) *QAbstractItemView {
	var n = new(QAbstractItemView)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QAbstractItemView_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAbstractItemView) QAbstractItemView_PTR() *QAbstractItemView {
	return ptr
}

//QAbstractItemView::DragDropMode
type QAbstractItemView__DragDropMode int64

const (
	QAbstractItemView__NoDragDrop   = QAbstractItemView__DragDropMode(0)
	QAbstractItemView__DragOnly     = QAbstractItemView__DragDropMode(1)
	QAbstractItemView__DropOnly     = QAbstractItemView__DragDropMode(2)
	QAbstractItemView__DragDrop     = QAbstractItemView__DragDropMode(3)
	QAbstractItemView__InternalMove = QAbstractItemView__DragDropMode(4)
)

//QAbstractItemView::EditTrigger
type QAbstractItemView__EditTrigger int64

const (
	QAbstractItemView__NoEditTriggers  = QAbstractItemView__EditTrigger(0)
	QAbstractItemView__CurrentChanged  = QAbstractItemView__EditTrigger(1)
	QAbstractItemView__DoubleClicked   = QAbstractItemView__EditTrigger(2)
	QAbstractItemView__SelectedClicked = QAbstractItemView__EditTrigger(4)
	QAbstractItemView__EditKeyPressed  = QAbstractItemView__EditTrigger(8)
	QAbstractItemView__AnyKeyPressed   = QAbstractItemView__EditTrigger(16)
	QAbstractItemView__AllEditTriggers = QAbstractItemView__EditTrigger(31)
)

//QAbstractItemView::ScrollHint
type QAbstractItemView__ScrollHint int64

const (
	QAbstractItemView__EnsureVisible    = QAbstractItemView__ScrollHint(0)
	QAbstractItemView__PositionAtTop    = QAbstractItemView__ScrollHint(1)
	QAbstractItemView__PositionAtBottom = QAbstractItemView__ScrollHint(2)
	QAbstractItemView__PositionAtCenter = QAbstractItemView__ScrollHint(3)
)

//QAbstractItemView::ScrollMode
type QAbstractItemView__ScrollMode int64

const (
	QAbstractItemView__ScrollPerItem  = QAbstractItemView__ScrollMode(0)
	QAbstractItemView__ScrollPerPixel = QAbstractItemView__ScrollMode(1)
)

//QAbstractItemView::SelectionBehavior
type QAbstractItemView__SelectionBehavior int64

const (
	QAbstractItemView__SelectItems   = QAbstractItemView__SelectionBehavior(0)
	QAbstractItemView__SelectRows    = QAbstractItemView__SelectionBehavior(1)
	QAbstractItemView__SelectColumns = QAbstractItemView__SelectionBehavior(2)
)

//QAbstractItemView::SelectionMode
type QAbstractItemView__SelectionMode int64

const (
	QAbstractItemView__NoSelection         = QAbstractItemView__SelectionMode(0)
	QAbstractItemView__SingleSelection     = QAbstractItemView__SelectionMode(1)
	QAbstractItemView__MultiSelection      = QAbstractItemView__SelectionMode(2)
	QAbstractItemView__ExtendedSelection   = QAbstractItemView__SelectionMode(3)
	QAbstractItemView__ContiguousSelection = QAbstractItemView__SelectionMode(4)
)

func (ptr *QAbstractItemView) AlternatingRowColors() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractItemView_AlternatingRowColors(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractItemView) AutoScrollMargin() int {
	if ptr.Pointer() != nil {
		return int(C.QAbstractItemView_AutoScrollMargin(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractItemView) DefaultDropAction() core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return core.Qt__DropAction(C.QAbstractItemView_DefaultDropAction(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractItemView) DragDropMode() QAbstractItemView__DragDropMode {
	if ptr.Pointer() != nil {
		return QAbstractItemView__DragDropMode(C.QAbstractItemView_DragDropMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractItemView) DragDropOverwriteMode() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractItemView_DragDropOverwriteMode(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractItemView) DragEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractItemView_DragEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractItemView) EditTriggers() QAbstractItemView__EditTrigger {
	if ptr.Pointer() != nil {
		return QAbstractItemView__EditTrigger(C.QAbstractItemView_EditTriggers(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractItemView) HasAutoScroll() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractItemView_HasAutoScroll(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractItemView) HorizontalScrollMode() QAbstractItemView__ScrollMode {
	if ptr.Pointer() != nil {
		return QAbstractItemView__ScrollMode(C.QAbstractItemView_HorizontalScrollMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractItemView) SelectionBehavior() QAbstractItemView__SelectionBehavior {
	if ptr.Pointer() != nil {
		return QAbstractItemView__SelectionBehavior(C.QAbstractItemView_SelectionBehavior(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractItemView) SelectionMode() QAbstractItemView__SelectionMode {
	if ptr.Pointer() != nil {
		return QAbstractItemView__SelectionMode(C.QAbstractItemView_SelectionMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractItemView) SetAlternatingRowColors(enable bool) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetAlternatingRowColors(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QAbstractItemView) SetAutoScroll(enable bool) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetAutoScroll(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QAbstractItemView) SetAutoScrollMargin(margin int) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetAutoScrollMargin(ptr.Pointer(), C.int(margin))
	}
}

func (ptr *QAbstractItemView) SetDefaultDropAction(dropAction core.Qt__DropAction) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetDefaultDropAction(ptr.Pointer(), C.int(dropAction))
	}
}

func (ptr *QAbstractItemView) SetDragDropMode(behavior QAbstractItemView__DragDropMode) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetDragDropMode(ptr.Pointer(), C.int(behavior))
	}
}

func (ptr *QAbstractItemView) SetDragDropOverwriteMode(overwrite bool) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetDragDropOverwriteMode(ptr.Pointer(), C.int(qt.GoBoolToInt(overwrite)))
	}
}

func (ptr *QAbstractItemView) SetDragEnabled(enable bool) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetDragEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QAbstractItemView) SetDropIndicatorShown(enable bool) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetDropIndicatorShown(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QAbstractItemView) SetEditTriggers(triggers QAbstractItemView__EditTrigger) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetEditTriggers(ptr.Pointer(), C.int(triggers))
	}
}

func (ptr *QAbstractItemView) SetHorizontalScrollMode(mode QAbstractItemView__ScrollMode) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetHorizontalScrollMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QAbstractItemView) SetIconSize(size core.QSize_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetIconSize(ptr.Pointer(), core.PointerFromQSize(size))
	}
}

func (ptr *QAbstractItemView) SetSelectionBehavior(behavior QAbstractItemView__SelectionBehavior) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetSelectionBehavior(ptr.Pointer(), C.int(behavior))
	}
}

func (ptr *QAbstractItemView) SetSelectionMode(mode QAbstractItemView__SelectionMode) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetSelectionMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QAbstractItemView) SetTabKeyNavigation(enable bool) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetTabKeyNavigation(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QAbstractItemView) SetTextElideMode(mode core.Qt__TextElideMode) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetTextElideMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QAbstractItemView) SetVerticalScrollMode(mode QAbstractItemView__ScrollMode) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetVerticalScrollMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QAbstractItemView) ShowDropIndicator() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractItemView_ShowDropIndicator(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractItemView) TabKeyNavigation() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractItemView_TabKeyNavigation(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractItemView) TextElideMode() core.Qt__TextElideMode {
	if ptr.Pointer() != nil {
		return core.Qt__TextElideMode(C.QAbstractItemView_TextElideMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractItemView) VerticalScrollMode() QAbstractItemView__ScrollMode {
	if ptr.Pointer() != nil {
		return QAbstractItemView__ScrollMode(C.QAbstractItemView_VerticalScrollMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractItemView) ConnectActivated(f func(index *core.QModelIndex)) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_ConnectActivated(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "activated", f)
	}
}

func (ptr *QAbstractItemView) DisconnectActivated() {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_DisconnectActivated(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "activated")
	}
}

//export callbackQAbstractItemViewActivated
func callbackQAbstractItemViewActivated(ptrName *C.char, index unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "activated").(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
}

func (ptr *QAbstractItemView) ClearSelection() {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_ClearSelection(ptr.Pointer())
	}
}

func (ptr *QAbstractItemView) ConnectClicked(f func(index *core.QModelIndex)) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_ConnectClicked(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "clicked", f)
	}
}

func (ptr *QAbstractItemView) DisconnectClicked() {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_DisconnectClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "clicked")
	}
}

//export callbackQAbstractItemViewClicked
func callbackQAbstractItemViewClicked(ptrName *C.char, index unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "clicked").(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
}

func (ptr *QAbstractItemView) ClosePersistentEditor(index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_ClosePersistentEditor(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QAbstractItemView) CurrentIndex() *core.QModelIndex {
	if ptr.Pointer() != nil {
		return core.NewQModelIndexFromPointer(C.QAbstractItemView_CurrentIndex(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractItemView) ConnectDoubleClicked(f func(index *core.QModelIndex)) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_ConnectDoubleClicked(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "doubleClicked", f)
	}
}

func (ptr *QAbstractItemView) DisconnectDoubleClicked() {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_DisconnectDoubleClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "doubleClicked")
	}
}

//export callbackQAbstractItemViewDoubleClicked
func callbackQAbstractItemViewDoubleClicked(ptrName *C.char, index unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "doubleClicked").(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
}

func (ptr *QAbstractItemView) Edit(index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_Edit(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QAbstractItemView) ConnectEntered(f func(index *core.QModelIndex)) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_ConnectEntered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "entered", f)
	}
}

func (ptr *QAbstractItemView) DisconnectEntered() {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_DisconnectEntered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "entered")
	}
}

//export callbackQAbstractItemViewEntered
func callbackQAbstractItemViewEntered(ptrName *C.char, index unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "entered").(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
}

func (ptr *QAbstractItemView) IndexAt(point core.QPoint_ITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		return core.NewQModelIndexFromPointer(C.QAbstractItemView_IndexAt(ptr.Pointer(), core.PointerFromQPoint(point)))
	}
	return nil
}

func (ptr *QAbstractItemView) IndexWidget(index core.QModelIndex_ITF) *QWidget {
	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QAbstractItemView_IndexWidget(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return nil
}

func (ptr *QAbstractItemView) InputMethodQuery(query core.Qt__InputMethodQuery) *core.QVariant {
	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QAbstractItemView_InputMethodQuery(ptr.Pointer(), C.int(query)))
	}
	return nil
}

func (ptr *QAbstractItemView) ItemDelegate() *QAbstractItemDelegate {
	if ptr.Pointer() != nil {
		return NewQAbstractItemDelegateFromPointer(C.QAbstractItemView_ItemDelegate(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractItemView) ItemDelegate2(index core.QModelIndex_ITF) *QAbstractItemDelegate {
	if ptr.Pointer() != nil {
		return NewQAbstractItemDelegateFromPointer(C.QAbstractItemView_ItemDelegate2(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return nil
}

func (ptr *QAbstractItemView) ItemDelegateForColumn(column int) *QAbstractItemDelegate {
	if ptr.Pointer() != nil {
		return NewQAbstractItemDelegateFromPointer(C.QAbstractItemView_ItemDelegateForColumn(ptr.Pointer(), C.int(column)))
	}
	return nil
}

func (ptr *QAbstractItemView) ItemDelegateForRow(row int) *QAbstractItemDelegate {
	if ptr.Pointer() != nil {
		return NewQAbstractItemDelegateFromPointer(C.QAbstractItemView_ItemDelegateForRow(ptr.Pointer(), C.int(row)))
	}
	return nil
}

func (ptr *QAbstractItemView) KeyboardSearch(search string) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_KeyboardSearch(ptr.Pointer(), C.CString(search))
	}
}

func (ptr *QAbstractItemView) Model() *core.QAbstractItemModel {
	if ptr.Pointer() != nil {
		return core.NewQAbstractItemModelFromPointer(C.QAbstractItemView_Model(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractItemView) OpenPersistentEditor(index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_OpenPersistentEditor(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QAbstractItemView) ConnectPressed(f func(index *core.QModelIndex)) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_ConnectPressed(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "pressed", f)
	}
}

func (ptr *QAbstractItemView) DisconnectPressed() {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_DisconnectPressed(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "pressed")
	}
}

//export callbackQAbstractItemViewPressed
func callbackQAbstractItemViewPressed(ptrName *C.char, index unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "pressed").(func(*core.QModelIndex))(core.NewQModelIndexFromPointer(index))
}

func (ptr *QAbstractItemView) Reset() {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_Reset(ptr.Pointer())
	}
}

func (ptr *QAbstractItemView) RootIndex() *core.QModelIndex {
	if ptr.Pointer() != nil {
		return core.NewQModelIndexFromPointer(C.QAbstractItemView_RootIndex(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractItemView) ScrollTo(index core.QModelIndex_ITF, hint QAbstractItemView__ScrollHint) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_ScrollTo(ptr.Pointer(), core.PointerFromQModelIndex(index), C.int(hint))
	}
}

func (ptr *QAbstractItemView) ScrollToBottom() {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_ScrollToBottom(ptr.Pointer())
	}
}

func (ptr *QAbstractItemView) ScrollToTop() {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_ScrollToTop(ptr.Pointer())
	}
}

func (ptr *QAbstractItemView) SelectAll() {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_SelectAll(ptr.Pointer())
	}
}

func (ptr *QAbstractItemView) SelectionModel() *core.QItemSelectionModel {
	if ptr.Pointer() != nil {
		return core.NewQItemSelectionModelFromPointer(C.QAbstractItemView_SelectionModel(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractItemView) SetCurrentIndex(index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetCurrentIndex(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QAbstractItemView) SetIndexWidget(index core.QModelIndex_ITF, widget QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetIndexWidget(ptr.Pointer(), core.PointerFromQModelIndex(index), PointerFromQWidget(widget))
	}
}

func (ptr *QAbstractItemView) SetItemDelegate(delegate QAbstractItemDelegate_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetItemDelegate(ptr.Pointer(), PointerFromQAbstractItemDelegate(delegate))
	}
}

func (ptr *QAbstractItemView) SetItemDelegateForColumn(column int, delegate QAbstractItemDelegate_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetItemDelegateForColumn(ptr.Pointer(), C.int(column), PointerFromQAbstractItemDelegate(delegate))
	}
}

func (ptr *QAbstractItemView) SetItemDelegateForRow(row int, delegate QAbstractItemDelegate_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetItemDelegateForRow(ptr.Pointer(), C.int(row), PointerFromQAbstractItemDelegate(delegate))
	}
}

func (ptr *QAbstractItemView) SetModel(model core.QAbstractItemModel_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetModel(ptr.Pointer(), core.PointerFromQAbstractItemModel(model))
	}
}

func (ptr *QAbstractItemView) SetRootIndex(index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetRootIndex(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QAbstractItemView) SetSelectionModel(selectionModel core.QItemSelectionModel_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetSelectionModel(ptr.Pointer(), core.PointerFromQItemSelectionModel(selectionModel))
	}
}

func (ptr *QAbstractItemView) SizeHintForColumn(column int) int {
	if ptr.Pointer() != nil {
		return int(C.QAbstractItemView_SizeHintForColumn(ptr.Pointer(), C.int(column)))
	}
	return 0
}

func (ptr *QAbstractItemView) SizeHintForRow(row int) int {
	if ptr.Pointer() != nil {
		return int(C.QAbstractItemView_SizeHintForRow(ptr.Pointer(), C.int(row)))
	}
	return 0
}

func (ptr *QAbstractItemView) Update(index core.QModelIndex_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_Update(ptr.Pointer(), core.PointerFromQModelIndex(index))
	}
}

func (ptr *QAbstractItemView) ConnectViewportEntered(f func()) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_ConnectViewportEntered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "viewportEntered", f)
	}
}

func (ptr *QAbstractItemView) DisconnectViewportEntered() {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_DisconnectViewportEntered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "viewportEntered")
	}
}

//export callbackQAbstractItemViewViewportEntered
func callbackQAbstractItemViewViewportEntered(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "viewportEntered").(func())()
}

func (ptr *QAbstractItemView) DestroyQAbstractItemView() {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_DestroyQAbstractItemView(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
