package sensors

//#include "qsensorgesturemanager.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

type QSensorGestureManager struct {
	core.QObject
}

type QSensorGestureManagerITF interface {
	core.QObjectITF
	QSensorGestureManagerPTR() *QSensorGestureManager
}

func PointerFromQSensorGestureManager(ptr QSensorGestureManagerITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorGestureManagerPTR().Pointer()
	}
	return nil
}

func QSensorGestureManagerFromPointer(ptr unsafe.Pointer) *QSensorGestureManager {
	var n = new(QSensorGestureManager)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QSensorGestureManager_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QSensorGestureManager) QSensorGestureManagerPTR() *QSensorGestureManager {
	return ptr
}

func NewQSensorGestureManager(parent core.QObjectITF) *QSensorGestureManager {
	return QSensorGestureManagerFromPointer(unsafe.Pointer(C.QSensorGestureManager_NewQSensorGestureManager(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QSensorGestureManager) GestureIds() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QSensorGestureManager_GestureIds(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QSensorGestureManager) ConnectNewSensorGestureAvailable(f func()) {
	if ptr.Pointer() != nil {
		C.QSensorGestureManager_ConnectNewSensorGestureAvailable(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "newSensorGestureAvailable", f)
	}
}

func (ptr *QSensorGestureManager) DisconnectNewSensorGestureAvailable() {
	if ptr.Pointer() != nil {
		C.QSensorGestureManager_DisconnectNewSensorGestureAvailable(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "newSensorGestureAvailable")
	}
}

//export callbackQSensorGestureManagerNewSensorGestureAvailable
func callbackQSensorGestureManagerNewSensorGestureAvailable(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "newSensorGestureAvailable").(func())()
}

func (ptr *QSensorGestureManager) RecognizerSignals(gestureId string) []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QSensorGestureManager_RecognizerSignals(C.QtObjectPtr(ptr.Pointer()), C.CString(gestureId))), "|")
	}
	return make([]string, 0)
}

func (ptr *QSensorGestureManager) RegisterSensorGestureRecognizer(recognizer QSensorGestureRecognizerITF) bool {
	if ptr.Pointer() != nil {
		return C.QSensorGestureManager_RegisterSensorGestureRecognizer(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQSensorGestureRecognizer(recognizer))) != 0
	}
	return false
}

func QSensorGestureManager_SensorGestureRecognizer(id string) *QSensorGestureRecognizer {
	return QSensorGestureRecognizerFromPointer(unsafe.Pointer(C.QSensorGestureManager_QSensorGestureManager_SensorGestureRecognizer(C.CString(id))))
}

func (ptr *QSensorGestureManager) DestroyQSensorGestureManager() {
	if ptr.Pointer() != nil {
		C.QSensorGestureManager_DestroyQSensorGestureManager(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
