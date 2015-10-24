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

type QAbstractItemViewITF interface {
	QAbstractScrollAreaITF
	QAbstractItemViewPTR() *QAbstractItemView
}

func PointerFromQAbstractItemView(ptr QAbstractItemViewITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractItemViewPTR().Pointer()
	}
	return nil
}

func QAbstractItemViewFromPointer(ptr unsafe.Pointer) *QAbstractItemView {
	var n = new(QAbstractItemView)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QAbstractItemView_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAbstractItemView) QAbstractItemViewPTR() *QAbstractItemView {
	return ptr
}

//QAbstractItemView::DragDropMode
type QAbstractItemView__DragDropMode int

var (
	QAbstractItemView__NoDragDrop   = QAbstractItemView__DragDropMode(0)
	QAbstractItemView__DragOnly     = QAbstractItemView__DragDropMode(1)
	QAbstractItemView__DropOnly     = QAbstractItemView__DragDropMode(2)
	QAbstractItemView__DragDrop     = QAbstractItemView__DragDropMode(3)
	QAbstractItemView__InternalMove = QAbstractItemView__DragDropMode(4)
)

//QAbstractItemView::EditTrigger
type QAbstractItemView__EditTrigger int

var (
	QAbstractItemView__NoEditTriggers  = QAbstractItemView__EditTrigger(0)
	QAbstractItemView__CurrentChanged  = QAbstractItemView__EditTrigger(1)
	QAbstractItemView__DoubleClicked   = QAbstractItemView__EditTrigger(2)
	QAbstractItemView__SelectedClicked = QAbstractItemView__EditTrigger(4)
	QAbstractItemView__EditKeyPressed  = QAbstractItemView__EditTrigger(8)
	QAbstractItemView__AnyKeyPressed   = QAbstractItemView__EditTrigger(16)
	QAbstractItemView__AllEditTriggers = QAbstractItemView__EditTrigger(31)
)

//QAbstractItemView::ScrollHint
type QAbstractItemView__ScrollHint int

var (
	QAbstractItemView__EnsureVisible    = QAbstractItemView__ScrollHint(0)
	QAbstractItemView__PositionAtTop    = QAbstractItemView__ScrollHint(1)
	QAbstractItemView__PositionAtBottom = QAbstractItemView__ScrollHint(2)
	QAbstractItemView__PositionAtCenter = QAbstractItemView__ScrollHint(3)
)

//QAbstractItemView::ScrollMode
type QAbstractItemView__ScrollMode int

var (
	QAbstractItemView__ScrollPerItem  = QAbstractItemView__ScrollMode(0)
	QAbstractItemView__ScrollPerPixel = QAbstractItemView__ScrollMode(1)
)

//QAbstractItemView::SelectionBehavior
type QAbstractItemView__SelectionBehavior int

var (
	QAbstractItemView__SelectItems   = QAbstractItemView__SelectionBehavior(0)
	QAbstractItemView__SelectRows    = QAbstractItemView__SelectionBehavior(1)
	QAbstractItemView__SelectColumns = QAbstractItemView__SelectionBehavior(2)
)

//QAbstractItemView::SelectionMode
type QAbstractItemView__SelectionMode int

var (
	QAbstractItemView__NoSelection         = QAbstractItemView__SelectionMode(0)
	QAbstractItemView__SingleSelection     = QAbstractItemView__SelectionMode(1)
	QAbstractItemView__MultiSelection      = QAbstractItemView__SelectionMode(2)
	QAbstractItemView__ExtendedSelection   = QAbstractItemView__SelectionMode(3)
	QAbstractItemView__ContiguousSelection = QAbstractItemView__SelectionMode(4)
)

func (ptr *QAbstractItemView) AlternatingRowColors() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractItemView_AlternatingRowColors(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QAbstractItemView) AutoScrollMargin() int {
	if ptr.Pointer() != nil {
		return int(C.QAbstractItemView_AutoScrollMargin(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAbstractItemView) DefaultDropAction() core.Qt__DropAction {
	if ptr.Pointer() != nil {
		return core.Qt__DropAction(C.QAbstractItemView_DefaultDropAction(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAbstractItemView) DragDropMode() QAbstractItemView__DragDropMode {
	if ptr.Pointer() != nil {
		return QAbstractItemView__DragDropMode(C.QAbstractItemView_DragDropMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAbstractItemView) DragDropOverwriteMode() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractItemView_DragDropOverwriteMode(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QAbstractItemView) DragEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractItemView_DragEnabled(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QAbstractItemView) EditTriggers() QAbstractItemView__EditTrigger {
	if ptr.Pointer() != nil {
		return QAbstractItemView__EditTrigger(C.QAbstractItemView_EditTriggers(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAbstractItemView) HasAutoScroll() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractItemView_HasAutoScroll(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QAbstractItemView) HorizontalScrollMode() QAbstractItemView__ScrollMode {
	if ptr.Pointer() != nil {
		return QAbstractItemView__ScrollMode(C.QAbstractItemView_HorizontalScrollMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAbstractItemView) SelectionBehavior() QAbstractItemView__SelectionBehavior {
	if ptr.Pointer() != nil {
		return QAbstractItemView__SelectionBehavior(C.QAbstractItemView_SelectionBehavior(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAbstractItemView) SelectionMode() QAbstractItemView__SelectionMode {
	if ptr.Pointer() != nil {
		return QAbstractItemView__SelectionMode(C.QAbstractItemView_SelectionMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAbstractItemView) SetAlternatingRowColors(enable bool) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetAlternatingRowColors(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QAbstractItemView) SetAutoScroll(enable bool) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetAutoScroll(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QAbstractItemView) SetAutoScrollMargin(margin int) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetAutoScrollMargin(C.QtObjectPtr(ptr.Pointer()), C.int(margin))
	}
}

func (ptr *QAbstractItemView) SetDefaultDropAction(dropAction core.Qt__DropAction) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetDefaultDropAction(C.QtObjectPtr(ptr.Pointer()), C.int(dropAction))
	}
}

func (ptr *QAbstractItemView) SetDragDropMode(behavior QAbstractItemView__DragDropMode) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetDragDropMode(C.QtObjectPtr(ptr.Pointer()), C.int(behavior))
	}
}

func (ptr *QAbstractItemView) SetDragDropOverwriteMode(overwrite bool) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetDragDropOverwriteMode(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(overwrite)))
	}
}

func (ptr *QAbstractItemView) SetDragEnabled(enable bool) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetDragEnabled(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QAbstractItemView) SetDropIndicatorShown(enable bool) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetDropIndicatorShown(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QAbstractItemView) SetEditTriggers(triggers QAbstractItemView__EditTrigger) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetEditTriggers(C.QtObjectPtr(ptr.Pointer()), C.int(triggers))
	}
}

func (ptr *QAbstractItemView) SetHorizontalScrollMode(mode QAbstractItemView__ScrollMode) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetHorizontalScrollMode(C.QtObjectPtr(ptr.Pointer()), C.int(mode))
	}
}

func (ptr *QAbstractItemView) SetIconSize(size core.QSizeITF) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetIconSize(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQSize(size)))
	}
}

func (ptr *QAbstractItemView) SetSelectionBehavior(behavior QAbstractItemView__SelectionBehavior) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetSelectionBehavior(C.QtObjectPtr(ptr.Pointer()), C.int(behavior))
	}
}

func (ptr *QAbstractItemView) SetSelectionMode(mode QAbstractItemView__SelectionMode) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetSelectionMode(C.QtObjectPtr(ptr.Pointer()), C.int(mode))
	}
}

func (ptr *QAbstractItemView) SetTabKeyNavigation(enable bool) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetTabKeyNavigation(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QAbstractItemView) SetTextElideMode(mode core.Qt__TextElideMode) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetTextElideMode(C.QtObjectPtr(ptr.Pointer()), C.int(mode))
	}
}

func (ptr *QAbstractItemView) SetVerticalScrollMode(mode QAbstractItemView__ScrollMode) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetVerticalScrollMode(C.QtObjectPtr(ptr.Pointer()), C.int(mode))
	}
}

func (ptr *QAbstractItemView) ShowDropIndicator() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractItemView_ShowDropIndicator(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QAbstractItemView) TabKeyNavigation() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractItemView_TabKeyNavigation(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QAbstractItemView) TextElideMode() core.Qt__TextElideMode {
	if ptr.Pointer() != nil {
		return core.Qt__TextElideMode(C.QAbstractItemView_TextElideMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAbstractItemView) VerticalScrollMode() QAbstractItemView__ScrollMode {
	if ptr.Pointer() != nil {
		return QAbstractItemView__ScrollMode(C.QAbstractItemView_VerticalScrollMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAbstractItemView) ConnectActivated(f func(index core.QModelIndexITF)) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_ConnectActivated(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "activated", f)
	}
}

func (ptr *QAbstractItemView) DisconnectActivated() {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_DisconnectActivated(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "activated")
	}
}

//export callbackQAbstractItemViewActivated
func callbackQAbstractItemViewActivated(ptrName *C.char, index unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "activated").(func(*core.QModelIndex))(core.QModelIndexFromPointer(index))
}

func (ptr *QAbstractItemView) ClearSelection() {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_ClearSelection(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QAbstractItemView) ConnectClicked(f func(index core.QModelIndexITF)) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_ConnectClicked(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "clicked", f)
	}
}

func (ptr *QAbstractItemView) DisconnectClicked() {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_DisconnectClicked(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "clicked")
	}
}

//export callbackQAbstractItemViewClicked
func callbackQAbstractItemViewClicked(ptrName *C.char, index unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "clicked").(func(*core.QModelIndex))(core.QModelIndexFromPointer(index))
}

func (ptr *QAbstractItemView) ClosePersistentEditor(index core.QModelIndexITF) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_ClosePersistentEditor(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQModelIndex(index)))
	}
}

func (ptr *QAbstractItemView) CurrentIndex() *core.QModelIndex {
	if ptr.Pointer() != nil {
		return core.QModelIndexFromPointer(unsafe.Pointer(C.QAbstractItemView_CurrentIndex(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QAbstractItemView) ConnectDoubleClicked(f func(index core.QModelIndexITF)) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_ConnectDoubleClicked(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "doubleClicked", f)
	}
}

func (ptr *QAbstractItemView) DisconnectDoubleClicked() {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_DisconnectDoubleClicked(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "doubleClicked")
	}
}

//export callbackQAbstractItemViewDoubleClicked
func callbackQAbstractItemViewDoubleClicked(ptrName *C.char, index unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "doubleClicked").(func(*core.QModelIndex))(core.QModelIndexFromPointer(index))
}

func (ptr *QAbstractItemView) Edit(index core.QModelIndexITF) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_Edit(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQModelIndex(index)))
	}
}

func (ptr *QAbstractItemView) ConnectEntered(f func(index core.QModelIndexITF)) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_ConnectEntered(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "entered", f)
	}
}

func (ptr *QAbstractItemView) DisconnectEntered() {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_DisconnectEntered(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "entered")
	}
}

//export callbackQAbstractItemViewEntered
func callbackQAbstractItemViewEntered(ptrName *C.char, index unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "entered").(func(*core.QModelIndex))(core.QModelIndexFromPointer(index))
}

func (ptr *QAbstractItemView) IndexAt(point core.QPointITF) *core.QModelIndex {
	if ptr.Pointer() != nil {
		return core.QModelIndexFromPointer(unsafe.Pointer(C.QAbstractItemView_IndexAt(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPoint(point)))))
	}
	return nil
}

func (ptr *QAbstractItemView) IndexWidget(index core.QModelIndexITF) *QWidget {
	if ptr.Pointer() != nil {
		return QWidgetFromPointer(unsafe.Pointer(C.QAbstractItemView_IndexWidget(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQModelIndex(index)))))
	}
	return nil
}

func (ptr *QAbstractItemView) InputMethodQuery(query core.Qt__InputMethodQuery) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAbstractItemView_InputMethodQuery(C.QtObjectPtr(ptr.Pointer()), C.int(query)))
	}
	return ""
}

func (ptr *QAbstractItemView) ItemDelegate() *QAbstractItemDelegate {
	if ptr.Pointer() != nil {
		return QAbstractItemDelegateFromPointer(unsafe.Pointer(C.QAbstractItemView_ItemDelegate(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QAbstractItemView) ItemDelegate2(index core.QModelIndexITF) *QAbstractItemDelegate {
	if ptr.Pointer() != nil {
		return QAbstractItemDelegateFromPointer(unsafe.Pointer(C.QAbstractItemView_ItemDelegate2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQModelIndex(index)))))
	}
	return nil
}

func (ptr *QAbstractItemView) ItemDelegateForColumn(column int) *QAbstractItemDelegate {
	if ptr.Pointer() != nil {
		return QAbstractItemDelegateFromPointer(unsafe.Pointer(C.QAbstractItemView_ItemDelegateForColumn(C.QtObjectPtr(ptr.Pointer()), C.int(column))))
	}
	return nil
}

func (ptr *QAbstractItemView) ItemDelegateForRow(row int) *QAbstractItemDelegate {
	if ptr.Pointer() != nil {
		return QAbstractItemDelegateFromPointer(unsafe.Pointer(C.QAbstractItemView_ItemDelegateForRow(C.QtObjectPtr(ptr.Pointer()), C.int(row))))
	}
	return nil
}

func (ptr *QAbstractItemView) KeyboardSearch(search string) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_KeyboardSearch(C.QtObjectPtr(ptr.Pointer()), C.CString(search))
	}
}

func (ptr *QAbstractItemView) Model() *core.QAbstractItemModel {
	if ptr.Pointer() != nil {
		return core.QAbstractItemModelFromPointer(unsafe.Pointer(C.QAbstractItemView_Model(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QAbstractItemView) OpenPersistentEditor(index core.QModelIndexITF) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_OpenPersistentEditor(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQModelIndex(index)))
	}
}

func (ptr *QAbstractItemView) ConnectPressed(f func(index core.QModelIndexITF)) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_ConnectPressed(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "pressed", f)
	}
}

func (ptr *QAbstractItemView) DisconnectPressed() {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_DisconnectPressed(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "pressed")
	}
}

//export callbackQAbstractItemViewPressed
func callbackQAbstractItemViewPressed(ptrName *C.char, index unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "pressed").(func(*core.QModelIndex))(core.QModelIndexFromPointer(index))
}

func (ptr *QAbstractItemView) Reset() {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_Reset(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QAbstractItemView) RootIndex() *core.QModelIndex {
	if ptr.Pointer() != nil {
		return core.QModelIndexFromPointer(unsafe.Pointer(C.QAbstractItemView_RootIndex(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QAbstractItemView) ScrollTo(index core.QModelIndexITF, hint QAbstractItemView__ScrollHint) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_ScrollTo(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQModelIndex(index)), C.int(hint))
	}
}

func (ptr *QAbstractItemView) ScrollToBottom() {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_ScrollToBottom(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QAbstractItemView) ScrollToTop() {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_ScrollToTop(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QAbstractItemView) SelectAll() {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_SelectAll(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QAbstractItemView) SelectionModel() *core.QItemSelectionModel {
	if ptr.Pointer() != nil {
		return core.QItemSelectionModelFromPointer(unsafe.Pointer(C.QAbstractItemView_SelectionModel(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QAbstractItemView) SetCurrentIndex(index core.QModelIndexITF) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetCurrentIndex(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQModelIndex(index)))
	}
}

func (ptr *QAbstractItemView) SetIndexWidget(index core.QModelIndexITF, widget QWidgetITF) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetIndexWidget(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQModelIndex(index)), C.QtObjectPtr(PointerFromQWidget(widget)))
	}
}

func (ptr *QAbstractItemView) SetItemDelegate(delegate QAbstractItemDelegateITF) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetItemDelegate(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAbstractItemDelegate(delegate)))
	}
}

func (ptr *QAbstractItemView) SetItemDelegateForColumn(column int, delegate QAbstractItemDelegateITF) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetItemDelegateForColumn(C.QtObjectPtr(ptr.Pointer()), C.int(column), C.QtObjectPtr(PointerFromQAbstractItemDelegate(delegate)))
	}
}

func (ptr *QAbstractItemView) SetItemDelegateForRow(row int, delegate QAbstractItemDelegateITF) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetItemDelegateForRow(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.QtObjectPtr(PointerFromQAbstractItemDelegate(delegate)))
	}
}

func (ptr *QAbstractItemView) SetModel(model core.QAbstractItemModelITF) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetModel(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQAbstractItemModel(model)))
	}
}

func (ptr *QAbstractItemView) SetRootIndex(index core.QModelIndexITF) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetRootIndex(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQModelIndex(index)))
	}
}

func (ptr *QAbstractItemView) SetSelectionModel(selectionModel core.QItemSelectionModelITF) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_SetSelectionModel(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQItemSelectionModel(selectionModel)))
	}
}

func (ptr *QAbstractItemView) SizeHintForColumn(column int) int {
	if ptr.Pointer() != nil {
		return int(C.QAbstractItemView_SizeHintForColumn(C.QtObjectPtr(ptr.Pointer()), C.int(column)))
	}
	return 0
}

func (ptr *QAbstractItemView) SizeHintForRow(row int) int {
	if ptr.Pointer() != nil {
		return int(C.QAbstractItemView_SizeHintForRow(C.QtObjectPtr(ptr.Pointer()), C.int(row)))
	}
	return 0
}

func (ptr *QAbstractItemView) Update(index core.QModelIndexITF) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_Update(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQModelIndex(index)))
	}
}

func (ptr *QAbstractItemView) ConnectViewportEntered(f func()) {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_ConnectViewportEntered(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "viewportEntered", f)
	}
}

func (ptr *QAbstractItemView) DisconnectViewportEntered() {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_DisconnectViewportEntered(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "viewportEntered")
	}
}

//export callbackQAbstractItemViewViewportEntered
func callbackQAbstractItemViewViewportEntered(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "viewportEntered").(func())()
}

func (ptr *QAbstractItemView) DestroyQAbstractItemView() {
	if ptr.Pointer() != nil {
		C.QAbstractItemView_DestroyQAbstractItemView(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
