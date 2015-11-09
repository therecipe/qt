package xml

//#include "qdomprocessinginstruction.h"
import "C"
import (
	"unsafe"
)

type QDomProcessingInstruction struct {
	QDomNode
}

type QDomProcessingInstruction_ITF interface {
	QDomNode_ITF
	QDomProcessingInstruction_PTR() *QDomProcessingInstruction
}

func PointerFromQDomProcessingInstruction(ptr QDomProcessingInstruction_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomProcessingInstruction_PTR().Pointer()
	}
	return nil
}

func NewQDomProcessingInstructionFromPointer(ptr unsafe.Pointer) *QDomProcessingInstruction {
	var n = new(QDomProcessingInstruction)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDomProcessingInstruction) QDomProcessingInstruction_PTR() *QDomProcessingInstruction {
	return ptr
}

func NewQDomProcessingInstruction() *QDomProcessingInstruction {
	return NewQDomProcessingInstructionFromPointer(C.QDomProcessingInstruction_NewQDomProcessingInstruction())
}

func NewQDomProcessingInstruction2(x QDomProcessingInstruction_ITF) *QDomProcessingInstruction {
	return NewQDomProcessingInstructionFromPointer(C.QDomProcessingInstruction_NewQDomProcessingInstruction2(PointerFromQDomProcessingInstruction(x)))
}

func (ptr *QDomProcessingInstruction) Data() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDomProcessingInstruction_Data(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDomProcessingInstruction) NodeType() QDomNode__NodeType {
	if ptr.Pointer() != nil {
		return QDomNode__NodeType(C.QDomProcessingInstruction_NodeType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDomProcessingInstruction) SetData(d string) {
	if ptr.Pointer() != nil {
		C.QDomProcessingInstruction_SetData(ptr.Pointer(), C.CString(d))
	}
}

func (ptr *QDomProcessingInstruction) Target() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDomProcessingInstruction_Target(ptr.Pointer()))
	}
	return ""
}
