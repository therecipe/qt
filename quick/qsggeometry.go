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

type QSGGeometryITF interface {
	QSGGeometryPTR() *QSGGeometry
}

func (p *QSGGeometry) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSGGeometry) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSGGeometry(ptr QSGGeometryITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGGeometryPTR().Pointer()
	}
	return nil
}

func QSGGeometryFromPointer(ptr unsafe.Pointer) *QSGGeometry {
	var n = new(QSGGeometry)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSGGeometry) QSGGeometryPTR() *QSGGeometry {
	return ptr
}

//QSGGeometry::DataPattern
type QSGGeometry__DataPattern int

var (
	QSGGeometry__AlwaysUploadPattern = QSGGeometry__DataPattern(0)
	QSGGeometry__StreamPattern       = QSGGeometry__DataPattern(1)
	QSGGeometry__DynamicPattern      = QSGGeometry__DataPattern(2)
	QSGGeometry__StaticPattern       = QSGGeometry__DataPattern(3)
)

func (ptr *QSGGeometry) Allocate(vertexCount int, indexCount int) {
	if ptr.Pointer() != nil {
		C.QSGGeometry_Allocate(C.QtObjectPtr(ptr.Pointer()), C.int(vertexCount), C.int(indexCount))
	}
}

func (ptr *QSGGeometry) AttributeCount() int {
	if ptr.Pointer() != nil {
		return int(C.QSGGeometry_AttributeCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSGGeometry) IndexCount() int {
	if ptr.Pointer() != nil {
		return int(C.QSGGeometry_IndexCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSGGeometry) IndexData() {
	if ptr.Pointer() != nil {
		C.QSGGeometry_IndexData(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QSGGeometry) IndexData2() {
	if ptr.Pointer() != nil {
		C.QSGGeometry_IndexData2(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QSGGeometry) IndexDataPattern() QSGGeometry__DataPattern {
	if ptr.Pointer() != nil {
		return QSGGeometry__DataPattern(C.QSGGeometry_IndexDataPattern(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSGGeometry) IndexType() int {
	if ptr.Pointer() != nil {
		return int(C.QSGGeometry_IndexType(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSGGeometry) MarkIndexDataDirty() {
	if ptr.Pointer() != nil {
		C.QSGGeometry_MarkIndexDataDirty(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QSGGeometry) MarkVertexDataDirty() {
	if ptr.Pointer() != nil {
		C.QSGGeometry_MarkVertexDataDirty(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QSGGeometry) SetIndexDataPattern(p QSGGeometry__DataPattern) {
	if ptr.Pointer() != nil {
		C.QSGGeometry_SetIndexDataPattern(C.QtObjectPtr(ptr.Pointer()), C.int(p))
	}
}

func (ptr *QSGGeometry) SetVertexDataPattern(p QSGGeometry__DataPattern) {
	if ptr.Pointer() != nil {
		C.QSGGeometry_SetVertexDataPattern(C.QtObjectPtr(ptr.Pointer()), C.int(p))
	}
}

func (ptr *QSGGeometry) SizeOfIndex() int {
	if ptr.Pointer() != nil {
		return int(C.QSGGeometry_SizeOfIndex(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSGGeometry) SizeOfVertex() int {
	if ptr.Pointer() != nil {
		return int(C.QSGGeometry_SizeOfVertex(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func QSGGeometry_UpdateRectGeometry(g QSGGeometryITF, rect core.QRectFITF) {
	C.QSGGeometry_QSGGeometry_UpdateRectGeometry(C.QtObjectPtr(PointerFromQSGGeometry(g)), C.QtObjectPtr(core.PointerFromQRectF(rect)))
}

func QSGGeometry_UpdateTexturedRectGeometry(g QSGGeometryITF, rect core.QRectFITF, textureRect core.QRectFITF) {
	C.QSGGeometry_QSGGeometry_UpdateTexturedRectGeometry(C.QtObjectPtr(PointerFromQSGGeometry(g)), C.QtObjectPtr(core.PointerFromQRectF(rect)), C.QtObjectPtr(core.PointerFromQRectF(textureRect)))
}

func (ptr *QSGGeometry) VertexCount() int {
	if ptr.Pointer() != nil {
		return int(C.QSGGeometry_VertexCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSGGeometry) VertexData() {
	if ptr.Pointer() != nil {
		C.QSGGeometry_VertexData(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QSGGeometry) VertexData2() {
	if ptr.Pointer() != nil {
		C.QSGGeometry_VertexData2(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QSGGeometry) VertexDataPattern() QSGGeometry__DataPattern {
	if ptr.Pointer() != nil {
		return QSGGeometry__DataPattern(C.QSGGeometry_VertexDataPattern(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSGGeometry) DestroyQSGGeometry() {
	if ptr.Pointer() != nil {
		C.QSGGeometry_DestroyQSGGeometry(C.QtObjectPtr(ptr.Pointer()))
	}
}
