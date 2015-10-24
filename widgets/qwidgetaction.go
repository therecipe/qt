package widgets

//#include "qwidgetaction.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QWidgetAction struct {
	QAction
}

type QWidgetActionITF interface {
	QActionITF
	QWidgetActionPTR() *QWidgetAction
}

func PointerFromQWidgetAction(ptr QWidgetActionITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWidgetActionPTR().Pointer()
	}
	return nil
}

func QWidgetActionFromPointer(ptr unsafe.Pointer) *QWidgetAction {
	var n = new(QWidgetAction)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QWidgetAction_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QWidgetAction) QWidgetActionPTR() *QWidgetAction {
	return ptr
}

func NewQWidgetAction(parent core.QObjectITF) *QWidgetAction {
	return QWidgetActionFromPointer(unsafe.Pointer(C.QWidgetAction_NewQWidgetAction(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QWidgetAction) DefaultWidget() *QWidget {
	if ptr.Pointer() != nil {
		return QWidgetFromPointer(unsafe.Pointer(C.QWidgetAction_DefaultWidget(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QWidgetAction) ReleaseWidget(widget QWidgetITF) {
	if ptr.Pointer() != nil {
		C.QWidgetAction_ReleaseWidget(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(widget)))
	}
}

func (ptr *QWidgetAction) RequestWidget(parent QWidgetITF) *QWidget {
	if ptr.Pointer() != nil {
		return QWidgetFromPointer(unsafe.Pointer(C.QWidgetAction_RequestWidget(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(parent)))))
	}
	return nil
}

func (ptr *QWidgetAction) SetDefaultWidget(widget QWidgetITF) {
	if ptr.Pointer() != nil {
		C.QWidgetAction_SetDefaultWidget(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(widget)))
	}
}

func (ptr *QWidgetAction) DestroyQWidgetAction() {
	if ptr.Pointer() != nil {
		C.QWidgetAction_DestroyQWidgetAction(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
