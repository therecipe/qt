package quick

//#include "qsggeometry.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QSGGeometry struct {
	ptr unsafe.Pointer
}

type QSGGeometry_ITF interface {
	QSGGeometry_PTR() *QSGGeometry
}

func (p *QSGGeometry) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSGGeometry) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSGGeometry(ptr QSGGeometry_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGGeometry_PTR().Pointer()
	}
	return nil
}

func NewQSGGeometryFromPointer(ptr unsafe.Pointer) *QSGGeometry {
	var n = new(QSGGeometry)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSGGeometry) QSGGeometry_PTR() *QSGGeometry {
	return ptr
}

//QSGGeometry::DataPattern
type QSGGeometry__DataPattern int64

const (
	QSGGeometry__AlwaysUploadPattern = QSGGeometry__DataPattern(0)
	QSGGeometry__StreamPattern       = QSGGeometry__DataPattern(1)
	QSGGeometry__DynamicPattern      = QSGGeometry__DataPattern(2)
	QSGGeometry__StaticPattern       = QSGGeometry__DataPattern(3)
)

func (ptr *QSGGeometry) Allocate(vertexCount int, indexCount int) {
	if ptr.Pointer() != nil {
		C.QSGGeometry_Allocate(ptr.Pointer(), C.int(vertexCount), C.int(indexCount))
	}
}

func (ptr *QSGGeometry) AttributeCount() int {
	if ptr.Pointer() != nil {
		return int(C.QSGGeometry_AttributeCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGGeometry) IndexCount() int {
	if ptr.Pointer() != nil {
		return int(C.QSGGeometry_IndexCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGGeometry) IndexData() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QSGGeometry_IndexData(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGGeometry) IndexData2() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QSGGeometry_IndexData2(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGGeometry) IndexDataPattern() QSGGeometry__DataPattern {
	if ptr.Pointer() != nil {
		return QSGGeometry__DataPattern(C.QSGGeometry_IndexDataPattern(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGGeometry) IndexType() int {
	if ptr.Pointer() != nil {
		return int(C.QSGGeometry_IndexType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGGeometry) MarkIndexDataDirty() {
	if ptr.Pointer() != nil {
		C.QSGGeometry_MarkIndexDataDirty(ptr.Pointer())
	}
}

func (ptr *QSGGeometry) MarkVertexDataDirty() {
	if ptr.Pointer() != nil {
		C.QSGGeometry_MarkVertexDataDirty(ptr.Pointer())
	}
}

func (ptr *QSGGeometry) SetIndexDataPattern(p QSGGeometry__DataPattern) {
	if ptr.Pointer() != nil {
		C.QSGGeometry_SetIndexDataPattern(ptr.Pointer(), C.int(p))
	}
}

func (ptr *QSGGeometry) SetVertexDataPattern(p QSGGeometry__DataPattern) {
	if ptr.Pointer() != nil {
		C.QSGGeometry_SetVertexDataPattern(ptr.Pointer(), C.int(p))
	}
}

func (ptr *QSGGeometry) SizeOfIndex() int {
	if ptr.Pointer() != nil {
		return int(C.QSGGeometry_SizeOfIndex(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGGeometry) SizeOfVertex() int {
	if ptr.Pointer() != nil {
		return int(C.QSGGeometry_SizeOfVertex(ptr.Pointer()))
	}
	return 0
}

func QSGGeometry_UpdateRectGeometry(g QSGGeometry_ITF, rect core.QRectF_ITF) {
	C.QSGGeometry_QSGGeometry_UpdateRectGeometry(PointerFromQSGGeometry(g), core.PointerFromQRectF(rect))
}

func QSGGeometry_UpdateTexturedRectGeometry(g QSGGeometry_ITF, rect core.QRectF_ITF, textureRect core.QRectF_ITF) {
	C.QSGGeometry_QSGGeometry_UpdateTexturedRectGeometry(PointerFromQSGGeometry(g), core.PointerFromQRectF(rect), core.PointerFromQRectF(textureRect))
}

func (ptr *QSGGeometry) VertexCount() int {
	if ptr.Pointer() != nil {
		return int(C.QSGGeometry_VertexCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGGeometry) VertexData() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QSGGeometry_VertexData(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGGeometry) VertexData2() unsafe.Pointer {
	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QSGGeometry_VertexData2(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGGeometry) VertexDataPattern() QSGGeometry__DataPattern {
	if ptr.Pointer() != nil {
		return QSGGeometry__DataPattern(C.QSGGeometry_VertexDataPattern(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGGeometry) DestroyQSGGeometry() {
	if ptr.Pointer() != nil {
		C.QSGGeometry_DestroyQSGGeometry(ptr.Pointer())
	}
}
