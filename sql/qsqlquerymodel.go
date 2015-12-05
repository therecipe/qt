package sql

//#include "sql.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QSqlQueryModel struct {
	core.QAbstractTableModel
}

type QSqlQueryModel_ITF interface {
	core.QAbstractTableModel_ITF
	QSqlQueryModel_PTR() *QSqlQueryModel
}

func PointerFromQSqlQueryModel(ptr QSqlQueryModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlQueryModel_PTR().Pointer()
	}
	return nil
}

func NewQSqlQueryModelFromPointer(ptr unsafe.Pointer) *QSqlQueryModel {
	var n = new(QSqlQueryModel)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QSqlQueryModel_") {
		n.SetObjectName("QSqlQueryModel_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QSqlQueryModel) QSqlQueryModel_PTR() *QSqlQueryModel {
	return ptr
}

func (ptr *QSqlQueryModel) RowCount(parent core.QModelIndex_ITF) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlQueryModel::rowCount")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QSqlQueryModel_RowCount(ptr.Pointer(), core.PointerFromQModelIndex(parent)))
	}
	return 0
}

func (ptr *QSqlQueryModel) Data(item core.QModelIndex_ITF, role int) *core.QVariant {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlQueryModel::data")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QSqlQueryModel_Data(ptr.Pointer(), core.PointerFromQModelIndex(item), C.int(role)))
	}
	return nil
}

func (ptr *QSqlQueryModel) CanFetchMore(parent core.QModelIndex_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlQueryModel::canFetchMore")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSqlQueryModel_CanFetchMore(ptr.Pointer(), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QSqlQueryModel) Clear() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlQueryModel::clear")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSqlQueryModel_Clear(ptr.Pointer())
	}
}

func (ptr *QSqlQueryModel) ColumnCount(index core.QModelIndex_ITF) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlQueryModel::columnCount")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QSqlQueryModel_ColumnCount(ptr.Pointer(), core.PointerFromQModelIndex(index)))
	}
	return 0
}

func (ptr *QSqlQueryModel) FetchMore(parent core.QModelIndex_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlQueryModel::fetchMore")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSqlQueryModel_FetchMore(ptr.Pointer(), core.PointerFromQModelIndex(parent))
	}
}

func (ptr *QSqlQueryModel) HeaderData(section int, orientation core.Qt__Orientation, role int) *core.QVariant {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlQueryModel::headerData")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QSqlQueryModel_HeaderData(ptr.Pointer(), C.int(section), C.int(orientation), C.int(role)))
	}
	return nil
}

func (ptr *QSqlQueryModel) InsertColumns(column int, count int, parent core.QModelIndex_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlQueryModel::insertColumns")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSqlQueryModel_InsertColumns(ptr.Pointer(), C.int(column), C.int(count), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QSqlQueryModel) RemoveColumns(column int, count int, parent core.QModelIndex_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlQueryModel::removeColumns")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSqlQueryModel_RemoveColumns(ptr.Pointer(), C.int(column), C.int(count), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QSqlQueryModel) SetHeaderData(section int, orientation core.Qt__Orientation, value core.QVariant_ITF, role int) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlQueryModel::setHeaderData")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSqlQueryModel_SetHeaderData(ptr.Pointer(), C.int(section), C.int(orientation), core.PointerFromQVariant(value), C.int(role)) != 0
	}
	return false
}

func (ptr *QSqlQueryModel) SetQuery(query QSqlQuery_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlQueryModel::setQuery")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSqlQueryModel_SetQuery(ptr.Pointer(), PointerFromQSqlQuery(query))
	}
}

func (ptr *QSqlQueryModel) SetQuery2(query string, db QSqlDatabase_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlQueryModel::setQuery")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSqlQueryModel_SetQuery2(ptr.Pointer(), C.CString(query), PointerFromQSqlDatabase(db))
	}
}

func (ptr *QSqlQueryModel) DestroyQSqlQueryModel() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlQueryModel::~QSqlQueryModel")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSqlQueryModel_DestroyQSqlQueryModel(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
