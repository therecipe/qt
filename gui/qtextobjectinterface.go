package gui

//#include "qtextobjectinterface.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QTextObjectInterface struct {
	ptr unsafe.Pointer
}

type QTextObjectInterfaceITF interface {
	QTextObjectInterfacePTR() *QTextObjectInterface
}

func (p *QTextObjectInterface) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QTextObjectInterface) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQTextObjectInterface(ptr QTextObjectInterfaceITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextObjectInterfacePTR().Pointer()
	}
	return nil
}

func QTextObjectInterfaceFromPointer(ptr unsafe.Pointer) *QTextObjectInterface {
	var n = new(QTextObjectInterface)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTextObjectInterface) QTextObjectInterfacePTR() *QTextObjectInterface {
	return ptr
}

func (ptr *QTextObjectInterface) DrawObject(painter QPainterITF, rect core.QRectFITF, doc QTextDocumentITF, posInDocument int, format QTextFormatITF) {
	if ptr.Pointer() != nil {
		C.QTextObjectInterface_DrawObject(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPainter(painter)), C.QtObjectPtr(core.PointerFromQRectF(rect)), C.QtObjectPtr(PointerFromQTextDocument(doc)), C.int(posInDocument), C.QtObjectPtr(PointerFromQTextFormat(format)))
	}
}

func (ptr *QTextObjectInterface) DestroyQTextObjectInterface() {
	if ptr.Pointer() != nil {
		C.QTextObjectInterface_DestroyQTextObjectInterface(C.QtObjectPtr(ptr.Pointer()))
	}
}
