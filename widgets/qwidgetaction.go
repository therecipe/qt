package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QWidgetAction struct {
	QAction
}

type QWidgetAction_ITF interface {
	QAction_ITF
	QWidgetAction_PTR() *QWidgetAction
}

func PointerFromQWidgetAction(ptr QWidgetAction_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWidgetAction_PTR().Pointer()
	}
	return nil
}

func NewQWidgetActionFromPointer(ptr unsafe.Pointer) *QWidgetAction {
	var n = new(QWidgetAction)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QWidgetAction_") {
		n.SetObjectName("QWidgetAction_" + qt.Identifier())
	}
	return n
}

func (ptr *QWidgetAction) QWidgetAction_PTR() *QWidgetAction {
	return ptr
}

func NewQWidgetAction(parent core.QObject_ITF) *QWidgetAction {
	defer qt.Recovering("QWidgetAction::QWidgetAction")

	return NewQWidgetActionFromPointer(C.QWidgetAction_NewQWidgetAction(core.PointerFromQObject(parent)))
}

func (ptr *QWidgetAction) DefaultWidget() *QWidget {
	defer qt.Recovering("QWidgetAction::defaultWidget")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QWidgetAction_DefaultWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWidgetAction) ConnectDeleteWidget(f func(widget *QWidget)) {
	defer qt.Recovering("connect QWidgetAction::deleteWidget")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteWidget", f)
	}
}

func (ptr *QWidgetAction) DisconnectDeleteWidget() {
	defer qt.Recovering("disconnect QWidgetAction::deleteWidget")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteWidget")
	}
}

//export callbackQWidgetActionDeleteWidget
func callbackQWidgetActionDeleteWidget(ptrName *C.char, widget unsafe.Pointer) bool {
	defer qt.Recovering("callback QWidgetAction::deleteWidget")

	var signal = qt.GetSignal(C.GoString(ptrName), "deleteWidget")
	if signal != nil {
		defer signal.(func(*QWidget))(NewQWidgetFromPointer(widget))
		return true
	}
	return false

}

func (ptr *QWidgetAction) ReleaseWidget(widget QWidget_ITF) {
	defer qt.Recovering("QWidgetAction::releaseWidget")

	if ptr.Pointer() != nil {
		C.QWidgetAction_ReleaseWidget(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QWidgetAction) RequestWidget(parent QWidget_ITF) *QWidget {
	defer qt.Recovering("QWidgetAction::requestWidget")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QWidgetAction_RequestWidget(ptr.Pointer(), PointerFromQWidget(parent)))
	}
	return nil
}

func (ptr *QWidgetAction) SetDefaultWidget(widget QWidget_ITF) {
	defer qt.Recovering("QWidgetAction::setDefaultWidget")

	if ptr.Pointer() != nil {
		C.QWidgetAction_SetDefaultWidget(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QWidgetAction) DestroyQWidgetAction() {
	defer qt.Recovering("QWidgetAction::~QWidgetAction")

	if ptr.Pointer() != nil {
		C.QWidgetAction_DestroyQWidgetAction(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
