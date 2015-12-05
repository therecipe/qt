package widgets

//#include "widgets.h"
import "C"
import (
	"log"
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
	return n
}

func (ptr *QUndoCommand) QUndoCommand_PTR() *QUndoCommand {
	return ptr
}

func NewQUndoCommand(parent QUndoCommand_ITF) *QUndoCommand {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoCommand::QUndoCommand")
		}
	}()

	return NewQUndoCommandFromPointer(C.QUndoCommand_NewQUndoCommand(PointerFromQUndoCommand(parent)))
}

func NewQUndoCommand2(text string, parent QUndoCommand_ITF) *QUndoCommand {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoCommand::QUndoCommand")
		}
	}()

	return NewQUndoCommandFromPointer(C.QUndoCommand_NewQUndoCommand2(C.CString(text), PointerFromQUndoCommand(parent)))
}

func (ptr *QUndoCommand) ActionText() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoCommand::actionText")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QUndoCommand_ActionText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QUndoCommand) Child(index int) *QUndoCommand {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoCommand::child")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQUndoCommandFromPointer(C.QUndoCommand_Child(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QUndoCommand) ChildCount() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoCommand::childCount")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QUndoCommand_ChildCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QUndoCommand) Id() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoCommand::id")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QUndoCommand_Id(ptr.Pointer()))
	}
	return 0
}

func (ptr *QUndoCommand) MergeWith(command QUndoCommand_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoCommand::mergeWith")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QUndoCommand_MergeWith(ptr.Pointer(), PointerFromQUndoCommand(command)) != 0
	}
	return false
}

func (ptr *QUndoCommand) Redo() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoCommand::redo")
		}
	}()

	if ptr.Pointer() != nil {
		C.QUndoCommand_Redo(ptr.Pointer())
	}
}

func (ptr *QUndoCommand) SetText(text string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoCommand::setText")
		}
	}()

	if ptr.Pointer() != nil {
		C.QUndoCommand_SetText(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QUndoCommand) Text() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoCommand::text")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QUndoCommand_Text(ptr.Pointer()))
	}
	return ""
}

func (ptr *QUndoCommand) Undo() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoCommand::undo")
		}
	}()

	if ptr.Pointer() != nil {
		C.QUndoCommand_Undo(ptr.Pointer())
	}
}

func (ptr *QUndoCommand) DestroyQUndoCommand() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QUndoCommand::~QUndoCommand")
		}
	}()

	if ptr.Pointer() != nil {
		C.QUndoCommand_DestroyQUndoCommand(ptr.Pointer())
	}
}
