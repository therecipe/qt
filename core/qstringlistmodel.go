package core

//#include "qstringlistmodel.h"
import "C"
import (
	"github.com/therecipe/qt"
	"strings"
	"unsafe"
)

type QStringListModel struct {
	QAbstractListModel
}

type QStringListModelITF interface {
	QAbstractListModelITF
	QStringListModelPTR() *QStringListModel
}

func PointerFromQStringListModel(ptr QStringListModelITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStringListModelPTR().Pointer()
	}
	return nil
}

func QStringListModelFromPointer(ptr unsafe.Pointer) *QStringListModel {
	var n = new(QStringListModel)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QStringListModel_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QStringListModel) QStringListModelPTR() *QStringListModel {
	return ptr
}

func (ptr *QStringListModel) Data(index QModelIndexITF, role int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QStringListModel_Data(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQModelIndex(index)), C.int(role)))
	}
	return ""
}

func (ptr *QStringListModel) Flags(index QModelIndexITF) Qt__ItemFlag {
	if ptr.Pointer() != nil {
		return Qt__ItemFlag(C.QStringListModel_Flags(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQModelIndex(index))))
	}
	return 0
}

func (ptr *QStringListModel) InsertRows(row int, count int, parent QModelIndexITF) bool {
	if ptr.Pointer() != nil {
		return C.QStringListModel_InsertRows(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.int(count), C.QtObjectPtr(PointerFromQModelIndex(parent))) != 0
	}
	return false
}

func (ptr *QStringListModel) RemoveRows(row int, count int, parent QModelIndexITF) bool {
	if ptr.Pointer() != nil {
		return C.QStringListModel_RemoveRows(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.int(count), C.QtObjectPtr(PointerFromQModelIndex(parent))) != 0
	}
	return false
}

func (ptr *QStringListModel) RowCount(parent QModelIndexITF) int {
	if ptr.Pointer() != nil {
		return int(C.QStringListModel_RowCount(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQModelIndex(parent))))
	}
	return 0
}

func (ptr *QStringListModel) SetData(index QModelIndexITF, value string, role int) bool {
	if ptr.Pointer() != nil {
		return C.QStringListModel_SetData(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQModelIndex(index)), C.CString(value), C.int(role)) != 0
	}
	return false
}

func (ptr *QStringListModel) SetStringList(strin []string) {
	if ptr.Pointer() != nil {
		C.QStringListModel_SetStringList(C.QtObjectPtr(ptr.Pointer()), C.CString(strings.Join(strin, "|")))
	}
}

func (ptr *QStringListModel) Sibling(row int, column int, idx QModelIndexITF) *QModelIndex {
	if ptr.Pointer() != nil {
		return QModelIndexFromPointer(unsafe.Pointer(C.QStringListModel_Sibling(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.int(column), C.QtObjectPtr(PointerFromQModelIndex(idx)))))
	}
	return nil
}

func (ptr *QStringListModel) Sort(column int, order Qt__SortOrder) {
	if ptr.Pointer() != nil {
		C.QStringListModel_Sort(C.QtObjectPtr(ptr.Pointer()), C.int(column), C.int(order))
	}
}

func (ptr *QStringListModel) StringList() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QStringListModel_StringList(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QStringListModel) SupportedDropActions() Qt__DropAction {
	if ptr.Pointer() != nil {
		return Qt__DropAction(C.QStringListModel_SupportedDropActions(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}
