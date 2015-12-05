package gui

//#include "gui.h"
import "C"
import (
	"log"
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
	return n
}

func (ptr *QAccessibleEditableTextInterface) QAccessibleEditableTextInterface_PTR() *QAccessibleEditableTextInterface {
	return ptr
}

func (ptr *QAccessibleEditableTextInterface) DeleteText(startOffset int, endOffset int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleEditableTextInterface::deleteText")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAccessibleEditableTextInterface_DeleteText(ptr.Pointer(), C.int(startOffset), C.int(endOffset))
	}
}

func (ptr *QAccessibleEditableTextInterface) InsertText(offset int, text string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleEditableTextInterface::insertText")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAccessibleEditableTextInterface_InsertText(ptr.Pointer(), C.int(offset), C.CString(text))
	}
}

func (ptr *QAccessibleEditableTextInterface) ReplaceText(startOffset int, endOffset int, text string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleEditableTextInterface::replaceText")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAccessibleEditableTextInterface_ReplaceText(ptr.Pointer(), C.int(startOffset), C.int(endOffset), C.CString(text))
	}
}

func (ptr *QAccessibleEditableTextInterface) DestroyQAccessibleEditableTextInterface() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleEditableTextInterface::~QAccessibleEditableTextInterface")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAccessibleEditableTextInterface_DestroyQAccessibleEditableTextInterface(ptr.Pointer())
	}
}
