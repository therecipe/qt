package xmlpatterns

//#include "qsimplexmlnodemodel.h"
import "C"
import (
	"unsafe"
)

type QSimpleXmlNodeModel struct {
	QAbstractXmlNodeModel
}

type QSimpleXmlNodeModelITF interface {
	QAbstractXmlNodeModelITF
	QSimpleXmlNodeModelPTR() *QSimpleXmlNodeModel
}

func PointerFromQSimpleXmlNodeModel(ptr QSimpleXmlNodeModelITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSimpleXmlNodeModelPTR().Pointer()
	}
	return nil
}

func QSimpleXmlNodeModelFromPointer(ptr unsafe.Pointer) *QSimpleXmlNodeModel {
	var n = new(QSimpleXmlNodeModel)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSimpleXmlNodeModel) QSimpleXmlNodeModelPTR() *QSimpleXmlNodeModel {
	return ptr
}

func (ptr *QSimpleXmlNodeModel) BaseUri(node QXmlNodeModelIndexITF) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSimpleXmlNodeModel_BaseUri(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQXmlNodeModelIndex(node))))
	}
	return ""
}

func (ptr *QSimpleXmlNodeModel) StringValue(node QXmlNodeModelIndexITF) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSimpleXmlNodeModel_StringValue(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQXmlNodeModelIndex(node))))
	}
	return ""
}

func (ptr *QSimpleXmlNodeModel) DestroyQSimpleXmlNodeModel() {
	if ptr.Pointer() != nil {
		C.QSimpleXmlNodeModel_DestroyQSimpleXmlNodeModel(C.QtObjectPtr(ptr.Pointer()))
	}
}
