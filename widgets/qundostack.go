package widgets

//#include "qundostack.h"
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
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QUndoStack_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QUndoStack) QUndoStack_PTR() *QUndoStack {
	return ptr
}

func (ptr *QUndoStack) IsActive() bool {
	if ptr.Pointer() != nil {
		return C.QUndoStack_IsActive(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QUndoStack) SetActive(active bool) {
	if ptr.Pointer() != nil {
		C.QUndoStack_SetActive(ptr.Pointer(), C.int(qt.GoBoolToInt(active)))
	}
}

func (ptr *QUndoStack) SetUndoLimit(limit int) {
	if ptr.Pointer() != nil {
		C.QUndoStack_SetUndoLimit(ptr.Pointer(), C.int(limit))
	}
}

func (ptr *QUndoStack) UndoLimit() int {
	if ptr.Pointer() != nil {
		return int(C.QUndoStack_UndoLimit(ptr.Pointer()))
	}
	return 0
}

func NewQUndoStack(parent core.QObject_ITF) *QUndoStack {
	return NewQUndoStackFromPointer(C.QUndoStack_NewQUndoStack(core.PointerFromQObject(parent)))
}

func (ptr *QUndoStack) BeginMacro(text string) {
	if ptr.Pointer() != nil {
		C.QUndoStack_BeginMacro(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QUndoStack) CanRedo() bool {
	if ptr.Pointer() != nil {
		return C.QUndoStack_CanRedo(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QUndoStack) ConnectCanRedoChanged(f func(canRedo bool)) {
	if ptr.Pointer() != nil {
		C.QUndoStack_ConnectCanRedoChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "canRedoChanged", f)
	}
}

func (ptr *QUndoStack) DisconnectCanRedoChanged() {
	if ptr.Pointer() != nil {
		C.QUndoStack_DisconnectCanRedoChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "canRedoChanged")
	}
}

//export callbackQUndoStackCanRedoChanged
func callbackQUndoStackCanRedoChanged(ptrName *C.char, canRedo C.int) {
	qt.GetSignal(C.GoString(ptrName), "canRedoChanged").(func(bool))(int(canRedo) != 0)
}

func (ptr *QUndoStack) CanUndo() bool {
	if ptr.Pointer() != nil {
		return C.QUndoStack_CanUndo(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QUndoStack) ConnectCanUndoChanged(f func(canUndo bool)) {
	if ptr.Pointer() != nil {
		C.QUndoStack_ConnectCanUndoChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "canUndoChanged", f)
	}
}

func (ptr *QUndoStack) DisconnectCanUndoChanged() {
	if ptr.Pointer() != nil {
		C.QUndoStack_DisconnectCanUndoChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "canUndoChanged")
	}
}

//export callbackQUndoStackCanUndoChanged
func callbackQUndoStackCanUndoChanged(ptrName *C.char, canUndo C.int) {
	qt.GetSignal(C.GoString(ptrName), "canUndoChanged").(func(bool))(int(canUndo) != 0)
}

func (ptr *QUndoStack) ConnectCleanChanged(f func(clean bool)) {
	if ptr.Pointer() != nil {
		C.QUndoStack_ConnectCleanChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "cleanChanged", f)
	}
}

func (ptr *QUndoStack) DisconnectCleanChanged() {
	if ptr.Pointer() != nil {
		C.QUndoStack_DisconnectCleanChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "cleanChanged")
	}
}

//export callbackQUndoStackCleanChanged
func callbackQUndoStackCleanChanged(ptrName *C.char, clean C.int) {
	qt.GetSignal(C.GoString(ptrName), "cleanChanged").(func(bool))(int(clean) != 0)
}

func (ptr *QUndoStack) CleanIndex() int {
	if ptr.Pointer() != nil {
		return int(C.QUndoStack_CleanIndex(ptr.Pointer()))
	}
	return 0
}

func (ptr *QUndoStack) Clear() {
	if ptr.Pointer() != nil {
		C.QUndoStack_Clear(ptr.Pointer())
	}
}

func (ptr *QUndoStack) Command(index int) *QUndoCommand {
	if ptr.Pointer() != nil {
		return NewQUndoCommandFromPointer(C.QUndoStack_Command(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QUndoStack) Count() int {
	if ptr.Pointer() != nil {
		return int(C.QUndoStack_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QUndoStack) CreateRedoAction(parent core.QObject_ITF, prefix string) *QAction {
	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QUndoStack_CreateRedoAction(ptr.Pointer(), core.PointerFromQObject(parent), C.CString(prefix)))
	}
	return nil
}

func (ptr *QUndoStack) CreateUndoAction(parent core.QObject_ITF, prefix string) *QAction {
	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QUndoStack_CreateUndoAction(ptr.Pointer(), core.PointerFromQObject(parent), C.CString(prefix)))
	}
	return nil
}

func (ptr *QUndoStack) EndMacro() {
	if ptr.Pointer() != nil {
		C.QUndoStack_EndMacro(ptr.Pointer())
	}
}

func (ptr *QUndoStack) Index() int {
	if ptr.Pointer() != nil {
		return int(C.QUndoStack_Index(ptr.Pointer()))
	}
	return 0
}

func (ptr *QUndoStack) ConnectIndexChanged(f func(idx int)) {
	if ptr.Pointer() != nil {
		C.QUndoStack_ConnectIndexChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "indexChanged", f)
	}
}

func (ptr *QUndoStack) DisconnectIndexChanged() {
	if ptr.Pointer() != nil {
		C.QUndoStack_DisconnectIndexChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "indexChanged")
	}
}

//export callbackQUndoStackIndexChanged
func callbackQUndoStackIndexChanged(ptrName *C.char, idx C.int) {
	qt.GetSignal(C.GoString(ptrName), "indexChanged").(func(int))(int(idx))
}

func (ptr *QUndoStack) IsClean() bool {
	if ptr.Pointer() != nil {
		return C.QUndoStack_IsClean(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QUndoStack) Push(cmd QUndoCommand_ITF) {
	if ptr.Pointer() != nil {
		C.QUndoStack_Push(ptr.Pointer(), PointerFromQUndoCommand(cmd))
	}
}

func (ptr *QUndoStack) Redo() {
	if ptr.Pointer() != nil {
		C.QUndoStack_Redo(ptr.Pointer())
	}
}

func (ptr *QUndoStack) RedoText() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QUndoStack_RedoText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QUndoStack) ConnectRedoTextChanged(f func(redoText string)) {
	if ptr.Pointer() != nil {
		C.QUndoStack_ConnectRedoTextChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "redoTextChanged", f)
	}
}

func (ptr *QUndoStack) DisconnectRedoTextChanged() {
	if ptr.Pointer() != nil {
		C.QUndoStack_DisconnectRedoTextChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "redoTextChanged")
	}
}

//export callbackQUndoStackRedoTextChanged
func callbackQUndoStackRedoTextChanged(ptrName *C.char, redoText *C.char) {
	qt.GetSignal(C.GoString(ptrName), "redoTextChanged").(func(string))(C.GoString(redoText))
}

func (ptr *QUndoStack) SetClean() {
	if ptr.Pointer() != nil {
		C.QUndoStack_SetClean(ptr.Pointer())
	}
}

func (ptr *QUndoStack) SetIndex(idx int) {
	if ptr.Pointer() != nil {
		C.QUndoStack_SetIndex(ptr.Pointer(), C.int(idx))
	}
}

func (ptr *QUndoStack) Text(idx int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QUndoStack_Text(ptr.Pointer(), C.int(idx)))
	}
	return ""
}

func (ptr *QUndoStack) Undo() {
	if ptr.Pointer() != nil {
		C.QUndoStack_Undo(ptr.Pointer())
	}
}

func (ptr *QUndoStack) UndoText() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QUndoStack_UndoText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QUndoStack) ConnectUndoTextChanged(f func(undoText string)) {
	if ptr.Pointer() != nil {
		C.QUndoStack_ConnectUndoTextChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "undoTextChanged", f)
	}
}

func (ptr *QUndoStack) DisconnectUndoTextChanged() {
	if ptr.Pointer() != nil {
		C.QUndoStack_DisconnectUndoTextChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "undoTextChanged")
	}
}

//export callbackQUndoStackUndoTextChanged
func callbackQUndoStackUndoTextChanged(ptrName *C.char, undoText *C.char) {
	qt.GetSignal(C.GoString(ptrName), "undoTextChanged").(func(string))(C.GoString(undoText))
}

func (ptr *QUndoStack) DestroyQUndoStack() {
	if ptr.Pointer() != nil {
		C.QUndoStack_DestroyQUndoStack(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
