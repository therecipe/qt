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

type QUndoStackITF interface {
	core.QObjectITF
	QUndoStackPTR() *QUndoStack
}

func PointerFromQUndoStack(ptr QUndoStackITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QUndoStackPTR().Pointer()
	}
	return nil
}

func QUndoStackFromPointer(ptr unsafe.Pointer) *QUndoStack {
	var n = new(QUndoStack)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QUndoStack_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QUndoStack) QUndoStackPTR() *QUndoStack {
	return ptr
}

func (ptr *QUndoStack) IsActive() bool {
	if ptr.Pointer() != nil {
		return C.QUndoStack_IsActive(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QUndoStack) SetActive(active bool) {
	if ptr.Pointer() != nil {
		C.QUndoStack_SetActive(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(active)))
	}
}

func (ptr *QUndoStack) SetUndoLimit(limit int) {
	if ptr.Pointer() != nil {
		C.QUndoStack_SetUndoLimit(C.QtObjectPtr(ptr.Pointer()), C.int(limit))
	}
}

func (ptr *QUndoStack) UndoLimit() int {
	if ptr.Pointer() != nil {
		return int(C.QUndoStack_UndoLimit(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func NewQUndoStack(parent core.QObjectITF) *QUndoStack {
	return QUndoStackFromPointer(unsafe.Pointer(C.QUndoStack_NewQUndoStack(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QUndoStack) BeginMacro(text string) {
	if ptr.Pointer() != nil {
		C.QUndoStack_BeginMacro(C.QtObjectPtr(ptr.Pointer()), C.CString(text))
	}
}

func (ptr *QUndoStack) CanRedo() bool {
	if ptr.Pointer() != nil {
		return C.QUndoStack_CanRedo(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QUndoStack) ConnectCanRedoChanged(f func(canRedo bool)) {
	if ptr.Pointer() != nil {
		C.QUndoStack_ConnectCanRedoChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "canRedoChanged", f)
	}
}

func (ptr *QUndoStack) DisconnectCanRedoChanged() {
	if ptr.Pointer() != nil {
		C.QUndoStack_DisconnectCanRedoChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "canRedoChanged")
	}
}

//export callbackQUndoStackCanRedoChanged
func callbackQUndoStackCanRedoChanged(ptrName *C.char, canRedo C.int) {
	qt.GetSignal(C.GoString(ptrName), "canRedoChanged").(func(bool))(int(canRedo) != 0)
}

func (ptr *QUndoStack) CanUndo() bool {
	if ptr.Pointer() != nil {
		return C.QUndoStack_CanUndo(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QUndoStack) ConnectCanUndoChanged(f func(canUndo bool)) {
	if ptr.Pointer() != nil {
		C.QUndoStack_ConnectCanUndoChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "canUndoChanged", f)
	}
}

func (ptr *QUndoStack) DisconnectCanUndoChanged() {
	if ptr.Pointer() != nil {
		C.QUndoStack_DisconnectCanUndoChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "canUndoChanged")
	}
}

//export callbackQUndoStackCanUndoChanged
func callbackQUndoStackCanUndoChanged(ptrName *C.char, canUndo C.int) {
	qt.GetSignal(C.GoString(ptrName), "canUndoChanged").(func(bool))(int(canUndo) != 0)
}

func (ptr *QUndoStack) ConnectCleanChanged(f func(clean bool)) {
	if ptr.Pointer() != nil {
		C.QUndoStack_ConnectCleanChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "cleanChanged", f)
	}
}

func (ptr *QUndoStack) DisconnectCleanChanged() {
	if ptr.Pointer() != nil {
		C.QUndoStack_DisconnectCleanChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "cleanChanged")
	}
}

//export callbackQUndoStackCleanChanged
func callbackQUndoStackCleanChanged(ptrName *C.char, clean C.int) {
	qt.GetSignal(C.GoString(ptrName), "cleanChanged").(func(bool))(int(clean) != 0)
}

func (ptr *QUndoStack) CleanIndex() int {
	if ptr.Pointer() != nil {
		return int(C.QUndoStack_CleanIndex(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QUndoStack) Clear() {
	if ptr.Pointer() != nil {
		C.QUndoStack_Clear(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QUndoStack) Command(index int) *QUndoCommand {
	if ptr.Pointer() != nil {
		return QUndoCommandFromPointer(unsafe.Pointer(C.QUndoStack_Command(C.QtObjectPtr(ptr.Pointer()), C.int(index))))
	}
	return nil
}

func (ptr *QUndoStack) Count() int {
	if ptr.Pointer() != nil {
		return int(C.QUndoStack_Count(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QUndoStack) CreateRedoAction(parent core.QObjectITF, prefix string) *QAction {
	if ptr.Pointer() != nil {
		return QActionFromPointer(unsafe.Pointer(C.QUndoStack_CreateRedoAction(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQObject(parent)), C.CString(prefix))))
	}
	return nil
}

func (ptr *QUndoStack) CreateUndoAction(parent core.QObjectITF, prefix string) *QAction {
	if ptr.Pointer() != nil {
		return QActionFromPointer(unsafe.Pointer(C.QUndoStack_CreateUndoAction(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQObject(parent)), C.CString(prefix))))
	}
	return nil
}

func (ptr *QUndoStack) EndMacro() {
	if ptr.Pointer() != nil {
		C.QUndoStack_EndMacro(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QUndoStack) Index() int {
	if ptr.Pointer() != nil {
		return int(C.QUndoStack_Index(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QUndoStack) ConnectIndexChanged(f func(idx int)) {
	if ptr.Pointer() != nil {
		C.QUndoStack_ConnectIndexChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "indexChanged", f)
	}
}

func (ptr *QUndoStack) DisconnectIndexChanged() {
	if ptr.Pointer() != nil {
		C.QUndoStack_DisconnectIndexChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "indexChanged")
	}
}

//export callbackQUndoStackIndexChanged
func callbackQUndoStackIndexChanged(ptrName *C.char, idx C.int) {
	qt.GetSignal(C.GoString(ptrName), "indexChanged").(func(int))(int(idx))
}

func (ptr *QUndoStack) IsClean() bool {
	if ptr.Pointer() != nil {
		return C.QUndoStack_IsClean(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QUndoStack) Push(cmd QUndoCommandITF) {
	if ptr.Pointer() != nil {
		C.QUndoStack_Push(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQUndoCommand(cmd)))
	}
}

func (ptr *QUndoStack) Redo() {
	if ptr.Pointer() != nil {
		C.QUndoStack_Redo(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QUndoStack) RedoText() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QUndoStack_RedoText(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QUndoStack) ConnectRedoTextChanged(f func(redoText string)) {
	if ptr.Pointer() != nil {
		C.QUndoStack_ConnectRedoTextChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "redoTextChanged", f)
	}
}

func (ptr *QUndoStack) DisconnectRedoTextChanged() {
	if ptr.Pointer() != nil {
		C.QUndoStack_DisconnectRedoTextChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "redoTextChanged")
	}
}

//export callbackQUndoStackRedoTextChanged
func callbackQUndoStackRedoTextChanged(ptrName *C.char, redoText *C.char) {
	qt.GetSignal(C.GoString(ptrName), "redoTextChanged").(func(string))(C.GoString(redoText))
}

func (ptr *QUndoStack) SetClean() {
	if ptr.Pointer() != nil {
		C.QUndoStack_SetClean(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QUndoStack) SetIndex(idx int) {
	if ptr.Pointer() != nil {
		C.QUndoStack_SetIndex(C.QtObjectPtr(ptr.Pointer()), C.int(idx))
	}
}

func (ptr *QUndoStack) Text(idx int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QUndoStack_Text(C.QtObjectPtr(ptr.Pointer()), C.int(idx)))
	}
	return ""
}

func (ptr *QUndoStack) Undo() {
	if ptr.Pointer() != nil {
		C.QUndoStack_Undo(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QUndoStack) UndoText() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QUndoStack_UndoText(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QUndoStack) ConnectUndoTextChanged(f func(undoText string)) {
	if ptr.Pointer() != nil {
		C.QUndoStack_ConnectUndoTextChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "undoTextChanged", f)
	}
}

func (ptr *QUndoStack) DisconnectUndoTextChanged() {
	if ptr.Pointer() != nil {
		C.QUndoStack_DisconnectUndoTextChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "undoTextChanged")
	}
}

//export callbackQUndoStackUndoTextChanged
func callbackQUndoStackUndoTextChanged(ptrName *C.char, undoText *C.char) {
	qt.GetSignal(C.GoString(ptrName), "undoTextChanged").(func(string))(C.GoString(undoText))
}

func (ptr *QUndoStack) DestroyQUndoStack() {
	if ptr.Pointer() != nil {
		C.QUndoStack_DestroyQUndoStack(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
