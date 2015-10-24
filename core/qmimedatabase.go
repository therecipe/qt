package core

//#include "qmimedatabase.h"
import "C"
import (
	"unsafe"
)

type QMimeDatabase struct {
	ptr unsafe.Pointer
}

type QMimeDatabaseITF interface {
	QMimeDatabasePTR() *QMimeDatabase
}

func (p *QMimeDatabase) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMimeDatabase) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMimeDatabase(ptr QMimeDatabaseITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMimeDatabasePTR().Pointer()
	}
	return nil
}

func QMimeDatabaseFromPointer(ptr unsafe.Pointer) *QMimeDatabase {
	var n = new(QMimeDatabase)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMimeDatabase) QMimeDatabasePTR() *QMimeDatabase {
	return ptr
}

//QMimeDatabase::MatchMode
type QMimeDatabase__MatchMode int

var (
	QMimeDatabase__MatchDefault   = QMimeDatabase__MatchMode(0x0)
	QMimeDatabase__MatchExtension = QMimeDatabase__MatchMode(0x1)
	QMimeDatabase__MatchContent   = QMimeDatabase__MatchMode(0x2)
)

func NewQMimeDatabase() *QMimeDatabase {
	return QMimeDatabaseFromPointer(unsafe.Pointer(C.QMimeDatabase_NewQMimeDatabase()))
}

func (ptr *QMimeDatabase) DestroyQMimeDatabase() {
	if ptr.Pointer() != nil {
		C.QMimeDatabase_DestroyQMimeDatabase(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QMimeDatabase) SuffixForFileName(fileName string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QMimeDatabase_SuffixForFileName(C.QtObjectPtr(ptr.Pointer()), C.CString(fileName)))
	}
	return ""
}
