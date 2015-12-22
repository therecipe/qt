package widgets

//#include "widgets.h"
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

type QShortcut_ITF interface {
	core.QObject_ITF
	QShortcut_PTR() *QShortcut
}

func PointerFromQShortcut(ptr QShortcut_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QShortcut_PTR().Pointer()
	}
	return nil
}

func NewQShortcutFromPointer(ptr unsafe.Pointer) *QShortcut {
	var n = new(QShortcut)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QShortcut_") {
		n.SetObjectName("QShortcut_" + qt.Identifier())
	}
	return n
}

func (ptr *QShortcut) QShortcut_PTR() *QShortcut {
	return ptr
}

func (ptr *QShortcut) AutoRepeat() bool {
	defer qt.Recovering("QShortcut::autoRepeat")

	if ptr.Pointer() != nil {
		return C.QShortcut_AutoRepeat(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QShortcut) Context() core.Qt__ShortcutContext {
	defer qt.Recovering("QShortcut::context")

	if ptr.Pointer() != nil {
		return core.Qt__ShortcutContext(C.QShortcut_Context(ptr.Pointer()))
	}
	return 0
}

func (ptr *QShortcut) IsEnabled() bool {
	defer qt.Recovering("QShortcut::isEnabled")

	if ptr.Pointer() != nil {
		return C.QShortcut_IsEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QShortcut) SetAutoRepeat(on bool) {
	defer qt.Recovering("QShortcut::setAutoRepeat")

	if ptr.Pointer() != nil {
		C.QShortcut_SetAutoRepeat(ptr.Pointer(), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QShortcut) SetContext(context core.Qt__ShortcutContext) {
	defer qt.Recovering("QShortcut::setContext")

	if ptr.Pointer() != nil {
		C.QShortcut_SetContext(ptr.Pointer(), C.int(context))
	}
}

func (ptr *QShortcut) SetEnabled(enable bool) {
	defer qt.Recovering("QShortcut::setEnabled")

	if ptr.Pointer() != nil {
		C.QShortcut_SetEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QShortcut) SetKey(key gui.QKeySequence_ITF) {
	defer qt.Recovering("QShortcut::setKey")

	if ptr.Pointer() != nil {
		C.QShortcut_SetKey(ptr.Pointer(), gui.PointerFromQKeySequence(key))
	}
}

func (ptr *QShortcut) SetWhatsThis(text string) {
	defer qt.Recovering("QShortcut::setWhatsThis")

	if ptr.Pointer() != nil {
		C.QShortcut_SetWhatsThis(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QShortcut) WhatsThis() string {
	defer qt.Recovering("QShortcut::whatsThis")

	if ptr.Pointer() != nil {
		return C.GoString(C.QShortcut_WhatsThis(ptr.Pointer()))
	}
	return ""
}

func NewQShortcut(parent QWidget_ITF) *QShortcut {
	defer qt.Recovering("QShortcut::QShortcut")

	return NewQShortcutFromPointer(C.QShortcut_NewQShortcut(PointerFromQWidget(parent)))
}

func NewQShortcut2(key gui.QKeySequence_ITF, parent QWidget_ITF, member string, ambiguousMember string, context core.Qt__ShortcutContext) *QShortcut {
	defer qt.Recovering("QShortcut::QShortcut")

	return NewQShortcutFromPointer(C.QShortcut_NewQShortcut2(gui.PointerFromQKeySequence(key), PointerFromQWidget(parent), C.CString(member), C.CString(ambiguousMember), C.int(context)))
}

func (ptr *QShortcut) ConnectActivated(f func()) {
	defer qt.Recovering("connect QShortcut::activated")

	if ptr.Pointer() != nil {
		C.QShortcut_ConnectActivated(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "activated", f)
	}
}

func (ptr *QShortcut) DisconnectActivated() {
	defer qt.Recovering("disconnect QShortcut::activated")

	if ptr.Pointer() != nil {
		C.QShortcut_DisconnectActivated(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "activated")
	}
}

//export callbackQShortcutActivated
func callbackQShortcutActivated(ptrName *C.char) {
	defer qt.Recovering("callback QShortcut::activated")

	if signal := qt.GetSignal(C.GoString(ptrName), "activated"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QShortcut) ConnectActivatedAmbiguously(f func()) {
	defer qt.Recovering("connect QShortcut::activatedAmbiguously")

	if ptr.Pointer() != nil {
		C.QShortcut_ConnectActivatedAmbiguously(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "activatedAmbiguously", f)
	}
}

func (ptr *QShortcut) DisconnectActivatedAmbiguously() {
	defer qt.Recovering("disconnect QShortcut::activatedAmbiguously")

	if ptr.Pointer() != nil {
		C.QShortcut_DisconnectActivatedAmbiguously(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "activatedAmbiguously")
	}
}

//export callbackQShortcutActivatedAmbiguously
func callbackQShortcutActivatedAmbiguously(ptrName *C.char) {
	defer qt.Recovering("callback QShortcut::activatedAmbiguously")

	if signal := qt.GetSignal(C.GoString(ptrName), "activatedAmbiguously"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QShortcut) Id() int {
	defer qt.Recovering("QShortcut::id")

	if ptr.Pointer() != nil {
		return int(C.QShortcut_Id(ptr.Pointer()))
	}
	return 0
}

func (ptr *QShortcut) ParentWidget() *QWidget {
	defer qt.Recovering("QShortcut::parentWidget")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QShortcut_ParentWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QShortcut) DestroyQShortcut() {
	defer qt.Recovering("QShortcut::~QShortcut")

	if ptr.Pointer() != nil {
		C.QShortcut_DestroyQShortcut(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
