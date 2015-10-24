package xmlpatterns

//#include "qxmlnodemodelindex.h"
import "C"
import (
	"unsafe"
)

type QXmlNodeModelIndex struct {
	ptr unsafe.Pointer
}

type QXmlNodeModelIndexITF interface {
	QXmlNodeModelIndexPTR() *QXmlNodeModelIndex
}

func (p *QXmlNodeModelIndex) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlNodeModelIndex) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlNodeModelIndex(ptr QXmlNodeModelIndexITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlNodeModelIndexPTR().Pointer()
	}
	return nil
}

func QXmlNodeModelIndexFromPointer(ptr unsafe.Pointer) *QXmlNodeModelIndex {
	var n = new(QXmlNodeModelIndex)
	n.SetPointer(ptr)
	return n
}

func (ptr *QXmlNodeModelIndex) QXmlNodeModelIndexPTR() *QXmlNodeModelIndex {
	return ptr
}

//QXmlNodeModelIndex::DocumentOrder
type QXmlNodeModelIndex__DocumentOrder int

var (
	QXmlNodeModelIndex__Precedes = QXmlNodeModelIndex__DocumentOrder(-1)
	QXmlNodeModelIndex__Is       = QXmlNodeModelIndex__DocumentOrder(0)
	QXmlNodeModelIndex__Follows  = QXmlNodeModelIndex__DocumentOrder(1)
)

//QXmlNodeModelIndex::NodeKind
type QXmlNodeModelIndex__NodeKind int

var (
	QXmlNodeModelIndex__Attribute             = QXmlNodeModelIndex__NodeKind(1)
	QXmlNodeModelIndex__Comment               = QXmlNodeModelIndex__NodeKind(2)
	QXmlNodeModelIndex__Document              = QXmlNodeModelIndex__NodeKind(4)
	QXmlNodeModelIndex__Element               = QXmlNodeModelIndex__NodeKind(8)
	QXmlNodeModelIndex__Namespace             = QXmlNodeModelIndex__NodeKind(16)
	QXmlNodeModelIndex__ProcessingInstruction = QXmlNodeModelIndex__NodeKind(32)
	QXmlNodeModelIndex__Text                  = QXmlNodeModelIndex__NodeKind(64)
)

func NewQXmlNodeModelIndex() *QXmlNodeModelIndex {
	return QXmlNodeModelIndexFromPointer(unsafe.Pointer(C.QXmlNodeModelIndex_NewQXmlNodeModelIndex()))
}

func NewQXmlNodeModelIndex2(other QXmlNodeModelIndexITF) *QXmlNodeModelIndex {
	return QXmlNodeModelIndexFromPointer(unsafe.Pointer(C.QXmlNodeModelIndex_NewQXmlNodeModelIndex2(C.QtObjectPtr(PointerFromQXmlNodeModelIndex(other)))))
}

func (ptr *QXmlNodeModelIndex) InternalPointer() {
	if ptr.Pointer() != nil {
		C.QXmlNodeModelIndex_InternalPointer(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QXmlNodeModelIndex) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QXmlNodeModelIndex_IsNull(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QXmlNodeModelIndex) Model() *QAbstractXmlNodeModel {
	if ptr.Pointer() != nil {
		return QAbstractXmlNodeModelFromPointer(unsafe.Pointer(C.QXmlNodeModelIndex_Model(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}
