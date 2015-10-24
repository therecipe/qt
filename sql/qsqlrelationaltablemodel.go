package sql

//#include "qsqlrelationaltablemodel.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QSqlRelationalTableModel struct {
	QSqlTableModel
}

type QSqlRelationalTableModelITF interface {
	QSqlTableModelITF
	QSqlRelationalTableModelPTR() *QSqlRelationalTableModel
}

func PointerFromQSqlRelationalTableModel(ptr QSqlRelationalTableModelITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlRelationalTableModelPTR().Pointer()
	}
	return nil
}

func QSqlRelationalTableModelFromPointer(ptr unsafe.Pointer) *QSqlRelationalTableModel {
	var n = new(QSqlRelationalTableModel)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QSqlRelationalTableModel_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QSqlRelationalTableModel) QSqlRelationalTableModelPTR() *QSqlRelationalTableModel {
	return ptr
}

//QSqlRelationalTableModel::JoinMode
type QSqlRelationalTableModel__JoinMode int

var (
	QSqlRelationalTableModel__InnerJoin = QSqlRelationalTableModel__JoinMode(0)
	QSqlRelationalTableModel__LeftJoin  = QSqlRelationalTableModel__JoinMode(1)
)

func (ptr *QSqlRelationalTableModel) Clear() {
	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_Clear(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QSqlRelationalTableModel) Data(index core.QModelIndexITF, role int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlRelationalTableModel_Data(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQModelIndex(index)), C.int(role)))
	}
	return ""
}

func (ptr *QSqlRelationalTableModel) RelationModel(column int) *QSqlTableModel {
	if ptr.Pointer() != nil {
		return QSqlTableModelFromPointer(unsafe.Pointer(C.QSqlRelationalTableModel_RelationModel(C.QtObjectPtr(ptr.Pointer()), C.int(column))))
	}
	return nil
}

func (ptr *QSqlRelationalTableModel) RemoveColumns(column int, count int, parent core.QModelIndexITF) bool {
	if ptr.Pointer() != nil {
		return C.QSqlRelationalTableModel_RemoveColumns(C.QtObjectPtr(ptr.Pointer()), C.int(column), C.int(count), C.QtObjectPtr(core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

func (ptr *QSqlRelationalTableModel) RevertRow(row int) {
	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_RevertRow(C.QtObjectPtr(ptr.Pointer()), C.int(row))
	}
}

func (ptr *QSqlRelationalTableModel) Select() bool {
	if ptr.Pointer() != nil {
		return C.QSqlRelationalTableModel_Select(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSqlRelationalTableModel) SetData(index core.QModelIndexITF, value string, role int) bool {
	if ptr.Pointer() != nil {
		return C.QSqlRelationalTableModel_SetData(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQModelIndex(index)), C.CString(value), C.int(role)) != 0
	}
	return false
}

func (ptr *QSqlRelationalTableModel) SetJoinMode(joinMode QSqlRelationalTableModel__JoinMode) {
	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_SetJoinMode(C.QtObjectPtr(ptr.Pointer()), C.int(joinMode))
	}
}

func (ptr *QSqlRelationalTableModel) SetRelation(column int, relation QSqlRelationITF) {
	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_SetRelation(C.QtObjectPtr(ptr.Pointer()), C.int(column), C.QtObjectPtr(PointerFromQSqlRelation(relation)))
	}
}

func (ptr *QSqlRelationalTableModel) SetTable(table string) {
	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_SetTable(C.QtObjectPtr(ptr.Pointer()), C.CString(table))
	}
}

func (ptr *QSqlRelationalTableModel) DestroyQSqlRelationalTableModel() {
	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_DestroyQSqlRelationalTableModel(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
