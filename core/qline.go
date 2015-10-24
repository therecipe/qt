package core

//#include "qline.h"
import "C"
import (
	"unsafe"
)

type QLine struct {
	ptr unsafe.Pointer
}

type QLineITF interface {
	QLinePTR() *QLine
}

func (p *QLine) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QLine) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQLine(ptr QLineITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLinePTR().Pointer()
	}
	return nil
}

func QLineFromPointer(ptr unsafe.Pointer) *QLine {
	var n = new(QLine)
	n.SetPointer(ptr)
	return n
}

func (ptr *QLine) QLinePTR() *QLine {
	return ptr
}

func NewQLine() *QLine {
	return QLineFromPointer(unsafe.Pointer(C.QLine_NewQLine()))
}

func NewQLine2(p1 QPointITF, p2 QPointITF) *QLine {
	return QLineFromPointer(unsafe.Pointer(C.QLine_NewQLine2(C.QtObjectPtr(PointerFromQPoint(p1)), C.QtObjectPtr(PointerFromQPoint(p2)))))
}

func NewQLine3(x1 int, y1 int, x2 int, y2 int) *QLine {
	return QLineFromPointer(unsafe.Pointer(C.QLine_NewQLine3(C.int(x1), C.int(y1), C.int(x2), C.int(y2))))
}

func (ptr *QLine) Dx() int {
	if ptr.Pointer() != nil {
		return int(C.QLine_Dx(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QLine) Dy() int {
	if ptr.Pointer() != nil {
		return int(C.QLine_Dy(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QLine) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QLine_IsNull(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QLine) SetLine(x1 int, y1 int, x2 int, y2 int) {
	if ptr.Pointer() != nil {
		C.QLine_SetLine(C.QtObjectPtr(ptr.Pointer()), C.int(x1), C.int(y1), C.int(x2), C.int(y2))
	}
}

func (ptr *QLine) SetP1(p1 QPointITF) {
	if ptr.Pointer() != nil {
		C.QLine_SetP1(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPoint(p1)))
	}
}

func (ptr *QLine) SetP2(p2 QPointITF) {
	if ptr.Pointer() != nil {
		C.QLine_SetP2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPoint(p2)))
	}
}

func (ptr *QLine) SetPoints(p1 QPointITF, p2 QPointITF) {
	if ptr.Pointer() != nil {
		C.QLine_SetPoints(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPoint(p1)), C.QtObjectPtr(PointerFromQPoint(p2)))
	}
}

func (ptr *QLine) Translate(offset QPointITF) {
	if ptr.Pointer() != nil {
		C.QLine_Translate(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPoint(offset)))
	}
}

func (ptr *QLine) Translate2(dx int, dy int) {
	if ptr.Pointer() != nil {
		C.QLine_Translate2(C.QtObjectPtr(ptr.Pointer()), C.int(dx), C.int(dy))
	}
}

func (ptr *QLine) X1() int {
	if ptr.Pointer() != nil {
		return int(C.QLine_X1(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QLine) X2() int {
	if ptr.Pointer() != nil {
		return int(C.QLine_X2(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QLine) Y1() int {
	if ptr.Pointer() != nil {
		return int(C.QLine_Y1(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QLine) Y2() int {
	if ptr.Pointer() != nil {
		return int(C.QLine_Y2(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}
