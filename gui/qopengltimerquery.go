package gui

//#include "qopengltimerquery.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QOpenGLTimerQuery struct {
	core.QObject
}

type QOpenGLTimerQuery_ITF interface {
	core.QObject_ITF
	QOpenGLTimerQuery_PTR() *QOpenGLTimerQuery
}

func PointerFromQOpenGLTimerQuery(ptr QOpenGLTimerQuery_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLTimerQuery_PTR().Pointer()
	}
	return nil
}

func NewQOpenGLTimerQueryFromPointer(ptr unsafe.Pointer) *QOpenGLTimerQuery {
	var n = new(QOpenGLTimerQuery)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QOpenGLTimerQuery_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QOpenGLTimerQuery) QOpenGLTimerQuery_PTR() *QOpenGLTimerQuery {
	return ptr
}
