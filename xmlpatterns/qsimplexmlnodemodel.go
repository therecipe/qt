package xmlpatterns

//#include "xmlpatterns.h"
import "C"
import (
	"log"
	"unsafe"
)

type QSimpleXmlNodeModel struct {
	QAbstractXmlNodeModel
}

type QSimpleXmlNodeModel_ITF interface {
	QAbstractXmlNodeModel_ITF
	QSimpleXmlNodeModel_PTR() *QSimpleXmlNodeModel
}

func PointerFromQSimpleXmlNodeModel(ptr QSimpleXmlNodeModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSimpleXmlNodeModel_PTR().Pointer()
	}
	return nil
}

func NewQSimpleXmlNodeModelFromPointer(ptr unsafe.Pointer) *QSimpleXmlNodeModel {
	var n = new(QSimpleXmlNodeModel)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSimpleXmlNodeModel) QSimpleXmlNodeModel_PTR() *QSimpleXmlNodeModel {
	return ptr
}

func (ptr *QSimpleXmlNodeModel) StringValue(node QXmlNodeModelIndex_ITF) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSimpleXmlNodeModel::stringValue")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QSimpleXmlNodeModel_StringValue(ptr.Pointer(), PointerFromQXmlNodeModelIndex(node)))
	}
	return ""
}

func (ptr *QSimpleXmlNodeModel) DestroyQSimpleXmlNodeModel() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSimpleXmlNodeModel::~QSimpleXmlNodeModel")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSimpleXmlNodeModel_DestroyQSimpleXmlNodeModel(ptr.Pointer())
	}
}
