package qt

//#include "qabstractitemview.h"
import "C"

type qabstractitemview struct {
	qabstractscrollarea
}

type QAbstractItemView interface {
	QAbstractScrollArea
	AlternatingRowColors() bool
	AutoScrollMargin() int
	DefaultDropAction() DropAction
	DragDropOverwriteMode() bool
	DragEnabled() bool
	HasAutoScroll() bool
	KeyboardSearch(search string)
	SelectionMode() SelectionMode
	SetAlternatingRowColors(enable bool)
	SetAutoScroll(enable bool)
	SetAutoScrollMargin(margin int)
	SetDefaultDropAction(dropAction DropAction)
	SetDragDropOverwriteMode(overwrite bool)
	SetDragEnabled(enable bool)
	SetDropIndicatorShown(enable bool)
	SetSelectionMode(mode SelectionMode)
	SetTabKeyNavigation(enable bool)
	SetTextElideMode(mode TextElideMode)
	ShowDropIndicator() bool
	SizeHintForColumn(column int) int
	SizeHintForRow(row int) int
	TabKeyNavigation() bool
	TextElideMode() TextElideMode
	ConnectSlotClearSelection()
	DisconnectSlotClearSelection()
	SlotClearSelection()
	ConnectSlotReset()
	DisconnectSlotReset()
	SlotReset()
	ConnectSlotScrollToBottom()
	DisconnectSlotScrollToBottom()
	SlotScrollToBottom()
	ConnectSlotScrollToTop()
	DisconnectSlotScrollToTop()
	SlotScrollToTop()
	ConnectSlotSelectAll()
	DisconnectSlotSelectAll()
	SlotSelectAll()
	ConnectSignalActivated(f func())
	DisconnectSignalActivated()
	SignalActivated() func()
	ConnectSignalClicked(f func())
	DisconnectSignalClicked()
	SignalClicked() func()
	ConnectSignalDoubleClicked(f func())
	DisconnectSignalDoubleClicked()
	SignalDoubleClicked() func()
	ConnectSignalEntered(f func())
	DisconnectSignalEntered()
	SignalEntered() func()
	ConnectSignalPressed(f func())
	DisconnectSignalPressed()
	SignalPressed() func()
	ConnectSignalViewportEntered(f func())
	DisconnectSignalViewportEntered()
	SignalViewportEntered() func()
}

func (p *qabstractitemview) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qabstractitemview) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

//DragDropMode
type DragDropMode int

var (
	NODRAGDROP   = DragDropMode(C.QAbstractItemView_NoDragDrop())
	DRAGONLY     = DragDropMode(C.QAbstractItemView_DragOnly())
	DROPONLY     = DragDropMode(C.QAbstractItemView_DropOnly())
	DRAGDROP     = DragDropMode(C.QAbstractItemView_DragDrop())
	INTERNALMOVE = DragDropMode(C.QAbstractItemView_InternalMove())
)

//EditTrigger
type EditTrigger int

var (
	NOEDITTRIGGERS  = EditTrigger(C.QAbstractItemView_NoEditTriggers())
	CURRENTCHANGED  = EditTrigger(C.QAbstractItemView_CurrentChanged())
	DOUBLECLICKED   = EditTrigger(C.QAbstractItemView_DoubleClicked())
	SELECTEDCLICKED = EditTrigger(C.QAbstractItemView_SelectedClicked())
	EDITKEYPRESSED  = EditTrigger(C.QAbstractItemView_EditKeyPressed())
	ANYKEYPRESSED   = EditTrigger(C.QAbstractItemView_AnyKeyPressed())
	ALLEDITTRIGGERS = EditTrigger(C.QAbstractItemView_AllEditTriggers())
)

//ScrollHint
type ScrollHint int

var (
	ENSUREVISIBLE    = ScrollHint(C.QAbstractItemView_EnsureVisible())
	POSITIONATTOP    = ScrollHint(C.QAbstractItemView_PositionAtTop())
	POSITIONATBOTTOM = ScrollHint(C.QAbstractItemView_PositionAtBottom())
	POSITIONATCENTER = ScrollHint(C.QAbstractItemView_PositionAtCenter())
)

//ScrollMode
type ScrollMode int

var (
	SCROLLPERITEM  = ScrollMode(C.QAbstractItemView_ScrollPerItem())
	SCROLLPERPIXEL = ScrollMode(C.QAbstractItemView_ScrollPerPixel())
)

//SelectionBehavior
type SelectionBehavior int

var (
	SELECTITEMS   = SelectionBehavior(C.QAbstractItemView_SelectItems())
	SELECTROWS    = SelectionBehavior(C.QAbstractItemView_SelectRows())
	SELECTCOLUMNS = SelectionBehavior(C.QAbstractItemView_SelectColumns())
)

//SelectionMode
type SelectionMode int

var (
	SINGLESELECTION     = SelectionMode(C.QAbstractItemView_SingleSelection())
	CONTIGUOUSSELECTION = SelectionMode(C.QAbstractItemView_ContiguousSelection())
	EXTENDEDSELECTION   = SelectionMode(C.QAbstractItemView_ExtendedSelection())
	MULTISELECTION      = SelectionMode(C.QAbstractItemView_MultiSelection())
	NOSELECTION         = SelectionMode(C.QAbstractItemView_NoSelection())
)

func (p *qabstractitemview) Destroy() {
	if p.Pointer() != nil {
		getSignal(p.ObjectName(), "destroyed")()
		C.QAbstractItemView_Destroy(p.Pointer())
		p.SetPointer(nil)
	}
}

func (p *qabstractitemview) AlternatingRowColors() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QAbstractItemView_AlternatingRowColors(p.Pointer()) != 0
}

func (p *qabstractitemview) AutoScrollMargin() int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QAbstractItemView_AutoScrollMargin(p.Pointer()))
}

func (p *qabstractitemview) DefaultDropAction() DropAction {
	if p.Pointer() == nil {
		return 0
	}
	return DropAction(C.QAbstractItemView_DefaultDropAction(p.Pointer()))
}

func (p *qabstractitemview) DragDropOverwriteMode() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QAbstractItemView_DragDropOverwriteMode(p.Pointer()) != 0
}

func (p *qabstractitemview) DragEnabled() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QAbstractItemView_DragEnabled(p.Pointer()) != 0
}

func (p *qabstractitemview) HasAutoScroll() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QAbstractItemView_HasAutoScroll(p.Pointer()) != 0
}

func (p *qabstractitemview) KeyboardSearch(search string) {
	if p.Pointer() != nil {
		C.QAbstractItemView_KeyboardSearch_String(p.Pointer(), C.CString(search))
	}
}

func (p *qabstractitemview) SelectionMode() SelectionMode {
	if p.Pointer() == nil {
		return 0
	}
	return SelectionMode(C.QAbstractItemView_SelectionMode(p.Pointer()))
}

func (p *qabstractitemview) SetAlternatingRowColors(enable bool) {
	if p.Pointer() != nil {
		C.QAbstractItemView_SetAlternatingRowColors_Bool(p.Pointer(), goBoolToCInt(enable))
	}
}

func (p *qabstractitemview) SetAutoScroll(enable bool) {
	if p.Pointer() != nil {
		C.QAbstractItemView_SetAutoScroll_Bool(p.Pointer(), goBoolToCInt(enable))
	}
}

func (p *qabstractitemview) SetAutoScrollMargin(margin int) {
	if p.Pointer() != nil {
		C.QAbstractItemView_SetAutoScrollMargin_Int(p.Pointer(), C.int(margin))
	}
}

func (p *qabstractitemview) SetDefaultDropAction(dropAction DropAction) {
	if p.Pointer() != nil {
		C.QAbstractItemView_SetDefaultDropAction_DropAction(p.Pointer(), C.int(dropAction))
	}
}

func (p *qabstractitemview) SetDragDropOverwriteMode(overwrite bool) {
	if p.Pointer() != nil {
		C.QAbstractItemView_SetDragDropOverwriteMode_Bool(p.Pointer(), goBoolToCInt(overwrite))
	}
}

func (p *qabstractitemview) SetDragEnabled(enable bool) {
	if p.Pointer() != nil {
		C.QAbstractItemView_SetDragEnabled_Bool(p.Pointer(), goBoolToCInt(enable))
	}
}

func (p *qabstractitemview) SetDropIndicatorShown(enable bool) {
	if p.Pointer() != nil {
		C.QAbstractItemView_SetDropIndicatorShown_Bool(p.Pointer(), goBoolToCInt(enable))
	}
}

func (p *qabstractitemview) SetSelectionMode(mode SelectionMode) {
	if p.Pointer() != nil {
		C.QAbstractItemView_SetSelectionMode_SelectionMode(p.Pointer(), C.int(mode))
	}
}

func (p *qabstractitemview) SetTabKeyNavigation(enable bool) {
	if p.Pointer() != nil {
		C.QAbstractItemView_SetTabKeyNavigation_Bool(p.Pointer(), goBoolToCInt(enable))
	}
}

func (p *qabstractitemview) SetTextElideMode(mode TextElideMode) {
	if p.Pointer() != nil {
		C.QAbstractItemView_SetTextElideMode_TextElideMode(p.Pointer(), C.int(mode))
	}
}

func (p *qabstractitemview) ShowDropIndicator() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QAbstractItemView_ShowDropIndicator(p.Pointer()) != 0
}

func (p *qabstractitemview) SizeHintForColumn(column int) int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QAbstractItemView_SizeHintForColumn_Int(p.Pointer(), C.int(column)))
}

func (p *qabstractitemview) SizeHintForRow(row int) int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QAbstractItemView_SizeHintForRow_Int(p.Pointer(), C.int(row)))
}

func (p *qabstractitemview) TabKeyNavigation() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QAbstractItemView_TabKeyNavigation(p.Pointer()) != 0
}

func (p *qabstractitemview) TextElideMode() TextElideMode {
	if p.Pointer() == nil {
		return 0
	}
	return TextElideMode(C.QAbstractItemView_TextElideMode(p.Pointer()))
}

func (p *qabstractitemview) ConnectSlotClearSelection() {
	C.QAbstractItemView_ConnectSlotClearSelection(p.Pointer())
}

func (p *qabstractitemview) DisconnectSlotClearSelection() {
	C.QAbstractItemView_DisconnectSlotClearSelection(p.Pointer())
}

func (p *qabstractitemview) SlotClearSelection() {
	if p.Pointer() != nil {
		C.QAbstractItemView_ClearSelection(p.Pointer())
	}
}

func (p *qabstractitemview) ConnectSlotReset() {
	C.QAbstractItemView_ConnectSlotReset(p.Pointer())
}

func (p *qabstractitemview) DisconnectSlotReset() {
	C.QAbstractItemView_DisconnectSlotReset(p.Pointer())
}

func (p *qabstractitemview) SlotReset() {
	if p.Pointer() != nil {
		C.QAbstractItemView_Reset(p.Pointer())
	}
}

func (p *qabstractitemview) ConnectSlotScrollToBottom() {
	C.QAbstractItemView_ConnectSlotScrollToBottom(p.Pointer())
}

func (p *qabstractitemview) DisconnectSlotScrollToBottom() {
	C.QAbstractItemView_DisconnectSlotScrollToBottom(p.Pointer())
}

func (p *qabstractitemview) SlotScrollToBottom() {
	if p.Pointer() != nil {
		C.QAbstractItemView_ScrollToBottom(p.Pointer())
	}
}

func (p *qabstractitemview) ConnectSlotScrollToTop() {
	C.QAbstractItemView_ConnectSlotScrollToTop(p.Pointer())
}

func (p *qabstractitemview) DisconnectSlotScrollToTop() {
	C.QAbstractItemView_DisconnectSlotScrollToTop(p.Pointer())
}

func (p *qabstractitemview) SlotScrollToTop() {
	if p.Pointer() != nil {
		C.QAbstractItemView_ScrollToTop(p.Pointer())
	}
}

func (p *qabstractitemview) ConnectSlotSelectAll() {
	C.QAbstractItemView_ConnectSlotSelectAll(p.Pointer())
}

func (p *qabstractitemview) DisconnectSlotSelectAll() {
	C.QAbstractItemView_DisconnectSlotSelectAll(p.Pointer())
}

func (p *qabstractitemview) SlotSelectAll() {
	if p.Pointer() != nil {
		C.QAbstractItemView_SelectAll(p.Pointer())
	}
}

func (p *qabstractitemview) ConnectSignalActivated(f func()) {
	C.QAbstractItemView_ConnectSignalActivated(p.Pointer())
	connectSignal(p.ObjectName(), "activated", f)
}

func (p *qabstractitemview) DisconnectSignalActivated() {
	C.QAbstractItemView_DisconnectSignalActivated(p.Pointer())
	disconnectSignal(p.ObjectName(), "activated")
}

func (p *qabstractitemview) SignalActivated() func() {
	return getSignal(p.ObjectName(), "activated")
}

func (p *qabstractitemview) ConnectSignalClicked(f func()) {
	C.QAbstractItemView_ConnectSignalClicked(p.Pointer())
	connectSignal(p.ObjectName(), "clicked", f)
}

func (p *qabstractitemview) DisconnectSignalClicked() {
	C.QAbstractItemView_DisconnectSignalClicked(p.Pointer())
	disconnectSignal(p.ObjectName(), "clicked")
}

func (p *qabstractitemview) SignalClicked() func() {
	return getSignal(p.ObjectName(), "clicked")
}

func (p *qabstractitemview) ConnectSignalDoubleClicked(f func()) {
	C.QAbstractItemView_ConnectSignalDoubleClicked(p.Pointer())
	connectSignal(p.ObjectName(), "doubleClicked", f)
}

func (p *qabstractitemview) DisconnectSignalDoubleClicked() {
	C.QAbstractItemView_DisconnectSignalDoubleClicked(p.Pointer())
	disconnectSignal(p.ObjectName(), "doubleClicked")
}

func (p *qabstractitemview) SignalDoubleClicked() func() {
	return getSignal(p.ObjectName(), "doubleClicked")
}

func (p *qabstractitemview) ConnectSignalEntered(f func()) {
	C.QAbstractItemView_ConnectSignalEntered(p.Pointer())
	connectSignal(p.ObjectName(), "entered", f)
}

func (p *qabstractitemview) DisconnectSignalEntered() {
	C.QAbstractItemView_DisconnectSignalEntered(p.Pointer())
	disconnectSignal(p.ObjectName(), "entered")
}

func (p *qabstractitemview) SignalEntered() func() {
	return getSignal(p.ObjectName(), "entered")
}

func (p *qabstractitemview) ConnectSignalPressed(f func()) {
	C.QAbstractItemView_ConnectSignalPressed(p.Pointer())
	connectSignal(p.ObjectName(), "pressed", f)
}

func (p *qabstractitemview) DisconnectSignalPressed() {
	C.QAbstractItemView_DisconnectSignalPressed(p.Pointer())
	disconnectSignal(p.ObjectName(), "pressed")
}

func (p *qabstractitemview) SignalPressed() func() {
	return getSignal(p.ObjectName(), "pressed")
}

func (p *qabstractitemview) ConnectSignalViewportEntered(f func()) {
	C.QAbstractItemView_ConnectSignalViewportEntered(p.Pointer())
	connectSignal(p.ObjectName(), "viewportEntered", f)
}

func (p *qabstractitemview) DisconnectSignalViewportEntered() {
	C.QAbstractItemView_DisconnectSignalViewportEntered(p.Pointer())
	disconnectSignal(p.ObjectName(), "viewportEntered")
}

func (p *qabstractitemview) SignalViewportEntered() func() {
	return getSignal(p.ObjectName(), "viewportEntered")
}

//export callbackQAbstractItemView
func callbackQAbstractItemView(ptr C.QtObjectPtr, msg *C.char) {
	var qabstractitemview = new(qabstractitemview)
	qabstractitemview.SetPointer(ptr)
	getSignal(qabstractitemview.ObjectName(), C.GoString(msg))()
}
