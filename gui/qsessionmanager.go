package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
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
		n.SetObjectName("QSessionManager_" + qt.Identifier())
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
	defer qt.Recovering("QSessionManager::restartHint")

	if ptr.Pointer() != nil {
		return QSessionManager__RestartHint(C.QSessionManager_RestartHint(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSessionManager) SessionKey() string {
	defer qt.Recovering("QSessionManager::sessionKey")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSessionManager_SessionKey(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSessionManager) AllowsErrorInteraction() bool {
	defer qt.Recovering("QSessionManager::allowsErrorInteraction")

	if ptr.Pointer() != nil {
		return C.QSessionManager_AllowsErrorInteraction(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSessionManager) AllowsInteraction() bool {
	defer qt.Recovering("QSessionManager::allowsInteraction")

	if ptr.Pointer() != nil {
		return C.QSessionManager_AllowsInteraction(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSessionManager) Cancel() {
	defer qt.Recovering("QSessionManager::cancel")

	if ptr.Pointer() != nil {
		C.QSessionManager_Cancel(ptr.Pointer())
	}
}

func (ptr *QSessionManager) DiscardCommand() []string {
	defer qt.Recovering("QSessionManager::discardCommand")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QSessionManager_DiscardCommand(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QSessionManager) IsPhase2() bool {
	defer qt.Recovering("QSessionManager::isPhase2")

	if ptr.Pointer() != nil {
		return C.QSessionManager_IsPhase2(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSessionManager) Release() {
	defer qt.Recovering("QSessionManager::release")

	if ptr.Pointer() != nil {
		C.QSessionManager_Release(ptr.Pointer())
	}
}

func (ptr *QSessionManager) RequestPhase2() {
	defer qt.Recovering("QSessionManager::requestPhase2")

	if ptr.Pointer() != nil {
		C.QSessionManager_RequestPhase2(ptr.Pointer())
	}
}

func (ptr *QSessionManager) RestartCommand() []string {
	defer qt.Recovering("QSessionManager::restartCommand")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QSessionManager_RestartCommand(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QSessionManager) SessionId() string {
	defer qt.Recovering("QSessionManager::sessionId")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSessionManager_SessionId(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSessionManager) SetDiscardCommand(command []string) {
	defer qt.Recovering("QSessionManager::setDiscardCommand")

	if ptr.Pointer() != nil {
		C.QSessionManager_SetDiscardCommand(ptr.Pointer(), C.CString(strings.Join(command, ",,,")))
	}
}

func (ptr *QSessionManager) SetManagerProperty2(name string, value string) {
	defer qt.Recovering("QSessionManager::setManagerProperty")

	if ptr.Pointer() != nil {
		C.QSessionManager_SetManagerProperty2(ptr.Pointer(), C.CString(name), C.CString(value))
	}
}

func (ptr *QSessionManager) SetManagerProperty(name string, value []string) {
	defer qt.Recovering("QSessionManager::setManagerProperty")

	if ptr.Pointer() != nil {
		C.QSessionManager_SetManagerProperty(ptr.Pointer(), C.CString(name), C.CString(strings.Join(value, ",,,")))
	}
}

func (ptr *QSessionManager) SetRestartCommand(command []string) {
	defer qt.Recovering("QSessionManager::setRestartCommand")

	if ptr.Pointer() != nil {
		C.QSessionManager_SetRestartCommand(ptr.Pointer(), C.CString(strings.Join(command, ",,,")))
	}
}

func (ptr *QSessionManager) SetRestartHint(hint QSessionManager__RestartHint) {
	defer qt.Recovering("QSessionManager::setRestartHint")

	if ptr.Pointer() != nil {
		C.QSessionManager_SetRestartHint(ptr.Pointer(), C.int(hint))
	}
}
