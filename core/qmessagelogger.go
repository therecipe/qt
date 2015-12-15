package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QMessageLogger struct {
	ptr unsafe.Pointer
}

type QMessageLogger_ITF interface {
	QMessageLogger_PTR() *QMessageLogger
}

func (p *QMessageLogger) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMessageLogger) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMessageLogger(ptr QMessageLogger_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMessageLogger_PTR().Pointer()
	}
	return nil
}

func NewQMessageLoggerFromPointer(ptr unsafe.Pointer) *QMessageLogger {
	var n = new(QMessageLogger)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMessageLogger) QMessageLogger_PTR() *QMessageLogger {
	return ptr
}

func NewQMessageLogger() *QMessageLogger {
	defer qt.Recovering("QMessageLogger::QMessageLogger")

	return NewQMessageLoggerFromPointer(C.QMessageLogger_NewQMessageLogger())
}

func NewQMessageLogger2(file string, line int, function string) *QMessageLogger {
	defer qt.Recovering("QMessageLogger::QMessageLogger")

	return NewQMessageLoggerFromPointer(C.QMessageLogger_NewQMessageLogger2(C.CString(file), C.int(line), C.CString(function)))
}

func NewQMessageLogger3(file string, line int, function string, category string) *QMessageLogger {
	defer qt.Recovering("QMessageLogger::QMessageLogger")

	return NewQMessageLoggerFromPointer(C.QMessageLogger_NewQMessageLogger3(C.CString(file), C.int(line), C.CString(function), C.CString(category)))
}
