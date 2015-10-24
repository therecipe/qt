package quick

//#include "qquickitemgrabresult.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QQuickItemGrabResult struct {
	core.QObject
}

type QQuickItemGrabResultITF interface {
	core.QObjectITF
	QQuickItemGrabResultPTR() *QQuickItemGrabResult
}

func PointerFromQQuickItemGrabResult(ptr QQuickItemGrabResultITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickItemGrabResultPTR().Pointer()
	}
	return nil
}

func QQuickItemGrabResultFromPointer(ptr unsafe.Pointer) *QQuickItemGrabResult {
	var n = new(QQuickItemGrabResult)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QQuickItemGrabResult_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QQuickItemGrabResult) QQuickItemGrabResultPTR() *QQuickItemGrabResult {
	return ptr
}

func (ptr *QQuickItemGrabResult) Url() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QQuickItemGrabResult_Url(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QQuickItemGrabResult) ConnectReady(f func()) {
	if ptr.Pointer() != nil {
		C.QQuickItemGrabResult_ConnectReady(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "ready", f)
	}
}

func (ptr *QQuickItemGrabResult) DisconnectReady() {
	if ptr.Pointer() != nil {
		C.QQuickItemGrabResult_DisconnectReady(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "ready")
	}
}

//export callbackQQuickItemGrabResultReady
func callbackQQuickItemGrabResultReady(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "ready").(func())()
}

func (ptr *QQuickItemGrabResult) SaveToFile(fileName string) bool {
	if ptr.Pointer() != nil {
		return C.QQuickItemGrabResult_SaveToFile(C.QtObjectPtr(ptr.Pointer()), C.CString(fileName)) != 0
	}
	return false
}
