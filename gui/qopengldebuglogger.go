package gui

//#include "qopengldebuglogger.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QOpenGLDebugLogger struct {
	core.QObject
}

type QOpenGLDebugLoggerITF interface {
	core.QObjectITF
	QOpenGLDebugLoggerPTR() *QOpenGLDebugLogger
}

func PointerFromQOpenGLDebugLogger(ptr QOpenGLDebugLoggerITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLDebugLoggerPTR().Pointer()
	}
	return nil
}

func QOpenGLDebugLoggerFromPointer(ptr unsafe.Pointer) *QOpenGLDebugLogger {
	var n = new(QOpenGLDebugLogger)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QOpenGLDebugLogger_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QOpenGLDebugLogger) QOpenGLDebugLoggerPTR() *QOpenGLDebugLogger {
	return ptr
}

//QOpenGLDebugLogger::LoggingMode
type QOpenGLDebugLogger__LoggingMode int

var (
	QOpenGLDebugLogger__AsynchronousLogging = QOpenGLDebugLogger__LoggingMode(0)
	QOpenGLDebugLogger__SynchronousLogging  = QOpenGLDebugLogger__LoggingMode(1)
)

func NewQOpenGLDebugLogger(parent core.QObjectITF) *QOpenGLDebugLogger {
	return QOpenGLDebugLoggerFromPointer(unsafe.Pointer(C.QOpenGLDebugLogger_NewQOpenGLDebugLogger(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QOpenGLDebugLogger) DisableMessages(sources QOpenGLDebugMessage__Source, types QOpenGLDebugMessage__Type, severities QOpenGLDebugMessage__Severity) {
	if ptr.Pointer() != nil {
		C.QOpenGLDebugLogger_DisableMessages(C.QtObjectPtr(ptr.Pointer()), C.int(sources), C.int(types), C.int(severities))
	}
}

func (ptr *QOpenGLDebugLogger) EnableMessages(sources QOpenGLDebugMessage__Source, types QOpenGLDebugMessage__Type, severities QOpenGLDebugMessage__Severity) {
	if ptr.Pointer() != nil {
		C.QOpenGLDebugLogger_EnableMessages(C.QtObjectPtr(ptr.Pointer()), C.int(sources), C.int(types), C.int(severities))
	}
}

func (ptr *QOpenGLDebugLogger) Initialize() bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLDebugLogger_Initialize(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QOpenGLDebugLogger) IsLogging() bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLDebugLogger_IsLogging(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QOpenGLDebugLogger) LogMessage(debugMessage QOpenGLDebugMessageITF) {
	if ptr.Pointer() != nil {
		C.QOpenGLDebugLogger_LogMessage(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQOpenGLDebugMessage(debugMessage)))
	}
}

func (ptr *QOpenGLDebugLogger) LoggingMode() QOpenGLDebugLogger__LoggingMode {
	if ptr.Pointer() != nil {
		return QOpenGLDebugLogger__LoggingMode(C.QOpenGLDebugLogger_LoggingMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QOpenGLDebugLogger) PopGroup() {
	if ptr.Pointer() != nil {
		C.QOpenGLDebugLogger_PopGroup(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLDebugLogger) StartLogging(loggingMode QOpenGLDebugLogger__LoggingMode) {
	if ptr.Pointer() != nil {
		C.QOpenGLDebugLogger_StartLogging(C.QtObjectPtr(ptr.Pointer()), C.int(loggingMode))
	}
}

func (ptr *QOpenGLDebugLogger) StopLogging() {
	if ptr.Pointer() != nil {
		C.QOpenGLDebugLogger_StopLogging(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLDebugLogger) DestroyQOpenGLDebugLogger() {
	if ptr.Pointer() != nil {
		C.QOpenGLDebugLogger_DestroyQOpenGLDebugLogger(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
