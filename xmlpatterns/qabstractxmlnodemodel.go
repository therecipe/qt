package xmlpatterns

//#include "qabstractxmlnodemodel.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QAbstractXmlNodeModel struct {
	core.QSharedData
}

type QAbstractXmlNodeModelITF interface {
	core.QSharedDataITF
	QAbstractXmlNodeModelPTR() *QAbstractXmlNodeModel
}

func PointerFromQAbstractXmlNodeModel(ptr QAbstractXmlNodeModelITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractXmlNodeModelPTR().Pointer()
	}
	return nil
}

func QAbstractXmlNodeModelFromPointer(ptr unsafe.Pointer) *QAbstractXmlNodeModel {
	var n = new(QAbstractXmlNodeModel)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAbstractXmlNodeModel) QAbstractXmlNodeModelPTR() *QAbstractXmlNodeModel {
	return ptr
}

//QAbstractXmlNodeModel::SimpleAxis
type QAbstractXmlNodeModel__SimpleAxis int

var (
	QAbstractXmlNodeModel__Parent          = QAbstractXmlNodeModel__SimpleAxis(0)
	QAbstractXmlNodeModel__FirstChild      = QAbstractXmlNodeModel__SimpleAxis(1)
	QAbstractXmlNodeModel__PreviousSibling = QAbstractXmlNodeModel__SimpleAxis(2)
	QAbstractXmlNodeModel__NextSibling     = QAbstractXmlNodeModel__SimpleAxis(3)
)

func (ptr *QAbstractXmlNodeModel) BaseUri(n QXmlNodeModelIndexITF) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAbstractXmlNodeModel_BaseUri(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQXmlNodeModelIndex(n))))
	}
	return ""
}

func (ptr *QAbstractXmlNodeModel) CompareOrder(ni1 QXmlNodeModelIndexITF, ni2 QXmlNodeModelIndexITF) QXmlNodeModelIndex__DocumentOrder {
	if ptr.Pointer() != nil {
		return QXmlNodeModelIndex__DocumentOrder(C.QAbstractXmlNodeModel_CompareOrder(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQXmlNodeModelIndex(ni1)), C.QtObjectPtr(PointerFromQXmlNodeModelIndex(ni2))))
	}
	return 0
}

func (ptr *QAbstractXmlNodeModel) DocumentUri(n QXmlNodeModelIndexITF) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAbstractXmlNodeModel_DocumentUri(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQXmlNodeModelIndex(n))))
	}
	return ""
}

func (ptr *QAbstractXmlNodeModel) Kind(ni QXmlNodeModelIndexITF) QXmlNodeModelIndex__NodeKind {
	if ptr.Pointer() != nil {
		return QXmlNodeModelIndex__NodeKind(C.QAbstractXmlNodeModel_Kind(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQXmlNodeModelIndex(ni))))
	}
	return 0
}

func (ptr *QAbstractXmlNodeModel) StringValue(n QXmlNodeModelIndexITF) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAbstractXmlNodeModel_StringValue(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQXmlNodeModelIndex(n))))
	}
	return ""
}

func (ptr *QAbstractXmlNodeModel) TypedValue(node QXmlNodeModelIndexITF) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAbstractXmlNodeModel_TypedValue(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQXmlNodeModelIndex(node))))
	}
	return ""
}

func (ptr *QAbstractXmlNodeModel) DestroyQAbstractXmlNodeModel() {
	if ptr.Pointer() != nil {
		C.QAbstractXmlNodeModel_DestroyQAbstractXmlNodeModel(C.QtObjectPtr(ptr.Pointer()))
	}
}
