package sql

//#include "sql.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
	"unsafe"
)

type QSqlIndex struct {
	QSqlRecord
}

type QSqlIndex_ITF interface {
	QSqlRecord_ITF
	QSqlIndex_PTR() *QSqlIndex
}

func PointerFromQSqlIndex(ptr QSqlIndex_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlIndex_PTR().Pointer()
	}
	return nil
}

func NewQSqlIndexFromPointer(ptr unsafe.Pointer) *QSqlIndex {
	var n = new(QSqlIndex)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSqlIndex) QSqlIndex_PTR() *QSqlIndex {
	return ptr
}

func NewQSqlIndex2(other QSqlIndex_ITF) *QSqlIndex {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlIndex::QSqlIndex")
		}
	}()

	return NewQSqlIndexFromPointer(C.QSqlIndex_NewQSqlIndex2(PointerFromQSqlIndex(other)))
}

func NewQSqlIndex(cursorname string, name string) *QSqlIndex {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlIndex::QSqlIndex")
		}
	}()

	return NewQSqlIndexFromPointer(C.QSqlIndex_NewQSqlIndex(C.CString(cursorname), C.CString(name)))
}

func (ptr *QSqlIndex) Append(field QSqlField_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlIndex::append")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSqlIndex_Append(ptr.Pointer(), PointerFromQSqlField(field))
	}
}

func (ptr *QSqlIndex) Append2(field QSqlField_ITF, desc bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlIndex::append")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSqlIndex_Append2(ptr.Pointer(), PointerFromQSqlField(field), C.int(qt.GoBoolToInt(desc)))
	}
}

func (ptr *QSqlIndex) CursorName() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlIndex::cursorName")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlIndex_CursorName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlIndex) IsDescending(i int) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlIndex::isDescending")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSqlIndex_IsDescending(ptr.Pointer(), C.int(i)) != 0
	}
	return false
}

func (ptr *QSqlIndex) Name() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlIndex::name")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlIndex_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlIndex) SetCursorName(cursorName string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlIndex::setCursorName")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSqlIndex_SetCursorName(ptr.Pointer(), C.CString(cursorName))
	}
}

func (ptr *QSqlIndex) SetDescending(i int, desc bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlIndex::setDescending")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSqlIndex_SetDescending(ptr.Pointer(), C.int(i), C.int(qt.GoBoolToInt(desc)))
	}
}

func (ptr *QSqlIndex) SetName(name string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlIndex::setName")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSqlIndex_SetName(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QSqlIndex) DestroyQSqlIndex() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSqlIndex::~QSqlIndex")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSqlIndex_DestroyQSqlIndex(ptr.Pointer())
	}
}
