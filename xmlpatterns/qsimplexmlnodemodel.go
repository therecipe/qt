package xmlpatterns

//#include "xmlpatterns.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
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
	for len(n.ObjectNameAbs()) < len("QSimpleXmlNodeModel_") {
		n.SetObjectNameAbs("QSimpleXmlNodeModel_" + qt.Identifier())
	}
	return n
}

func (ptr *QSimpleXmlNodeModel) QSimpleXmlNodeModel_PTR() *QSimpleXmlNodeModel {
	return ptr
}

func (ptr *QSimpleXmlNodeModel) BaseUri(node QXmlNodeModelIndex_ITF) *core.QUrl {
	defer qt.Recovering("QSimpleXmlNodeModel::baseUri")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QSimpleXmlNodeModel_BaseUri(ptr.Pointer(), PointerFromQXmlNodeModelIndex(node)))
	}
	return nil
}

func (ptr *QSimpleXmlNodeModel) StringValue(node QXmlNodeModelIndex_ITF) string {
	defer qt.Recovering("QSimpleXmlNodeModel::stringValue")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSimpleXmlNodeModel_StringValue(ptr.Pointer(), PointerFromQXmlNodeModelIndex(node)))
	}
	return ""
}

func (ptr *QSimpleXmlNodeModel) DestroyQSimpleXmlNodeModel() {
	defer qt.Recovering("QSimpleXmlNodeModel::~QSimpleXmlNodeModel")

	if ptr.Pointer() != nil {
		C.QSimpleXmlNodeModel_DestroyQSimpleXmlNodeModel(ptr.Pointer())
	}
}

func (ptr *QSimpleXmlNodeModel) ObjectNameAbs() string {
	defer qt.Recovering("QSimpleXmlNodeModel::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSimpleXmlNodeModel_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSimpleXmlNodeModel) SetObjectNameAbs(name string) {
	defer qt.Recovering("QSimpleXmlNodeModel::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QSimpleXmlNodeModel_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
