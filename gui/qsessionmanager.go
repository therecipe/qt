package gui

//#include "qsessionmanager.h"
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

type QSessionManagerITF interface {
	core.QObjectITF
	QSessionManagerPTR() *QSessionManager
}

func PointerFromQSessionManager(ptr QSessionManagerITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSessionManagerPTR().Pointer()
	}
	return nil
}

func QSessionManagerFromPointer(ptr unsafe.Pointer) *QSessionManager {
	var n = new(QSessionManager)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QSessionManager_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QSessionManager) QSessionManagerPTR() *QSessionManager {
	return ptr
}

//QSessionManager::RestartHint
type QSessionManager__RestartHint int

var (
	QSessionManager__RestartIfRunning   = QSessionManager__RestartHint(0)
	QSessionManager__RestartAnyway      = QSessionManager__RestartHint(1)
	QSessionManager__RestartImmediately = QSessionManager__RestartHint(2)
	QSessionManager__RestartNever       = QSessionManager__RestartHint(3)
)

func (ptr *QSessionManager) RestartHint() QSessionManager__RestartHint {
	if ptr.Pointer() != nil {
		return QSessionManager__RestartHint(C.QSessionManager_RestartHint(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSessionManager) SessionKey() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSessionManager_SessionKey(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QSessionManager) AllowsErrorInteraction() bool {
	if ptr.Pointer() != nil {
		return C.QSessionManager_AllowsErrorInteraction(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSessionManager) AllowsInteraction() bool {
	if ptr.Pointer() != nil {
		return C.QSessionManager_AllowsInteraction(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSessionManager) Cancel() {
	if ptr.Pointer() != nil {
		C.QSessionManager_Cancel(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QSessionManager) DiscardCommand() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QSessionManager_DiscardCommand(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QSessionManager) IsPhase2() bool {
	if ptr.Pointer() != nil {
		return C.QSessionManager_IsPhase2(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSessionManager) Release() {
	if ptr.Pointer() != nil {
		C.QSessionManager_Release(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QSessionManager) RequestPhase2() {
	if ptr.Pointer() != nil {
		C.QSessionManager_RequestPhase2(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QSessionManager) RestartCommand() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QSessionManager_RestartCommand(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QSessionManager) SessionId() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSessionManager_SessionId(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QSessionManager) SetDiscardCommand(command []string) {
	if ptr.Pointer() != nil {
		C.QSessionManager_SetDiscardCommand(C.QtObjectPtr(ptr.Pointer()), C.CString(strings.Join(command, "|")))
	}
}

func (ptr *QSessionManager) SetManagerProperty2(name string, value string) {
	if ptr.Pointer() != nil {
		C.QSessionManager_SetManagerProperty2(C.QtObjectPtr(ptr.Pointer()), C.CString(name), C.CString(value))
	}
}

func (ptr *QSessionManager) SetManagerProperty(name string, value []string) {
	if ptr.Pointer() != nil {
		C.QSessionManager_SetManagerProperty(C.QtObjectPtr(ptr.Pointer()), C.CString(name), C.CString(strings.Join(value, "|")))
	}
}

func (ptr *QSessionManager) SetRestartCommand(command []string) {
	if ptr.Pointer() != nil {
		C.QSessionManager_SetRestartCommand(C.QtObjectPtr(ptr.Pointer()), C.CString(strings.Join(command, "|")))
	}
}

func (ptr *QSessionManager) SetRestartHint(hint QSessionManager__RestartHint) {
	if ptr.Pointer() != nil {
		C.QSessionManager_SetRestartHint(C.QtObjectPtr(ptr.Pointer()), C.int(hint))
	}
}
