package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QMimeDatabase struct {
	ptr unsafe.Pointer
}

type QMimeDatabase_ITF interface {
	QMimeDatabase_PTR() *QMimeDatabase
}

func (p *QMimeDatabase) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMimeDatabase) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMimeDatabase(ptr QMimeDatabase_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMimeDatabase_PTR().Pointer()
	}
	return nil
}

func NewQMimeDatabaseFromPointer(ptr unsafe.Pointer) *QMimeDatabase {
	var n = new(QMimeDatabase)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMimeDatabase) QMimeDatabase_PTR() *QMimeDatabase {
	return ptr
}

//QMimeDatabase::MatchMode
type QMimeDatabase__MatchMode int64

const (
	QMimeDatabase__MatchDefault   = QMimeDatabase__MatchMode(0x0)
	QMimeDatabase__MatchExtension = QMimeDatabase__MatchMode(0x1)
	QMimeDatabase__MatchContent   = QMimeDatabase__MatchMode(0x2)
)

func NewQMimeDatabase() *QMimeDatabase {
	defer qt.Recovering("QMimeDatabase::QMimeDatabase")

	return NewQMimeDatabaseFromPointer(C.QMimeDatabase_NewQMimeDatabase())
}

func (ptr *QMimeDatabase) DestroyQMimeDatabase() {
	defer qt.Recovering("QMimeDatabase::~QMimeDatabase")

	if ptr.Pointer() != nil {
		C.QMimeDatabase_DestroyQMimeDatabase(ptr.Pointer())
	}
}

func (ptr *QMimeDatabase) SuffixForFileName(fileName string) string {
	defer qt.Recovering("QMimeDatabase::suffixForFileName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QMimeDatabase_SuffixForFileName(ptr.Pointer(), C.CString(fileName)))
	}
	return ""
}
