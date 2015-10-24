package sql

//#include "qsqlquerymodel.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QSqlQueryModel struct {
	core.QAbstractTableModel
}

type QSqlQueryModelITF interface {
	core.QAbstractTableModelITF
	QSqlQueryModelPTR() *QSqlQueryModel
}

func PointerFromQSqlQueryModel(ptr QSqlQueryModelITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlQueryModelPTR().Pointer()
	}
	return nil
}

func QSqlQueryModelFromPointer(ptr unsafe.Pointer) *QSqlQueryModel {
	var n = new(QSqlQueryModel)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QSqlQueryModel_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QSqlQueryModel) QSqlQueryModelPTR() *QSqlQueryModel {
	return ptr
}

func (ptr *QSqlQueryModel) RowCount(parent core.QModelIndexITF) int {
	if ptr.Pointer() != nil {
		return int(C.QSqlQueryModel_RowCount(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQModelIndex(parent))))
	}
	return 0
}

func (ptr *QSqlQueryModel) Data(item core.QModelIndexITF, role int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlQueryModel_Data(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQModelIndex(item)), C.int(role)))
	}
	return ""
}

func (ptr *QSqlQueryModel) CanFetchMore(parent core.QModelIndexITF) bool {
	if ptr.Pointer() != nil {
		return C.QSqlQueryModel_CanFetchMore(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

func (ptr *QSqlQueryModel) Clear() {
	if ptr.Pointer() != nil {
		C.QSqlQueryModel_Clear(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QSqlQueryModel) ColumnCount(index core.QModelIndexITF) int {
	if ptr.Pointer() != nil {
		return int(C.QSqlQueryModel_ColumnCount(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQModelIndex(index))))
	}
	return 0
}

func (ptr *QSqlQueryModel) FetchMore(parent core.QModelIndexITF) {
	if ptr.Pointer() != nil {
		C.QSqlQueryModel_FetchMore(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQModelIndex(parent)))
	}
}

func (ptr *QSqlQueryModel) HeaderData(section int, orientation core.Qt__Orientation, role int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlQueryModel_HeaderData(C.QtObjectPtr(ptr.Pointer()), C.int(section), C.int(orientation), C.int(role)))
	}
	return ""
}

func (ptr *QSqlQueryModel) InsertColumns(column int, count int, parent core.QModelIndexITF) bool {
	if ptr.Pointer() != nil {
		return C.QSqlQueryModel_InsertColumns(C.QtObjectPtr(ptr.Pointer()), C.int(column), C.int(count), C.QtObjectPtr(core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

func (ptr *QSqlQueryModel) RemoveColumns(column int, count int, parent core.QModelIndexITF) bool {
	if ptr.Pointer() != nil {
		return C.QSqlQueryModel_RemoveColumns(C.QtObjectPtr(ptr.Pointer()), C.int(column), C.int(count), C.QtObjectPtr(core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

func (ptr *QSqlQueryModel) SetHeaderData(section int, orientation core.Qt__Orientation, value string, role int) bool {
	if ptr.Pointer() != nil {
		return C.QSqlQueryModel_SetHeaderData(C.QtObjectPtr(ptr.Pointer()), C.int(section), C.int(orientation), C.CString(value), C.int(role)) != 0
	}
	return false
}

func (ptr *QSqlQueryModel) SetQuery(query QSqlQueryITF) {
	if ptr.Pointer() != nil {
		C.QSqlQueryModel_SetQuery(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQSqlQuery(query)))
	}
}

func (ptr *QSqlQueryModel) SetQuery2(query string, db QSqlDatabaseITF) {
	if ptr.Pointer() != nil {
		C.QSqlQueryModel_SetQuery2(C.QtObjectPtr(ptr.Pointer()), C.CString(query), C.QtObjectPtr(PointerFromQSqlDatabase(db)))
	}
}

func (ptr *QSqlQueryModel) DestroyQSqlQueryModel() {
	if ptr.Pointer() != nil {
		C.QSqlQueryModel_DestroyQSqlQueryModel(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
