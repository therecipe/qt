package xmlpatterns

//#include "xmlpatterns.h"
import "C"
import (
	"log"
	"unsafe"
)

type QXmlNodeModelIndex struct {
	ptr unsafe.Pointer
}

type QXmlNodeModelIndex_ITF interface {
	QXmlNodeModelIndex_PTR() *QXmlNodeModelIndex
}

func (p *QXmlNodeModelIndex) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlNodeModelIndex) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlNodeModelIndex(ptr QXmlNodeModelIndex_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlNodeModelIndex_PTR().Pointer()
	}
	return nil
}

func NewQXmlNodeModelIndexFromPointer(ptr unsafe.Pointer) *QXmlNodeModelIndex {
	var n = new(QXmlNodeModelIndex)
	n.SetPointer(ptr)
	return n
}

func (ptr *QXmlNodeModelIndex) QXmlNodeModelIndex_PTR() *QXmlNodeModelIndex {
	return ptr
}

//QXmlNodeModelIndex::DocumentOrder
type QXmlNodeModelIndex__DocumentOrder int64

const (
	QXmlNodeModelIndex__Precedes = QXmlNodeModelIndex__DocumentOrder(-1)
	QXmlNodeModelIndex__Is       = QXmlNodeModelIndex__DocumentOrder(0)
	QXmlNodeModelIndex__Follows  = QXmlNodeModelIndex__DocumentOrder(1)
)

//QXmlNodeModelIndex::NodeKind
type QXmlNodeModelIndex__NodeKind int64

const (
	QXmlNodeModelIndex__Attribute             = QXmlNodeModelIndex__NodeKind(1)
	QXmlNodeModelIndex__Comment               = QXmlNodeModelIndex__NodeKind(2)
	QXmlNodeModelIndex__Document              = QXmlNodeModelIndex__NodeKind(4)
	QXmlNodeModelIndex__Element               = QXmlNodeModelIndex__NodeKind(8)
	QXmlNodeModelIndex__Namespace             = QXmlNodeModelIndex__NodeKind(16)
	QXmlNodeModelIndex__ProcessingInstruction = QXmlNodeModelIndex__NodeKind(32)
	QXmlNodeModelIndex__Text                  = QXmlNodeModelIndex__NodeKind(64)
)

func NewQXmlNodeModelIndex() *QXmlNodeModelIndex {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlNodeModelIndex::QXmlNodeModelIndex")
		}
	}()

	return NewQXmlNodeModelIndexFromPointer(C.QXmlNodeModelIndex_NewQXmlNodeModelIndex())
}

func NewQXmlNodeModelIndex2(other QXmlNodeModelIndex_ITF) *QXmlNodeModelIndex {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlNodeModelIndex::QXmlNodeModelIndex")
		}
	}()

	return NewQXmlNodeModelIndexFromPointer(C.QXmlNodeModelIndex_NewQXmlNodeModelIndex2(PointerFromQXmlNodeModelIndex(other)))
}

func (ptr *QXmlNodeModelIndex) InternalPointer() unsafe.Pointer {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlNodeModelIndex::internalPointer")
		}
	}()

	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QXmlNodeModelIndex_InternalPointer(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlNodeModelIndex) IsNull() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlNodeModelIndex::isNull")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QXmlNodeModelIndex_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlNodeModelIndex) Model() *QAbstractXmlNodeModel {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlNodeModelIndex::model")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQAbstractXmlNodeModelFromPointer(C.QXmlNodeModelIndex_Model(ptr.Pointer()))
	}
	return nil
}
