package xml

//#include "qdomprocessinginstruction.h"
import "C"
import (
	"unsafe"
)

type QDomProcessingInstruction struct {
	QDomNode
}

type QDomProcessingInstructionITF interface {
	QDomNodeITF
	QDomProcessingInstructionPTR() *QDomProcessingInstruction
}

func PointerFromQDomProcessingInstruction(ptr QDomProcessingInstructionITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomProcessingInstructionPTR().Pointer()
	}
	return nil
}

func QDomProcessingInstructionFromPointer(ptr unsafe.Pointer) *QDomProcessingInstruction {
	var n = new(QDomProcessingInstruction)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDomProcessingInstruction) QDomProcessingInstructionPTR() *QDomProcessingInstruction {
	return ptr
}

func NewQDomProcessingInstruction() *QDomProcessingInstruction {
	return QDomProcessingInstructionFromPointer(unsafe.Pointer(C.QDomProcessingInstruction_NewQDomProcessingInstruction()))
}

func NewQDomProcessingInstruction2(x QDomProcessingInstructionITF) *QDomProcessingInstruction {
	return QDomProcessingInstructionFromPointer(unsafe.Pointer(C.QDomProcessingInstruction_NewQDomProcessingInstruction2(C.QtObjectPtr(PointerFromQDomProcessingInstruction(x)))))
}

func (ptr *QDomProcessingInstruction) Data() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDomProcessingInstruction_Data(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QDomProcessingInstruction) NodeType() QDomNode__NodeType {
	if ptr.Pointer() != nil {
		return QDomNode__NodeType(C.QDomProcessingInstruction_NodeType(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QDomProcessingInstruction) SetData(d string) {
	if ptr.Pointer() != nil {
		C.QDomProcessingInstruction_SetData(C.QtObjectPtr(ptr.Pointer()), C.CString(d))
	}
}

func (ptr *QDomProcessingInstruction) Target() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDomProcessingInstruction_Target(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}
