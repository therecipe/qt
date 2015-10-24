package widgets

//#include "qshortcut.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QShortcut struct {
	core.QObject
}

type QShortcutITF interface {
	core.QObjectITF
	QShortcutPTR() *QShortcut
}

func PointerFromQShortcut(ptr QShortcutITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QShortcutPTR().Pointer()
	}
	return nil
}

func QShortcutFromPointer(ptr unsafe.Pointer) *QShortcut {
	var n = new(QShortcut)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QShortcut_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QShortcut) QShortcutPTR() *QShortcut {
	return ptr
}

func (ptr *QShortcut) AutoRepeat() bool {
	if ptr.Pointer() != nil {
		return C.QShortcut_AutoRepeat(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QShortcut) Context() core.Qt__ShortcutContext {
	if ptr.Pointer() != nil {
		return core.Qt__ShortcutContext(C.QShortcut_Context(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QShortcut) IsEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QShortcut_IsEnabled(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QShortcut) SetAutoRepeat(on bool) {
	if ptr.Pointer() != nil {
		C.QShortcut_SetAutoRepeat(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QShortcut) SetContext(context core.Qt__ShortcutContext) {
	if ptr.Pointer() != nil {
		C.QShortcut_SetContext(C.QtObjectPtr(ptr.Pointer()), C.int(context))
	}
}

func (ptr *QShortcut) SetEnabled(enable bool) {
	if ptr.Pointer() != nil {
		C.QShortcut_SetEnabled(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QShortcut) SetKey(key gui.QKeySequenceITF) {
	if ptr.Pointer() != nil {
		C.QShortcut_SetKey(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQKeySequence(key)))
	}
}

func (ptr *QShortcut) SetWhatsThis(text string) {
	if ptr.Pointer() != nil {
		C.QShortcut_SetWhatsThis(C.QtObjectPtr(ptr.Pointer()), C.CString(text))
	}
}

func (ptr *QShortcut) WhatsThis() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QShortcut_WhatsThis(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func NewQShortcut(parent QWidgetITF) *QShortcut {
	return QShortcutFromPointer(unsafe.Pointer(C.QShortcut_NewQShortcut(C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func NewQShortcut2(key gui.QKeySequenceITF, parent QWidgetITF, member string, ambiguousMember string, context core.Qt__ShortcutContext) *QShortcut {
	return QShortcutFromPointer(unsafe.Pointer(C.QShortcut_NewQShortcut2(C.QtObjectPtr(gui.PointerFromQKeySequence(key)), C.QtObjectPtr(PointerFromQWidget(parent)), C.CString(member), C.CString(ambiguousMember), C.int(context))))
}

func (ptr *QShortcut) ConnectActivated(f func()) {
	if ptr.Pointer() != nil {
		C.QShortcut_ConnectActivated(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "activated", f)
	}
}

func (ptr *QShortcut) DisconnectActivated() {
	if ptr.Pointer() != nil {
		C.QShortcut_DisconnectActivated(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "activated")
	}
}

//export callbackQShortcutActivated
func callbackQShortcutActivated(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "activated").(func())()
}

func (ptr *QShortcut) ConnectActivatedAmbiguously(f func()) {
	if ptr.Pointer() != nil {
		C.QShortcut_ConnectActivatedAmbiguously(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "activatedAmbiguously", f)
	}
}

func (ptr *QShortcut) DisconnectActivatedAmbiguously() {
	if ptr.Pointer() != nil {
		C.QShortcut_DisconnectActivatedAmbiguously(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "activatedAmbiguously")
	}
}

//export callbackQShortcutActivatedAmbiguously
func callbackQShortcutActivatedAmbiguously(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "activatedAmbiguously").(func())()
}

func (ptr *QShortcut) Id() int {
	if ptr.Pointer() != nil {
		return int(C.QShortcut_Id(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QShortcut) ParentWidget() *QWidget {
	if ptr.Pointer() != nil {
		return QWidgetFromPointer(unsafe.Pointer(C.QShortcut_ParentWidget(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QShortcut) DestroyQShortcut() {
	if ptr.Pointer() != nil {
		C.QShortcut_DestroyQShortcut(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
