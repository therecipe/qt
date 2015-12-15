package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QAccessibleEditableTextInterface struct {
	ptr unsafe.Pointer
}

type QAccessibleEditableTextInterface_ITF interface {
	QAccessibleEditableTextInterface_PTR() *QAccessibleEditableTextInterface
}

func (p *QAccessibleEditableTextInterface) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QAccessibleEditableTextInterface) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQAccessibleEditableTextInterface(ptr QAccessibleEditableTextInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAccessibleEditableTextInterface_PTR().Pointer()
	}
	return nil
}

func NewQAccessibleEditableTextInterfaceFromPointer(ptr unsafe.Pointer) *QAccessibleEditableTextInterface {
	var n = new(QAccessibleEditableTextInterface)
	n.SetPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QAccessibleEditableTextInterface_") {
		n.SetObjectNameAbs("QAccessibleEditableTextInterface_" + qt.Identifier())
	}
	return n
}

func (ptr *QAccessibleEditableTextInterface) QAccessibleEditableTextInterface_PTR() *QAccessibleEditableTextInterface {
	return ptr
}

func (ptr *QAccessibleEditableTextInterface) DeleteText(startOffset int, endOffset int) {
	defer qt.Recovering("QAccessibleEditableTextInterface::deleteText")

	if ptr.Pointer() != nil {
		C.QAccessibleEditableTextInterface_DeleteText(ptr.Pointer(), C.int(startOffset), C.int(endOffset))
	}
}

func (ptr *QAccessibleEditableTextInterface) InsertText(offset int, text string) {
	defer qt.Recovering("QAccessibleEditableTextInterface::insertText")

	if ptr.Pointer() != nil {
		C.QAccessibleEditableTextInterface_InsertText(ptr.Pointer(), C.int(offset), C.CString(text))
	}
}

func (ptr *QAccessibleEditableTextInterface) ReplaceText(startOffset int, endOffset int, text string) {
	defer qt.Recovering("QAccessibleEditableTextInterface::replaceText")

	if ptr.Pointer() != nil {
		C.QAccessibleEditableTextInterface_ReplaceText(ptr.Pointer(), C.int(startOffset), C.int(endOffset), C.CString(text))
	}
}

func (ptr *QAccessibleEditableTextInterface) DestroyQAccessibleEditableTextInterface() {
	defer qt.Recovering("QAccessibleEditableTextInterface::~QAccessibleEditableTextInterface")

	if ptr.Pointer() != nil {
		C.QAccessibleEditableTextInterface_DestroyQAccessibleEditableTextInterface(ptr.Pointer())
	}
}

func (ptr *QAccessibleEditableTextInterface) ObjectNameAbs() string {
	defer qt.Recovering("QAccessibleEditableTextInterface::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAccessibleEditableTextInterface_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAccessibleEditableTextInterface) SetObjectNameAbs(name string) {
	defer qt.Recovering("QAccessibleEditableTextInterface::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QAccessibleEditableTextInterface_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
