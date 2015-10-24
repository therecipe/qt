package widgets

//#include "qmaccocoaviewcontainer.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QMacCocoaViewContainer struct {
	QWidget
}

type QMacCocoaViewContainerITF interface {
	QWidgetITF
	QMacCocoaViewContainerPTR() *QMacCocoaViewContainer
}

func PointerFromQMacCocoaViewContainer(ptr QMacCocoaViewContainerITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMacCocoaViewContainerPTR().Pointer()
	}
	return nil
}

func QMacCocoaViewContainerFromPointer(ptr unsafe.Pointer) *QMacCocoaViewContainer {
	var n = new(QMacCocoaViewContainer)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QMacCocoaViewContainer_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QMacCocoaViewContainer) QMacCocoaViewContainerPTR() *QMacCocoaViewContainer {
	return ptr
}

func (ptr *QMacCocoaViewContainer) DestroyQMacCocoaViewContainer() {
	if ptr.Pointer() != nil {
		C.QMacCocoaViewContainer_DestroyQMacCocoaViewContainer(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
