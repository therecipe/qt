package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QUndoGroup struct {
	core.QObject
}

type QUndoGroup_ITF interface {
	core.QObject_ITF
	QUndoGroup_PTR() *QUndoGroup
}

func PointerFromQUndoGroup(ptr QUndoGroup_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QUndoGroup_PTR().Pointer()
	}
	return nil
}

func NewQUndoGroupFromPointer(ptr unsafe.Pointer) *QUndoGroup {
	var n = new(QUndoGroup)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QUndoGroup_") {
		n.SetObjectName("QUndoGroup_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QUndoGroup) QUndoGroup_PTR() *QUndoGroup {
	return ptr
}

func NewQUndoGroup(parent core.QObject_ITF) *QUndoGroup {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoGroup::QUndoGroup")
		}
	}()

	return NewQUndoGroupFromPointer(C.QUndoGroup_NewQUndoGroup(core.PointerFromQObject(parent)))
}

func (ptr *QUndoGroup) ActiveStack() *QUndoStack {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoGroup::activeStack")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQUndoStackFromPointer(C.QUndoGroup_ActiveStack(ptr.Pointer()))
	}
	return nil
}

func (ptr *QUndoGroup) ConnectActiveStackChanged(f func(stack *QUndoStack)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoGroup::activeStackChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QUndoGroup_ConnectActiveStackChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "activeStackChanged", f)
	}
}

func (ptr *QUndoGroup) DisconnectActiveStackChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoGroup::activeStackChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QUndoGroup_DisconnectActiveStackChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "activeStackChanged")
	}
}

//export callbackQUndoGroupActiveStackChanged
func callbackQUndoGroupActiveStackChanged(ptrName *C.char, stack unsafe.Pointer) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoGroup::activeStackChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "activeStackChanged").(func(*QUndoStack))(NewQUndoStackFromPointer(stack))
}

func (ptr *QUndoGroup) AddStack(stack QUndoStack_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoGroup::addStack")
		}
	}()

	if ptr.Pointer() != nil {
		C.QUndoGroup_AddStack(ptr.Pointer(), PointerFromQUndoStack(stack))
	}
}

func (ptr *QUndoGroup) CanRedo() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoGroup::canRedo")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QUndoGroup_CanRedo(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QUndoGroup) ConnectCanRedoChanged(f func(canRedo bool)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoGroup::canRedoChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QUndoGroup_ConnectCanRedoChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "canRedoChanged", f)
	}
}

func (ptr *QUndoGroup) DisconnectCanRedoChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoGroup::canRedoChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QUndoGroup_DisconnectCanRedoChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "canRedoChanged")
	}
}

//export callbackQUndoGroupCanRedoChanged
func callbackQUndoGroupCanRedoChanged(ptrName *C.char, canRedo C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoGroup::canRedoChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "canRedoChanged").(func(bool))(int(canRedo) != 0)
}

func (ptr *QUndoGroup) CanUndo() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoGroup::canUndo")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QUndoGroup_CanUndo(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QUndoGroup) ConnectCanUndoChanged(f func(canUndo bool)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoGroup::canUndoChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QUndoGroup_ConnectCanUndoChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "canUndoChanged", f)
	}
}

func (ptr *QUndoGroup) DisconnectCanUndoChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoGroup::canUndoChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QUndoGroup_DisconnectCanUndoChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "canUndoChanged")
	}
}

//export callbackQUndoGroupCanUndoChanged
func callbackQUndoGroupCanUndoChanged(ptrName *C.char, canUndo C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoGroup::canUndoChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "canUndoChanged").(func(bool))(int(canUndo) != 0)
}

func (ptr *QUndoGroup) ConnectCleanChanged(f func(clean bool)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoGroup::cleanChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QUndoGroup_ConnectCleanChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "cleanChanged", f)
	}
}

func (ptr *QUndoGroup) DisconnectCleanChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoGroup::cleanChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QUndoGroup_DisconnectCleanChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "cleanChanged")
	}
}

//export callbackQUndoGroupCleanChanged
func callbackQUndoGroupCleanChanged(ptrName *C.char, clean C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoGroup::cleanChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "cleanChanged").(func(bool))(int(clean) != 0)
}

func (ptr *QUndoGroup) CreateRedoAction(parent core.QObject_ITF, prefix string) *QAction {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoGroup::createRedoAction")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QUndoGroup_CreateRedoAction(ptr.Pointer(), core.PointerFromQObject(parent), C.CString(prefix)))
	}
	return nil
}

func (ptr *QUndoGroup) CreateUndoAction(parent core.QObject_ITF, prefix string) *QAction {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoGroup::createUndoAction")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QUndoGroup_CreateUndoAction(ptr.Pointer(), core.PointerFromQObject(parent), C.CString(prefix)))
	}
	return nil
}

func (ptr *QUndoGroup) ConnectIndexChanged(f func(idx int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoGroup::indexChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QUndoGroup_ConnectIndexChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "indexChanged", f)
	}
}

func (ptr *QUndoGroup) DisconnectIndexChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoGroup::indexChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QUndoGroup_DisconnectIndexChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "indexChanged")
	}
}

//export callbackQUndoGroupIndexChanged
func callbackQUndoGroupIndexChanged(ptrName *C.char, idx C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoGroup::indexChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "indexChanged").(func(int))(int(idx))
}

func (ptr *QUndoGroup) IsClean() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoGroup::isClean")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QUndoGroup_IsClean(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QUndoGroup) Redo() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoGroup::redo")
		}
	}()

	if ptr.Pointer() != nil {
		C.QUndoGroup_Redo(ptr.Pointer())
	}
}

func (ptr *QUndoGroup) RedoText() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoGroup::redoText")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QUndoGroup_RedoText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QUndoGroup) ConnectRedoTextChanged(f func(redoText string)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoGroup::redoTextChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QUndoGroup_ConnectRedoTextChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "redoTextChanged", f)
	}
}

func (ptr *QUndoGroup) DisconnectRedoTextChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoGroup::redoTextChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QUndoGroup_DisconnectRedoTextChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "redoTextChanged")
	}
}

//export callbackQUndoGroupRedoTextChanged
func callbackQUndoGroupRedoTextChanged(ptrName *C.char, redoText *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoGroup::redoTextChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "redoTextChanged").(func(string))(C.GoString(redoText))
}

func (ptr *QUndoGroup) RemoveStack(stack QUndoStack_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoGroup::removeStack")
		}
	}()

	if ptr.Pointer() != nil {
		C.QUndoGroup_RemoveStack(ptr.Pointer(), PointerFromQUndoStack(stack))
	}
}

func (ptr *QUndoGroup) SetActiveStack(stack QUndoStack_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoGroup::setActiveStack")
		}
	}()

	if ptr.Pointer() != nil {
		C.QUndoGroup_SetActiveStack(ptr.Pointer(), PointerFromQUndoStack(stack))
	}
}

func (ptr *QUndoGroup) Undo() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoGroup::undo")
		}
	}()

	if ptr.Pointer() != nil {
		C.QUndoGroup_Undo(ptr.Pointer())
	}
}

func (ptr *QUndoGroup) UndoText() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoGroup::undoText")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QUndoGroup_UndoText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QUndoGroup) ConnectUndoTextChanged(f func(undoText string)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoGroup::undoTextChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QUndoGroup_ConnectUndoTextChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "undoTextChanged", f)
	}
}

func (ptr *QUndoGroup) DisconnectUndoTextChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoGroup::undoTextChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QUndoGroup_DisconnectUndoTextChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "undoTextChanged")
	}
}

//export callbackQUndoGroupUndoTextChanged
func callbackQUndoGroupUndoTextChanged(ptrName *C.char, undoText *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoGroup::undoTextChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "undoTextChanged").(func(string))(C.GoString(undoText))
}

func (ptr *QUndoGroup) DestroyQUndoGroup() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoGroup::~QUndoGroup")
		}
	}()

	if ptr.Pointer() != nil {
		C.QUndoGroup_DestroyQUndoGroup(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
