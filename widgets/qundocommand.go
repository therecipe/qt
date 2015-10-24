package widgets

//#include "qundocommand.h"
import "C"
import (
	"unsafe"
)

type QUndoCommand struct {
	ptr unsafe.Pointer
}

type QUndoCommandITF interface {
	QUndoCommandPTR() *QUndoCommand
}

func (p *QUndoCommand) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QUndoCommand) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQUndoCommand(ptr QUndoCommandITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QUndoCommandPTR().Pointer()
	}
	return nil
}

func QUndoCommandFromPointer(ptr unsafe.Pointer) *QUndoCommand {
	var n = new(QUndoCommand)
	n.SetPointer(ptr)
	return n
}

func (ptr *QUndoCommand) QUndoCommandPTR() *QUndoCommand {
	return ptr
}

func NewQUndoCommand(parent QUndoCommandITF) *QUndoCommand {
	return QUndoCommandFromPointer(unsafe.Pointer(C.QUndoCommand_NewQUndoCommand(C.QtObjectPtr(PointerFromQUndoCommand(parent)))))
}

func NewQUndoCommand2(text string, parent QUndoCommandITF) *QUndoCommand {
	return QUndoCommandFromPointer(unsafe.Pointer(C.QUndoCommand_NewQUndoCommand2(C.CString(text), C.QtObjectPtr(PointerFromQUndoCommand(parent)))))
}

func (ptr *QUndoCommand) ActionText() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QUndoCommand_ActionText(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QUndoCommand) Child(index int) *QUndoCommand {
	if ptr.Pointer() != nil {
		return QUndoCommandFromPointer(unsafe.Pointer(C.QUndoCommand_Child(C.QtObjectPtr(ptr.Pointer()), C.int(index))))
	}
	return nil
}

func (ptr *QUndoCommand) ChildCount() int {
	if ptr.Pointer() != nil {
		return int(C.QUndoCommand_ChildCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QUndoCommand) Id() int {
	if ptr.Pointer() != nil {
		return int(C.QUndoCommand_Id(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QUndoCommand) MergeWith(command QUndoCommandITF) bool {
	if ptr.Pointer() != nil {
		return C.QUndoCommand_MergeWith(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQUndoCommand(command))) != 0
	}
	return false
}

func (ptr *QUndoCommand) Redo() {
	if ptr.Pointer() != nil {
		C.QUndoCommand_Redo(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QUndoCommand) SetText(text string) {
	if ptr.Pointer() != nil {
		C.QUndoCommand_SetText(C.QtObjectPtr(ptr.Pointer()), C.CString(text))
	}
}

func (ptr *QUndoCommand) Text() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QUndoCommand_Text(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QUndoCommand) Undo() {
	if ptr.Pointer() != nil {
		C.QUndoCommand_Undo(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QUndoCommand) DestroyQUndoCommand() {
	if ptr.Pointer() != nil {
		C.QUndoCommand_DestroyQUndoCommand(C.QtObjectPtr(ptr.Pointer()))
	}
}
