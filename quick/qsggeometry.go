package quick

//#include "quick.h"
import "C"
import (
	"github.com/therecipe/qt"
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
	for len(n.ObjectNameAbs()) < len("QSGGeometry_") {
		n.SetObjectNameAbs("QSGGeometry_" + qt.Identifier())
	}
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
	defer qt.Recovering("QSGGeometry::allocate")

	if ptr.Pointer() != nil {
		C.QSGGeometry_Allocate(ptr.Pointer(), C.int(vertexCount), C.int(indexCount))
	}
}

func (ptr *QSGGeometry) AttributeCount() int {
	defer qt.Recovering("QSGGeometry::attributeCount")

	if ptr.Pointer() != nil {
		return int(C.QSGGeometry_AttributeCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGGeometry) IndexCount() int {
	defer qt.Recovering("QSGGeometry::indexCount")

	if ptr.Pointer() != nil {
		return int(C.QSGGeometry_IndexCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGGeometry) IndexData() unsafe.Pointer {
	defer qt.Recovering("QSGGeometry::indexData")

	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QSGGeometry_IndexData(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGGeometry) IndexData2() unsafe.Pointer {
	defer qt.Recovering("QSGGeometry::indexData")

	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QSGGeometry_IndexData2(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGGeometry) IndexDataPattern() QSGGeometry__DataPattern {
	defer qt.Recovering("QSGGeometry::indexDataPattern")

	if ptr.Pointer() != nil {
		return QSGGeometry__DataPattern(C.QSGGeometry_IndexDataPattern(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGGeometry) IndexType() int {
	defer qt.Recovering("QSGGeometry::indexType")

	if ptr.Pointer() != nil {
		return int(C.QSGGeometry_IndexType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGGeometry) MarkIndexDataDirty() {
	defer qt.Recovering("QSGGeometry::markIndexDataDirty")

	if ptr.Pointer() != nil {
		C.QSGGeometry_MarkIndexDataDirty(ptr.Pointer())
	}
}

func (ptr *QSGGeometry) MarkVertexDataDirty() {
	defer qt.Recovering("QSGGeometry::markVertexDataDirty")

	if ptr.Pointer() != nil {
		C.QSGGeometry_MarkVertexDataDirty(ptr.Pointer())
	}
}

func (ptr *QSGGeometry) SetIndexDataPattern(p QSGGeometry__DataPattern) {
	defer qt.Recovering("QSGGeometry::setIndexDataPattern")

	if ptr.Pointer() != nil {
		C.QSGGeometry_SetIndexDataPattern(ptr.Pointer(), C.int(p))
	}
}

func (ptr *QSGGeometry) SetVertexDataPattern(p QSGGeometry__DataPattern) {
	defer qt.Recovering("QSGGeometry::setVertexDataPattern")

	if ptr.Pointer() != nil {
		C.QSGGeometry_SetVertexDataPattern(ptr.Pointer(), C.int(p))
	}
}

func (ptr *QSGGeometry) SizeOfIndex() int {
	defer qt.Recovering("QSGGeometry::sizeOfIndex")

	if ptr.Pointer() != nil {
		return int(C.QSGGeometry_SizeOfIndex(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGGeometry) SizeOfVertex() int {
	defer qt.Recovering("QSGGeometry::sizeOfVertex")

	if ptr.Pointer() != nil {
		return int(C.QSGGeometry_SizeOfVertex(ptr.Pointer()))
	}
	return 0
}

func QSGGeometry_UpdateRectGeometry(g QSGGeometry_ITF, rect core.QRectF_ITF) {
	defer qt.Recovering("QSGGeometry::updateRectGeometry")

	C.QSGGeometry_QSGGeometry_UpdateRectGeometry(PointerFromQSGGeometry(g), core.PointerFromQRectF(rect))
}

func QSGGeometry_UpdateTexturedRectGeometry(g QSGGeometry_ITF, rect core.QRectF_ITF, textureRect core.QRectF_ITF) {
	defer qt.Recovering("QSGGeometry::updateTexturedRectGeometry")

	C.QSGGeometry_QSGGeometry_UpdateTexturedRectGeometry(PointerFromQSGGeometry(g), core.PointerFromQRectF(rect), core.PointerFromQRectF(textureRect))
}

func (ptr *QSGGeometry) VertexCount() int {
	defer qt.Recovering("QSGGeometry::vertexCount")

	if ptr.Pointer() != nil {
		return int(C.QSGGeometry_VertexCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGGeometry) VertexData() unsafe.Pointer {
	defer qt.Recovering("QSGGeometry::vertexData")

	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QSGGeometry_VertexData(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGGeometry) VertexData2() unsafe.Pointer {
	defer qt.Recovering("QSGGeometry::vertexData")

	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QSGGeometry_VertexData2(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGGeometry) VertexDataPattern() QSGGeometry__DataPattern {
	defer qt.Recovering("QSGGeometry::vertexDataPattern")

	if ptr.Pointer() != nil {
		return QSGGeometry__DataPattern(C.QSGGeometry_VertexDataPattern(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGGeometry) DestroyQSGGeometry() {
	defer qt.Recovering("QSGGeometry::~QSGGeometry")

	if ptr.Pointer() != nil {
		C.QSGGeometry_DestroyQSGGeometry(ptr.Pointer())
	}
}

func (ptr *QSGGeometry) ObjectNameAbs() string {
	defer qt.Recovering("QSGGeometry::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSGGeometry_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSGGeometry) SetObjectNameAbs(name string) {
	defer qt.Recovering("QSGGeometry::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QSGGeometry_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
