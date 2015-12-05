package xmlpatterns

//#include "xmlpatterns.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QAbstractXmlNodeModel struct {
	core.QSharedData
}

type QAbstractXmlNodeModel_ITF interface {
	core.QSharedData_ITF
	QAbstractXmlNodeModel_PTR() *QAbstractXmlNodeModel
}

func PointerFromQAbstractXmlNodeModel(ptr QAbstractXmlNodeModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractXmlNodeModel_PTR().Pointer()
	}
	return nil
}

func NewQAbstractXmlNodeModelFromPointer(ptr unsafe.Pointer) *QAbstractXmlNodeModel {
	var n = new(QAbstractXmlNodeModel)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAbstractXmlNodeModel) QAbstractXmlNodeModel_PTR() *QAbstractXmlNodeModel {
	return ptr
}

//QAbstractXmlNodeModel::SimpleAxis
type QAbstractXmlNodeModel__SimpleAxis int64

const (
	QAbstractXmlNodeModel__Parent          = QAbstractXmlNodeModel__SimpleAxis(0)
	QAbstractXmlNodeModel__FirstChild      = QAbstractXmlNodeModel__SimpleAxis(1)
	QAbstractXmlNodeModel__PreviousSibling = QAbstractXmlNodeModel__SimpleAxis(2)
	QAbstractXmlNodeModel__NextSibling     = QAbstractXmlNodeModel__SimpleAxis(3)
)

func (ptr *QAbstractXmlNodeModel) CompareOrder(ni1 QXmlNodeModelIndex_ITF, ni2 QXmlNodeModelIndex_ITF) QXmlNodeModelIndex__DocumentOrder {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractXmlNodeModel::compareOrder")
		}
	}()

	if ptr.Pointer() != nil {
		return QXmlNodeModelIndex__DocumentOrder(C.QAbstractXmlNodeModel_CompareOrder(ptr.Pointer(), PointerFromQXmlNodeModelIndex(ni1), PointerFromQXmlNodeModelIndex(ni2)))
	}
	return 0
}

func (ptr *QAbstractXmlNodeModel) Kind(ni QXmlNodeModelIndex_ITF) QXmlNodeModelIndex__NodeKind {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractXmlNodeModel::kind")
		}
	}()

	if ptr.Pointer() != nil {
		return QXmlNodeModelIndex__NodeKind(C.QAbstractXmlNodeModel_Kind(ptr.Pointer(), PointerFromQXmlNodeModelIndex(ni)))
	}
	return 0
}

func (ptr *QAbstractXmlNodeModel) StringValue(n QXmlNodeModelIndex_ITF) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractXmlNodeModel::stringValue")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QAbstractXmlNodeModel_StringValue(ptr.Pointer(), PointerFromQXmlNodeModelIndex(n)))
	}
	return ""
}

func (ptr *QAbstractXmlNodeModel) TypedValue(node QXmlNodeModelIndex_ITF) *core.QVariant {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractXmlNodeModel::typedValue")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QAbstractXmlNodeModel_TypedValue(ptr.Pointer(), PointerFromQXmlNodeModelIndex(node)))
	}
	return nil
}

func (ptr *QAbstractXmlNodeModel) DestroyQAbstractXmlNodeModel() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAbstractXmlNodeModel::~QAbstractXmlNodeModel")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAbstractXmlNodeModel_DestroyQAbstractXmlNodeModel(ptr.Pointer())
	}
}
