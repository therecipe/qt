package sensors

//#include "sensors.h"
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

type QSensorGestureManager_ITF interface {
	core.QObject_ITF
	QSensorGestureManager_PTR() *QSensorGestureManager
}

func PointerFromQSensorGestureManager(ptr QSensorGestureManager_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorGestureManager_PTR().Pointer()
	}
	return nil
}

func NewQSensorGestureManagerFromPointer(ptr unsafe.Pointer) *QSensorGestureManager {
	var n = new(QSensorGestureManager)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QSensorGestureManager_") {
		n.SetObjectName("QSensorGestureManager_" + qt.Identifier())
	}
	return n
}

func (ptr *QSensorGestureManager) QSensorGestureManager_PTR() *QSensorGestureManager {
	return ptr
}

func NewQSensorGestureManager(parent core.QObject_ITF) *QSensorGestureManager {
	defer qt.Recovering("QSensorGestureManager::QSensorGestureManager")

	return NewQSensorGestureManagerFromPointer(C.QSensorGestureManager_NewQSensorGestureManager(core.PointerFromQObject(parent)))
}

func (ptr *QSensorGestureManager) GestureIds() []string {
	defer qt.Recovering("QSensorGestureManager::gestureIds")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QSensorGestureManager_GestureIds(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QSensorGestureManager) ConnectNewSensorGestureAvailable(f func()) {
	defer qt.Recovering("connect QSensorGestureManager::newSensorGestureAvailable")

	if ptr.Pointer() != nil {
		C.QSensorGestureManager_ConnectNewSensorGestureAvailable(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "newSensorGestureAvailable", f)
	}
}

func (ptr *QSensorGestureManager) DisconnectNewSensorGestureAvailable() {
	defer qt.Recovering("disconnect QSensorGestureManager::newSensorGestureAvailable")

	if ptr.Pointer() != nil {
		C.QSensorGestureManager_DisconnectNewSensorGestureAvailable(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "newSensorGestureAvailable")
	}
}

//export callbackQSensorGestureManagerNewSensorGestureAvailable
func callbackQSensorGestureManagerNewSensorGestureAvailable(ptrName *C.char) {
	defer qt.Recovering("callback QSensorGestureManager::newSensorGestureAvailable")

	if signal := qt.GetSignal(C.GoString(ptrName), "newSensorGestureAvailable"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QSensorGestureManager) RecognizerSignals(gestureId string) []string {
	defer qt.Recovering("QSensorGestureManager::recognizerSignals")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QSensorGestureManager_RecognizerSignals(ptr.Pointer(), C.CString(gestureId))), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QSensorGestureManager) RegisterSensorGestureRecognizer(recognizer QSensorGestureRecognizer_ITF) bool {
	defer qt.Recovering("QSensorGestureManager::registerSensorGestureRecognizer")

	if ptr.Pointer() != nil {
		return C.QSensorGestureManager_RegisterSensorGestureRecognizer(ptr.Pointer(), PointerFromQSensorGestureRecognizer(recognizer)) != 0
	}
	return false
}

func QSensorGestureManager_SensorGestureRecognizer(id string) *QSensorGestureRecognizer {
	defer qt.Recovering("QSensorGestureManager::sensorGestureRecognizer")

	return NewQSensorGestureRecognizerFromPointer(C.QSensorGestureManager_QSensorGestureManager_SensorGestureRecognizer(C.CString(id)))
}

func (ptr *QSensorGestureManager) DestroyQSensorGestureManager() {
	defer qt.Recovering("QSensorGestureManager::~QSensorGestureManager")

	if ptr.Pointer() != nil {
		C.QSensorGestureManager_DestroyQSensorGestureManager(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
