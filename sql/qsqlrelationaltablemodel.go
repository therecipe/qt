package sql

//#include "sql.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QSqlRelationalTableModel struct {
	QSqlTableModel
}

type QSqlRelationalTableModel_ITF interface {
	QSqlTableModel_ITF
	QSqlRelationalTableModel_PTR() *QSqlRelationalTableModel
}

func PointerFromQSqlRelationalTableModel(ptr QSqlRelationalTableModel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlRelationalTableModel_PTR().Pointer()
	}
	return nil
}

func NewQSqlRelationalTableModelFromPointer(ptr unsafe.Pointer) *QSqlRelationalTableModel {
	var n = new(QSqlRelationalTableModel)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QSqlRelationalTableModel_") {
		n.SetObjectName("QSqlRelationalTableModel_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QSqlRelationalTableModel) QSqlRelationalTableModel_PTR() *QSqlRelationalTableModel {
	return ptr
}

//QSqlRelationalTableModel::JoinMode
type QSqlRelationalTableModel__JoinMode int64

const (
	QSqlRelationalTableModel__InnerJoin = QSqlRelationalTableModel__JoinMode(0)
	QSqlRelationalTableModel__LeftJoin  = QSqlRelationalTableModel__JoinMode(1)
)

func (ptr *QSqlRelationalTableModel) Clear() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlRelationalTableModel::clear")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_Clear(ptr.Pointer())
	}
}

func (ptr *QSqlRelationalTableModel) Data(index core.QModelIndex_ITF, role int) *core.QVariant {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlRelationalTableModel::data")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QSqlRelationalTableModel_Data(ptr.Pointer(), core.PointerFromQModelIndex(index), C.int(role)))
	}
	return nil
}

func (ptr *QSqlRelationalTableModel) RelationModel(column int) *QSqlTableModel {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlRelationalTableModel::relationModel")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQSqlTableModelFromPointer(C.QSqlRelationalTableModel_RelationModel(ptr.Pointer(), C.int(column)))
	}
	return nil
}

func (ptr *QSqlRelationalTableModel) RemoveColumns(column int, count int, parent core.QModelIndex_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlRelationalTableModel::removeColumns")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSqlRelationalTableModel_RemoveColumns(ptr.Pointer(), C.int(column), C.int(count), core.PointerFromQModelIndex(parent)) != 0
	}
	return false
}

func (ptr *QSqlRelationalTableModel) RevertRow(row int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlRelationalTableModel::revertRow")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_RevertRow(ptr.Pointer(), C.int(row))
	}
}

func (ptr *QSqlRelationalTableModel) Select() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlRelationalTableModel::select")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSqlRelationalTableModel_Select(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlRelationalTableModel) SetData(index core.QModelIndex_ITF, value core.QVariant_ITF, role int) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlRelationalTableModel::setData")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSqlRelationalTableModel_SetData(ptr.Pointer(), core.PointerFromQModelIndex(index), core.PointerFromQVariant(value), C.int(role)) != 0
	}
	return false
}

func (ptr *QSqlRelationalTableModel) SetJoinMode(joinMode QSqlRelationalTableModel__JoinMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlRelationalTableModel::setJoinMode")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_SetJoinMode(ptr.Pointer(), C.int(joinMode))
	}
}

func (ptr *QSqlRelationalTableModel) SetRelation(column int, relation QSqlRelation_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlRelationalTableModel::setRelation")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_SetRelation(ptr.Pointer(), C.int(column), PointerFromQSqlRelation(relation))
	}
}

func (ptr *QSqlRelationalTableModel) SetTable(table string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlRelationalTableModel::setTable")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_SetTable(ptr.Pointer(), C.CString(table))
	}
}

func (ptr *QSqlRelationalTableModel) DestroyQSqlRelationalTableModel() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlRelationalTableModel::~QSqlRelationalTableModel")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSqlRelationalTableModel_DestroyQSqlRelationalTableModel(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
