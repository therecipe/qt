package sql

//#include "qsqltablemodel.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QSqlTableModel struct {
	QSqlQueryModel
}

type QSqlTableModelITF interface {
	QSqlQueryModelITF
	QSqlTableModelPTR() *QSqlTableModel
}

func PointerFromQSqlTableModel(ptr QSqlTableModelITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlTableModelPTR().Pointer()
	}
	return nil
}

func QSqlTableModelFromPointer(ptr unsafe.Pointer) *QSqlTableModel {
	var n = new(QSqlTableModel)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QSqlTableModel_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QSqlTableModel) QSqlTableModelPTR() *QSqlTableModel {
	return ptr
}

//QSqlTableModel::EditStrategy
type QSqlTableModel__EditStrategy int

var (
	QSqlTableModel__OnFieldChange  = QSqlTableModel__EditStrategy(0)
	QSqlTableModel__OnRowChange    = QSqlTableModel__EditStrategy(1)
	QSqlTableModel__OnManualSubmit = QSqlTableModel__EditStrategy(2)
)

func (ptr *QSqlTableModel) ConnectBeforeDelete(f func(row int)) {
	if ptr.Pointer() != nil {
		C.QSqlTableModel_ConnectBeforeDelete(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "beforeDelete", f)
	}
}

func (ptr *QSqlTableModel) DisconnectBeforeDelete() {
	if ptr.Pointer() != nil {
		C.QSqlTableModel_DisconnectBeforeDelete(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "beforeDelete")
	}
}

//export callbackQSqlTableModelBeforeDelete
func callbackQSqlTableModelBeforeDelete(ptrName *C.char, row C.int) {
	qt.GetSignal(C.GoString(ptrName), "beforeDelete").(func(int))(int(row))
}

func (ptr *QSqlTableModel) Clear() {
	if ptr.Pointer() != nil {
		C.QSqlTableModel_Clear(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QSqlTableModel) Data(index core.QModelIndexITF, role int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlTableModel_Data(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQModelIndex(index)), C.int(role)))
	}
	return ""
}

func (ptr *QSqlTableModel) EditStrategy() QSqlTableModel__EditStrategy {
	if ptr.Pointer() != nil {
		return QSqlTableModel__EditStrategy(C.QSqlTableModel_EditStrategy(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSqlTableModel) FieldIndex(fieldName string) int {
	if ptr.Pointer() != nil {
		return int(C.QSqlTableModel_FieldIndex(C.QtObjectPtr(ptr.Pointer()), C.CString(fieldName)))
	}
	return 0
}

func (ptr *QSqlTableModel) Filter() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlTableModel_Filter(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QSqlTableModel) Flags(index core.QModelIndexITF) core.Qt__ItemFlag {
	if ptr.Pointer() != nil {
		return core.Qt__ItemFlag(C.QSqlTableModel_Flags(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQModelIndex(index))))
	}
	return 0
}

func (ptr *QSqlTableModel) HeaderData(section int, orientation core.Qt__Orientation, role int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlTableModel_HeaderData(C.QtObjectPtr(ptr.Pointer()), C.int(section), C.int(orientation), C.int(role)))
	}
	return ""
}

func (ptr *QSqlTableModel) InsertRecord(row int, record QSqlRecordITF) bool {
	if ptr.Pointer() != nil {
		return C.QSqlTableModel_InsertRecord(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.QtObjectPtr(PointerFromQSqlRecord(record))) != 0
	}
	return false
}

func (ptr *QSqlTableModel) InsertRows(row int, count int, parent core.QModelIndexITF) bool {
	if ptr.Pointer() != nil {
		return C.QSqlTableModel_InsertRows(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.int(count), C.QtObjectPtr(core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

func (ptr *QSqlTableModel) IsDirty2() bool {
	if ptr.Pointer() != nil {
		return C.QSqlTableModel_IsDirty2(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSqlTableModel) IsDirty(index core.QModelIndexITF) bool {
	if ptr.Pointer() != nil {
		return C.QSqlTableModel_IsDirty(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQModelIndex(index))) != 0
	}
	return false
}

func (ptr *QSqlTableModel) RemoveColumns(column int, count int, parent core.QModelIndexITF) bool {
	if ptr.Pointer() != nil {
		return C.QSqlTableModel_RemoveColumns(C.QtObjectPtr(ptr.Pointer()), C.int(column), C.int(count), C.QtObjectPtr(core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

func (ptr *QSqlTableModel) RemoveRows(row int, count int, parent core.QModelIndexITF) bool {
	if ptr.Pointer() != nil {
		return C.QSqlTableModel_RemoveRows(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.int(count), C.QtObjectPtr(core.PointerFromQModelIndex(parent))) != 0
	}
	return false
}

func (ptr *QSqlTableModel) Revert() {
	if ptr.Pointer() != nil {
		C.QSqlTableModel_Revert(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QSqlTableModel) RevertAll() {
	if ptr.Pointer() != nil {
		C.QSqlTableModel_RevertAll(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QSqlTableModel) RevertRow(row int) {
	if ptr.Pointer() != nil {
		C.QSqlTableModel_RevertRow(C.QtObjectPtr(ptr.Pointer()), C.int(row))
	}
}

func (ptr *QSqlTableModel) RowCount(parent core.QModelIndexITF) int {
	if ptr.Pointer() != nil {
		return int(C.QSqlTableModel_RowCount(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQModelIndex(parent))))
	}
	return 0
}

func (ptr *QSqlTableModel) Select() bool {
	if ptr.Pointer() != nil {
		return C.QSqlTableModel_Select(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSqlTableModel) SelectRow(row int) bool {
	if ptr.Pointer() != nil {
		return C.QSqlTableModel_SelectRow(C.QtObjectPtr(ptr.Pointer()), C.int(row)) != 0
	}
	return false
}

func (ptr *QSqlTableModel) SetData(index core.QModelIndexITF, value string, role int) bool {
	if ptr.Pointer() != nil {
		return C.QSqlTableModel_SetData(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQModelIndex(index)), C.CString(value), C.int(role)) != 0
	}
	return false
}

func (ptr *QSqlTableModel) SetEditStrategy(strategy QSqlTableModel__EditStrategy) {
	if ptr.Pointer() != nil {
		C.QSqlTableModel_SetEditStrategy(C.QtObjectPtr(ptr.Pointer()), C.int(strategy))
	}
}

func (ptr *QSqlTableModel) SetFilter(filter string) {
	if ptr.Pointer() != nil {
		C.QSqlTableModel_SetFilter(C.QtObjectPtr(ptr.Pointer()), C.CString(filter))
	}
}

func (ptr *QSqlTableModel) SetRecord(row int, values QSqlRecordITF) bool {
	if ptr.Pointer() != nil {
		return C.QSqlTableModel_SetRecord(C.QtObjectPtr(ptr.Pointer()), C.int(row), C.QtObjectPtr(PointerFromQSqlRecord(values))) != 0
	}
	return false
}

func (ptr *QSqlTableModel) SetSort(column int, order core.Qt__SortOrder) {
	if ptr.Pointer() != nil {
		C.QSqlTableModel_SetSort(C.QtObjectPtr(ptr.Pointer()), C.int(column), C.int(order))
	}
}

func (ptr *QSqlTableModel) SetTable(tableName string) {
	if ptr.Pointer() != nil {
		C.QSqlTableModel_SetTable(C.QtObjectPtr(ptr.Pointer()), C.CString(tableName))
	}
}

func (ptr *QSqlTableModel) Sort(column int, order core.Qt__SortOrder) {
	if ptr.Pointer() != nil {
		C.QSqlTableModel_Sort(C.QtObjectPtr(ptr.Pointer()), C.int(column), C.int(order))
	}
}

func (ptr *QSqlTableModel) Submit() bool {
	if ptr.Pointer() != nil {
		return C.QSqlTableModel_Submit(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSqlTableModel) SubmitAll() bool {
	if ptr.Pointer() != nil {
		return C.QSqlTableModel_SubmitAll(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSqlTableModel) TableName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlTableModel_TableName(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QSqlTableModel) DestroyQSqlTableModel() {
	if ptr.Pointer() != nil {
		C.QSqlTableModel_DestroyQSqlTableModel(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
