package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"strings"
	"unsafe"
)

type QSessionManager struct {
	core.QObject
}

type QSessionManager_ITF interface {
	core.QObject_ITF
	QSessionManager_PTR() *QSessionManager
}

func PointerFromQSessionManager(ptr QSessionManager_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSessionManager_PTR().Pointer()
	}
	return nil
}

func NewQSessionManagerFromPointer(ptr unsafe.Pointer) *QSessionManager {
	var n = new(QSessionManager)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QSessionManager_") {
		n.SetObjectName("QSessionManager_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QSessionManager) QSessionManager_PTR() *QSessionManager {
	return ptr
}

//QSessionManager::RestartHint
type QSessionManager__RestartHint int64

const (
	QSessionManager__RestartIfRunning   = QSessionManager__RestartHint(0)
	QSessionManager__RestartAnyway      = QSessionManager__RestartHint(1)
	QSessionManager__RestartImmediately = QSessionManager__RestartHint(2)
	QSessionManager__RestartNever       = QSessionManager__RestartHint(3)
)

func (ptr *QSessionManager) RestartHint() QSessionManager__RestartHint {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSessionManager::restartHint")
		}
	}()

	if ptr.Pointer() != nil {
		return QSessionManager__RestartHint(C.QSessionManager_RestartHint(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSessionManager) SessionKey() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSessionManager::sessionKey")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QSessionManager_SessionKey(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSessionManager) AllowsErrorInteraction() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSessionManager::allowsErrorInteraction")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSessionManager_AllowsErrorInteraction(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSessionManager) AllowsInteraction() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSessionManager::allowsInteraction")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSessionManager_AllowsInteraction(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSessionManager) Cancel() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSessionManager::cancel")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSessionManager_Cancel(ptr.Pointer())
	}
}

func (ptr *QSessionManager) DiscardCommand() []string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSessionManager::discardCommand")
		}
	}()

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QSessionManager_DiscardCommand(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QSessionManager) IsPhase2() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSessionManager::isPhase2")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSessionManager_IsPhase2(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSessionManager) Release() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSessionManager::release")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSessionManager_Release(ptr.Pointer())
	}
}

func (ptr *QSessionManager) RequestPhase2() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSessionManager::requestPhase2")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSessionManager_RequestPhase2(ptr.Pointer())
	}
}

func (ptr *QSessionManager) RestartCommand() []string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSessionManager::restartCommand")
		}
	}()

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QSessionManager_RestartCommand(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QSessionManager) SessionId() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSessionManager::sessionId")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QSessionManager_SessionId(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSessionManager) SetDiscardCommand(command []string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSessionManager::setDiscardCommand")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSessionManager_SetDiscardCommand(ptr.Pointer(), C.CString(strings.Join(command, ",,,")))
	}
}

func (ptr *QSessionManager) SetManagerProperty2(name string, value string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSessionManager::setManagerProperty")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSessionManager_SetManagerProperty2(ptr.Pointer(), C.CString(name), C.CString(value))
	}
}

func (ptr *QSessionManager) SetManagerProperty(name string, value []string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSessionManager::setManagerProperty")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSessionManager_SetManagerProperty(ptr.Pointer(), C.CString(name), C.CString(strings.Join(value, ",,,")))
	}
}

func (ptr *QSessionManager) SetRestartCommand(command []string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSessionManager::setRestartCommand")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSessionManager_SetRestartCommand(ptr.Pointer(), C.CString(strings.Join(command, ",,,")))
	}
}

func (ptr *QSessionManager) SetRestartHint(hint QSessionManager__RestartHint) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSessionManager::setRestartHint")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSessionManager_SetRestartHint(ptr.Pointer(), C.int(hint))
	}
}
