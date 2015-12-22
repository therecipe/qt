package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QUndoCommand struct {
	ptr unsafe.Pointer
}

type QUndoCommand_ITF interface {
	QUndoCommand_PTR() *QUndoCommand
}

func (p *QUndoCommand) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QUndoCommand) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQUndoCommand(ptr QUndoCommand_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QUndoCommand_PTR().Pointer()
	}
	return nil
}

func NewQUndoCommandFromPointer(ptr unsafe.Pointer) *QUndoCommand {
	var n = new(QUndoCommand)
	n.SetPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QUndoCommand_") {
		n.SetObjectNameAbs("QUndoCommand_" + qt.Identifier())
	}
	return n
}

func (ptr *QUndoCommand) QUndoCommand_PTR() *QUndoCommand {
	return ptr
}

func NewQUndoCommand(parent QUndoCommand_ITF) *QUndoCommand {
	defer qt.Recovering("QUndoCommand::QUndoCommand")

	return NewQUndoCommandFromPointer(C.QUndoCommand_NewQUndoCommand(PointerFromQUndoCommand(parent)))
}

func NewQUndoCommand2(text string, parent QUndoCommand_ITF) *QUndoCommand {
	defer qt.Recovering("QUndoCommand::QUndoCommand")

	return NewQUndoCommandFromPointer(C.QUndoCommand_NewQUndoCommand2(C.CString(text), PointerFromQUndoCommand(parent)))
}

func (ptr *QUndoCommand) ActionText() string {
	defer qt.Recovering("QUndoCommand::actionText")

	if ptr.Pointer() != nil {
		return C.GoString(C.QUndoCommand_ActionText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QUndoCommand) Child(index int) *QUndoCommand {
	defer qt.Recovering("QUndoCommand::child")

	if ptr.Pointer() != nil {
		return NewQUndoCommandFromPointer(C.QUndoCommand_Child(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QUndoCommand) ChildCount() int {
	defer qt.Recovering("QUndoCommand::childCount")

	if ptr.Pointer() != nil {
		return int(C.QUndoCommand_ChildCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QUndoCommand) Id() int {
	defer qt.Recovering("QUndoCommand::id")

	if ptr.Pointer() != nil {
		return int(C.QUndoCommand_Id(ptr.Pointer()))
	}
	return 0
}

func (ptr *QUndoCommand) MergeWith(command QUndoCommand_ITF) bool {
	defer qt.Recovering("QUndoCommand::mergeWith")

	if ptr.Pointer() != nil {
		return C.QUndoCommand_MergeWith(ptr.Pointer(), PointerFromQUndoCommand(command)) != 0
	}
	return false
}

func (ptr *QUndoCommand) ConnectRedo(f func()) {
	defer qt.Recovering("connect QUndoCommand::redo")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "redo", f)
	}
}

func (ptr *QUndoCommand) DisconnectRedo() {
	defer qt.Recovering("disconnect QUndoCommand::redo")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "redo")
	}
}

//export callbackQUndoCommandRedo
func callbackQUndoCommandRedo(ptrName *C.char) bool {
	defer qt.Recovering("callback QUndoCommand::redo")

	if signal := qt.GetSignal(C.GoString(ptrName), "redo"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QUndoCommand) SetText(text string) {
	defer qt.Recovering("QUndoCommand::setText")

	if ptr.Pointer() != nil {
		C.QUndoCommand_SetText(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QUndoCommand) Text() string {
	defer qt.Recovering("QUndoCommand::text")

	if ptr.Pointer() != nil {
		return C.GoString(C.QUndoCommand_Text(ptr.Pointer()))
	}
	return ""
}

func (ptr *QUndoCommand) ConnectUndo(f func()) {
	defer qt.Recovering("connect QUndoCommand::undo")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "undo", f)
	}
}

func (ptr *QUndoCommand) DisconnectUndo() {
	defer qt.Recovering("disconnect QUndoCommand::undo")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "undo")
	}
}

//export callbackQUndoCommandUndo
func callbackQUndoCommandUndo(ptrName *C.char) bool {
	defer qt.Recovering("callback QUndoCommand::undo")

	if signal := qt.GetSignal(C.GoString(ptrName), "undo"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QUndoCommand) DestroyQUndoCommand() {
	defer qt.Recovering("QUndoCommand::~QUndoCommand")

	if ptr.Pointer() != nil {
		C.QUndoCommand_DestroyQUndoCommand(ptr.Pointer())
	}
}

func (ptr *QUndoCommand) ObjectNameAbs() string {
	defer qt.Recovering("QUndoCommand::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QUndoCommand_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QUndoCommand) SetObjectNameAbs(name string) {
	defer qt.Recovering("QUndoCommand::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QUndoCommand_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
