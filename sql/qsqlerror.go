package sql

//#include "qsqlerror.h"
import "C"
import (
	"unsafe"
)

type QSqlError struct {
	ptr unsafe.Pointer
}

type QSqlErrorITF interface {
	QSqlErrorPTR() *QSqlError
}

func (p *QSqlError) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSqlError) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSqlError(ptr QSqlErrorITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlErrorPTR().Pointer()
	}
	return nil
}

func QSqlErrorFromPointer(ptr unsafe.Pointer) *QSqlError {
	var n = new(QSqlError)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSqlError) QSqlErrorPTR() *QSqlError {
	return ptr
}

//QSqlError::ErrorType
type QSqlError__ErrorType int

var (
	QSqlError__NoError          = QSqlError__ErrorType(0)
	QSqlError__ConnectionError  = QSqlError__ErrorType(1)
	QSqlError__StatementError   = QSqlError__ErrorType(2)
	QSqlError__TransactionError = QSqlError__ErrorType(3)
	QSqlError__UnknownError     = QSqlError__ErrorType(4)
)

func NewQSqlError3(other QSqlErrorITF) *QSqlError {
	return QSqlErrorFromPointer(unsafe.Pointer(C.QSqlError_NewQSqlError3(C.QtObjectPtr(PointerFromQSqlError(other)))))
}

func NewQSqlError(driverText string, databaseText string, ty QSqlError__ErrorType, code string) *QSqlError {
	return QSqlErrorFromPointer(unsafe.Pointer(C.QSqlError_NewQSqlError(C.CString(driverText), C.CString(databaseText), C.int(ty), C.CString(code))))
}

func (ptr *QSqlError) DatabaseText() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlError_DatabaseText(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QSqlError) DriverText() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlError_DriverText(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QSqlError) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QSqlError_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSqlError) NativeErrorCode() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlError_NativeErrorCode(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QSqlError) Text() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlError_Text(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QSqlError) Type() QSqlError__ErrorType {
	if ptr.Pointer() != nil {
		return QSqlError__ErrorType(C.QSqlError_Type(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSqlError) DestroyQSqlError() {
	if ptr.Pointer() != nil {
		C.QSqlError_DestroyQSqlError(C.QtObjectPtr(ptr.Pointer()))
	}
}
