package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QTextObjectInterface struct {
	ptr unsafe.Pointer
}

type QTextObjectInterface_ITF interface {
	QTextObjectInterface_PTR() *QTextObjectInterface
}

func (p *QTextObjectInterface) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QTextObjectInterface) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQTextObjectInterface(ptr QTextObjectInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextObjectInterface_PTR().Pointer()
	}
	return nil
}

func NewQTextObjectInterfaceFromPointer(ptr unsafe.Pointer) *QTextObjectInterface {
	var n = new(QTextObjectInterface)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTextObjectInterface) QTextObjectInterface_PTR() *QTextObjectInterface {
	return ptr
}

func (ptr *QTextObjectInterface) DrawObject(painter QPainter_ITF, rect core.QRectF_ITF, doc QTextDocument_ITF, posInDocument int, format QTextFormat_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextObjectInterface::drawObject")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextObjectInterface_DrawObject(ptr.Pointer(), PointerFromQPainter(painter), core.PointerFromQRectF(rect), PointerFromQTextDocument(doc), C.int(posInDocument), PointerFromQTextFormat(format))
	}
}

func (ptr *QTextObjectInterface) DestroyQTextObjectInterface() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextObjectInterface::~QTextObjectInterface")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextObjectInterface_DestroyQTextObjectInterface(ptr.Pointer())
	}
}
