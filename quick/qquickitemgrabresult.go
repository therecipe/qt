package quick

//#include "quick.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
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
		n.SetObjectName("QQuickItemGrabResult_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QQuickItemGrabResult) QQuickItemGrabResult_PTR() *QQuickItemGrabResult {
	return ptr
}

func (ptr *QQuickItemGrabResult) ConnectReady(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickItemGrabResult::ready")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickItemGrabResult_ConnectReady(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "ready", f)
	}
}

func (ptr *QQuickItemGrabResult) DisconnectReady() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickItemGrabResult::ready")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickItemGrabResult_DisconnectReady(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "ready")
	}
}

//export callbackQQuickItemGrabResultReady
func callbackQQuickItemGrabResultReady(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickItemGrabResult::ready")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "ready").(func())()
}

func (ptr *QQuickItemGrabResult) SaveToFile(fileName string) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickItemGrabResult::saveToFile")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QQuickItemGrabResult_SaveToFile(ptr.Pointer(), C.CString(fileName)) != 0
	}
	return false
}
