package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
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
		n.SetObjectName("QUndoStack_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QUndoStack) QUndoStack_PTR() *QUndoStack {
	return ptr
}

func (ptr *QUndoStack) IsActive() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoStack::isActive")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QUndoStack_IsActive(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QUndoStack) SetActive(active bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoStack::setActive")
		}
	}()

	if ptr.Pointer() != nil {
		C.QUndoStack_SetActive(ptr.Pointer(), C.int(qt.GoBoolToInt(active)))
	}
}

func (ptr *QUndoStack) SetUndoLimit(limit int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoStack::setUndoLimit")
		}
	}()

	if ptr.Pointer() != nil {
		C.QUndoStack_SetUndoLimit(ptr.Pointer(), C.int(limit))
	}
}

func (ptr *QUndoStack) UndoLimit() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoStack::undoLimit")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QUndoStack_UndoLimit(ptr.Pointer()))
	}
	return 0
}

func NewQUndoStack(parent core.QObject_ITF) *QUndoStack {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoStack::QUndoStack")
		}
	}()

	return NewQUndoStackFromPointer(C.QUndoStack_NewQUndoStack(core.PointerFromQObject(parent)))
}

func (ptr *QUndoStack) BeginMacro(text string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoStack::beginMacro")
		}
	}()

	if ptr.Pointer() != nil {
		C.QUndoStack_BeginMacro(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QUndoStack) CanRedo() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoStack::canRedo")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QUndoStack_CanRedo(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QUndoStack) ConnectCanRedoChanged(f func(canRedo bool)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoStack::canRedoChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QUndoStack_ConnectCanRedoChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "canRedoChanged", f)
	}
}

func (ptr *QUndoStack) DisconnectCanRedoChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoStack::canRedoChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QUndoStack_DisconnectCanRedoChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "canRedoChanged")
	}
}

//export callbackQUndoStackCanRedoChanged
func callbackQUndoStackCanRedoChanged(ptrName *C.char, canRedo C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoStack::canRedoChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "canRedoChanged").(func(bool))(int(canRedo) != 0)
}

func (ptr *QUndoStack) CanUndo() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoStack::canUndo")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QUndoStack_CanUndo(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QUndoStack) ConnectCanUndoChanged(f func(canUndo bool)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoStack::canUndoChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QUndoStack_ConnectCanUndoChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "canUndoChanged", f)
	}
}

func (ptr *QUndoStack) DisconnectCanUndoChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoStack::canUndoChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QUndoStack_DisconnectCanUndoChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "canUndoChanged")
	}
}

//export callbackQUndoStackCanUndoChanged
func callbackQUndoStackCanUndoChanged(ptrName *C.char, canUndo C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoStack::canUndoChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "canUndoChanged").(func(bool))(int(canUndo) != 0)
}

func (ptr *QUndoStack) ConnectCleanChanged(f func(clean bool)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoStack::cleanChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QUndoStack_ConnectCleanChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "cleanChanged", f)
	}
}

func (ptr *QUndoStack) DisconnectCleanChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoStack::cleanChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QUndoStack_DisconnectCleanChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "cleanChanged")
	}
}

//export callbackQUndoStackCleanChanged
func callbackQUndoStackCleanChanged(ptrName *C.char, clean C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoStack::cleanChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "cleanChanged").(func(bool))(int(clean) != 0)
}

func (ptr *QUndoStack) CleanIndex() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoStack::cleanIndex")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QUndoStack_CleanIndex(ptr.Pointer()))
	}
	return 0
}

func (ptr *QUndoStack) Clear() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoStack::clear")
		}
	}()

	if ptr.Pointer() != nil {
		C.QUndoStack_Clear(ptr.Pointer())
	}
}

func (ptr *QUndoStack) Command(index int) *QUndoCommand {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoStack::command")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQUndoCommandFromPointer(C.QUndoStack_Command(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QUndoStack) Count() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoStack::count")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QUndoStack_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QUndoStack) CreateRedoAction(parent core.QObject_ITF, prefix string) *QAction {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoStack::createRedoAction")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QUndoStack_CreateRedoAction(ptr.Pointer(), core.PointerFromQObject(parent), C.CString(prefix)))
	}
	return nil
}

func (ptr *QUndoStack) CreateUndoAction(parent core.QObject_ITF, prefix string) *QAction {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoStack::createUndoAction")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QUndoStack_CreateUndoAction(ptr.Pointer(), core.PointerFromQObject(parent), C.CString(prefix)))
	}
	return nil
}

func (ptr *QUndoStack) EndMacro() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoStack::endMacro")
		}
	}()

	if ptr.Pointer() != nil {
		C.QUndoStack_EndMacro(ptr.Pointer())
	}
}

func (ptr *QUndoStack) Index() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoStack::index")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QUndoStack_Index(ptr.Pointer()))
	}
	return 0
}

func (ptr *QUndoStack) ConnectIndexChanged(f func(idx int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoStack::indexChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QUndoStack_ConnectIndexChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "indexChanged", f)
	}
}

func (ptr *QUndoStack) DisconnectIndexChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoStack::indexChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QUndoStack_DisconnectIndexChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "indexChanged")
	}
}

//export callbackQUndoStackIndexChanged
func callbackQUndoStackIndexChanged(ptrName *C.char, idx C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoStack::indexChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "indexChanged").(func(int))(int(idx))
}

func (ptr *QUndoStack) IsClean() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoStack::isClean")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QUndoStack_IsClean(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QUndoStack) Push(cmd QUndoCommand_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoStack::push")
		}
	}()

	if ptr.Pointer() != nil {
		C.QUndoStack_Push(ptr.Pointer(), PointerFromQUndoCommand(cmd))
	}
}

func (ptr *QUndoStack) Redo() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoStack::redo")
		}
	}()

	if ptr.Pointer() != nil {
		C.QUndoStack_Redo(ptr.Pointer())
	}
}

func (ptr *QUndoStack) RedoText() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoStack::redoText")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QUndoStack_RedoText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QUndoStack) ConnectRedoTextChanged(f func(redoText string)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoStack::redoTextChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QUndoStack_ConnectRedoTextChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "redoTextChanged", f)
	}
}

func (ptr *QUndoStack) DisconnectRedoTextChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoStack::redoTextChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QUndoStack_DisconnectRedoTextChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "redoTextChanged")
	}
}

//export callbackQUndoStackRedoTextChanged
func callbackQUndoStackRedoTextChanged(ptrName *C.char, redoText *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoStack::redoTextChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "redoTextChanged").(func(string))(C.GoString(redoText))
}

func (ptr *QUndoStack) SetClean() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoStack::setClean")
		}
	}()

	if ptr.Pointer() != nil {
		C.QUndoStack_SetClean(ptr.Pointer())
	}
}

func (ptr *QUndoStack) SetIndex(idx int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoStack::setIndex")
		}
	}()

	if ptr.Pointer() != nil {
		C.QUndoStack_SetIndex(ptr.Pointer(), C.int(idx))
	}
}

func (ptr *QUndoStack) Text(idx int) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoStack::text")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QUndoStack_Text(ptr.Pointer(), C.int(idx)))
	}
	return ""
}

func (ptr *QUndoStack) Undo() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoStack::undo")
		}
	}()

	if ptr.Pointer() != nil {
		C.QUndoStack_Undo(ptr.Pointer())
	}
}

func (ptr *QUndoStack) UndoText() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoStack::undoText")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QUndoStack_UndoText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QUndoStack) ConnectUndoTextChanged(f func(undoText string)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoStack::undoTextChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QUndoStack_ConnectUndoTextChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "undoTextChanged", f)
	}
}

func (ptr *QUndoStack) DisconnectUndoTextChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoStack::undoTextChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QUndoStack_DisconnectUndoTextChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "undoTextChanged")
	}
}

//export callbackQUndoStackUndoTextChanged
func callbackQUndoStackUndoTextChanged(ptrName *C.char, undoText *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoStack::undoTextChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "undoTextChanged").(func(string))(C.GoString(undoText))
}

func (ptr *QUndoStack) DestroyQUndoStack() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoStack::~QUndoStack")
		}
	}()

	if ptr.Pointer() != nil {
		C.QUndoStack_DestroyQUndoStack(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
