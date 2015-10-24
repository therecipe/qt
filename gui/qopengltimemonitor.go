package gui

//#include "qopengltimemonitor.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QOpenGLTimeMonitor struct {
	core.QObject
}

type QOpenGLTimeMonitorITF interface {
	core.QObjectITF
	QOpenGLTimeMonitorPTR() *QOpenGLTimeMonitor
}

func PointerFromQOpenGLTimeMonitor(ptr QOpenGLTimeMonitorITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLTimeMonitorPTR().Pointer()
	}
	return nil
}

func QOpenGLTimeMonitorFromPointer(ptr unsafe.Pointer) *QOpenGLTimeMonitor {
	var n = new(QOpenGLTimeMonitor)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QOpenGLTimeMonitor_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QOpenGLTimeMonitor) QOpenGLTimeMonitorPTR() *QOpenGLTimeMonitor {
	return ptr
}

func NewQOpenGLTimeMonitor(parent core.QObjectITF) *QOpenGLTimeMonitor {
	return QOpenGLTimeMonitorFromPointer(unsafe.Pointer(C.QOpenGLTimeMonitor_NewQOpenGLTimeMonitor(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QOpenGLTimeMonitor) Create() bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLTimeMonitor_Create(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QOpenGLTimeMonitor) Destroy() {
	if ptr.Pointer() != nil {
		C.QOpenGLTimeMonitor_Destroy(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLTimeMonitor) IsCreated() bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLTimeMonitor_IsCreated(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QOpenGLTimeMonitor) IsResultAvailable() bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLTimeMonitor_IsResultAvailable(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QOpenGLTimeMonitor) RecordSample() int {
	if ptr.Pointer() != nil {
		return int(C.QOpenGLTimeMonitor_RecordSample(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QOpenGLTimeMonitor) Reset() {
	if ptr.Pointer() != nil {
		C.QOpenGLTimeMonitor_Reset(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLTimeMonitor) SampleCount() int {
	if ptr.Pointer() != nil {
		return int(C.QOpenGLTimeMonitor_SampleCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QOpenGLTimeMonitor) SetSampleCount(sampleCount int) {
	if ptr.Pointer() != nil {
		C.QOpenGLTimeMonitor_SetSampleCount(C.QtObjectPtr(ptr.Pointer()), C.int(sampleCount))
	}
}

func (ptr *QOpenGLTimeMonitor) DestroyQOpenGLTimeMonitor() {
	if ptr.Pointer() != nil {
		C.QOpenGLTimeMonitor_DestroyQOpenGLTimeMonitor(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
