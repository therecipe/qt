package sql

//#include "sql.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QSqlTableModel struct {
	QSqlQueryModel
}

type QSqlTableModel_ITF interface {
	QSqlQueryModel_ITF
	QSqlTableModel_PTR() *QSqlTableModel
}

func PointerFromQSqlTableModel(ptr QSqlTableModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlTableModel_PTR().Pointer()
	}
	return nil
}

func NewQSqlTableModelFromPointer(ptr unsafe.Pointer) *QSqlTableModel {
	var n = new(QSqlTableModel)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QSqlTableModel_") {
		n.SetObjectName("QSqlTableModel_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QSqlTableModel) QSqlTableModel_PTR() *QSqlTableModel {
	return ptr
}

//QSqlTableModel::EditStrategy
type QSqlTableModel__EditStrategy int64

const (
	QSqlTableModel__OnFieldChange  = QSqlTableModel__EditStrategy(0)
	QSqlTableModel__OnRowChange    = QSqlTableModel__EditStrategy(1)
	QSqlTableModel__OnManualSubmit = QSqlTableModel__EditStrategy(2)
)

func (ptr *QSqlTableModel) ConnectBeforeDelete(f func(row int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlTableModel::beforeDelete")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSqlTableModel_ConnectBeforeDelete(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "beforeDelete", f)
	}
}

func (ptr *QSqlTableModel) DisconnectBeforeDelete() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlTableModel::beforeDelete")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSqlTableModel_DisconnectBeforeDelete(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "beforeDelete")
	}
}

//export callbackQSqlTableModelBeforeDelete
func callbackQSqlTableModelBeforeDelete(ptrName *C.char, row C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlTableModel::beforeDelete")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "beforeDelete").(func(int))(int(row))
}

func (ptr *QSqlTableModel) Clear() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlTableModel::clear")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSqlTableModel_Clear(ptr.Pointer())
	}
}

func (ptr *QSqlTableModel) Data(index core.QModelIndex_ITF, role int) *core.QVariant {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlTableModel::data")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QSqlTableModel_Data(ptr.Pointer(), core.PointerFromQModelIndex(index), C.int(role)))
	}
	return nil
}

func (ptr *QSqlTableModel) EditStrategy() QSqlTableModel__EditStrategy {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlTableModel::editStrategy")
		}
	}()

	if ptr.Pointer() != nil {
		return QSqlTableModel__EditStrategy(C.QSqlTableModel_EditStrategy(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSqlTableModel) FieldIndex(fieldName string) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlTableModel::fieldIndex")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QSqlTableModel_FieldIndex(ptr.Pointer(), C.CString(fieldName)))
	}
	return 0
}

func (ptr *QSqlTableModel) Filter() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlTableModel::filter")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlTableModel_Filter(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlTableModel) Flags(index core.QModelIndex_ITF) core.Qt__ItemFlag {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlTableModel::flags")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__ItemFlag(C.QSqlTableModel_Flags(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return 0
}

func (ptr *QSqlTableModel) HeaderData(section int, orientation core.Qt__Orientation, role int) *core.QVariant {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlTableModel::headerData")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QSqlTableModel_HeaderData(ptr.Pointer(), C.int(section), C.int(orientation), C.int(role)))
	}
	return nil
}

func (ptr *QSqlTableModel) InsertRecord(row int, record QSqlRecord_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlTableModel::insertRecord")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_InsertRecord(ptr.Pointer(), C.int(row), PointerFromQSqlRecord(record)) != 0
	}
	return false
}

func (ptr *QSqlTableModel) InsertRows(row int, count int, parent core.QModelIndex_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlTableModel::insertRows")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_InsertRows(ptr.Pointer(), C.int(row), C.int(count), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QSqlTableModel) IsDirty2() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlTableModel::isDirty")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_IsDirty2(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlTableModel) IsDirty(index core.QModelIndex_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlTableModel::isDirty")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_IsDirty(ptr.Pointer(), core.PointerFromQModelIndex(index)) != 0
	}
	return false
}

func (ptr *QSqlTableModel) RemoveColumns(column int, count int, parent core.QModelIndex_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlTableModel::removeColumns")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_RemoveColumns(ptr.Pointer(), C.int(column), C.int(count), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QSqlTableModel) RemoveRows(row int, count int, parent core.QModelIndex_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlTableModel::removeRows")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_RemoveRows(ptr.Pointer(), C.int(row), C.int(count), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QSqlTableModel) Revert() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlTableModel::revert")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSqlTableModel_Revert(ptr.Pointer())
	}
}

func (ptr *QSqlTableModel) RevertAll() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlTableModel::revertAll")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSqlTableModel_RevertAll(ptr.Pointer())
	}
}

func (ptr *QSqlTableModel) RevertRow(row int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlTableModel::revertRow")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSqlTableModel_RevertRow(ptr.Pointer(), C.int(row))
	}
}

func (ptr *QSqlTableModel) RowCount(parent core.QModelIndex_ITF) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlTableModel::rowCount")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QSqlTableModel_RowCount(ptr.Pointer(), core.PointerFromQModelIndex(parent)))
	}
	return 0
}

func (ptr *QSqlTableModel) Select() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlTableModel::select")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_Select(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlTableModel) SelectRow(row int) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlTableModel::selectRow")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_SelectRow(ptr.Pointer(), C.int(row)) != 0
	}
	return false
}

func (ptr *QSqlTableModel) SetData(index core.QModelIndex_ITF, value core.QVariant_ITF, role int) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlTableModel::setData")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_SetData(ptr.Pointer(), core.PointerFromQModelIndex(index), core.PointerFromQVariant(value), C.int(role)) != 0
	}
	return false
}

func (ptr *QSqlTableModel) SetEditStrategy(strategy QSqlTableModel__EditStrategy) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlTableModel::setEditStrategy")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSqlTableModel_SetEditStrategy(ptr.Pointer(), C.int(strategy))
	}
}

func (ptr *QSqlTableModel) SetFilter(filter string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlTableModel::setFilter")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSqlTableModel_SetFilter(ptr.Pointer(), C.CString(filter))
	}
}

func (ptr *QSqlTableModel) SetRecord(row int, values QSqlRecord_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlTableModel::setRecord")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_SetRecord(ptr.Pointer(), C.int(row), PointerFromQSqlRecord(values)) != 0
	}
	return false
}

func (ptr *QSqlTableModel) SetSort(column int, order core.Qt__SortOrder) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlTableModel::setSort")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSqlTableModel_SetSort(ptr.Pointer(), C.int(column), C.int(order))
	}
}

func (ptr *QSqlTableModel) SetTable(tableName string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlTableModel::setTable")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSqlTableModel_SetTable(ptr.Pointer(), C.CString(tableName))
	}
}

func (ptr *QSqlTableModel) Sort(column int, order core.Qt__SortOrder) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlTableModel::sort")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSqlTableModel_Sort(ptr.Pointer(), C.int(column), C.int(order))
	}
}

func (ptr *QSqlTableModel) Submit() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlTableModel::submit")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_Submit(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlTableModel) SubmitAll() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlTableModel::submitAll")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSqlTableModel_SubmitAll(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlTableModel) TableName() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlTableModel::tableName")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlTableModel_TableName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlTableModel) DestroyQSqlTableModel() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlTableModel::~QSqlTableModel")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSqlTableModel_DestroyQSqlTableModel(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
