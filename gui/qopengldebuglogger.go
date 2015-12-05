package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QOpenGLDebugLogger struct {
	core.QObject
}

type QOpenGLDebugLogger_ITF interface {
	core.QObject_ITF
	QOpenGLDebugLogger_PTR() *QOpenGLDebugLogger
}

func PointerFromQOpenGLDebugLogger(ptr QOpenGLDebugLogger_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLDebugLogger_PTR().Pointer()
	}
	return nil
}

func NewQOpenGLDebugLoggerFromPointer(ptr unsafe.Pointer) *QOpenGLDebugLogger {
	var n = new(QOpenGLDebugLogger)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QOpenGLDebugLogger_") {
		n.SetObjectName("QOpenGLDebugLogger_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QOpenGLDebugLogger) QOpenGLDebugLogger_PTR() *QOpenGLDebugLogger {
	return ptr
}

//QOpenGLDebugLogger::LoggingMode
type QOpenGLDebugLogger__LoggingMode int64

const (
	QOpenGLDebugLogger__AsynchronousLogging = QOpenGLDebugLogger__LoggingMode(0)
	QOpenGLDebugLogger__SynchronousLogging  = QOpenGLDebugLogger__LoggingMode(1)
)
