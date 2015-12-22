package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QUndoStack struct {
	core.QObject
}

type QUndoStack_ITF interface {
	core.QObject_ITF
	QUndoStack_PTR() *QUndoStack
}

func PointerFromQUndoStack(ptr QUndoStack_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QUndoStack_PTR().Pointer()
	}
	return nil
}

func NewQUndoStackFromPointer(ptr unsafe.Pointer) *QUndoStack {
	var n = new(QUndoStack)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QUndoStack_") {
		n.SetObjectName("QUndoStack_" + qt.Identifier())
	}
	return n
}

func (ptr *QUndoStack) QUndoStack_PTR() *QUndoStack {
	return ptr
}

func (ptr *QUndoStack) IsActive() bool {
	defer qt.Recovering("QUndoStack::isActive")

	if ptr.Pointer() != nil {
		return C.QUndoStack_IsActive(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QUndoStack) SetActive(active bool) {
	defer qt.Recovering("QUndoStack::setActive")

	if ptr.Pointer() != nil {
		C.QUndoStack_SetActive(ptr.Pointer(), C.int(qt.GoBoolToInt(active)))
	}
}

func (ptr *QUndoStack) SetUndoLimit(limit int) {
	defer qt.Recovering("QUndoStack::setUndoLimit")

	if ptr.Pointer() != nil {
		C.QUndoStack_SetUndoLimit(ptr.Pointer(), C.int(limit))
	}
}

func (ptr *QUndoStack) UndoLimit() int {
	defer qt.Recovering("QUndoStack::undoLimit")

	if ptr.Pointer() != nil {
		return int(C.QUndoStack_UndoLimit(ptr.Pointer()))
	}
	return 0
}

func NewQUndoStack(parent core.QObject_ITF) *QUndoStack {
	defer qt.Recovering("QUndoStack::QUndoStack")

	return NewQUndoStackFromPointer(C.QUndoStack_NewQUndoStack(core.PointerFromQObject(parent)))
}

func (ptr *QUndoStack) BeginMacro(text string) {
	defer qt.Recovering("QUndoStack::beginMacro")

	if ptr.Pointer() != nil {
		C.QUndoStack_BeginMacro(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QUndoStack) CanRedo() bool {
	defer qt.Recovering("QUndoStack::canRedo")

	if ptr.Pointer() != nil {
		return C.QUndoStack_CanRedo(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QUndoStack) ConnectCanRedoChanged(f func(canRedo bool)) {
	defer qt.Recovering("connect QUndoStack::canRedoChanged")

	if ptr.Pointer() != nil {
		C.QUndoStack_ConnectCanRedoChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "canRedoChanged", f)
	}
}

func (ptr *QUndoStack) DisconnectCanRedoChanged() {
	defer qt.Recovering("disconnect QUndoStack::canRedoChanged")

	if ptr.Pointer() != nil {
		C.QUndoStack_DisconnectCanRedoChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "canRedoChanged")
	}
}

//export callbackQUndoStackCanRedoChanged
func callbackQUndoStackCanRedoChanged(ptrName *C.char, canRedo C.int) {
	defer qt.Recovering("callback QUndoStack::canRedoChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "canRedoChanged"); signal != nil {
		signal.(func(bool))(int(canRedo) != 0)
	}

}

func (ptr *QUndoStack) CanUndo() bool {
	defer qt.Recovering("QUndoStack::canUndo")

	if ptr.Pointer() != nil {
		return C.QUndoStack_CanUndo(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QUndoStack) ConnectCanUndoChanged(f func(canUndo bool)) {
	defer qt.Recovering("connect QUndoStack::canUndoChanged")

	if ptr.Pointer() != nil {
		C.QUndoStack_ConnectCanUndoChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "canUndoChanged", f)
	}
}

func (ptr *QUndoStack) DisconnectCanUndoChanged() {
	defer qt.Recovering("disconnect QUndoStack::canUndoChanged")

	if ptr.Pointer() != nil {
		C.QUndoStack_DisconnectCanUndoChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "canUndoChanged")
	}
}

//export callbackQUndoStackCanUndoChanged
func callbackQUndoStackCanUndoChanged(ptrName *C.char, canUndo C.int) {
	defer qt.Recovering("callback QUndoStack::canUndoChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "canUndoChanged"); signal != nil {
		signal.(func(bool))(int(canUndo) != 0)
	}

}

func (ptr *QUndoStack) ConnectCleanChanged(f func(clean bool)) {
	defer qt.Recovering("connect QUndoStack::cleanChanged")

	if ptr.Pointer() != nil {
		C.QUndoStack_ConnectCleanChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "cleanChanged", f)
	}
}

func (ptr *QUndoStack) DisconnectCleanChanged() {
	defer qt.Recovering("disconnect QUndoStack::cleanChanged")

	if ptr.Pointer() != nil {
		C.QUndoStack_DisconnectCleanChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "cleanChanged")
	}
}

//export callbackQUndoStackCleanChanged
func callbackQUndoStackCleanChanged(ptrName *C.char, clean C.int) {
	defer qt.Recovering("callback QUndoStack::cleanChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "cleanChanged"); signal != nil {
		signal.(func(bool))(int(clean) != 0)
	}

}

func (ptr *QUndoStack) CleanIndex() int {
	defer qt.Recovering("QUndoStack::cleanIndex")

	if ptr.Pointer() != nil {
		return int(C.QUndoStack_CleanIndex(ptr.Pointer()))
	}
	return 0
}

func (ptr *QUndoStack) Clear() {
	defer qt.Recovering("QUndoStack::clear")

	if ptr.Pointer() != nil {
		C.QUndoStack_Clear(ptr.Pointer())
	}
}

func (ptr *QUndoStack) Command(index int) *QUndoCommand {
	defer qt.Recovering("QUndoStack::command")

	if ptr.Pointer() != nil {
		return NewQUndoCommandFromPointer(C.QUndoStack_Command(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QUndoStack) Count() int {
	defer qt.Recovering("QUndoStack::count")

	if ptr.Pointer() != nil {
		return int(C.QUndoStack_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QUndoStack) CreateRedoAction(parent core.QObject_ITF, prefix string) *QAction {
	defer qt.Recovering("QUndoStack::createRedoAction")

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QUndoStack_CreateRedoAction(ptr.Pointer(), core.PointerFromQObject(parent), C.CString(prefix)))
	}
	return nil
}

func (ptr *QUndoStack) CreateUndoAction(parent core.QObject_ITF, prefix string) *QAction {
	defer qt.Recovering("QUndoStack::createUndoAction")

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QUndoStack_CreateUndoAction(ptr.Pointer(), core.PointerFromQObject(parent), C.CString(prefix)))
	}
	return nil
}

func (ptr *QUndoStack) EndMacro() {
	defer qt.Recovering("QUndoStack::endMacro")

	if ptr.Pointer() != nil {
		C.QUndoStack_EndMacro(ptr.Pointer())
	}
}

func (ptr *QUndoStack) Index() int {
	defer qt.Recovering("QUndoStack::index")

	if ptr.Pointer() != nil {
		return int(C.QUndoStack_Index(ptr.Pointer()))
	}
	return 0
}

func (ptr *QUndoStack) ConnectIndexChanged(f func(idx int)) {
	defer qt.Recovering("connect QUndoStack::indexChanged")

	if ptr.Pointer() != nil {
		C.QUndoStack_ConnectIndexChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "indexChanged", f)
	}
}

func (ptr *QUndoStack) DisconnectIndexChanged() {
	defer qt.Recovering("disconnect QUndoStack::indexChanged")

	if ptr.Pointer() != nil {
		C.QUndoStack_DisconnectIndexChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "indexChanged")
	}
}

//export callbackQUndoStackIndexChanged
func callbackQUndoStackIndexChanged(ptrName *C.char, idx C.int) {
	defer qt.Recovering("callback QUndoStack::indexChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "indexChanged"); signal != nil {
		signal.(func(int))(int(idx))
	}

}

func (ptr *QUndoStack) IsClean() bool {
	defer qt.Recovering("QUndoStack::isClean")

	if ptr.Pointer() != nil {
		return C.QUndoStack_IsClean(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QUndoStack) Push(cmd QUndoCommand_ITF) {
	defer qt.Recovering("QUndoStack::push")

	if ptr.Pointer() != nil {
		C.QUndoStack_Push(ptr.Pointer(), PointerFromQUndoCommand(cmd))
	}
}

func (ptr *QUndoStack) Redo() {
	defer qt.Recovering("QUndoStack::redo")

	if ptr.Pointer() != nil {
		C.QUndoStack_Redo(ptr.Pointer())
	}
}

func (ptr *QUndoStack) RedoText() string {
	defer qt.Recovering("QUndoStack::redoText")

	if ptr.Pointer() != nil {
		return C.GoString(C.QUndoStack_RedoText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QUndoStack) ConnectRedoTextChanged(f func(redoText string)) {
	defer qt.Recovering("connect QUndoStack::redoTextChanged")

	if ptr.Pointer() != nil {
		C.QUndoStack_ConnectRedoTextChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "redoTextChanged", f)
	}
}

func (ptr *QUndoStack) DisconnectRedoTextChanged() {
	defer qt.Recovering("disconnect QUndoStack::redoTextChanged")

	if ptr.Pointer() != nil {
		C.QUndoStack_DisconnectRedoTextChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "redoTextChanged")
	}
}

//export callbackQUndoStackRedoTextChanged
func callbackQUndoStackRedoTextChanged(ptrName *C.char, redoText *C.char) {
	defer qt.Recovering("callback QUndoStack::redoTextChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "redoTextChanged"); signal != nil {
		signal.(func(string))(C.GoString(redoText))
	}

}

func (ptr *QUndoStack) SetClean() {
	defer qt.Recovering("QUndoStack::setClean")

	if ptr.Pointer() != nil {
		C.QUndoStack_SetClean(ptr.Pointer())
	}
}

func (ptr *QUndoStack) SetIndex(idx int) {
	defer qt.Recovering("QUndoStack::setIndex")

	if ptr.Pointer() != nil {
		C.QUndoStack_SetIndex(ptr.Pointer(), C.int(idx))
	}
}

func (ptr *QUndoStack) Text(idx int) string {
	defer qt.Recovering("QUndoStack::text")

	if ptr.Pointer() != nil {
		return C.GoString(C.QUndoStack_Text(ptr.Pointer(), C.int(idx)))
	}
	return ""
}

func (ptr *QUndoStack) Undo() {
	defer qt.Recovering("QUndoStack::undo")

	if ptr.Pointer() != nil {
		C.QUndoStack_Undo(ptr.Pointer())
	}
}

func (ptr *QUndoStack) UndoText() string {
	defer qt.Recovering("QUndoStack::undoText")

	if ptr.Pointer() != nil {
		return C.GoString(C.QUndoStack_UndoText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QUndoStack) ConnectUndoTextChanged(f func(undoText string)) {
	defer qt.Recovering("connect QUndoStack::undoTextChanged")

	if ptr.Pointer() != nil {
		C.QUndoStack_ConnectUndoTextChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "undoTextChanged", f)
	}
}

func (ptr *QUndoStack) DisconnectUndoTextChanged() {
	defer qt.Recovering("disconnect QUndoStack::undoTextChanged")

	if ptr.Pointer() != nil {
		C.QUndoStack_DisconnectUndoTextChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "undoTextChanged")
	}
}

//export callbackQUndoStackUndoTextChanged
func callbackQUndoStackUndoTextChanged(ptrName *C.char, undoText *C.char) {
	defer qt.Recovering("callback QUndoStack::undoTextChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "undoTextChanged"); signal != nil {
		signal.(func(string))(C.GoString(undoText))
	}

}

func (ptr *QUndoStack) DestroyQUndoStack() {
	defer qt.Recovering("QUndoStack::~QUndoStack")

	if ptr.Pointer() != nil {
		C.QUndoStack_DestroyQUndoStack(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
