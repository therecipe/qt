package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QMacCocoaViewContainer struct {
	QWidget
}

type QMacCocoaViewContainer_ITF interface {
	QWidget_ITF
	QMacCocoaViewContainer_PTR() *QMacCocoaViewContainer
}

func PointerFromQMacCocoaViewContainer(ptr QMacCocoaViewContainer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMacCocoaViewContainer_PTR().Pointer()
	}
	return nil
}

func NewQMacCocoaViewContainerFromPointer(ptr unsafe.Pointer) *QMacCocoaViewContainer {
	var n = new(QMacCocoaViewContainer)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QMacCocoaViewContainer_") {
		n.SetObjectName("QMacCocoaViewContainer_" + qt.Identifier())
	}
	return n
}

func (ptr *QMacCocoaViewContainer) QMacCocoaViewContainer_PTR() *QMacCocoaViewContainer {
	return ptr
}

func (ptr *QMacCocoaViewContainer) DestroyQMacCocoaViewContainer() {
	defer qt.Recovering("QMacCocoaViewContainer::~QMacCocoaViewContainer")

	if ptr.Pointer() != nil {
		C.QMacCocoaViewContainer_DestroyQMacCocoaViewContainer(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
