package gui

//#include "qaccessibleeditabletextinterface.h"
import "C"
import (
	"unsafe"
)

type QAccessibleEditableTextInterface struct {
	ptr unsafe.Pointer
}

type QAccessibleEditableTextInterfaceITF interface {
	QAccessibleEditableTextInterfacePTR() *QAccessibleEditableTextInterface
}

func (p *QAccessibleEditableTextInterface) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QAccessibleEditableTextInterface) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQAccessibleEditableTextInterface(ptr QAccessibleEditableTextInterfaceITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAccessibleEditableTextInterfacePTR().Pointer()
	}
	return nil
}

func QAccessibleEditableTextInterfaceFromPointer(ptr unsafe.Pointer) *QAccessibleEditableTextInterface {
	var n = new(QAccessibleEditableTextInterface)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAccessibleEditableTextInterface) QAccessibleEditableTextInterfacePTR() *QAccessibleEditableTextInterface {
	return ptr
}

func (ptr *QAccessibleEditableTextInterface) DeleteText(startOffset int, endOffset int) {
	if ptr.Pointer() != nil {
		C.QAccessibleEditableTextInterface_DeleteText(C.QtObjectPtr(ptr.Pointer()), C.int(startOffset), C.int(endOffset))
	}
}

func (ptr *QAccessibleEditableTextInterface) InsertText(offset int, text string) {
	if ptr.Pointer() != nil {
		C.QAccessibleEditableTextInterface_InsertText(C.QtObjectPtr(ptr.Pointer()), C.int(offset), C.CString(text))
	}
}

func (ptr *QAccessibleEditableTextInterface) ReplaceText(startOffset int, endOffset int, text string) {
	if ptr.Pointer() != nil {
		C.QAccessibleEditableTextInterface_ReplaceText(C.QtObjectPtr(ptr.Pointer()), C.int(startOffset), C.int(endOffset), C.CString(text))
	}
}

func (ptr *QAccessibleEditableTextInterface) DestroyQAccessibleEditableTextInterface() {
	if ptr.Pointer() != nil {
		C.QAccessibleEditableTextInterface_DestroyQAccessibleEditableTextInterface(C.QtObjectPtr(ptr.Pointer()))
	}
}
