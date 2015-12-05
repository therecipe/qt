package core

//#include "core.h"
import "C"
import (
	"log"
	"unsafe"
)

type QLine struct {
	ptr unsafe.Pointer
}

type QLine_ITF interface {
	QLine_PTR() *QLine
}

func (p *QLine) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QLine) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQLine(ptr QLine_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLine_PTR().Pointer()
	}
	return nil
}

func NewQLineFromPointer(ptr unsafe.Pointer) *QLine {
	var n = new(QLine)
	n.SetPointer(ptr)
	return n
}

func (ptr *QLine) QLine_PTR() *QLine {
	return ptr
}

func NewQLine() *QLine {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLine::QLine")
		}
	}()

	return NewQLineFromPointer(C.QLine_NewQLine())
}

func NewQLine2(p1 QPoint_ITF, p2 QPoint_ITF) *QLine {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLine::QLine")
		}
	}()

	return NewQLineFromPointer(C.QLine_NewQLine2(PointerFromQPoint(p1), PointerFromQPoint(p2)))
}

func NewQLine3(x1 int, y1 int, x2 int, y2 int) *QLine {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLine::QLine")
		}
	}()

	return NewQLineFromPointer(C.QLine_NewQLine3(C.int(x1), C.int(y1), C.int(x2), C.int(y2)))
}

func (ptr *QLine) Dx() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLine::dx")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QLine_Dx(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLine) Dy() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLine::dy")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QLine_Dy(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLine) IsNull() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLine::isNull")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QLine_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLine) SetLine(x1 int, y1 int, x2 int, y2 int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLine::setLine")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLine_SetLine(ptr.Pointer(), C.int(x1), C.int(y1), C.int(x2), C.int(y2))
	}
}

func (ptr *QLine) SetP1(p1 QPoint_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLine::setP1")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLine_SetP1(ptr.Pointer(), PointerFromQPoint(p1))
	}
}

func (ptr *QLine) SetP2(p2 QPoint_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLine::setP2")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLine_SetP2(ptr.Pointer(), PointerFromQPoint(p2))
	}
}

func (ptr *QLine) SetPoints(p1 QPoint_ITF, p2 QPoint_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLine::setPoints")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLine_SetPoints(ptr.Pointer(), PointerFromQPoint(p1), PointerFromQPoint(p2))
	}
}

func (ptr *QLine) Translate(offset QPoint_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLine::translate")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLine_Translate(ptr.Pointer(), PointerFromQPoint(offset))
	}
}

func (ptr *QLine) Translate2(dx int, dy int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLine::translate")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLine_Translate2(ptr.Pointer(), C.int(dx), C.int(dy))
	}
}

func (ptr *QLine) X1() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLine::x1")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QLine_X1(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLine) X2() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLine::x2")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QLine_X2(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLine) Y1() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLine::y1")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QLine_Y1(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLine) Y2() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLine::y2")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QLine_Y2(ptr.Pointer()))
	}
	return 0
}
