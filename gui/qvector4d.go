package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QVector4D struct {
	ptr unsafe.Pointer
}

type QVector4D_ITF interface {
	QVector4D_PTR() *QVector4D
}

func (p *QVector4D) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QVector4D) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQVector4D(ptr QVector4D_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVector4D_PTR().Pointer()
	}
	return nil
}

func NewQVector4DFromPointer(ptr unsafe.Pointer) *QVector4D {
	var n = new(QVector4D)
	n.SetPointer(ptr)
	return n
}

func (ptr *QVector4D) QVector4D_PTR() *QVector4D {
	return ptr
}

func NewQVector4D() *QVector4D {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVector4D::QVector4D")
		}
	}()

	return NewQVector4DFromPointer(C.QVector4D_NewQVector4D())
}

func NewQVector4D4(point core.QPoint_ITF) *QVector4D {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVector4D::QVector4D")
		}
	}()

	return NewQVector4DFromPointer(C.QVector4D_NewQVector4D4(core.PointerFromQPoint(point)))
}

func NewQVector4D5(point core.QPointF_ITF) *QVector4D {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVector4D::QVector4D")
		}
	}()

	return NewQVector4DFromPointer(C.QVector4D_NewQVector4D5(core.PointerFromQPointF(point)))
}

func NewQVector4D6(vector QVector2D_ITF) *QVector4D {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVector4D::QVector4D")
		}
	}()

	return NewQVector4DFromPointer(C.QVector4D_NewQVector4D6(PointerFromQVector2D(vector)))
}

func NewQVector4D8(vector QVector3D_ITF) *QVector4D {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVector4D::QVector4D")
		}
	}()

	return NewQVector4DFromPointer(C.QVector4D_NewQVector4D8(PointerFromQVector3D(vector)))
}

func (ptr *QVector4D) IsNull() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVector4D::isNull")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QVector4D_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QVector4D) Normalize() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVector4D::normalize")
		}
	}()

	if ptr.Pointer() != nil {
		C.QVector4D_Normalize(ptr.Pointer())
	}
}
