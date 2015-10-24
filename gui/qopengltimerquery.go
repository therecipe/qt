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

type QOpenGLTimerQueryITF interface {
	core.QObjectITF
	QOpenGLTimerQueryPTR() *QOpenGLTimerQuery
}

func PointerFromQOpenGLTimerQuery(ptr QOpenGLTimerQueryITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLTimerQueryPTR().Pointer()
	}
	return nil
}

func QOpenGLTimerQueryFromPointer(ptr unsafe.Pointer) *QOpenGLTimerQuery {
	var n = new(QOpenGLTimerQuery)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QOpenGLTimerQuery_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QOpenGLTimerQuery) QOpenGLTimerQueryPTR() *QOpenGLTimerQuery {
	return ptr
}

func NewQOpenGLTimerQuery(parent core.QObjectITF) *QOpenGLTimerQuery {
	return QOpenGLTimerQueryFromPointer(unsafe.Pointer(C.QOpenGLTimerQuery_NewQOpenGLTimerQuery(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QOpenGLTimerQuery) Begin() {
	if ptr.Pointer() != nil {
		C.QOpenGLTimerQuery_Begin(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLTimerQuery) Create() bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLTimerQuery_Create(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QOpenGLTimerQuery) Destroy() {
	if ptr.Pointer() != nil {
		C.QOpenGLTimerQuery_Destroy(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLTimerQuery) End() {
	if ptr.Pointer() != nil {
		C.QOpenGLTimerQuery_End(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLTimerQuery) IsCreated() bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLTimerQuery_IsCreated(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QOpenGLTimerQuery) IsResultAvailable() bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLTimerQuery_IsResultAvailable(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QOpenGLTimerQuery) RecordTimestamp() {
	if ptr.Pointer() != nil {
		C.QOpenGLTimerQuery_RecordTimestamp(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLTimerQuery) DestroyQOpenGLTimerQuery() {
	if ptr.Pointer() != nil {
		C.QOpenGLTimerQuery_DestroyQOpenGLTimerQuery(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
