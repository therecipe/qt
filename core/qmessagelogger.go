package core

//#include "qmessagelogger.h"
import "C"
import (
	"unsafe"
)

type QMessageLogger struct {
	ptr unsafe.Pointer
}

type QMessageLoggerITF interface {
	QMessageLoggerPTR() *QMessageLogger
}

func (p *QMessageLogger) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMessageLogger) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMessageLogger(ptr QMessageLoggerITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMessageLoggerPTR().Pointer()
	}
	return nil
}

func QMessageLoggerFromPointer(ptr unsafe.Pointer) *QMessageLogger {
	var n = new(QMessageLogger)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMessageLogger) QMessageLoggerPTR() *QMessageLogger {
	return ptr
}

func NewQMessageLogger() *QMessageLogger {
	return QMessageLoggerFromPointer(unsafe.Pointer(C.QMessageLogger_NewQMessageLogger()))
}

func NewQMessageLogger2(file string, line int, function string) *QMessageLogger {
	return QMessageLoggerFromPointer(unsafe.Pointer(C.QMessageLogger_NewQMessageLogger2(C.CString(file), C.int(line), C.CString(function))))
}

func NewQMessageLogger3(file string, line int, function string, category string) *QMessageLogger {
	return QMessageLoggerFromPointer(unsafe.Pointer(C.QMessageLogger_NewQMessageLogger3(C.CString(file), C.int(line), C.CString(function), C.CString(category))))
}
