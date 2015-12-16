package quick

//#include "quick.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QQuickItemGrabResult struct {
	core.QObject
}

type QQuickItemGrabResult_ITF interface {
	core.QObject_ITF
	QQuickItemGrabResult_PTR() *QQuickItemGrabResult
}

func PointerFromQQuickItemGrabResult(ptr QQuickItemGrabResult_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickItemGrabResult_PTR().Pointer()
	}
	return nil
}

func NewQQuickItemGrabResultFromPointer(ptr unsafe.Pointer) *QQuickItemGrabResult {
	var n = new(QQuickItemGrabResult)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QQuickItemGrabResult_") {
		n.SetObjectName("QQuickItemGrabResult_" + qt.Identifier())
	}
	return n
}

func (ptr *QQuickItemGrabResult) QQuickItemGrabResult_PTR() *QQuickItemGrabResult {
	return ptr
}

func (ptr *QQuickItemGrabResult) Url() *core.QUrl {
	defer qt.Recovering("QQuickItemGrabResult::url")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QQuickItemGrabResult_Url(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickItemGrabResult) ConnectReady(f func()) {
	defer qt.Recovering("connect QQuickItemGrabResult::ready")

	if ptr.Pointer() != nil {
		C.QQuickItemGrabResult_ConnectReady(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "ready", f)
	}
}

func (ptr *QQuickItemGrabResult) DisconnectReady() {
	defer qt.Recovering("disconnect QQuickItemGrabResult::ready")

	if ptr.Pointer() != nil {
		C.QQuickItemGrabResult_DisconnectReady(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "ready")
	}
}

//export callbackQQuickItemGrabResultReady
func callbackQQuickItemGrabResultReady(ptrName *C.char) {
	defer qt.Recovering("callback QQuickItemGrabResult::ready")

	var signal = qt.GetSignal(C.GoString(ptrName), "ready")
	if signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickItemGrabResult) SaveToFile(fileName string) bool {
	defer qt.Recovering("QQuickItemGrabResult::saveToFile")

	if ptr.Pointer() != nil {
		return C.QQuickItemGrabResult_SaveToFile(ptr.Pointer(), C.CString(fileName)) != 0
	}
	return false
}
